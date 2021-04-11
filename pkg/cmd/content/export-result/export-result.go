package export_result

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
	"os"
	"path/filepath"
)

func NewCmdExportResult() *cobra.Command {
	var (
		contentId   string
		jobId       string
		isAdminMode bool
		saveToFile  bool
		filePath    string
		fileName    string
	)

	cmd := &cobra.Command{
		Use:   "export-result",
		Short: "Gets results from content export job for the given job identifier",
		Run: func(cmd *cobra.Command, args []string) {
			exportResult(contentId, jobId, isAdminMode, saveToFile, filePath, fileName)
		},
	}
	cmd.Flags().StringVar(&contentId, "contentId", "", "Specify the id of the content item to export")
	cmd.Flags().StringVar(&jobId, "jobId", "", "Specify the job id for the export (returned from running sumocli content start-export)")
	cmd.Flags().BoolVar(&isAdminMode, "isAdminMode", false, "Set to true if you want to perform the request as a content administrator")
	cmd.Flags().BoolVar(&saveToFile, "saveToFile", false, "Saves the export results to a file")
	cmd.Flags().StringVar(&filePath, "filePath", "", "Folder path to save results to. If this is empty then the results will be saved to the users home directory")
	cmd.Flags().StringVar(&fileName, "fileName", "", "File name for the results.")
	cmd.MarkFlagRequired("contentId")
	cmd.MarkFlagRequired("jobId")
	return cmd
}

