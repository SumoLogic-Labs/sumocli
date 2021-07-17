package disable_mfa

import (
	"errors"
	"github.com/manifoldco/promptui"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
)

func NewCmdUserDisableMFA(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "disable-mfa",
		Short: "Disables MFA for a Sumo Logic user (this command only works interactively).",
		Run: func(cmd *cobra.Command, args []string) {
			userDisableMFA(client, log)
		},
	}
	return cmd
}

func userDisableMFA(client *cip.APIClient, log *zerolog.Logger) {
	validate := func(input string) error {
		if input == "" {
			return errors.New("Value is empty")
		}
		return nil
	}
	promptId := promptui.Prompt{
		Label:    "Please enter the Sumo Logic id for the user",
		Validate: validate,
	}
	promptEmail := promptui.Prompt{
		Label:    "Please enter the Sumo Logic users email address",
		Validate: validate,
	}
	promptPassword := promptui.Prompt{
		Label:    "Please enter the Sumo Logic users password",
		Mask:     '*',
		Validate: validate,
	}
	promptConfirm := promptui.Prompt{
		Label:     "Confirm that you want to disable MFA? Removing MFA can be a security risk!",
		IsConfirm: true,
	}
	idResult, err := promptId.Run()
	emailResult, err := promptEmail.Run()
	passwordResult, err := promptPassword.Run()
	_, err = promptConfirm.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to generate prompt")
	}
	httpResponse, errorResponse := client.DisableMfa(types.DisableMfaRequest{
		Email:    emailResult,
		Password: passwordResult,
	},
		idResult)
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to disable mfa")
	} else {
		cmdutils.Output(nil, httpResponse, errorResponse, "User's MFA was disabled successfully.")
	}
}
