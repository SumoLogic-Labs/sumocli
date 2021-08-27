package add

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
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
	apiResponse, httpResponse, errorResponse := client.AddAllowlistedCidrs(cmdutils.GenerateCidrList(
		ipAddresses, descriptions))
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
