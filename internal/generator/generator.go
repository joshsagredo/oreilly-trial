package generator

import (
	"errors"
	"github.com/bilalcaliskan/oreilly-trial/internal/logging"
	"github.com/bilalcaliskan/oreilly-trial/internal/mail"
	"github.com/bilalcaliskan/oreilly-trial/internal/oreilly"
	"github.com/bilalcaliskan/oreilly-trial/internal/random"
)

func RunGenerator() error {
	logger := logging.GetLogger()

	var password string
	var err error

	if password, err = random.GeneratePassword(); err != nil {
		logger.Error().Str("error", err.Error()).Msg("unable to generate password")
		return err
	}

	validDomains, err := mail.GetPossiblyValidDomains()
	if err != nil {
		logger.Error().Str("error", err.Error()).Msg("an error occurred while fetching valid domains")
		return err
	}

	for i, domain := range validDomains {
		email, err := mail.GenerateTempMail(domain)
		if err != nil {
			logger.Error().
				Str("error", err.Error()).
				Str("domain", domain).
				Msg("an error occurred while generating email with specific domain")
			continue
		}

		if err := oreilly.Generate(email, password, logger); err != nil {
			logger.Warn().
				Str("error", err.Error()).
				Str("domain", domain).
				Str("mail", email).
				Int("attempt", i+1).
				Msg("an error occurred while generating email with specific domain")
			continue
		}

		logger.Info().
			Str("email", email).
			Str("password", password).
			Msg("trial account successfully created")

		return nil
	}

	return errors.New("all attempts failed")
}
