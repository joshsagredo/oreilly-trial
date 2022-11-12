package generator

import (
	"errors"
	"github.com/bilalcaliskan/oreilly-trial/internal/logging"
	"github.com/bilalcaliskan/oreilly-trial/internal/mail"
	"github.com/bilalcaliskan/oreilly-trial/internal/oreilly"
	"github.com/bilalcaliskan/oreilly-trial/internal/random"
)

func RunGenerator() error {
	var password string
	var err error

	if password, err = random.GeneratePassword(); err != nil {
		logging.GetLogger().Errorw("unable to generate password", "error", err.Error())
		return err
	}

	validDomains, err := mail.GetPossiblyValidDomains()
	if err != nil {
		logging.GetLogger().Errorw("an error occurred while fetching valid domains",
			"error", err.Error())
		return err
	}

	for i, domain := range validDomains {
		email, err := mail.GenerateTempMail(domain)
		if err != nil {
			logging.GetLogger().Errorw("an error occurred while generating email with specific domain",
				"domain", domain, "error", err.Error())
			continue
		}

		if err := oreilly.Generate(email, password); err != nil {
			logging.GetLogger().Errorw("an error occurred while generating user with tempmail", "attempt", i+1,
				"mail", email, "domain", domain, "error", err.Error())
			continue
		}

		logging.GetLogger().Infow("trial account successfully created", "email", email, "password", password)
		return nil
	}

	return errors.New("all attempts failed")
}
