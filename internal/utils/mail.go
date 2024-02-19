package utils

import (
	"bytes"
	"crypto/tls"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/k3a/html2text"
	"github.com/xor111xor/s3-jwt-totp-backend/internal/domain"
	"gopkg.in/gomail.v2"
)

type EmailData struct {
	URL       string
	FirstName string
	Subject   string
}

// 👇 Email template parser
func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}

func SendEmail(user *domain.User, config domain.SysConfig) {
	// sender data
	from := config.EmailFrom
	smtpPass := config.SMTPPass
	smtpUser := config.SMTPUser
	to := user.Mail
	smtpHost := config.SMTPHost
	smtpPort := config.SMTPPort

	// enclosure data
	mailData := EmailData{
		// TODO: Disable this schema in production
		URL:       config.ServiceSchema + "//" + config.ServiceIP + ":" + config.ServicePort + "/verifyemail/" + user.VerifyString,
		FirstName: user.Mail,
		Subject:   "Your account verification code",
	}

	var body bytes.Buffer

	template, err := ParseTemplateDir("templates/mail")
	if err != nil {
		log.Fatal("Could not parse template", err)
	}

	err = template.ExecuteTemplate(&body, "verificationCode.html", &mailData)
	if err != nil {
		log.Fatal("Could not parse template", err)
	}

	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", mailData.Subject)
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send Email
	if err := d.DialAndSend(m); err != nil {
		log.Fatal("Could not send email: ", err)
	}
}
