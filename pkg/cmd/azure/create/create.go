package create

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/features"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
)

func NewCmdAzureCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create Azure infrastructure to collect logs or metrics",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logging.GetLoggerForCommand(cmd)
			logger.Debug().Msg("Create Azure infrastructure request started")
			ctx := context.Background()
			createResourceGroup(ctx, "sumocli-log-collection")
			logger.Debug().Msg("Create Azure infrastructure request finished.")
		},
	}
	return cmd
}

/*
func azureCreateLogCollection() {
}

func azureCreateMetricCollection() {
}
*/

func createResourceGroup(ctx context.Context, rgName string) {
	fmt.Println("Creating or updating resource group " + rgName)
	rgClient := factory.GetResourceGroupClient()
	_, err := rgClient.CreateOrUpdate(
		ctx,
		rgName,
		features.ResourceGroup{
			Name:     to.StringPtr(rgName),
			Location: to.StringPtr(factory.Location),
			Tags:     factory.AzureLogTags(),
		})

	if err != nil {
		fmt.Println("cannot create resource group " + rgName)
		fmt.Println(err)
	}
	ctx.Done()
	fmt.Println("Created or updated resource group " + rgName)
}
