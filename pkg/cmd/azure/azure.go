package azure

import (
	"github.com/spf13/cobra"
	cmdAzureCreate "github.com/wizedkyle/sumocli/pkg/cmd/azure/create"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
)

func NewCmdAzure() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "azure <command>",
		Short: "Creates and deletes Azure infrastructure for log and metric collection",
	}

	cmd.PersistentFlags().StringVar(&factory.ApplicationId, "appid", "", "Specify the Azure Application Registration Id")
	cmd.PersistentFlags().StringVar(&factory.ApplicationSecret, "appsecret", "", "Specify the Azure Application Registration secret")
	cmd.PersistentFlags().StringVar(&factory.Cloud, "cloud", "AzurePublicCloud", "Name of the Azure Cloud")
	cmd.PersistentFlags().StringVar(&factory.Location, "location", "", "Azure Cloud location")
	cmd.PersistentFlags().StringVar(&factory.SubscriptionId, "subscriptionid", "", "Specify the Azure Subscription Id")
	cmd.PersistentFlags().StringVar(&factory.TenantId, "tenantid", "", "Specify the Azure tenant id")

	cmd.AddCommand(cmdAzureCreate.NewCmdAzureCreate())
	return cmd
}
