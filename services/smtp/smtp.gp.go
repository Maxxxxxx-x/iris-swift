package smtp

import (
	"fmt"
	"github.com/Maxxxxxx-x/iris-swift/config"
	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
)

func NewClient(cfg config.SMTPConfig) error {
	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	client, err := smtp.Dial(addr)
	if err != nil {
		return err
	}
	auth := sasl.NewPlainClient("", cfg.Base_Sender_Email, cfg.Password)
	if err := client.Auth(auth); err != nil {
		
	}
}
