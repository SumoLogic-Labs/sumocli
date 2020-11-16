package create

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	util2 "github.com/wizedkyle/sumocli/pkg/cmdutil"
	"io/ioutil"
	"os"
	"strings"
)

func NewCmdRoleCreate() *cobra.Command {
	var (
		name         string
		description  string
		filter       string
		users        []string
		capabilities []string
		autofill     bool
		output       string
	)
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a Sumo Logic role",
		Run: func(cmd *cobra.Command, args []string) {
			createRole(name, description, filter, users, capabilities, autofill, output)
		},
	}

	cmd.Flags().StringVar(&name, "name", "", "Name of the role")
	cmd.Flags().StringVar(&description, "description", "", "Description for the role")
	cmd.Flags().StringVar(&filter, "filter", "", "Search filter for the role")
	cmd.Flags().StringSliceVar(&users, "users", []string{}, "Comma deliminated list of user ids to add to the role")
	cmd.Flags().StringSliceVar(&capabilities, "capabilities", []string{}, "Comma deliminated list of capabilities")
	cmd.Flags().BoolVar(&autofill, "autofill", true, "Is set to true by default.")
	cmd.Flags().StringVar(&output, "output", "", "Specify the field to export the value from")

	return cmd
}

func createRole(name string, description string, filter string, users []string, capabilities []string, autofill bool, output string) {
	var createRoleResponse api.RoleData

	if name == "" {
		fmt.Println("--name field needs to be specified.")
		os.Exit(0)
	}

	for i, capability := range capabilities {
		if validateCapabilities(capability) == false {
			fmt.Println(capability + " is not a valid Sumo Logic role capability.")
		}
		i++
	}

	requestBodySchema := &api.CreateRoleRequest{
		Name:                 name,
		Description:          description,
		FilterPredicate:      filter,
		Users:                users,
		Capabilities:         capabilities,
		AutoFillDependencies: autofill,
	}

	requestBody, _ := json.Marshal(requestBodySchema)
	client, request := factory.NewHttpRequestWithBody("POST", "v1/roles", requestBody)
	response, err := client.Do(request)
	util2.LogError(err)

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)

	jsonErr := json.Unmarshal(responseBody, &createRoleResponse)
	util2.LogError(jsonErr)

	createRoleResponseJson, err := json.MarshalIndent(createRoleResponse, "", "    ")
	util2.LogError(err)

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody)
	} else {
		if factory.ValidateRoleOutput(output) == true {
			value := gjson.Get(string(createRoleResponseJson), output)
			formattedValue := strings.Trim(value.String(), `"[]"`)
			fmt.Println(formattedValue)
		} else {
			fmt.Println(string(createRoleResponseJson))
			fmt.Println(createRoleResponse.Name + " role successfully created")
		}
	}
}

func validateCapabilities(capability string) bool {
	switch capability {
	case
		"viewCollectors",
		"manageCollectors",
		"manageBudgets",
		"manageDataVolumeFeed",
		"viewFieldExtraction",
		"manageFieldExtractionRules",
		"manageS3DataForwarding",
		"manageContent",
		"dataVolumeIndex",
		"viewConnections",
		"manageConnections",
		"viewScheduledViews",
		"manageScheduledViews",
		"viewPartitions",
		"managePartitions",
		"viewFields",
		"manageFields",
		"viewAccountOverview",
		"manageTokens",
		"manageDataStreams",
		"manageEntityTypeConfig",
		"manageMonitors",
		"metricsTransformation",
		"metricsExtraction",
		"metricsRules",
		"managePasswordPolicy",
		"ipWhitelisting",
		"createAccessKeys",
		"manageAccessKeys",
		"manageSupportAccountAccess",
		"manageAuditDataFeed",
		"manageSaml",
		"shareDashboardOutsideOrg",
		"manageOrgSettings",
		"changeDataAccessLevel",
		"shareDashboardWorld",
		"shareDashboardWhitelist",
		"manageUsersAndRoles",
		"searchAuditIndex",
		"auditEventIndex":
		return true
	}
	return false
}
