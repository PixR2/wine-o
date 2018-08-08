package main

import (
	machinery "github.com/RichardKnop/machinery/v1"
	config "github.com/RichardKnop/machinery/v1/config"
	"github.com/go-gomail/gomail"
	"github.com/pkg/errors"
)

func main() {
	var cnf = &config.Config{
		Broker:        "redis://localhost:6379",
		DefaultQueue:  "default_queue",
		ResultBackend: "redis://localhost:6379",
		NoUnixSignals: true,
	}

	machineryServer, err := machinery.NewServer(cnf)
	if err != nil {
		panic(errors.Wrap(err, "failed to initialize machinery server"))
	}

	machineryServer.RegisterTask("sendAccountVerificationMail", func(email string) error {
		m := gomail.NewMessage()
		m.SetHeader("From", "admin@wineo.com")
		m.SetHeader("To", email)
		m.SetHeader("Subject", "Verify your WineO account!")
		m.SetBody("text/html", `Click <href a="www.pyramics.com">here</a> to verify your WimeO account.`)

		d := gomail.NewDialer("smtp.gmail.com", 587, "hannesradmer@gmail.com", "Thedamnedstandready0913.GE")

		if err := d.DialAndSend(m); err != nil {
			return errors.Wrapf(err, "sending verification email to '%s' failed", email)
		}

		return nil
	})
}
