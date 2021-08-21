package start_import

import (
	"encoding/json"
	"github.com/antihax/optional"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
	"os"
)

func NewCmdStartImport(client *cip.APIClient) *cobra.Command {
	var (
		file        string
		folderId    string
		isAdminMode bool
		overwrite   bool
	)
	cmd := &cobra.Command{
		Use:   "start-import",
		Short: "Schedule an asynchronous import of content inside an existing folder with the given identifier. The start-import command can be used to create or update content within a folder.",
		Run: func(cmd *cobra.Command, args []string) {
			startImport(file, folderId, isAdminMode, overwrite, client)
		},
	}
	cmd.Flags().StringVar(&file, "file", "", "File path that contains Sumo Logic content in JSON format")
	cmd.Flags().StringVar(&folderId, "folderId", "", "Specify the folder ID to import into must be in hexadecimal format. Use sumocli content get-path to get the ID of a folder")
	cmd.Flags().BoolVar(&isAdminMode, "isAdminMode", false, "Set to true if you want to perform the request as a content administrator")
	cmd.Flags().BoolVar(&overwrite, "overwrite", false, "Set to true if you want to overwrite existing content with the same name")
	cmd.MarkFlagRequired("folderId")
	return cmd
}

func startImport(file string, folderId string, isAdminMode bool, overwrite bool, client *cip.APIClient) {
	var (
		contentType types.ContentSyncDefinition
		options     types.ContentImportOpts
	)
	if isAdminMode == true {
		options.IsAdminMode = optional.NewString("true")
	} else {
		options.IsAdminMode = optional.NewString("false")
	}
	options.Overwrite = optional.NewBool(overwrite)
	fileData, err := os.ReadFile(file)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to read file")
	}
	err = json.Unmarshal(fileData, &contentType)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to unmarshal file data")
	}
	if contentType.Type_ == "FolderSyncDefinition" {
		var folderContent types.FolderSyncDefinition
		err = json.Unmarshal(fileData, &folderContent)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to unmarshal file data")
		}
		apiResponse, httpResponse, errorResponse := client.BeginFolderAsyncImport(folderContent, folderId, &options)
		if errorResponse != nil {
			cmdutils.OutputError(httpResponse)
		} else {
			cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
		}
	} else if contentType.Type_ == "DashboardSyncDefinition" {
		var dashboardContent types.DashboardSyncDefinition
		err = json.Unmarshal(fileData, &dashboardContent)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to unmarshal file data")
		}
		apiResponse, httpResponse, errorResponse := client.BeginDashboardAsyncImport(dashboardContent, folderId, &options)
		if errorResponse != nil {
			cmdutils.OutputError(httpResponse)
		} else {
			cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
		}
	} else if contentType.Type_ == "MewboardSyncDefinition" {
		var mewboardContent types.MewboardSyncDefinition
		err = json.Unmarshal(fileData, &mewboardContent)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to unmarshal file data")
		}
		apiResponse, httpResponse, errorResponse := client.BeginMewboardAsyncImport(mewboardContent, folderId, &options)
		if errorResponse != nil {
			cmdutils.OutputError(httpResponse)
		} else {
			cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
		}
	} else if contentType.Type_ == "SavedSearchWithScheduleSyncDefinition" {
		var savedSearchContent types.SavedSearchWithScheduleSyncDefinition
		err = json.Unmarshal(fileData, &savedSearchContent)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to unmarshal file data")
		}
		apiResponse, httpResponse, errorResponse := client.BeginSavedSearchAsyncImport(savedSearchContent, folderId, &options)
		if errorResponse != nil {
			cmdutils.OutputError(httpResponse)
		} else {
			cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
		}
	} else if contentType.Type_ == "MetricsSavedSearchSyncDefinition" {
		var metricsSavedSearchContent types.MetricsSavedSearchSyncDefinition
		err = json.Unmarshal(fileData, &metricsSavedSearchContent)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to unmarshal file data")
		}
		apiResponse, httpResponse, errorResponse := client.BeginMetricsSavedSearchAsyncImport(metricsSavedSearchContent, folderId, &options)
		if errorResponse != nil {
			cmdutils.OutputError(httpResponse)
		} else {
			cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
		}
	} else if contentType.Type_ == "MetricsSearchSyncDefinition" {
		var metricsSearchContent types.MetricsSearchSyncDefinition
		err = json.Unmarshal(fileData, &metricsSearchContent)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to unmarshal file data")
		}
		apiResponse, httpResponse, errorResponse := client.BeginMetricsSearchAsyncImport(metricsSearchContent, folderId, &options)
		if errorResponse != nil {
			cmdutils.OutputError(httpResponse)
		} else {
			cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
		}
	} else if contentType.Type_ == "LookupTableSyncDefinition" {
		var lookupTableContent types.LookupTableSyncDefinition
		err = json.Unmarshal(fileData, &lookupTableContent)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to unmarshal file data")
		}
		apiResponse, httpResponse, errorResponse := client.BeginLookupTableAsyncImport(lookupTableContent, folderId, &options)
		if errorResponse != nil {
			cmdutils.OutputError(httpResponse)
		} else {
			cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
		}
	}
}
