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
		// skip if tag key says it should be ignored
		if name == "-" || name == "" {
			return ""
		}
		return name
	})

	validate.RegisterValidation("isUsername", func(fl validator.FieldLevel) bool {
		re := regexp.MustCompile(`^[a-zA-Z0-9_\-]+$`)
		return re.MatchString(fl.Field().String())
	})
	validate.RegisterTranslation("isUsername", trans, func(trans ut.Translator) error {
		return trans.Add("isUsername", "The {0} must be composed of alphanumeric or '-' and '_' character.", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("isUsername", fe.Field())
		return t
	})

	return validate, trans
}
