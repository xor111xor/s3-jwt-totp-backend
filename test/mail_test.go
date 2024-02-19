package test

// Test mail sender
import (
	"testing"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

// Some variables to connect and the body.
var (
	htmlBody = `<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
		<title>Hello Gophers!</title>
	</head>
	<body>
		<p>This is the <b>Go gopher</b>.</p>
		<p><img src="cid:Gopher.png" alt="Go gopher" /></p>
		<p>Image created by Renee French</p>
	</body>
</html>`
)

func sendEmail(htmlBody string, to string, smtpClient *mail.SMTPClient) error {
	//Create the email message
	email := mail.NewMSG()

	email.SetFrom("From Example <user@mail.com>").
		AddTo(to).
		SetSubject("New Go Email")

	//Get from each mail
	email.GetFrom()
	email.SetBody(mail.TextHTML, htmlBody)

	//Send with high priority
	email.SetPriority(mail.PriorityHigh)

	// always check error after send
	if email.Error != nil {
		return email.Error
	}

	//Pass the client to the email message to send it
	return email.Send(smtpClient)
}

// TestWithTLS using gmail port 587.
func TestWithTLS(t *testing.T) {
	client := mail.NewSMTPClient()

	//SMTP Client
	client.Host = "mail.com"
	client.Port = 587
	client.Username = "user@mail.com"
	client.Password = "pass"
	client.Encryption = mail.EncryptionSTARTTLS
	client.ConnectTimeout = 10 * time.Second
	client.SendTimeout = 10 * time.Second

	//KeepAlive is not settted because by default is false

	//Connect to client
	smtpClient, err := client.Connect()

	if err != nil {
		t.Error("Expected nil, got", err, "connecting to client")
	}

	err = sendEmail(htmlBody, "test@mail.com", smtpClient)
	if err != nil {
		t.Error("Expected nil, got", err, "sending email")
	}
}
