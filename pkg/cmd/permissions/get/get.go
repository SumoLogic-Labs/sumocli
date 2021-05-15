package get

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
	"net/url"
)

func NewCmdPermissionsGet() *cobra.Command {
	var (
		id           string
		explicitOnly bool
		isAdminMode  bool
	)

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Returns content permissions of a content item with the given identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			getPermissions(id, explicitOnly, isAdminMode)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of a content item")
	cmd.Flags().BoolVar(&explicitOnly, "explicitOnly", false, "There are two permission types: explicit and implicit. "+
		"Permissions specifically assigned to the content item are explicit. Permissions derived from a parent content item, like a folder are implicit. "+
		"To return only explicit permissions set this to true.")
	cmd.Flags().BoolVar(&isAdminMode, "isAdminMode", false, "Set to true if you want to perform the request as a content administrator")
	cmd.MarkFlagRequired("id")
	return cmd
}

func getPermissions(id string, explicitOnly bool, isAdminMode bool) {
	var permissionsResponse api.GetPermissions
	log := logging.GetConsoleLogger()
	requestUrl := "v2/content/" + id + "/permissions"
	client, request := factory.NewHttpRequest("GET", requestUrl)
	query := url.Values{}
	if explicitOnly == false {
		query.Add("explicitOnly", "false")
	} else {
		query.Add("explicitOnly", "true")
	}
	request.URL.RawQuery = query.Encode()
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

	err = json.Unmarshal(responseBody, &permissionsResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	permissionsResponseJson, err := json.MarshalIndent(permissionsResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal foldersResponse")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(permissionsResponseJson))
	}
}
