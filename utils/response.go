package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
	"rm/common"
	"rm/utils/validation"
	"runtime/debug"
)

func HttpSuccess(c *gin.Context, data ...interface{}) {
	c.GetHeader("Accept-Language")
	t, _ := GetUTrans().GetTranslator(c.GetHeader("Accept-Language"))

	if len(data) > 0 {
		c.JSON(http.StatusOK, common.OK.Translator(t).H(data[0]))
	} else {
		c.JSON(http.StatusOK, common.OK.Translator(t).H())
	}
}

func HttpErr(c *gin.Context, err interface{}) {
	c.GetHeader("Accept-Language")
	t, _ := GetUTrans().GetTranslator(c.GetHeader("Accept-Language"))

	switch err.(type) {
	case *common.ResStatus:
		c.JSON(http.StatusOK, err.(*common.ResStatus).Translator(t).H())
	case validator.ValidationErrors:
		for _, fieldError := range err.(validator.ValidationErrors) {
			if _, ok := validation.Validators[fieldError.Tag()]; ok {
				// 自定义验证器
				c.JSON(http.StatusOK, common.InvalidParam.Desc(common.ResDesc(fieldError.Tag())).Translator(t).H())
			} else {
				c.JSON(http.StatusOK, common.InvalidParam.Desc(common.ResDesc(fieldError.Translate(t))).H())
			}
			return
		}
	case *json.UnmarshalTypeError:
		utErr := err.(*json.UnmarshalTypeError)
		c.JSON(http.StatusOK, common.InvalidParam.Desc(common.FieldTypeError, utErr.Field, utErr.Type.String()).Translator(t).H())
	default:
		common.Log.Errorf("发生未捕获错误: type: %v err: %v \n stack: %s", reflect.TypeOf(err), err, string(debug.Stack()))
		c.JSON(http.StatusOK, common.Unknown.Translator(t).H())
	}
}
