package create

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/features"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/internal/clients"
	"github.com/wizedkyle/sumocli/pkg/logging"
)

func NewCmdAzureEventHubSourceCreate() *cobra.Command {

	cmd := &cobra.Command{}

	return cmd
}

func createEventHubSource() {
	// TODO: var sourceResponse api.AzureEventHubResponse
	rgName := "sumocli-eventhub-random" //TODO: fix this to be a real random value
	log := logging.GetConsoleLogger()
	rgClient := clients.GetResourceGroupClient()
	log.Info().Msg("creating or updating resource group " + rgName)
	rg, err := rgClient.CreateOrUpdate(
		context.TODO(),
		rgName,
		features.ResourceGroup{
			Name:     to.StringPtr(rgName),
			Location: nil,
			Tags:     nil,
		})
	if err != nil {
		log.Error().Err(err).Msg("failed to create or update resource group " + rgName)
	}

	log.Info().Msg("createing or update event hub namespace " + eventHubNamespace)
	eventHubNamespaceClient := clients.GetEventHubNamespaceClient()
	namespace, err := eventHubNamespaceClient.CreateOrUpdate()
}
