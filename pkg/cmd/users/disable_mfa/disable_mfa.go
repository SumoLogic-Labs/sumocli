package disable_mfa

import (
	"errors"
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip/types"
	"github.com/manifoldco/promptui"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func NewCmdUserDisableMFA(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "disable-mfa",
		Short: "Disables MFA for a Sumo Logic user (this command only works interactively).",
		Run: func(cmd *cobra.Command, args []string) {
			userDisableMFA(client)
		},
	}
	return cmd
}

func userDisableMFA(client *cip.APIClient) {
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
	response, err := client.DisableMfa(types.DisableMfaRequest{
		Email:    emailResult,
		Password: passwordResult,
	},
		idResult)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(nil, response, err, "User's MFA was disabled successfully.")
	}
}
