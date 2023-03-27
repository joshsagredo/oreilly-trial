package prompt

import (
	"errors"

	"github.com/bilalcaliskan/oreilly-trial/internal/mail"
	"github.com/manifoldco/promptui"
)

type SelectRunner interface {
	Run() (int, string, error)
}

func GetSelectRunner() *promptui.Select {
	return &promptui.Select{
		Label: "An error occurred while generating Oreilly account with temporary mail, would you like to provide your own valid email address?",
		Items: []string{"Yes please!", "No thanks!"},
	}
}

type PromptRunner interface {
	Run() (string, error)
}

func GetPromptRunner() *promptui.Prompt {
	return &promptui.Prompt{
		Label: "Your valid email address",
		Validate: func(s string) error {
			if !mail.IsValidEmail(s) {
				return errors.New("no valid email provided by user")
			}

			return nil
		},
	}
}
