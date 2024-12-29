package utils

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	utTrans "github.com/go-playground/validator/v10/translations/en"
)

func NewValidate() (*validator.Validate, ut.Translator) {
	validate := validator.New()
	english := en.New()
	trans, _ := ut.New(english, english).GetTranslator("en")
	utTrans.RegisterDefaultTranslations(validate, trans)

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" || name == "" {
			return ""
		}
		return name
	})

	registerCustomValidation(validate, trans, "isUsername",
		"The {0} must be composed of only alphanumeric, '-', or '_' character.",
		func(fl validator.FieldLevel) bool {
			re := regexp.MustCompile(`^[a-zA-Z0-9_\-]+$`)
			return re.MatchString(fl.Field().String())
		},
	)
	registerCustomValidation(validate, trans, "isLBName",
		"The {0} must be composed of only alphanumeric, '-', '_', or whitespace character.",
		func(fl validator.FieldLevel) bool {
			re := regexp.MustCompile(`^[a-zA-Z0-9_\- ]+$`)
			return re.MatchString(fl.Field().String())
		},
	)
	registerCustomValidation(validate, trans, "isSafeName",
		"The {0} must not contain `'` or `\"`.",
		func(fl validator.FieldLevel) bool {
			re := regexp.MustCompile(`['"]`)
			return !re.MatchString(fl.Field().String())
		},
	)

	return validate, trans
}

func registerCustomValidation(validate *validator.Validate, trans ut.Translator, name, translation string, validationFunc validator.Func) {
	validate.RegisterValidation(name, validationFunc)
	validate.RegisterTranslation(name, trans, func(trans ut.Translator) error {
		return trans.Add(name, translation, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(name, fe.Field())
		return t
	})
}
