package config

import (
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

func NewValidator(viper *viper.Viper) *validator.Validate {
	v := validator.New()
	_ = v.RegisterValidation("email_or_phone", ValidateCredentialValue)
	return v
}
func ValidateCredentialValue(fl validator.FieldLevel) bool {
	credentialType := fl.Parent().FieldByName("CredentialType").String()
	credentialValue := fl.Field().String()

	if credentialType == "email" {
		// validate email with regex
		emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		return regexp.MustCompile(emailRegex).MatchString(credentialValue)
	} else if credentialType == "phone" {
		// validate phone with regex
		phoneRegex := `^\+\d{7,13}$`
		return regexp.MustCompile(phoneRegex).MatchString(credentialValue)
	}
	return false
}
