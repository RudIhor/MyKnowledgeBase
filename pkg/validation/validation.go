package validation

import (
	"net/http"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator  *validator.Validate
	translator ut.Translator
}

func NewCustomValidator() *CustomValidator {
	en := en.New()
	uni := ut.New(en, en)

	trans, _ := uni.GetTranslator("en")
	validate := validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)

	return &CustomValidator{validator: validate, translator: trans}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, e.Translate(cv.translator))
		}
	}
	return nil
}
