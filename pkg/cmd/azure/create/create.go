package create

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/features"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/storage/mgmt/storage"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
)

func NewCmdAzureCreate() *cobra.Command {
	var (
		prefix     string
		diagnostic bool
		metrics    bool
		blob       bool
	)

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create Azure infrastructure to collect logs or metrics",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logging.GetLoggerForCommand(cmd)
			logger.Debug().Msg("Create Azure infrastructure request started")
			if blob == true {
				azureCreateBlobCollection(prefix)
			} else if metrics == true {

			} else {
				fmt.Println("Please select either --logs or --metrics")
			}
			logger.Debug().Msg("Create Azure infrastructure request finished.")
		},
	}

	cmd.Flags().StringVar(&prefix, "prefix", "", "Name of the resource")
	return cmd
}

func azureCreateBlobCollection(prefix string) {
	ctx := context.Background()
	logsName := "scliblob"
	rgName := logsName + prefix
	sgName := logsName + prefix + "logs"
	createResourceGroup(ctx, rgName)
	sgAccount, err := createStorageAccount(ctx, rgName, sgName)
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
		fmt.Errorf("cannot create resource group %v: %v", rgName, err)
	}
	fmt.Println("Created or updated resource group " + rgName)
}

func createStorageAccount(ctx context.Context, rgName string, sgName string) (storage.Account, error) {
	fmt.Println("Creating or updating storage account " + sgName)
	var sku *storage.Sku
	sku.Name = "Standard_LRS"
	sku.Tier = "Standard"
	sgClient := factory.GetStorageClient()
	sgAccount, err := sgClient.Create(
		ctx,
		rgName,
		sgName,
		storage.AccountCreateParameters{
			Sku:      sku,
			Location: to.StringPtr(factory.Location),
			Tags:     factory.AzureLogTags(),
		})

	if err != nil {
		fmt.Errorf("cannot create resource group %v: %v", sgName, err)
	}

	sgAccount.WaitForCompletionRef(ctx, sgClient.Client)
	if err != nil {
		fmt.Errorf("cannot create resource group %v: %v", sgName, err)
	}

	return sgAccount.Result(sgClient)
}
