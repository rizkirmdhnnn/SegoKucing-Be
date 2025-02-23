package model

type RegisterUserRequest struct {
	CredentialType  string `json:"credentialType" validate:"required,oneof=phone email"`
	CredentialValue string `json:"credentialValue" validate:"required,min=7,max=50,email_or_phone"`
	Name            string `json:"name" validate:"required,min=5,max=50"`
	Password        string `json:"password" validate:"required,min=5,max=15"`
}

type DataUserWithToken struct {
	Phone       string `json:"phone,omitempty"`
	Email       string `json:"email,omitempty"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}

type LoginUserRequest struct {
	CredentialType  string `json:"credentialType" validate:"required,oneof=phone email"`
	CredentialValue string `json:"credentialValue" validate:"required,min=7,max=50,email_or_phone"`
	Password        string `json:"password" validate:"required,min=5,max=15"`
}

type LinkEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type LinkPhoneRequest struct {
	Phone string `json:"phone" validate:"required,min=7,max=15"`
}
