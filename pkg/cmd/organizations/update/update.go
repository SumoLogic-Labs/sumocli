package update

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip/types"
	"github.com/spf13/cobra"
)

func NewCmdOrganizationsUpdate(client *cip.APIClient) *cobra.Command {
	var (
		continuousIngest   int64
		frequentIngest     int64
		infrequentIngest   int64
		metrics            int64
		organizationId     string
		parentDeploymentId string
		tracingIngest      int64
	)
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update an existing organization's subscription based on its identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			update(client, continuousIngest, frequentIngest, infrequentIngest, organizationId, metrics,
				parentDeploymentId, tracingIngest)
		},
	}
	cmd.Flags().Int64Var(&continuousIngest, "continuousIngest", 0, "Specify the average daily amount of continuous logs this child organization is expected to ingest, in GB.")
	cmd.Flags().Int64Var(&frequentIngest, "frequentIngest", 0, "Specify the average daily amount of frequent logs this child organization is expected to ingest, in GB.")
	cmd.Flags().Int64Var(&infrequentIngest, "infrequentIngest", 0, "Specify the average daily amount of infrequent logs this child organization is expected to ingest, in GB.")
	cmd.Flags().Int64Var(&metrics, "metrics", 0, "Specify the average daily amount of metrics this child organization is expected to ingest, in DPMs (Data Points per minute).")
	cmd.Flags().StringVar(&organizationId, "organizationId", "", "Specify the identifier of the organization to update.")
	cmd.Flags().StringVar(&parentDeploymentId, "parentDeploymentId", "", "Specify the identifier of the deployment on which the calling organization resides.")
	cmd.Flags().Int64Var(&tracingIngest, "tracingIngest", 0, "Specify the average daily amount of tracing data his child organization is expected to ingest, in GB.")
	cmd.MarkFlagRequired("organizationId")
	cmd.MarkFlagRequired("parentDeploymentId")
	return cmd
}

func update(client *cip.APIClient, continuousIngest int64, frequentIngest int64, infrequentIngest int64,
	organizationId string, metrics int64, parentDeploymentId string, tracingIngest int64) {
	data, response, err := client.UpdateOrganization(types.Baselines{
		ContinuousIngest: continuousIngest,
		FrequentIngest:   frequentIngest,
		InfrequentIngest: infrequentIngest,
		Metrics:          metrics,
		TracingIngest:    tracingIngest,
	},
		parentDeploymentId,
		organizationId)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
