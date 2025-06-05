package smtp

import (
	"bytes"
	"fmt"
	"github.com/Maxxxxxx-x/iris-swift/config"
	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
)

type SmtpClient interface {
	Send(to string, subject string, body string) error
	Close() error
}

type smtpClient struct {
	client *smtp.Client
	config config.SMTPConfig
}

func NewClient(cfg config.SMTPConfig) (SmtpClient, error) {
	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	client, err := smtp.Dial(addr)
	if err != nil {
		return nil, err
	}
	auth := sasl.NewPlainClient("", cfg.Base_Sender_Email, cfg.Password)
	if err := client.Auth(auth); err != nil {
		return nil, err
	}

	smtpclient := &smtpClient{
		client: client,
	}

	return smtpclient, nil
}

func (s *smtpClient) Send(to string, subject string, body string) error {
	if err := s.client.Mail(s.config.Base_Sender_Email, nil); err != nil {
		return err
	}
	if err := s.client.Rcpt(to, nil); err != nil {
		return err
	}
	w, err := s.client.Data()
	if err != nil {
		return err
	}
	defer w.Close()

	var msg bytes.Buffer
	msg.WriteString(fmt.Sprintf("From: %s\r\n", s.config.Base_Sender_Email))
	msg.WriteString(fmt.Sprintf("To: %s\r\n", to))
	msg.WriteString(fmt.Sprintf("Subject: %s\r\n", subject))
	msg.WriteString(fmt.Sprintf("Content-Type: text/html; charset=\"UTF-8\"\r\n"))
	msg.WriteString("\r\n")
	msg.WriteString(body)
	if _, err := msg.WriteTo(w); err != nil {
		return err
	}

	return nil
}

func (s *smtpClient) Close() error {
	if err := s.client.Quit(); err != nil {
		return err
	}
	return nil
}
