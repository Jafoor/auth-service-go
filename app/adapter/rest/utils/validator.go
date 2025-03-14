package utils

import (
	"log/slog"

	"auth-service/logger"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type Validator struct {
	en       locales.Translator
	uniTrans *ut.UniversalTranslator
	trans    ut.Translator
	validate *validator.Validate
}

var val *Validator

func InitValidator() {
	val = &Validator{}
	val.en = en.New()
	val.uniTrans = ut.New(val.en, val.en)
	val.trans, _ = val.uniTrans.GetTranslator("en")
	val.validate = validator.New()
	en_translations.RegisterDefaultTranslations(val.validate, val.trans)
}

func Validate(v interface{}) error {
	if err := val.validate.Struct(v); err != nil {
		slog.Error("Failed to validate", logger.Extra(map[string]any{
			"error": err.Error(),
		}))
		return err
	}
	return nil
}

func TranslateError(e error) validator.ValidationErrorsTranslations {
	valErr, ok := e.(validator.ValidationErrors)
	if !ok {
		return validator.ValidationErrorsTranslations{
			"error": e.Error(),
		}
	}
	return valErr.Translate(val.trans)
}
