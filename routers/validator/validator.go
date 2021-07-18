package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator"
)

type Validator struct {
	UserValidator        UserValidator
	AccountValidator     AccountValidator
	TransactionValidator TransactionValidator
}

// InitServiceFactory initialize services factory
func InitServiceFactory() *Validator {
	return &Validator{
		UserValidator:        newUserValidator(),
		AccountValidator:     newAccountValidator(),
		TransactionValidator: newTransactionValidator(),
	}
}

func ValidatorStruct(fn func(sl validator.StructLevel), dest, typeI interface{}) error {
	validate := validator.New()
	// register function to get tag name from json tags.
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	validate.RegisterStructValidation(fn, typeI)
	if err := validate.Struct(dest); err != nil {
		return err
	}
	return nil
}
