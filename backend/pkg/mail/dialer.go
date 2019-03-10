package mail

import (
	"monita/config"

	"gopkg.in/gomail.v2"
)

var d *gomail.Dialer

func init() {
	d = gomail.NewDialer(
		config.DialerAddress,
		config.DialerPort,
		config.DialerUser,
		config.DialerPassword,
	)
}
