package start_import

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
	"net/url"
	"os"
)

func NewCmdStartImport() *cobra.Command {
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
			startImport(file, folderId, isAdminMode, overwrite)
		},
	}
	cmd.Flags().StringVar(&file, "file", "", "File path that contains Sumo Logic content in JSON format")
	cmd.Flags().StringVar(&folderId, "folderId", "", "Specify the folder ID to import into must be in hexadecimal format. Use sumocli content get-path to get the ID of a folder")
	cmd.Flags().BoolVar(&isAdminMode, "isAdminMode", false, "Set to true if you want to perform the request as a content administrator")
	cmd.Flags().BoolVar(&overwrite, "overwrite", false, "Set to true if you want to overwrite existing content with the same name")
	cmd.MarkFlagRequired("folderId")
	return cmd
}

func startImport(file string, folderId string, isAdminMode bool, overwrite bool) {
	var responseType api.ResponseType
	var responseId api.StartExportResponse
	var responseStatusCode int
	var responseBodyData []byte
	log := logging.GetConsoleLogger()
	requestUrl := "v2/content/folders/" + folderId + "/import"
	fileData, err := os.ReadFile(file)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to read file")
	}
	err = json.Unmarshal(fileData, &responseType)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal file data")
	}
	if responseType.Type == "FolderSyncDefinition" {
		var folderSyncDefinition api.FolderSyncDefinition
		err = json.Unmarshal(fileData, &folderSyncDefinition)
		if err != nil {
			log.Error().Err(err).Msg("failed to unmarshal file data")
		}
		requestBody, err := json.Marshal(folderSyncDefinition)
		if err != nil {
			log.Error().Err(err).Msg("failed to marshal request body")
		}
		client, request := factory.NewHttpRequestWithBody("POST", requestUrl, requestBody)
		if isAdminMode == true {
			request.Header.Add("isAdminMode", "true")
		}
		query := url.Values{}
		if overwrite == true {
			query.Add("overwrite", "true")
		} else {
			query.Add("overwrite", "false")
		}
		request.URL.RawQuery = query.Encode()
		response, err := client.Do(request)
		if err != nil {
			log.Error().Err(err).Msg("failed to make http request to ")
		}
		defer response.Body.Close()
		responseBody, err := io.ReadAll(response.Body)
		responseBodyData = responseBody
		if err != nil {
			log.Error().Err(err).Msg("failed to read response body")
		}
		err = json.Unmarshal(responseBody, &responseId)
		if err != nil {
			log.Error().Err(err).Msg("failed to unmarshal response body")
		}
		responseStatusCode = response.StatusCode
	} else if responseType.Type == "DashboardSyncDefinition" {
		var dashboardSyncDefinition api.DashboardSyncDefinition
		err = json.Unmarshal(fileData, &dashboardSyncDefinition)
		if err != nil {
			log.Error().Err(err).Msg("failed to unmarshal file data")
		}
		requestBody, err := json.Marshal(dashboardSyncDefinition)
		if err != nil {
			log.Error().Err(err).Msg("failed to marshal request body")
		}
		client, request := factory.NewHttpRequestWithBody("POST", requestUrl, requestBody)
		if isAdminMode == true {
			request.Header.Add("isAdminMode", "true")
		}
		query := url.Values{}
		if overwrite == true {
			query.Add("overwrite", "true")
		} else {
			query.Add("overwrite", "false")
		}
		request.URL.RawQuery = query.Encode()
		response, err := client.Do(request)
		if err != nil {
			log.Error().Err(err).Msg("failed to make http request to ")
		}
		defer response.Body.Close()
		responseBody, err := io.ReadAll(response.Body)
		responseBodyData = responseBody
		if err != nil {
			log.Error().Err(err).Msg("failed to read response body")
		}
		err = json.Unmarshal(responseBody, &responseId)
		if err != nil {
			log.Error().Err(err).Msg("failed to unmarshal response body")
		}
		responseStatusCode = response.StatusCode
	} else if responseType.Type == "MewboardSyncDefinition" {
		var mewboardSyncDefinition api.MewboardSyncDefinition
		err = json.Unmarshal(fileData, &mewboardSyncDefinition)
		if err != nil {
			log.Error().Err(err).Msg("failed to unmarshal file data")
		}
		requestBody, err := json.Marshal(mewboardSyncDefinition)
		if err != nil {
			log.Error().Err(err).Msg("failed to marshal request body")
		}
		client, request := factory.NewHttpRequestWithBody("POST", requestUrl, requestBody)
		if isAdminMode == true {
			request.Header.Add("isAdminMode", "true")
		}
		query := url.Values{}
		if overwrite == true {
			query.Add("overwrite", "true")
		} else {
			query.Add("overwrite", "false")
		}
		request.URL.RawQuery = query.Encode()
		response, err := client.Do(request)
		if err != nil {
			log.Error().Err(err).Msg("failed to make http request to ")
		}
		defer response.Body.Close()
		responseBody, err := io.ReadAll(response.Body)
		responseBodyData = responseBody
		if err != nil {
			log.Error().Err(err).Msg("failed to read response body")
		}
		err = json.Unmarshal(responseBody, &responseId)
		if err != nil {
			log.Error().Err(err).Msg("failed to unmarshal response body")
		}
		responseStatusCode = response.StatusCode
	} else if responseType.Type == "SavedSearchWithScheduleSyncDefinition" {
		var savedSearchWithScheduleSyncDefinition api.SavedSearchWithScheduleSyncDefinition
		err = json.Unmarshal(fileData, &savedSearchWithScheduleSyncDefinition)
		if err != nil {
			log.Error().Err(err).Msg("failed to unmarshal file data")
		}
		requestBody, err := json.Marshal(savedSearchWithScheduleSyncDefinition)
		if err != nil {
			log.Error().Err(err).Msg("failed to marshal request body")
		}
		client, request := factory.NewHttpRequestWithBody("POST", requestUrl, requestBody)
		if isAdminMode == true {
			request.Header.Add("isAdminMode", "true")
		}
		query := url.Values{}
		if overwrite == true {
			query.Add("overwrite", "true")
		} else {
			query.Add("overwrite", "false")
		}
		request.URL.RawQuery = query.Encode()
		response, err := client.Do(request)
		if err != nil {
			log.Error().Err(err).Msg("failed to make http request to ")
		}
		defer response.Body.Close()
		responseBody, err := io.ReadAll(response.Body)
		responseBodyData = responseBody
		if err != nil {
			log.Error().Err(err).Msg("failed to read response body")
		}
		err = json.Unmarshal(responseBody, &responseId)
		if err != nil {
			log.Error().Err(err).Msg("failed to unmarshal response body")
		}
		responseStatusCode = response.StatusCode
	} else if responseType.Type == "MetricsSavedSearchSyncDefinition" {
		var metricsSavedSearchSyncDefinition api.MetricsSavedSearchSyncDefinition
		err = json.Unmarshal(fileData, &metricsSavedSearchSyncDefinition)
		if err != nil {
			log.Error().Err(err).Msg("failed to unmarshal file data")
		}
		requestBody, err := json.Marshal(metricsSavedSearchSyncDefinition)
		if err != nil {
			log.Error().Err(err).Msg("failed to marshal request body")
		}
		client, request := factory.NewHttpRequestWithBody("POST", requestUrl, requestBody)
		if isAdminMode == true {
			request.Header.Add("isAdminMode", "true")
		}
		query := url.Values{}
		if overwrite == true {
			query.Add("overwrite", "true")
		} else {
			query.Add("overwrite", "false")
		}
		request.URL.RawQuery = query.Encode()
		response, err := client.Do(request)
		if err != nil {
			log.Error().Err(err).Msg("failed to make http request to ")
		}
		defer response.Body.Close()
		responseBody, err := io.ReadAll(response.Body)
		responseBodyData = responseBody
		if err != nil {
			log.Error().Err(err).Msg("failed to read response body")
		}
		err = json.Unmarshal(responseBody, &responseId)
		if err != nil {
			log.Error().Err(err).Msg("failed to unmarshal response body")
		}
		responseStatusCode = response.StatusCode
	} else if responseType.Type == "MetricsSearchSyncDefinition" {
		var metricsSearchSyncDefinition api.MetricsSearchSyncDefinition
		err = json.Unmarshal(fileData, &metricsSearchSyncDefinition)
		if err != nil {
			log.Error().Err(err).Msg("failed to unmarshal file data")
		}
		requestBody, err := json.Marshal(metricsSearchSyncDefinition)
		if err != nil {
			log.Error().Err(err).Msg("failed to marshal request body")
		}
		client, request := factory.NewHttpRequestWithBody("POST", requestUrl, requestBody)
		if isAdminMode == true {
			request.Header.Add("isAdminMode", "true")
		}
		query := url.Values{}
		if overwrite == true {
			query.Add("overwrite", "true")
		} else {
			query.Add("overwrite", "false")
		}
		request.URL.RawQuery = query.Encode()
		response, err := client.Do(request)
		if err != nil {
			log.Error().Err(err).Msg("failed to make http request to ")
		}
		defer response.Body.Close()
		responseBody, err := io.ReadAll(response.Body)
		responseBodyData = responseBody
		if err != nil {
			log.Error().Err(err).Msg("failed to read response body")
		}
		err = json.Unmarshal(responseBody, &responseId)
		if err != nil {
			log.Error().Err(err).Msg("failed to unmarshal response body")
		}
		responseStatusCode = response.StatusCode
	} else if responseType.Type == "LookupTableSyncDefinition" {
		var lookupTableSyncDefinition api.LookupTableSyncDefinition
		err = json.Unmarshal(fileData, &lookupTableSyncDefinition)
		if err != nil {
			log.Error().Err(err).Msg("failed to unmarshal file data")
		}
		requestBody, err := json.Marshal(lookupTableSyncDefinition)
		if err != nil {
			log.Error().Err(err).Msg("failed to marshal request body")
		}
		client, request := factory.NewHttpRequestWithBody("POST", requestUrl, requestBody)
		if isAdminMode == true {
			request.Header.Add("isAdminMode", "true")
		}
		query := url.Values{}
		if overwrite == true {
			query.Add("overwrite", "true")
		} else {
			query.Add("overwrite", "false")
		}
		request.URL.RawQuery = query.Encode()
		response, err := client.Do(request)
		if err != nil {
			log.Error().Err(err).Msg("failed to make http request to ")
		}
		defer response.Body.Close()
		responseBody, err := io.ReadAll(response.Body)
		responseBodyData = responseBody
		if err != nil {
			log.Error().Err(err).Msg("failed to read response body")
		}
		err = json.Unmarshal(responseBody, &responseId)
		if err != nil {
			log.Error().Err(err).Msg("failed to unmarshal response body")
		}
		responseStatusCode = response.StatusCode
	}

	responseIdJson, err := json.MarshalIndent(responseId, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal responseId")
	}
	if responseStatusCode != 200 {
		factory.HttpError(responseStatusCode, responseBodyData, log)
	} else {
		fmt.Println(string(responseIdJson))
	}
}
