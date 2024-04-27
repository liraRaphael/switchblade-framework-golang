package validator

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func New() *Validator {
	validate := validator.New()
	return &Validator{
		validate: validate,
	}
}

func (v *Validator) Validate(s interface{}) error {
	return v.validate.Struct(s)
}

var (
	once     sync.Once
	instance *Validator
)

func Get() *Validator {
	once.Do(func() {
		instance = New()
	})
	return instance
}

func (v *Validator) AddCustomValidator(tag string, fn validator.Func) {
	v.validate.RegisterValidation(tag, fn)
}
