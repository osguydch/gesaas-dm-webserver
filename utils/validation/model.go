package validation

import (
	"github.com/go-playground/validator/v10"
)

type ICustomValidator interface {
	GetTag() string
	GetFunc() validator.Func
}

type CustomValidator struct {
	tag string
	validator.Func
}

func (p *CustomValidator) GetTag() string {
	return p.tag
}

func (p *CustomValidator) GetFunc() validator.Func {
	return p.Func
}

func newCustomValidator(tag string, f validator.Func) *CustomValidator {
	return &CustomValidator{tag, f}
}
