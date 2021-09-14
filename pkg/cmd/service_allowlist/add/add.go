package add

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdServiceAllowlistAdd(client *cip.APIClient) *cobra.Command {
	var (
		ipAddresses  []string
		descriptions []string
	)
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add CIDR notations and/or IP addresses to the allowlist of the organization if not already there.",
		Long:  "When service allowlisting functionality is enabled, CIDRs/IP addresses that are allowlisted will have access to Sumo Logic and/or content sharing.",
		Run: func(cmd *cobra.Command, args []string) {
			addServiceAllowlistCidr(ipAddresses, descriptions, client)
		},
	}
	cmd.Flags().StringSliceVar(&ipAddresses, "ipAddresses", []string{}, "Specify the IP addresses to add in CIDR format. Multiple addresses can be specified by comma separating them.")
	cmd.Flags().StringSliceVar(&descriptions, "description", []string{}, "Specify the descriptions for the IP addresses. Provide a comma separated list of descriptions relating to the IP addresses.")
	cmd.MarkFlagRequired("cidr")
	return cmd
}

func addServiceAllowlistCidr(ipAddresses []string, descriptions []string, client *cip.APIClient) {
	data, response, err := client.AddAllowlistedCidrs(cmdutils.GenerateCidrList(
		ipAddresses, descriptions))
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
