package api

type signUpInput struct {
	Mail            string `json:"mail" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required,min=8"`
}

type signInInput struct {
	Mail     string `json:"mail"  binding:"required"`
	Password string `json:"password"  binding:"required,min=8"`
	OtpPin   string `json:"otp_pin" binding:"required"`
}
