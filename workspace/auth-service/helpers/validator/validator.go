package utils

import (			
	"fmt"	
	"strings"
	"database/sql"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/validator/v10"	
	ut "github.com/go-playground/universal-translator"
	"auth-service/domain"
	"auth-service/user/repository/pgsql"
	// en_translations "github.com/go-playground/validator/v10/translations/en"
)

// use a single instance , it caches struct info
var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)
func NewValidator(conn *sql.DB)*validatorController{
	return &validatorController{
		userRepo : pgsql.NewPgsqlUserRepository(conn),
	}
}
type validatorController struct{
	userRepo domain.UserRepository
}
// InitValidate use for inizialitation validation request
func (c *validatorController) IsValidRequest(data interface{}) map[string][]string {

	// NOTE: ommitting allot of error checking for brevity

	en := en.New()
	uni = ut.New(en, en)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator("en")

	validate = validator.New()
	validate.RegisterValidation("is_phone_number_exists", func(fl validator.FieldLevel)bool{
		_, err := c.userRepo.GetByPhoneNumber(fl.Field().String())	
		if err == domain.ErrNotFound{
			return true
		}
		return false
	})
	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0}:{1} is required", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field(), fe.Field())

		return strings.ToLower(t)
	})
	validate.RegisterTranslation("min", trans, func(ut ut.Translator) error {
		return ut.Add("min", "{0}:{1} minimum length is {2}", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("min", fe.Field(), fe.Field(), fe.Param())

		return strings.ToLower(t)
	})
	validate.RegisterTranslation("max", trans, func(ut ut.Translator) error {
		return ut.Add("max", "{0}:{1} maximum length is {2}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("max", fe.Field(), fe.Field(), fe.Param())

		return strings.ToLower(t)
	})
	validate.RegisterTranslation("lte", trans, func(ut ut.Translator) error {
		return ut.Add("lte", "{0}:{1} maximum is {2}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("lt", fe.Field(), fe.Field(), fe.Param())

		return strings.ToLower(t)
	})
	validate.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "{0}:the format of the {1} address is not correct", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field(), fe.Field())

		return strings.ToLower(t)
	})
	validate.RegisterTranslation("numeric", trans, func(ut ut.Translator) error {
		return ut.Add("numeric", "{0}:{1} should be a number", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("numeric", fe.Field(), fe.Field())

		return strings.ToLower(t)
	})
	validate.RegisterTranslation("alpha", trans, func(ut ut.Translator) error {
		return ut.Add("alpha", "{0}:{1} should be a character", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("alpha", fe.Field(), fe.Field())

		return strings.ToLower(t)
	})
	validate.RegisterTranslation("uuid4", trans, func(ut ut.Translator) error {
		return ut.Add("uuid4", "{0}:{1} should be a uuid", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("uuid4", fe.Field(), fe.Field())

		return strings.ToLower(t)
	})
	validate.RegisterTranslation("is_phone_number_exists", trans, func(ut ut.Translator) error {
		return ut.Add("is_phone_number_exists", "{0}:{1} is already used", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("is_phone_number_exists", fe.Field(), fe.Field())

		return strings.ToLower(t)
	})
	
	// request validation
	if err := validate.Struct(data); err != nil {
		errStrings := make(map[string][]string)
		// translate all error at once
		errs := err.(validator.ValidationErrors)

		// returns a map with key = namespace & value = translated error
		// NOTICE: 2 errors are returned and you'll see something surprising
		// translations are i18n aware!!!!
		// eg. '10 characters' vs '1 character'
		for _, e := range errs {
			fmt.Println(e.Translate(trans))
			arrErr := strings.Split(e.Translate(trans), ":")			
			errStrings[arrErr[0]] = append(errStrings[arrErr[0]], arrErr[1])					
		}
		return errStrings
	}
	return nil
}
