package utils

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"rm/common"
)

var uTrans *ut.UniversalTranslator

func SetupTrans(configFolder string) {
	zhT := zh.New()
	enT := en.New()
	uTrans = ut.New(zhT, enT, zhT)

	err := uTrans.Import(ut.FormatJSON, configFolder+"/translations")
	if err != nil {
		common.Log.Panicf("[SetupTrans panic] %v", err)
	}

	err = uTrans.VerifyTranslations()
	if err != nil {
		common.Log.Panicf("[SetupTrans panic] %v", err)
	}
}

func GetUTrans() *ut.UniversalTranslator {
	return uTrans
}
