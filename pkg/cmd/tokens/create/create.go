package create

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
)

func NewCmdTokensCreate(client *cip.APIClient) *cobra.Command {
	var (
		description string
		inactive    bool
		name        string
	)
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a token in the token library.",
		Run: func(cmd *cobra.Command, args []string) {
			createToken(description, inactive, name, client)
		},
	}
	cmd.Flags().StringVar(&description, "description", "", "Specify a description for the token")
	cmd.Flags().BoolVar(&inactive, "inactive", false, "Set to true if you want the token to be inactive")
	cmd.Flags().StringVar(&name, "name", "", "Specify a name for the token")
	cmd.MarkFlagRequired("name")
	return cmd
}

func createToken(description string, inactive bool, name string, client *cip.APIClient) {
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
		cmdutils.OutputError(httpResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
