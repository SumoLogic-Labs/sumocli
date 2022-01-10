package get_export_result

import (
	"encoding/json"
	"fmt"
	"github.com/SumoLogic-Labs/sumocli/internal/authentication"
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/antihax/optional"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
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
			authentication.ConfirmCredentialsSet(client)
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
	data, response, err := client.GetAsyncExportResult(contentId, jobId, &options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		if data.Type_ == "FolderSyncDefinition" {
			results, response, err := client.GetFolderAsyncExportResult(contentId, jobId, &options)
			if err != nil {
				cmdutils.OutputError(response, err)
			} else {
				outputResults(saveToFile, filePath, fileName, results, response, err)
			}
		} else if data.Type_ == "DashboardSyncDefinition" {
			results, response, err := client.GetDashboardAsyncExportResult(contentId, jobId, &options)
			if err != nil {
				cmdutils.OutputError(response, err)
			} else {
				outputResults(saveToFile, filePath, fileName, results, response, err)
			}
		} else if data.Type_ == "MewboardSyncDefinition" {
			results, response, err := client.GetMewboardAsyncExportResult(contentId, jobId, &options)
			if err != nil {
				cmdutils.OutputError(response, err)
			} else {
				outputResults(saveToFile, filePath, fileName, results, response, err)
			}
		} else if data.Type_ == "SavedSearchWithScheduleSyncDefinition" {
			results, response, err := client.GetSavedSearchAsyncExportResult(contentId, jobId, &options)
			if err != nil {
				cmdutils.OutputError(response, err)
			} else {
				outputResults(saveToFile, filePath, fileName, results, response, err)
			}
		} else if data.Type_ == "MetricsSavedSearchSyncDefinition" {
			results, response, err := client.GetMetricsSearchAsyncExportResult(contentId, jobId, &options)
			if err != nil {
				cmdutils.OutputError(response, err)
			} else {
				outputResults(saveToFile, filePath, fileName, results, response, err)
			}
		} else if data.Type_ == "MetricsSearchSyncDefinition" {
			results, response, err := client.GetMetricsSearchAsyncExportResult(contentId, jobId, &options)
			if err != nil {
				cmdutils.OutputError(response, err)
			} else {
				outputResults(saveToFile, filePath, fileName, results, response, err)
			}
		} else if data.Type_ == "LookupTableSyncDefinition" {
			results, response, err := client.GetLookupTableAsyncExportResult(contentId, jobId, &options)
			if err != nil {
				cmdutils.OutputError(response, err)
			} else {
				outputResults(saveToFile, filePath, fileName, results, response, err)
			}
		}
	}
}

func outputResults(saveToFile bool, filePath string, fileName string, data interface{},
	response *http.Response, err error) {
	if saveToFile == true {
		dataBytes, err := json.Marshal(data)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to convert data to byte array")
		}
		if filePath != "" {
			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				err := os.MkdirAll(filePath, 0755)
				if err != nil {
					log.Fatal().Err(err).Msg("failed to create file path")
				}
			}
			resultFile := filepath.Join(filePath, fileName)
			err := os.WriteFile(resultFile, dataBytes, 0644)
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
			err = os.WriteFile(resultFile, dataBytes, 0644)
			if err != nil {
				log.Error().Err(err).Msg("failed to write results file " + resultFile)
			}
			fmt.Println("Results written to " + resultFile)
		}
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
