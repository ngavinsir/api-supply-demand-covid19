package handler

import (
	"errors"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

// SendPasswordResetConfirmationMail sends password reset confirmation mail.
func SendPasswordResetConfirmationMail(userEmail string, requestID string) error {
	frontendURL := "http://localhost/confirm/"
	if envFrontend := os.Getenv("RESET_PASSWORD_URL"); envFrontend != "" {
		frontendURL = envFrontend
	}

	err := SendEmail(
		"Password reset confirmation",
		"Hi\n\n" +
		"Click the link below to confirm your password reset request.\n" +
		frontendURL + requestID,
		userEmail,
	)

	return err
}

// SendEmail sends emailgi ven body and receiver.
func SendEmail(subject string, body string, receiver string) error {
	godotenv.Load()

	from := "dev.supplydemandcovid19@gmail.com"
	if envFrom := os.Getenv("EMAIL_FROM"); envFrom != "" {
		from = envFrom
	}

	password := os.Getenv("EMAIL_PASSWORD")
	if password == "" {
		return errors.New("email password not set")
	}

	auth := smtp.PlainAuth(
		"",
		from,
		password,
		"smtp.gmail.com",
	)

	//_ = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n";

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		from,
		[]string{receiver},
		[]byte(
			"Subject: " + subject + "\r\n" +
			"\r\n" +
			body,
		),
	)

	return err
}