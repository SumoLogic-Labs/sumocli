package create

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
)

func NewCmdTokensCreate(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	var (
		description string
		inactive    bool
		name        string
	)
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a token in the token library.",
		Run: func(cmd *cobra.Command, args []string) {
			createToken(description, inactive, name, client, log)
		},
	}
	cmd.Flags().StringVar(&description, "description", "", "Specify a description for the token")
	cmd.Flags().BoolVar(&inactive, "inactive", false, "Set to true if you want the token to be inactive")
	cmd.Flags().StringVar(&name, "name", "", "Specify a name for the token")
	cmd.MarkFlagRequired("name")
	return cmd
}

func createToken(description string, inactive bool, name string, client *cip.APIClient, log *zerolog.Logger) {
	var options types.TokenBaseDefinition
	if inactive == true {
		options.Status = "Inactive"
	} else {
		options.Status = "Active"
	}
	options.Name = name
	options.Description = description
	options.Type_ = "CollectorRegistration"
	apiResponse, httpResponse, errorResponse := client.CreateToken(options)
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to create token")
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
