package apollo

import (
	"rm/common"
)

type platformDefinition struct {
	PlatformKey         int32  `json:"platformKey"`
	PlatformDisplayName string `json:"platformDisplayName"`
	PlatformIndexMode   int32  `json:"platformIndexMode"`
}

type index1Mode struct {
	Index1ModeKey     int32  `json:"index1ModeKey"`
	Index1DisplayName string `json:"index1DisplayName"`
	Index1Length      int32  `json:"index1Length"`
}

type index2Mode struct {
	Index2ModeKey     int32  `json:"index2ModeKey"`
	Index2DisplayName string `json:"index2DisplayName"`
	Index2Length      int32  `json:"index2Length"`
}

type BKCommon struct {
	PlatformDefinitions []*platformDefinition `json:"platformDefinitions"`
	Index1Modes         []*index1Mode         `json:"index1Modes"`
	Index2Modes         []*index2Mode         `json:"index2Modes"`

	pdMap map[int32]*platformDefinition
	i1Map map[int32]*index1Mode
	i2Map map[int32]*index2Mode
}

func (p *BKCommon) setup() {
	p.pdMap = map[int32]*platformDefinition{}
	for _, v := range p.PlatformDefinitions {
		p.pdMap[v.PlatformKey] = v
	}
	p.i1Map = map[int32]*index1Mode{}
	for _, v := range p.Index1Modes {
		p.i1Map[v.Index1ModeKey] = v
	}
	p.i2Map = map[int32]*index2Mode{}
	for _, v := range p.Index2Modes {
		p.i2Map[v.Index2ModeKey] = v
	}
}

func (p *BKCommon) GetPlatform(platformKey int32) (*platformDefinition, bool) {
	if definition, ok := p.pdMap[platformKey]; ok {
		return definition, ok
	}
	return nil, false
}

func (p *BKCommon) GetIndex1(index1ModeKey int32) (*index1Mode, bool) {
	if mode, ok := p.i1Map[index1ModeKey]; ok {
		return mode, ok
	}
	return nil, false
}

func (p *BKCommon) GetIndex2(index2ModeKey int32) (*index2Mode, bool) {
	if mode, ok := p.i2Map[index2ModeKey]; ok {
		return mode, ok
	}
	return nil, false
}

func bkCommonSetup() {
	jsonUnmarshal(bkCommonNamespace, &BKCommonConfig, platformDefinitionKey, index1DefinitionKey, index2DefinitionKey)

	BKCommonConfig.setup()
	common.Log.Debugf("BKCommonConfig: %v", common.J(BKCommonConfig))
}
