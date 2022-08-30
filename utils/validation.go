package utils

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"rm/common"
	"rm/utils/validation"
)

func registerZh(v *validator.Validate) {
	trans, ok := GetUTrans().GetTranslator("zh")
	if !ok {
		common.Log.Panicf("uni.GetTranslator zh failed")
	}
	// 注册翻译器
	err := zhTranslations.RegisterDefaultTranslations(v, trans)
	if err != nil {
		common.Log.Panicf("SetupTrans err: %v", err)
	}
}

func registerEn(v *validator.Validate) {
	trans, ok := GetUTrans().GetTranslator("en")
	if !ok {
		common.Log.Panic("uni.GetTranslator en failed")
	}
	// 注册翻译器
	err := enTranslations.RegisterDefaultTranslations(v, trans)
	if err != nil {
		common.Log.Panicf("SetupTrans err: %v", err)
	}
}

// SetupValidator 初始化validator
func SetupValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		registerZh(v)
		registerEn(v)
		// 注册自定义验证器
		for _, validate := range validation.Validators {
			if err := v.RegisterValidation(validate.GetTag(), validate.GetFunc()); err != nil {
				common.Log.Panicf("SetupValidator err: %v", err)
			}
		}
		common.Log.Info("setup validator ok")
	}
}
