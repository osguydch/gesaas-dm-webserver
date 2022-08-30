package validation

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"rm/common"
	"rm/dto"
	"rm/utils/apollo"
	"testing"
)

func init() {
	fmt.Println("-------------test init---------------")
	common.SetUpConfig("../../config/config.toml")
	common.SetupLogger()
	apollo.SetupApollo()
}

func TestValidateSampleInfoRequest(t *testing.T) {
	requests := []struct {
		content string
		err     string
	}{
		{
			content: `{
				"name": "tes5t11",
				"creatorId": 1,
				"libPrepMethodKey": 1,
				"index1LengthKey": 0,
				"index1": false,
				"index2LengthKey": 1,
				"index2": true,
				"samples": [
					{"sampleID": "test1","description": "","index1ID": "","index1": "ATCGT","index2ID": "",
					"index2": "TGACT","read1IndexID":"","read1Index":"","refGenome": ""}
				]
			}`,
			err: "无效参数:index1",
		},
		{
			content: `{
				"name": "tes5t11",
				"creatorId": 1,
				"libPrepMethodKey": 1,
				"index1LengthKey": 0,
				"index1": true,
				"index2LengthKey": 1,
				"index2": true,
				"samples": [
					{"sampleID": "test1","description": "","index1ID": "","index1": "ATCGT","index2ID": "",
					"index2": "TGACT","read1IndexID":"","read1Index":"","refGenome": ""}
				]
			}`,
			err: "无效参数:index1LengthKey",
		},
		{
			content: `{
				"name": "tes5t11",
				"creatorId": 1,
				"libPrepMethodKey": 1,
				"index1LengthKey": 1,
				"index1": true,
				"index2LengthKey": 0,
				"index2": true,
				"samples": [
					{"sampleID": "test1","description": "","index1ID": "","index1": "ATCGT","index2ID": "",
					"index2": "TGACT","read1IndexID":"","read1Index":"","refGenome": ""}
				]
			}`,
			err: "无效参数:index2LengthKey",
		},
		{
			content: `{
				"name": "tes5t11",
				"creatorId": 1,
				"libPrepMethodKey": 1,
				"index1LengthKey": 1,
				"index1": true,
				"index2LengthKey": 1,
				"index2": false,
				"samples": [
					{"sampleID": "test1","description": "","index1ID": "","index1": "ATCGT","index2ID": "",
					"index2": "TGACT","read1IndexID":"","read1Index":"","refGenome": ""}
				]
			}`,
			err: "无效参数:index2LengthKey",
		},
		{
			content: `{
				"name": "tes5t11",
				"creatorId": 1,
				"libPrepMethodKey": 3,
				"index1LengthKey": 1,
				"index1": true,
				"index2LengthKey": 1,
				"index2": true,
				"samples": [
					{"sampleID": "test1","description": "","index1ID": "","index1": "ATCGT","index2ID": "",
					"index2": "TGACT","read1IndexID":"","read1Index":"","refGenome": ""}
				]
			}`,
			err: "无效参数:index1",
		},
		{
			content: `{
				"name": "tes5t11",
				"creatorId": 1,
				"libPrepMethodKey": 3,
				"index1LengthKey": 1,
				"index1": false,
				"index2LengthKey": 1,
				"index2": true,
				"samples": [
					{"sampleID": "test1","description": "","index1ID": "","index1": "ATCGT","index2ID": "",
					"index2": "TGACT","read1IndexID":"","read1Index":"","refGenome": ""}
				]
			}`,
			err: "无效参数:index2",
		},
		{
			content: `{
				"name": "tes5t11",
				"creatorId": 1,
				"libPrepMethodKey": 3,
				"index1LengthKey": 1,
				"index1": false,
				"index2LengthKey": 1,
				"index2": false,
				"samples": [
					{"sampleID": "test1","description": "","index1ID": "","index1": "ATCGT","index2ID": "",
					"index2": "TGACT","read1IndexID":"","read1Index":"","refGenome": ""}
				]
			}`,
			err: "无效参数:index1LengthKey",
		},
		{
			content: `{
				"name": "tes5t11",
				"creatorId": 1,
				"libPrepMethodKey": 3,
				"index1LengthKey": 0,
				"index1": false,
				"index2LengthKey": 1,
				"index2": false,
				"samples": [
					{"sampleID": "test1","description": "","index1ID": "","index1": "ATCGT","index2ID": "",
					"index2": "TGACT","read1IndexID":"","read1Index":"","refGenome": ""}
				]
			}`,
			err: "无效参数:index2LengthKey",
		},
		{
			content: `{
				"name": "tes5t11",
				"creatorId": 1,
				"libPrepMethodKey": 3,
				"index1LengthKey": 0,
				"index1": false,
				"index2LengthKey": 0,
				"index2": false,
				"samples": []
			}`,
			err: "无效参数:sample,至少有一个样本",
		},
		{
			content: `{
				"name": "tes5t11",
				"creatorId": 1,
				"libPrepMethodKey": 1,
				"index1LengthKey": 1,
				"index1": true,
				"index2LengthKey": 0,
				"index2": false,
				"samples": [
					{"sampleID": "@A","description": "","index1ID": "","index1": "ATCGT","index2ID": "",
					"index2": "","read1IndexID":"","read1Index":"","refGenome": ""}
				]
			}`,
			err: "无效参数:sampleID,样本第1行:'@A'不符合规则",
		},
		{
			content: `{
				"name": "tes5t11",
				"creatorId": 1,
				"libPrepMethodKey": 1,
				"index1LengthKey": 1,
				"index1": true,
				"index2LengthKey": 0,
				"index2": false,
				"samples": [
					{"sampleID": "ShortRead","description": "","index1ID": "","index1": "ATCGT","index2ID": "",
					"index2": "","read1IndexID":"","read1Index":"","refGenome": ""}
				]
			}`,
			err: "无效参数:sampleID,样本第1行:'ShortRead'为保留字",
		},
		{
			content: `{
				"name": "tes5t11",
				"creatorId": 1,
				"libPrepMethodKey": 1,
				"index1LengthKey": 1,
				"index1": true,
				"index2LengthKey": 0,
				"index2": false,
				"samples": [
					{"sampleID": "Sample_01","description": "","index1ID": "","index1": "ATCGT","index2ID": "",
					"index2": "","read1IndexID":"","read1Index":"","refGenome": ""},
					{"sampleID": "Sample_01","description": "","index1ID": "","index1": "ATCGT","index2ID": "",
					"index2": "","read1IndexID":"","read1Index":"","refGenome": ""}
				]
			}`,
			err: "无效参数:sampleID,样本第2行:'Sample_01'重复",
		},
		{
			content: `{
				"name": "tes5t11",
				"creatorId": 1,
				"libPrepMethodKey": 1,
				"index1LengthKey": 1,
				"index1": true,
				"index2LengthKey": 0,
				"index2": false,
				"samples": [
					{"sampleID": "Sample_01","description": "","index1ID": "","index1": "ATCGT","index2ID": "",
					"index2": "","read1IndexID":"","read1Index":"A","refGenome": ""}
				]
			}`,
			err: "无效参数:read1IndexID、read1Index,样本第1行",
		},
		{
			content: `{
				"name": "tes5t11",
				"creatorId": 1,
				"libPrepMethodKey": 1,
				"index1LengthKey": 1,
				"index1": true,
				"index2LengthKey": 0,
				"index2": false,
				"samples": [
					{"sampleID": "Sample_01","description": "","index1ID": "A1","index1": "ATCGT","index2ID": "",
					"index2": "","read1IndexID":"","read1Index":"","refGenome": ""},
					{"sampleID": "Sample_02","description": "","index1ID": "A1","index1": "ATCGT","index2ID": "",
					"index2": "","read1IndexID":"","read1Index":"","refGenome": ""}
				]
			}`,
			err: "无效参数:index1ID,样本第2行:'A1'重复",
		},
		{
			content: `{
				"name": "tes5t11",
				"creatorId": 1,
				"libPrepMethodKey": 1,
				"index1LengthKey": 1,
				"index1": true,
				"index2LengthKey": 0,
				"index2": false,
				"samples": [
					{"sampleID": "Sample_01","description": "","index1ID": "A1","index1": "ATCGAA","index2ID": "",
					"index2": "","read1IndexID":"","read1Index":"","refGenome": ""}
				]
			}`,
			err: "无效参数:index1,样本第1行:'ATCGAA'长度大于5",
		},
		{
			content: `{
				"name": "tes5t11",
				"creatorId": 1,
				"libPrepMethodKey": 1,
				"index1LengthKey": 1,
				"index1": true,
				"index2LengthKey": 0,
				"index2": false,
				"samples": [
					{"sampleID": "Sample_01","description": "","index1ID": "A1","index1": "ATCGQ","index2ID": "",
					"index2": "","read1IndexID":"","read1Index":"","refGenome": ""}
				]
			}`,
			err: "无效参数:index1,样本第1行:'ATCGQ'格式不正确",
		},
		{
			content: `{
				"name": "tes5t11",
				"creatorId": 1,
				"libPrepMethodKey": 1,
				"index1LengthKey": 1,
				"index1": true,
				"index2LengthKey": 0,
				"index2": false,
				"samples": [
					{"sampleID": "Sample_01","description": "","index1ID": "A1","index1": "ATCGT","index2ID": "",
					"index2": "","read1IndexID":"","read1Index":"","refGenome": ""},
					{"sampleID": "Sample_02","description": "","index1ID": "","index1": "ATCGT","index2ID": "",
					"index2": "","read1IndexID":"","read1Index":"","refGenome": ""}
				]
			}`,
			err: "无效参数:index1,样本第2行:'ATCGT'重复",
		},
		{
			content: `{
				"name": "tes5t11",
				"creatorId": 1,
				"libPrepMethodKey": 1,
				"index1LengthKey": 1,
				"index1": true,
				"index2LengthKey": 1,
				"index2": true,
				"samples": [
					{"sampleID": "Sample_01","description": "","index1ID": "A1","index1": "ATCGT","index2ID": "A1",
					"index2": "CTAGC","read1IndexID":"","read1Index":"","refGenome": ""},
					{"sampleID": "Sample_02","description": "","index1ID": "A2","index1": "ATCGC","index2ID": "A1",
					"index2": "CTAGC","read1IndexID":"","read1Index":"","refGenome": ""}
				]
			}`,
			err: "无效参数:index2ID,样本第2行:'A1'重复",
		},
		{
			content: `{
				"name": "tes5t11",
				"creatorId": 1,
				"libPrepMethodKey": 1,
				"index1LengthKey": 1,
				"index1": true,
				"index2LengthKey": 1,
				"index2": true,
				"samples": [
					{"sampleID": "Sample_01","description": "","index1ID": "A1","index1": "ATCGT","index2ID": "",
					"index2": "ATCGTA","read1IndexID":"","read1Index":"","refGenome": ""}
				]
			}`,
			err: "无效参数:index2,样本第1行:'ATCGTA'长度大于5",
		},
		{
			content: `{
				"name": "tes5t11",
				"creatorId": 1,
				"libPrepMethodKey": 1,
				"index1LengthKey": 1,
				"index1": true,
				"index2LengthKey": 1,
				"index2": true,
				"samples": [
					{"sampleID": "Sample_01","description": "","index1ID": "A1","index1": "ATCGT","index2ID": "",
					"index2": "ATCGQ","read1IndexID":"","read1Index":"","refGenome": ""}
				]
			}`,
			err: "无效参数:index2,样本第1行:'ATCGQ'格式不正确",
		},
		{
			content: `{
				"name": "tes5t11",
				"creatorId": 1,
				"libPrepMethodKey": 1,
				"index1LengthKey": 1,
				"index1": true,
				"index2LengthKey": 1,
				"index2": true,
				"samples": [
					{"sampleID": "Sample_01","description": "","index1ID": "A1","index1": "ATCGT","index2ID": "",
					"index2": "ATCGC","read1IndexID":"","read1Index":"","refGenome": ""},
					{"sampleID": "Sample_02","description": "","index1ID": "","index1": "ATCGC","index2ID": "",
					"index2": "ATCGC","read1IndexID":"","read1Index":"","refGenome": ""}
				]
			}`,
			err: "无效参数:index2,样本第2行:'ATCGC'重复",
		},
		{
			content: `{
				"name": "tes5t11",
				"creatorId": 1,
				"libPrepMethodKey": 3,
				"index1LengthKey": 0,
				"index1": false,
				"index2LengthKey": 0,
				"index2": false,
				"samples": [
					{"sampleID": "Sample_01","description": "","index1ID": "A1","index1": "ATCGT","index2ID": "",
					"index2": "ATCGC","read1IndexID":"","read1Index":"","refGenome": ""}
				]
			}`,
			err: "无效参数:index1ID、index1,样本第1行",
		},
		{
			content: `{
				"name": "tes5t11",
				"creatorId": 1,
				"libPrepMethodKey": 3,
				"index1LengthKey": 0,
				"index1": false,
				"index2LengthKey": 0,
				"index2": false,
				"samples": [
					{"sampleID": "Sample_01","description": "","index1ID": "","index1": "","index2ID": "",
					"index2": "ATCGC","read1IndexID":"","read1Index":"","refGenome": ""}
				]
			}`,
			err: "无效参数:index2ID、index2,样本第1行",
		},
		{
			content: `{
				"name": "tes5t11",
				"creatorId": 1,
				"libPrepMethodKey": 3,
				"index1LengthKey": 0,
				"index1": false,
				"index2LengthKey": 0,
				"index2": false,
				"samples": [
					{"sampleID": "Sample_01","description": "","index1ID": "","index1": "","index2ID": "",
					"index2": "ATCGC","read1IndexID":"","read1Index":"","refGenome": ""}
				]
			}`,
			err: "无效参数:index2ID、index2,样本第1行",
		},
		{
			content: `{
				"name": "tes5t11",
				"creatorId": 1,
				"libPrepMethodKey": 3,
				"index1LengthKey": 0,
				"index1": false,
				"index2LengthKey": 0,
				"index2": false,
				"samples": [
					{"sampleID": "Sample_01","description": "","index1ID": "","index1": "","index2ID": "",
					"index2": "","read1IndexID":"R1","read1Index":"ATCG","refGenome": ""},
					{"sampleID": "Sample_02","description": "","index1ID": "","index1": "","index2ID": "",
					"index2": "","read1IndexID":"R1","read1Index":"ATCG","refGenome": ""}
				]
			}`,
			err: "无效参数:read1IndexID,样本第2行:'R1'重复",
		},
		{
			content: `{
				"name": "tes5t11",
				"creatorId": 1,
				"libPrepMethodKey": 3,
				"index1LengthKey": 0,
				"index1": false,
				"index2LengthKey": 0,
				"index2": false,
				"samples": [
					{"sampleID": "Sample_01","description": "","index1ID": "","index1": "","index2ID": "",
					"index2": "","read1IndexID":"R1","read1Index":"Q","refGenome": ""}
				]
			}`,
			err: "无效参数:read1Index,样本第1行:'Q'格式不正确",
		},
		{
			content: `{
				"name": "tes5t11",
				"creatorId": 1,
				"libPrepMethodKey": 3,
				"index1LengthKey": 0,
				"index1": false,
				"index2LengthKey": 0,
				"index2": false,
				"samples": [
					{"sampleID": "Sample_01","description": "","index1ID": "","index1": "","index2ID": "",
					"index2": "","read1IndexID":"","read1Index":"A","refGenome": ""},
					{"sampleID": "Sample_02","description": "","index1ID": "","index1": "","index2ID": "",
					"index2": "","read1IndexID":"","read1Index":"A","refGenome": ""}
				]
			}`,
			err: "无效参数:read1Index,样本第2行:'A'重复",
		},
	}
	for i, request := range requests {
		var a dto.SampleInfoRequest
		if err := json.Unmarshal([]byte(request.content), &a); err != nil {
			t.Errorf("json unmarshal err: %v", err)
		}
		err := ValidateSampleInfoRequest(&a)
		fmt.Printf("test:%d\n", i)
		if len(request.err) > 0 {
			assert.ErrorContains(t, err, request.err)
		} else {
			assert.NoError(t, err)
		}
	}
}
