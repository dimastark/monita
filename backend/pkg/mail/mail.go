package mail

import (
	"bytes"
	"fmt"

	"monita/config"

	"gopkg.in/gomail.v2"
)

// TemplateData holds all data needed to send
type TemplateData struct {
	ReportType string
	Updates    []UpdateData
}

// UpdateData holds all data about observable update
type UpdateData struct {
	ID      uint
	Name    string
	OldData string
	NewData string
}

// Send send updates to `email`
func Send(email string, data TemplateData) error {
	var buffer bytes.Buffer

	err := t.Execute(&buffer, data)

	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", config.EmailFrom)
	m.SetHeader("To", email)
	m.SetHeader("Subject", fmt.Sprintf("Monita %s report", data.ReportType))
	m.SetBody("text/html", buffer.String())

	return d.DialAndSend(m)
}
