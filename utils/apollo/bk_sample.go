package apollo

import (
	"github.com/emirpasic/gods/maps/linkedhashmap"
	"github.com/philchia/agollo/v4"
	"rm/common"
)

// 样本信息基因组的配置
type geneGroup struct {
	GeneGroupKey         int32  `json:"geneGroupKey"`
	GeneGroupDisplayName string `json:"geneGroupDisplayName"`
}

// 建库方法与Index1序列展示映射关系以及Index1序列展示的大类名称。
type index1DisplayMap struct {
	Index1NameSelectionKey         int32              `json:"index1NameSelectionKey"`
	Index1NameSelectionDisplayName string             `json:"index1NameSelectionDisplayName"`
	PlatformKey                    int32              `json:"platformKey"`
	IndexMap                       *linkedhashmap.Map `json:"-"`
}

// 定义Index1序列展示定义和具体的Index1名称和序列文件之间的映射关系。
type index1SelectionFileMap struct {
	Index1NameSelectionKey int32
	FileName               string
}

// 定义建库方法与Index2序列展示映射关系以及Index2序列展示的大类名称。
type index2DisplayMap struct {
	Index2NameSelectionKey         int32              `json:"index2NameSelectionKey"`
	Index2NameSelectionDisplayName string             `json:"index2NameSelectionDisplayName"`
	PlatformKey                    int32              `json:"platformKey"`
	IndexMap                       *linkedhashmap.Map `json:"-"`
}

// 定义Index2序列展示定义和具体的Index2名称和序列文件之间的映射关系。
type index2SelectionFileMap struct {
	Index2NameSelectionKey int32
	FileName               string
}

type indexMap struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type BKSample struct {
	GeneGroups              []*geneGroup
	Index1DisplayMaps       []*index1DisplayMap
	Index1SelectionFileMaps []*index1SelectionFileMap
	Index2DisplayMaps       []*index2DisplayMap
	Index2SelectionFileMaps []*index2SelectionFileMap
	Read1IndexMaps          []*indexMap

	i1Map map[int32]*index1DisplayMap
	i2Map map[int32]*index2DisplayMap

	platformI1 map[int32][]*index1DisplayMap //根据平台id区分index1Desplay
	platformI2 map[int32][]*index2DisplayMap //根据平台id区分index2Desplay
}

func (p *BKSample) setup() {
	p.i1Map = map[int32]*index1DisplayMap{}
	p.platformI1 = map[int32][]*index1DisplayMap{}
	for _, v := range p.Index1DisplayMaps {
		p.i1Map[v.Index1NameSelectionKey] = v
		p.platformI1[v.PlatformKey] = append(p.platformI1[v.PlatformKey], v)
	}

	for _, v := range p.Index1SelectionFileMaps {
		records, err := csvUnmarshal(agollo.GetString(v.FileName, bkSampleNamespace))
		if err != nil {
			common.Log.Panicf("apollo config %s parse fail: %v", v.FileName, err)
		}
		index1Display := p.i1Map[v.Index1NameSelectionKey]
		index1Display.IndexMap = linkedhashmap.New()
		for _, record := range records {
			index1Display.IndexMap.Put(record[1], &indexMap{record[0], record[1]})
		}
		p.i1Map[v.Index1NameSelectionKey] = index1Display
	}

	p.i2Map = map[int32]*index2DisplayMap{}
	p.platformI2 = map[int32][]*index2DisplayMap{}
	for _, v := range p.Index2DisplayMaps {
		p.i2Map[v.Index2NameSelectionKey] = v
		p.platformI2[v.PlatformKey] = append(p.platformI2[v.PlatformKey], v)
	}

	for _, v := range BKSampleConfig.Index2SelectionFileMaps {
		records, err := csvUnmarshal(agollo.GetString(v.FileName, bkSampleNamespace))
		if err != nil {
			common.Log.Panicf("apollo config %s parse fail: %v", v.FileName, err)
		}
		index2Display := p.i2Map[v.Index2NameSelectionKey]
		index2Display.IndexMap = linkedhashmap.New()
		for _, record := range records {
			index2Display.IndexMap.Put(record[1], &indexMap{record[0], record[1]})
		}
		p.i2Map[v.Index2NameSelectionKey] = index2Display
	}

	records, err := csvUnmarshal(agollo.GetString(read1IndexKey, bkSampleNamespace))
	if err != nil {
		common.Log.Panicf("apollo config %s parse fail: %v", read1IndexKey, err)
	}
	for _, record := range records {
		BKSampleConfig.Read1IndexMaps = append(BKSampleConfig.Read1IndexMaps, &indexMap{record[0], record[1]})
	}
}

func (p *BKSample) GetIndex1DisplayMapByPlatformKey(platformKey int32) []*index1DisplayMap {
	if v, ok := p.platformI1[platformKey]; ok {
		return v
	}
	return []*index1DisplayMap{}
}

// 判断index1序列是否已预定义
func (p *BKSample) I1IndexMapContains(platformKey int32, index1 string) (string, bool) {
	if v, ok := p.platformI1[platformKey]; ok {
		for _, index1DisplayMap := range v {
			if index1, ok := index1DisplayMap.IndexMap.Get(index1); ok {
				return index1.(*indexMap).Key, true
			}
		}
	}
	return "", false
}

func (p *BKSample) GetIndex2DisplayMapByPlatformKey(platformKey int32) []*index2DisplayMap {
	if v, ok := p.platformI2[platformKey]; ok {
		return v
	}
	return []*index2DisplayMap{}
}

// 判断index2序列是否已预定义
func (p *BKSample) I2IndexMapContains(platformKey int32, index2 string) (string, bool) {
	if v, ok := p.platformI2[platformKey]; ok {
		for _, index2DisplayMap := range v {
			if index2, ok := index2DisplayMap.IndexMap.Get(index2); ok {
				return index2.(*indexMap).Key, true
			}
		}
	}
	return "", false
}

func bkSampleSetup() {
	jsonUnmarshal(bkSampleNamespace, &BKSampleConfig, geneGroupDefinitionKey, index1DisplayMapKey, index1SelectionFileMapKey, index2DisplayMapKey, index2SelectionFileMapKey)

	BKSampleConfig.setup()
	common.Log.Debugf("BKSampleConfig: %v", common.J(BKSampleConfig))
}
