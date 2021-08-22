package get_export_result

import (
	"encoding/json"
	"fmt"
	"github.com/antihax/optional"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
	"net/http"
	"os"
	"path/filepath"
)

func NewCmdGetExportResult(client *cip.APIClient) *cobra.Command {
	var (
		contentId   string
		jobId       string
		isAdminMode bool
		saveToFile  bool
		filePath    string
		fileName    string
	)
	cmd := &cobra.Command{
		Use:   "get-export-result",
		Short: "Gets results from content export job for the given job identifier",
		Run: func(cmd *cobra.Command, args []string) {
			exportResult(contentId, jobId, isAdminMode, saveToFile, filePath, fileName, client)
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

func exportResult(contentId string, jobId string, isAdminMode bool, saveToFile bool, filePath string, fileName string,
	client *cip.APIClient) {
	var options types.ContentOpts
	if isAdminMode == true {
		options.IsAdminMode = optional.NewString("true")
	} else {
		options.IsAdminMode = optional.NewString("false")
	}
	apiResponse, httpResponse, errorResponse := client.GetAsyncExportResult(contentId, jobId, &options)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		if apiResponse.Type_ == "FolderSyncDefinition" {
			results, httpResponse, errorResponse := client.GetFolderAsyncExportResult(contentId, jobId, &options)
			if errorResponse != nil {
				cmdutils.OutputError(httpResponse, errorResponse)
			} else {
				outputResults(saveToFile, filePath, fileName, results, httpResponse, errorResponse)
			}
		} else if apiResponse.Type_ == "DashboardSyncDefinition" {
			results, httpResponse, errorResponse := client.GetDashboardAsyncExportResult(contentId, jobId, &options)
			if errorResponse != nil {
				cmdutils.OutputError(httpResponse, errorResponse)
			} else {
				outputResults(saveToFile, filePath, fileName, results, httpResponse, errorResponse)
			}
		} else if apiResponse.Type_ == "MewboardSyncDefinition" {
			results, httpResponse, errorResponse := client.GetMewboardAsyncExportResult(contentId, jobId, &options)
			if errorResponse != nil {
				cmdutils.OutputError(httpResponse, errorResponse)
			} else {
				outputResults(saveToFile, filePath, fileName, results, httpResponse, errorResponse)
			}
		} else if apiResponse.Type_ == "SavedSearchWithScheduleSyncDefinition" {
			results, httpResponse, errorResponse := client.GetSavedSearchAsyncExportResult(contentId, jobId, &options)
			if errorResponse != nil {
				cmdutils.OutputError(httpResponse, errorResponse)
			} else {
				outputResults(saveToFile, filePath, fileName, results, httpResponse, errorResponse)
			}
		} else if apiResponse.Type_ == "MetricsSavedSearchSyncDefinition" {
			results, httpResponse, errorResponse := client.GetMetricsSearchAsyncExportResult(contentId, jobId, &options)
			if errorResponse != nil {
				cmdutils.OutputError(httpResponse, errorResponse)
			} else {
				outputResults(saveToFile, filePath, fileName, results, httpResponse, errorResponse)
			}
		} else if apiResponse.Type_ == "MetricsSearchSyncDefinition" {
			results, httpResponse, errorResponse := client.GetMetricsSearchAsyncExportResult(contentId, jobId, &options)
			if errorResponse != nil {
				cmdutils.OutputError(httpResponse, errorResponse)
			} else {
				outputResults(saveToFile, filePath, fileName, results, httpResponse, errorResponse)
			}
		} else if apiResponse.Type_ == "LookupTableSyncDefinition" {
			results, httpResponse, errorResponse := client.GetLookupTableAsyncExportResult(contentId, jobId, &options)
			if errorResponse != nil {
				cmdutils.OutputError(httpResponse, errorResponse)
			} else {
				outputResults(saveToFile, filePath, fileName, results, httpResponse, errorResponse)
			}
		}
	}
}

func outputResults(saveToFile bool, filePath string, fileName string, apiResponse interface{},
	httpResponse *http.Response, errorResponse error) {
	if saveToFile == true {
		apiResponseBytes, err := json.Marshal(apiResponse)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to convert apiResponse to byte array")
		}
		if filePath != "" {
			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				err := os.MkdirAll(filePath, 0755)
				if err != nil {
					log.Fatal().Err(err).Msg("failed to create file path")
				}
			}
			resultFile := filepath.Join(filePath, fileName)
			err := os.WriteFile(resultFile, apiResponseBytes, 0644)
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
			err = os.WriteFile(resultFile, apiResponseBytes, 0644)
			if err != nil {
				log.Error().Err(err).Msg("failed to write results file " + resultFile)
			}
			fmt.Println("Results written to " + resultFile)
		}
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
