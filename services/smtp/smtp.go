package smtp

import (
	"fmt"
	"github.com/Maxxxxxx-x/iris-swift/config"
	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
)

type SmtpClient interface {
	Send(to string, body string) error
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

func (s *smtpClient) Send(to string, body string) error {
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

	if _, err := w.Write([]byte(body)); err != nil {
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
