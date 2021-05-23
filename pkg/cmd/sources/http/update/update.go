package update

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"strconv"
)

func NewCmdUpdateHttpSource() *cobra.Command {
	var (
		category                   string
		collectorId                int
		fieldNames                 string
		fieldValues                string
		messagePerRequest          bool
		multilineProcessingEnabled bool
		name                       string
		sourceId                   string
		merge                      bool
	)

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a specific HTTP source on the specified Sumo Logic collector",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	cmd.Flags().StringVar(&category, "category", "", "Specify the sourceCategory for the source")
	cmd.Flags().IntVar(&collectorId, "collectorId", 0, "Specify the collector id the source is associated with")
	cmd.Flags().StringVar(&fieldNames, "fieldNames", "", "Specify the names of fields to add to the source "+
		"{names need to be comma separated e.g. field1,field2")
	cmd.Flags().StringVar(&fieldValues, "fieldValues", "", "Specify the values of fields to add to the source "+
		"(values need to be comma separated e.g. value1,value2")
	cmd.Flags().BoolVar(&messagePerRequest, "messagePerRequest", false, "Specify if there is one message per request")
	cmd.Flags().BoolVar(&multilineProcessingEnabled, "multilineProcessingEnabled", false, "Specify if multiline processing is enabled")
	cmd.Flags().StringVar(&name, "name", "", "Specify the name of the source")
	cmd.Flags().StringVar(&sourceId, "sourceId", "", "Specify the source Id to update")
	cmd.Flags().BoolVar(&merge, "merge", true, "")
	cmd.MarkFlagRequired("collectorId")
	cmd.MarkFlagRequired("name")
	return cmd
}

func updateHttpSource(category string, collectorId int, fieldName string, fieldValues string,
	messagePerRequest bool, multilineProcessingEnabled bool, name string, sourceId string, merge bool) {
	var sourceResponse api.SourcesResponse
	log := logging.GetConsoleLogger()
	if merge == true {
		requestUrl := "v1/collectors/" + strconv.Itoa(collectorId) + "/sources/" + sourceId

	} else {

	}
}
