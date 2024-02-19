package domain

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/thanhpk/randstr"
	"github.com/xlzd/gotp"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrNoUser                   = fmt.Errorf("no user")
	ErrUserExist                = fmt.Errorf("user with this mail exist")
	ErrUserVerified             = fmt.Errorf("User has verified with this mail string")
	ErrUserVerifyStringNotMatch = fmt.Errorf("verify string not match")
	ErrUserPasswordMismatch     = fmt.Errorf("user password mismatch")
	ErrUserOTPPinMismatch       = fmt.Errorf("OTP pin mismatch")
	ErrNoFile                   = fmt.Errorf("no file")
)

type User struct {
	UuidUser     uuid.UUID
	Mail         string
	PasswordHash string
	OtpCache     string
	CreatedAt    time.Time
	Verified     bool
	VerifyString string
	Files        *[]File
}

type File struct {
	UuidFile   uuid.UUID
	Name       string
	UploadDate time.Time
	Bucket     string
	UuidUser   uuid.UUID
	Size       int64
}

type RepoDB interface {
	UserGetByMail(mail string) (User, error)
	UserAddOnRegistration(user User) error
	UserCheckExistByMail(mail string) (bool, error)
	FileInsert(file File) error
	FileDelete(file File) error
	FileGetByID(string) (File, error)
	FilesGetByUserID(string) (*[]File, error)
}

type Cache interface {
	Add(user User) error
	Update(user User) error
	Get(mail string) (User, error)
	CheckVerifyRegString(checkMail string) (*User, error)
	LengthCache() (float64, error)
	LengthUnverifiedUsers() (float64, error)
}

type StorageS3 interface {
	Upload(File) error
	Download(File) error
	Remove(File) error
}

func NewUser(mail, password, passwordRepeat string) (User, error) {
	user := User{}
	var err error

	if password != passwordRepeat {
		return user, ErrUserPasswordMismatch
	}

	user.UuidUser = uuid.New()
	user.Mail = mail
	user.PasswordHash, err = HashPassword(password)
	if err != nil {
		return user, err
	}
	user.OtpCache = gotp.RandomSecret(16)

	user.VerifyString = randstr.String(20)

	return user, nil
}

func NewFile(name, bucket string, size int64, user_id uuid.UUID) File {
	return File{
		UuidFile:   uuid.New(),
		Name:       name,
		UploadDate: time.Now(),
		Bucket:     bucket,
		UuidUser:   user_id,
		Size:       size,
	}
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}

func GetOtpUrl(user *User) string {
	parseMail := strings.Split(user.Mail, "@")
	return gotp.NewDefaultTOTP(user.OtpCache).ProvisioningUri(parseMail[0], "store")
}

func CheckUserCredential(user *User, mail, pasword, pin string) error {
	// mail
	if user.Mail != mail {
		return ErrNoUser
	}
	// pasword
	if err := VerifyPassword(user.PasswordHash, pasword); err != nil {
		return ErrUserPasswordMismatch
	}
	// otp
	totp := gotp.NewDefaultTOTP(user.OtpCache)
	currentPin := totp.Now()
	if currentPin != pin {
		return ErrUserOTPPinMismatch
	}
	return nil
}
