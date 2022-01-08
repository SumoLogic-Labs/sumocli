package create_organization

import (
	"fmt"
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/spf13/cobra"
)

func NewCmdOrganizationsCreateOrganization(client *cip.APIClient) *cobra.Command {
	var (
		continuousIngest   int64
		deploymentId       string
		email              string
		firstName          string
		frequentIngest     int64
		infrequentIngest   int64
		lastName           string
		metrics            int64
		organizationName   string
		parentDeploymentId string
		tracingIngest      int64
		trialPlanPeriod    int32
	)
	cmd := &cobra.Command{
		Use:   "create-organization",
		Short: "Create a new child organization.",
		Run: func(cmd *cobra.Command, args []string) {
			createOrganization(client, continuousIngest, deploymentId, email, firstName, frequentIngest, infrequentIngest,
				lastName, metrics, organizationName, parentDeploymentId, tracingIngest, trialPlanPeriod)
		},
	}
	cmd.Flags().Int64Var(&continuousIngest, "continuousIngest", 0, "Specify the average daily amount of continuous logs this child organization is expected to ingest, in GB.")
	cmd.Flags().StringVar(&deploymentId, "deploymentId", "", "Specify the identifier of the deployment in which the organization should be created. Deployment ids can be found here: https://help.sumologic.com/APIs/General-API-Information/Sumo-Logic-Endpoints-and-Firewall-Security.")
	cmd.Flags().StringVar(&email, "email", "", "Specify the email address of the account owner.")
	cmd.Flags().StringVar(&firstName, "firstName", "", "Specify the first name of the account owner.")
	cmd.Flags().StringVar(&lastName, "lastName", "", "Specify the last name of the account owner.")
	cmd.Flags().Int64Var(&frequentIngest, "frequentIngest", 0, "Specify the average daily amount of frequent logs this child organization is expected to ingest, in GB.")
	cmd.Flags().Int64Var(&infrequentIngest, "infrequentIngest", 0, "Specify the average daily amount of infrequent logs this child organization is expected to ingest, in GB.")
	cmd.Flags().Int64Var(&metrics, "metrics", 0, "Specify the average daily amount of metrics this child organization is expected to ingest, in DPMs (Data Points per minute).")
	cmd.Flags().StringVar(&organizationName, "organizationName", "", "Specify the name of the organization.")
	cmd.Flags().StringVar(&parentDeploymentId, "parentDeploymentId", "", "Specify the identifier of the deployment on which the calling organization resides.")
	cmd.Flags().Int64Var(&tracingIngest, "tracingIngest", 0, "Specify the average daily amount of tracing data his child organization is expected to ingest, in GB.")
	cmd.Flags().Int32Var(&trialPlanPeriod, "trialPlanPeriod", 0, "Specify the the duration of the Trial plan. If not specified, your subscription plan will be used for thr created organization.")
	cmd.MarkFlagRequired("deploymentId")
	cmd.MarkFlagRequired("email")
	cmd.MarkFlagRequired("firstName")
	cmd.MarkFlagRequired("organizationName")
	cmd.MarkFlagRequired("parentDeploymentId")
	return cmd
}

func createOrganization(client *cip.APIClient, continuousIngest int64, deploymentId string, email string, firstName string,
	frequentIngest int64, infrequentIngest int64, lastName string, metrics int64, organizationName string,
	parentDeploymentId string, tracingIngest int64, trialPlanPeriod int32) {
	if cmdutils.ValidateDeploymentId(deploymentId) == false {
		fmt.Println(deploymentId + " is not a valid deployment, https://help.sumologic.com/APIs/General-API-Information/Sumo-Logic-Endpoints-and-Firewall-Security provides a list of valid deployments.")
	}
	if cmdutils.ValidateDeploymentId(parentDeploymentId) == false {
		fmt.Println(parentDeploymentId + " is not a valid deployment, https://help.sumologic.com/APIs/General-API-Information/Sumo-Logic-Endpoints-and-Firewall-Security provides a list of valid deployments.")
	}
	data, response, err := client.CreateOrganization(types.OrganizationWithSubscriptionDetails{
		DeploymentId:    deploymentId,
		TrialPlanPeriod: trialPlanPeriod,
		Baselines: &types.Baselines{
			ContinuousIngest: continuousIngest,
			FrequentIngest:   frequentIngest,
			InfrequentIngest: infrequentIngest,
			Metrics:          metrics,
			TracingIngest:    tracingIngest,
		},
		Email:            email,
		OrganizationName: organizationName,
		FirstName:        firstName,
		LastName:         lastName,
	},
		parentDeploymentId)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