func exportResult(contentId string, jobId string, isAdminMode bool, saveToFile bool, filePath string, fileName string) {
	var responseType api.ResponseType
	log := logging.GetConsoleLogger()
	requestUrl := "v2/content/" + contentId + "/export/" + jobId + "/result"
	client, request := factory.NewHttpRequest("GET", requestUrl)
	if isAdminMode == true {
		request.Header.Add("isAdminMode", "true")
	}
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request to " + requestUrl)
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read response body")
	}

	err = json.Unmarshal(responseBody, &responseType)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	if responseType.Type == "FolderSyncDefinition" {
		var folderSyncResponse api.FolderSyncDefinition
		err = json.Unmarshal(responseBody, &folderSyncResponse)
		if err != nil {
			log.Error().Err(err).Msg("failed to unmarshal response body into DashboardSyncDefinition")
		}
		folderSyncResponseJson, err := json.MarshalIndent(folderSyncResponse, "", "    ")
		if err != nil {
			log.Error().Err(err).Msg("failed to marshal dashboardSyncResponse")
		}
		writeResults(folderSyncResponseJson, saveToFile, filePath, fileName)
	} else if responseType.Type == "DashboardSyncDefinition" {
		var dashboardSyncResponse api.DashboardSyncDefinition
		err = json.Unmarshal(responseBody, &dashboardSyncResponse)
		if err != nil {
			log.Error().Err(err).Msg("failed to unmarshal response body into DashboardSyncDefinition")
		}
		dashboardSyncResponseJson, err := json.MarshalIndent(dashboardSyncResponse, "", "    ")
		if err != nil {
			log.Error().Err(err).Msg("failed to marshal dashboardSyncResponse")
		}
		writeResults(dashboardSyncResponseJson, saveToFile, filePath, fileName)
	} else if responseType.Type == "MewboardSyncDefinition" {
		var mewboardSyncResponse api.MewboardSyncDefinition
		err = json.Unmarshal(responseBody, &mewboardSyncResponse)
		if err != nil {
			log.Error().Err(err).Msg("failed to unmarshal response body into DashboardSyncDefinition")
		}
		mewboardSyncResponseJson, err := json.MarshalIndent(mewboardSyncResponse, "", "    ")
		if err != nil {
			log.Error().Err(err).Msg("failed to marshal dashboardSyncResponse")
		}
		writeResults(mewboardSyncResponseJson, saveToFile, filePath, fileName)
	} else if responseType.Type == "SavedSearchWithScheduleSyncDefinition" {
		var savedSearchWithScheduleSyncResponse api.SavedSearchWithScheduleSyncDefinition
		err = json.Unmarshal(responseBody, &savedSearchWithScheduleSyncResponse)
		if err != nil {
			log.Error().Err(err).Msg("failed to unmarshal response body into DashboardSyncDefinition")
		}
		savedSearchWithScheduleSyncResponseJson, err := json.MarshalIndent(savedSearchWithScheduleSyncResponse, "", "    ")
		if err != nil {
			log.Error().Err(err).Msg("failed to marshal dashboardSyncResponse")
		}
		writeResults(savedSearchWithScheduleSyncResponseJson, saveToFile, filePath, fileName)
	} else if responseType.Type == "MetricsSavedSearchSyncDefinition" {
		var metricsSavedSearchSyncResponse api.MetricsSavedSearchSyncDefinition
		err = json.Unmarshal(responseBody, &metricsSavedSearchSyncResponse)
		if err != nil {
			log.Error().Err(err).Msg("failed to unmarshal response body into DashboardSyncDefinition")
		}
		metricsSavedSearchSyncResponseJson, err := json.MarshalIndent(metricsSavedSearchSyncResponse, "", "    ")
		if err != nil {
			log.Error().Err(err).Msg("failed to marshal dashboardSyncResponse")
		}
		writeResults(metricsSavedSearchSyncResponseJson, saveToFile, filePath, fileName)
	} else if responseType.Type == "MetricsSearchSyncDefinition" {
		var metricsSearchSyncResponse api.MetricsSearchSyncDefinition
		err = json.Unmarshal(responseBody, &metricsSearchSyncResponse)
		if err != nil {
			log.Error().Err(err).Msg("failed to unmarshal response body into DashboardSyncDefinition")
		}
		metricsSearchSyncResponseJson, err := json.MarshalIndent(metricsSearchSyncResponse, "", "    ")
		if err != nil {
			log.Error().Err(err).Msg("failed to marshal dashboardSyncResponse")
		}
		writeResults(metricsSearchSyncResponseJson, saveToFile, filePath, fileName)
	} else if responseType.Type == "LookupTableSyncDefinition" {
		var lookupTableSyncResponse api.LookupTableSyncDefinition
		err = json.Unmarshal(responseBody, &lookupTableSyncResponse)
		if err != nil {
			log.Error().Err(err).Msg("failed to unmarshal response body into DashboardSyncDefinition")
		}
		lookupTableSyncResponseJson, err := json.MarshalIndent(lookupTableSyncResponse, "", "    ")
		if err != nil {
			log.Error().Err(err).Msg("failed to marshal dashboardSyncResponse")
		}
		writeResults(lookupTableSyncResponseJson, saveToFile, filePath, fileName)
	}
}

func writeResults(resultsData []byte, saveToFile bool, filePath string, fileName string) {
	if saveToFile == true {
		if filePath != "" {
			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				err := os.MkdirAll(filePath, 0755)
				if err != nil {
					log.Fatal().Err(err).Msg("failed to create file path")
				}
			}
			resultFile := filepath.Join(filePath, fileName)
			err := os.WriteFile(resultFile, resultsData, 0644)
			if err != nil {
				log.Error().Err(err).Msg("failed to write results file " + filePath)
			}
			fmt.Println("Results written to " + resultFile)
		} else {
			homeDirectory, err := os.UserHomeDir()
			if err != nil {
				log.Error().Err(err).Msg("failed to get user home directory")
			}
			resultFile := filepath.Join(homeDirectory, "results.json")
			err = os.WriteFile(resultFile, resultsData, 0644)
			if err != nil {
				log.Error().Err(err).Msg("failed to write results file " + resultFile)
			}
			fmt.Println("Results written to " + resultFile)
		}
	} else {
		fmt.Println(string(resultsData))
	}
}
