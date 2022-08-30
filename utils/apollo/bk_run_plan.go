package apollo

import (
	"rm/common"
)

//通量模式
type throughputDefinition struct {
	ThroughputModeKey  int32   `json:"throughputModeKey"`
	ThroughputModeName string  `json:"throughputModeName"`
	ScanDataKey        int32   `json:"scanDataKey"`
	SupportPlatformKey []int32 `json:"supportPlatformKey"`
}

//读长模式
type readlengthDefinition struct {
	ReadlengthModeKey         int32   `json:"readlengthModeKey"`
	ReadlengthModeDisplayName string  `json:"readlengthModeDisplayName"`
	ReadlengthModeType        int32   `json:"readlengthModeType"`
	SupportPlatformKey        []int32 `json:"supportPlatformKey"`
	SupportThroughputKey      []int32 `json:"supportThroughputKey"`
}

//内参
type referenceMap struct {
	ReferenceKey         int32  `json:"referenceKey"`
	ReferenceDisplayName string `json:"referenceDisplayName"`
	PlatformKey          int32  `json:"platformKey"`
}

type referenceDetail struct {
	ReferenceId  int32  `json:"referenceId"`
	SampleID     string `json:"sampleID"`
	Description  string `json:"description"`
	Index1ID     string `json:"index1ID" mapstructure:"Index1-ID"`
	Index1       string `json:"index1"`
	Index2ID     string `json:"index2ID" mapstructure:"Index2-ID"`
	Index2       string `json:"index2"`
	Read1IndexID string `json:"read1IndexID" mapstructure:"Read1Index-ID"`
	Read1Index   string `json:"read1Index"`
	RefGenome    string `json:"refGenome"`
}

type BKRunPlan struct {
	ThroughputDefinitions []*throughputDefinition
	ReadlengthDefinitions []*readlengthDefinition
	ReferenceMaps         []*referenceMap
	ReferenceDetailModels []*referenceDetail

	tdMap map[int32]*throughputDefinition
	rdMap map[int32]*readlengthDefinition
	rfMap map[int32]*referenceMap
	rmMap map[int32]*referenceDetail

	ptdMap  map[int32][]*throughputDefinition //map[platformKey][]*throughputDefinition 根据平台不同区分通量模式
	ptrdMap map[int32][]*readlengthDefinition //map[platformKey * 10 + throughputModeKey][]*readlengthDefinition 根据平台和通量不同区分读长模式
	prfMap  map[int32][]*referenceMap         //map[platformKey][]*throughputDefinition 根据平台不同区分内参
}

func (p *BKRunPlan) setup() {
	p.tdMap = map[int32]*throughputDefinition{}
	p.ptdMap = map[int32][]*throughputDefinition{}
	for _, v := range p.ThroughputDefinitions {
		p.tdMap[v.ThroughputModeKey] = v
		for _, platformKey := range v.SupportPlatformKey {
			p.ptdMap[platformKey] = append(p.ptdMap[platformKey], v)
		}
	}
	p.rdMap = map[int32]*readlengthDefinition{}
	p.ptrdMap = map[int32][]*readlengthDefinition{}
	for _, v := range p.ReadlengthDefinitions {
		p.rdMap[v.ReadlengthModeKey] = v
		for _, platformKey := range v.SupportPlatformKey {
			for _, throughputModeKey := range v.SupportThroughputKey {
				key := platformKey*10 + throughputModeKey
				p.ptrdMap[key] = append(p.ptrdMap[key], v)
			}
		}
	}
	p.rfMap = map[int32]*referenceMap{}
	p.prfMap = map[int32][]*referenceMap{}
	for _, v := range p.ReferenceMaps {
		p.rfMap[v.ReferenceKey] = v
		p.prfMap[v.PlatformKey] = append(p.prfMap[v.PlatformKey], v)
	}
	p.rmMap = map[int32]*referenceDetail{}
	for _, v := range p.ReferenceDetailModels {
		p.rmMap[v.ReferenceId] = v
	}
}

func (p *BKRunPlan) GetThroughputDefinitionsByPlatformKey(platformKey int32) []*throughputDefinition {
	if definitions, ok := p.ptdMap[platformKey]; ok {
		return definitions
	}
	return []*throughputDefinition{}
}

func (p *BKRunPlan) GetThroughputDefinition(throughputModeKey int32) (*throughputDefinition, bool) {
	if definition, ok := p.tdMap[throughputModeKey]; ok {
		return definition, true
	}
	return nil, false
}

func (p *BKRunPlan) GetReadlengthModeDefinitionsByPlatformKey(platformKey, throughputModeKey int32) ([]*readlengthDefinition, bool) {
	key := platformKey*10 + throughputModeKey
	if definitions, ok := p.ptrdMap[key]; ok {
		return definitions, true
	}
	return []*readlengthDefinition{}, false
}

func (p *BKRunPlan) GetReadlengthModeDefinition(readlengthModeKey int32) (*readlengthDefinition, bool) {
	if definition, ok := p.rdMap[readlengthModeKey]; ok {
		return definition, true
	}
	return nil, false
}

func (p *BKRunPlan) GetReferences(platformKey int32) []*referenceMap {
	if references, ok := p.prfMap[platformKey]; ok {
		return references
	}
	return []*referenceMap{}
}

func (p *BKRunPlan) GetReference(referenceKey int32) (*referenceMap, bool) {
	if reference, ok := p.rfMap[referenceKey]; ok {
		return reference, true
	}
	return nil, false
}

func (p *BKRunPlan) GetReferenceDisplayName(referenceKey int32) string {
	if reference, ok := p.rfMap[referenceKey]; ok {
		return reference.ReferenceDisplayName
	}
	return ""
}

func (p *BKRunPlan) GetReferenceDetail(referenceKey int32) (*referenceDetail, bool) {
	if detail, ok := p.rmMap[referenceKey]; ok {
		return detail, true
	}
	return nil, false
}

func bkRunPlanSetup() {
	jsonUnmarshal(bkRunPlanNamespace, &BKRunPlanConfig, throughputModeDefinitionKey, readlengthModeDefinitionsKey, referenceMapKey, referenceDetailKey)

	BKRunPlanConfig.setup()
	common.Log.Debugf("BKRunPlanConfig: %v", common.J(BKRunPlanConfig))
}
