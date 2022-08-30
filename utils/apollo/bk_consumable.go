package apollo

import (
	"github.com/emirpasic/gods/maps/linkedhashmap"
	"rm/common"
)

//耗材配置
type consumableSet struct {
	PlatformKey int32  `json:"platformKey"`
	Type        string `json:"type"`
	Variant     string `json:"variant"`
	Version     string `json:"version"`
}

type BKConsumable struct {
	ConsumableSets []*consumableSet

	csMap map[int32]*linkedhashmap.Map //map[platformKey]map[type]map[variant][]version
}

func (p *BKConsumable) setup() {
	p.csMap = map[int32]*linkedhashmap.Map{}
	for _, consumable := range p.ConsumableSets {
		if tv, ok := p.csMap[consumable.PlatformKey]; ok {
			if variant, ok := tv.Get(consumable.Type); ok {
				if versions, ok := variant.(*linkedhashmap.Map).Get(consumable.Variant); ok {
					variant.(*linkedhashmap.Map).Put(consumable.Variant, append(versions.([]string), consumable.Version))
				} else {
					variant.(*linkedhashmap.Map).Put(consumable.Variant, []string{consumable.Version})
				}
			} else {
				vv := linkedhashmap.New()
				vv.Put(consumable.Variant, []string{consumable.Version})
				tv.Put(consumable.Type, vv)
			}
		} else {
			tv, vv := linkedhashmap.New(), linkedhashmap.New()
			vv.Put(consumable.Variant, []string{consumable.Version})
			tv.Put(consumable.Type, vv)
			p.csMap[consumable.PlatformKey] = tv
		}
	}
}

func (p *BKConsumable) GetTypes(platformKey int32) []interface{} {
	if v, ok := p.csMap[platformKey]; ok {
		return v.Keys()
	}
	return []interface{}{}
}

func (p *BKConsumable) GetVariants(platformKey int32, bkType string) (interface{}, bool) {
	if v, ok := p.csMap[platformKey]; ok {
		if vv, ok := v.Get(bkType); ok {
			return vv.(*linkedhashmap.Map).Keys(), ok
		}
	}
	return []interface{}{}, false
}

func (p *BKConsumable) GetVersions(platformKey int32, bkType, bkVariant string) (interface{}, bool) {
	if v, ok := p.csMap[platformKey]; ok {
		if vv, ok := v.Get(bkType); ok {
			return vv.(*linkedhashmap.Map).Get(bkVariant)
		}
	}
	return []interface{}{}, false
}

func (p *BKConsumable) ContainsVersion(platformKey int32, bkType, bkVariant, bkVersion string) bool {
	if t, ok := p.csMap[platformKey]; ok {
		if v, ok1 := t.Get(bkType); ok1 {
			if versions, ok2 := v.(*linkedhashmap.Map).Get(bkVariant); ok2 {
				for _, version := range versions.([]string) {
					if version == bkVersion {
						return true
					}
				}
			}
		}
	}
	return false
}

func bkConsumableSetup() {
	jsonUnmarshal(bkConsumableNamespace, &BKConsumableConfig, consumableSetDefinitionKey)

	BKConsumableConfig.setup()
	common.Log.Debugf("BKConsumableConfig: %v", common.J(BKConsumableConfig))
}
