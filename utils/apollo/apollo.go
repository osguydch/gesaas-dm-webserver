package apollo

import (
	"encoding/csv"
	"encoding/json"
	"github.com/philchia/agollo/v4"
	"io"
	"rm/common"
	"strings"
)

const (
	platformDefinitionKey = "PlatformDefinition"
	index1DefinitionKey   = "Index1Definition"
	index2DefinitionKey   = "Index2Definition"

	geneGroupDefinitionKey    = "GeneGroupDefinition"
	index1DisplayMapKey       = "Index1DisplayMap"
	index1SelectionFileMapKey = "Index1SelectionFileMap"
	index2DisplayMapKey       = "Index2DisplayMap"
	index2SelectionFileMapKey = "Index2SelectionFileMap"
	read1IndexKey             = "Read1Index"

	throughputModeDefinitionKey  = "ThroughputModeDefinition"
	readlengthModeDefinitionsKey = "ReadlengthModeDefinitions"
	referenceMapKey              = "ReferenceMap"
	referenceDetailKey           = "ReferenceDetail"

	consumableSetDefinitionKey = "ConsumableSetDefinition"
)

var (
	bkCommonNamespace     agollo.OpOption
	bkSampleNamespace     agollo.OpOption
	bkRunPlanNamespace    agollo.OpOption
	bkConsumableNamespace agollo.OpOption

	BKCommonConfig     BKCommon
	BKSampleConfig     BKSample
	BKRunPlanConfig    BKRunPlan
	BKConsumableConfig BKConsumable
)

func SetupApollo() {
	common.Log.Info("setup apollo starting")
	bkCommonNamespace = agollo.WithNamespace(common.ApolloConfig.BKCommon)
	bkSampleNamespace = agollo.WithNamespace(common.ApolloConfig.BKSample)
	bkRunPlanNamespace = agollo.WithNamespace(common.ApolloConfig.BKRunPlan)
	bkConsumableNamespace = agollo.WithNamespace(common.ApolloConfig.BKConsumable)
	err := agollo.Start(&agollo.Conf{
		AppID:          common.ApolloConfig.AppID,
		Cluster:        common.ApolloConfig.Cluster,
		CacheDir:       ".apollo",
		NameSpaceNames: []string{common.ApolloConfig.BKCommon, common.ApolloConfig.BKSample, common.ApolloConfig.BKRunPlan, common.ApolloConfig.BKConsumable},
		MetaAddr:       common.ApolloConfig.MetaAddr,
	})
	if err != nil {
		panic(err)
	}

	bkCommonSetup()
	bkSampleSetup()
	bkRunPlanSetup()
	bkConsumableSetup()

	common.Log.Info("setup apollo ok")
}

func jsonUnmarshal(namespace agollo.OpOption, v interface{}, keys ...string) {
	for _, key := range keys {
		if err := json.Unmarshal([]byte(agollo.GetString(key, namespace)), &v); err != nil {
			common.Log.Panicf("apollo config %s parse fail: %v", key, err)
		}
	}
}

func csvUnmarshal(content string) ([][]string, error) {
	reader := csv.NewReader(strings.NewReader(content))
	var res [][]string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		res = append(res, record)
	}
	return res, nil
}
