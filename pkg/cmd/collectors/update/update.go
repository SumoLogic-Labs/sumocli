package update

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
)

func NewCmdCollectorUpdate(client *cip.APIClient) *cobra.Command {
	var (
		category           string
		cutoffTimestamp    int32
		description        string
		ephemeral          bool
		fieldNames         []string
		fieldValues        []string
		hostName           string
		id                 string
		installedCollector bool
		name               string
		sourceSyncMode     string
		timeZone           string
		targetCPU          int32
	)
	cmd := &cobra.Command{
		Use:   "update",
		Short: "updates a Sumo Logic collector settings",
		Run: func(cmd *cobra.Command, args []string) {
			updateCollector(category, id, cutoffTimestamp, description, ephemeral, fieldNames, fieldValues, hostName,
				installedCollector, name, sourceSyncMode, timeZone, targetCPU, client)
		},
	}
	cmd.Flags().StringVar(&category, "category", "", "Specify a category for the collector")
	cmd.Flags().StringVar(&id, "id", "", "Id of the collector you want to update")
	cmd.Flags().Int32Var(&cutoffTimestamp, "cutoffTimestamp", 0, "Specify a cutoff timestamp for the collector, specified as milliseconds since epoch")
	cmd.Flags().StringVar(&description, "description", "", "Specify a description for the collector")
	cmd.Flags().BoolVar(&ephemeral, "ephemeral", false, "When true the collector will be deleted after 12 hours of inactivity, defaults to false")
	cmd.Flags().StringSliceVar(&fieldNames, "fieldNames", []string{}, "Specify the names of fields to add to the collector "+
		"{names need to be comma separated e.g. field1,field2")
	cmd.Flags().StringSliceVar(&fieldValues, "fieldValues", []string{}, "Specify the values of fields to add to the collector "+
		"(values need to be comma separated e.g. value1,value2")
	cmd.Flags().StringVar(&hostName, "hostName", "", "Host name of the collector")
	cmd.Flags().BoolVar(&installedCollector, "installedCollector", false, "Set to true if you are wanting to update an installed collector")
	cmd.Flags().StringVar(&name, "name", "", "Name of the collector, it must be unique on your account")
	cmd.Flags().StringVar(&sourceSyncMode, "sourceSyncMode", "", "For installed collectors whether the Collector is using local source of cloud management"+
		"(\"Json\" for local source and \"UI\" for cloud source this is only configurable on installed collectors")
	cmd.Flags().StringVar(&timeZone, "timeZone", "", "Time zone of the Collector. Refer to the TZ column of this site: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones")
	cmd.Flags().Int32Var(&targetCPU, "targetCPU", 0, "When CPU utilisation exceeds this threshold, the Collector will slow down its rate of ingestion to lower its CPU utilisation"+
		"(only configurable on installable collectors)")
	cmd.MarkFlagRequired("id")
	return cmd
}

func updateCollector(category string, id string, cutoffTimestamp int32, description string, ephemeral bool,
	fieldNames []string, fieldValues []string, hostName string, installedCollector bool, name string, sourceSyncMode string,
	timeZone string, targetCPU int32, client *cip.APIClient) {
	fields := cmdutils.GenerateFieldsMap(fieldNames, fieldValues)
	if installedCollector == true {
		apiResponse, httpResponse, errorResponse := client.UpdateInstalledCollector(types.UpdateInstalledCollectorDefinition{
			Collector: types.UpdateInstalledCollectorModel{
				Id:              id,
				Name:            name,
				Description:     description,
				Category:        category,
				CutOffTimestamp: cutoffTimestamp,
				Fields:          fields,
				Ephemeral:       ephemeral,
				HostName:        hostName,
				SourceSyncMode:  sourceSyncMode,
				TimeZone:        timeZone,
				TargetCPU:       targetCPU,
			},
		},
			id)
		if errorResponse != nil {
			cmdutils.OutputError(httpResponse)
		} else {
			cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
		}
	} else {
		apiResponse, httpResponse, errorResponse := client.UpdateHostedCollector(types.UpdateHostedCollectorDefinition{
			Collector: types.UpdateHostedCollectorModel{
				Id:              id,
				Name:            name,
				Description:     description,
				Category:        category,
				CutOffTimestamp: cutoffTimestamp,
				Fields:          fields,
				TimeZone:        timeZone,
			},
		},
			id)
		if errorResponse != nil {
			cmdutils.OutputError(httpResponse)
		} else {
			cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
		}
	}
}
