package create

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/features"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/storage/mgmt/storage"
	"github.com/Azure/azure-service-bus-go"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/rs/zerolog"
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
			log := logging.GetConsoleLogger()
			logger.Debug().Msg("Create Azure infrastructure request started")
			if blob == true {
				azureCreateBlobCollection(prefix, log)
			} else if metrics == true {

			} else if diagnostic == true {

			} else {
				fmt.Println("Please select either --diagnostic, --logs or --metrics")
			}
			logger.Debug().Msg("Create Azure infrastructure request finished.")
		},
	}

	cmd.Flags().BoolVar(&blob, "blob", false, "Deploys infrastructure for Azure Blob collection.")
	cmd.Flags().StringVar(&prefix, "prefix", "", "Name of the resource")
	return cmd
}

func azureCreateBlobCollection(prefix string, log zerolog.Logger) {
	ctx := context.Background()
	logsName := "scliblob"
	rgName := logsName + prefix
	sgName := logsName + prefix + "logs"
	nsName := logsName + prefix + "logs"

	createResourceGroup(ctx, rgName, log)
	sgAccount, err := createStorageAccount(ctx, rgName, sgName, log)
	if err != nil {
		log.Error().Err(err).Msg("error creating storage account")
	}
	fmt.Println(to.String(sgAccount.Name)) // TODO: Remove this line
	createStorageAccountTable(ctx, rgName, sgName, log)
	createServiceBusNamespace(nsName, log)
}

/*
func azureCreateLogCollection() {
}

func azureCreateMetricCollection() {
}
*/

func createResourceGroup(ctx context.Context, rgName string, log zerolog.Logger) {
	log.Info().Msg("Creating or updating resource group " + rgName)
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
		log.Error().Err(err).Msg("cannot create resource group " + rgName)
	}
	log.Info().Msg("Created or updated resource group " + rgName)
}

func createStorageAccount(ctx context.Context, rgName string, sgName string, log zerolog.Logger) (storage.Account, error) {
	log.Info().Msg("Creating or updating storage account " + sgName)
	sgClient := factory.GetStorageClient()
	sgAccount, err := sgClient.Create(
		ctx,
		rgName,
		sgName,
		storage.AccountCreateParameters{
			Sku: &storage.Sku{
				Name: storage.StandardLRS,
				Tier: storage.Standard,
			},
			Kind:     storage.StorageV2,
			Location: to.StringPtr(factory.Location),
			Tags:     factory.AzureLogTags(),
		})

	if err != nil {
		log.Error().Err(err).Msg("cannot create resource group " + sgName)
	}

	err = sgAccount.WaitForCompletionRef(ctx, sgClient.Client)
	if err != nil {
		log.Error().Err(err).Msg("cannot create resource group " + sgName)
	}

	log.Info().Msg("Created or updated storage account " + rgName)
	return sgAccount.Result(sgClient)
}

func createStorageAccountTable(ctx context.Context, rgName string, sgName string, log zerolog.Logger) {
	log.Info().Msg("Creating FileOffsetMap table")
	tableClient := factory.GetStorageTableClient()
	_, err := tableClient.Create(
		ctx,
		rgName,
		sgName,
		"FileOffsetMap")

	if err != nil {
		log.Error().Err(err).Msg("cannot create FileOffsetMap table")
	}

	log.Info().Msg("Created FileOffsetMap table")
}

func createServiceBusNamespace(nsName string, log zerolog.Logger) {
	log.Info().Msg("Creating Service Bus namespace " + nsName)
	namespace, err := servicebus.NewNamespace(
		servicebus.NamespaceWithAzureEnvironment(
			nsName,
			factory.Cloud))
}
