package utils

import (
	"fmt"
	"strings"

	"github.com/xlzd/gotp"
	"github.com/xor111xor/s3-jwt-totp-backend/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

func GetOtpUrl(user *domain.User) string {
	parseMail := strings.Split(user.Mail, "@")
	return gotp.NewDefaultTOTP(user.OtpCache).ProvisioningUri(parseMail[0], "store")
}

func VerifyPassword(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}

func CheckUserCredential(user *domain.User, mail, pasword, pin string) error {
	// mail
	if user.Mail != mail {
		return domain.ErrNoUser
	}
	// pasword
	if err := VerifyPassword(user.PasswordHash, pasword); err != nil {
		return domain.ErrUserPasswordMismatch
	}
	// otp
	totp := gotp.NewDefaultTOTP(user.OtpCache)
	currentPin := totp.Now()
	if currentPin != pin {
		return domain.ErrUserOTPPinMismatch
	}
	return nil
}
