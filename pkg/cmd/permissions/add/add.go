package add

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdPermissionsAdd() *cobra.Command {
	var (
		id                  string
		isAdminMode         bool
		notifyRecipients    bool
		notificationMessage string
		permissionName      string
		sourceId            string
		sourceType          string
	)

	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add permissions to a content item with the given identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			addPermissions(id, isAdminMode, notifyRecipients, notificationMessage, permissionName, sourceId, sourceType)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of a content item")
	cmd.Flags().BoolVar(&isAdminMode, "isAdminMode", false, "Set to true if you want to perform the request as a content administrator")
	cmd.Flags().BoolVar(&notifyRecipients, "notifyRecipients", false, "Set to true if you want to notify recipients")
	cmd.Flags().StringVar(&notificationMessage, "notificationMessage", "", "Specify a notification message")
	cmd.Flags().StringVar(&permissionName, "permissionName", "", "Specify a content permission name."+
		"Valid values are: View, GrantView, Edit, GrantEdit, Manage, GrantManage")
	cmd.Flags().StringVar(&sourceId, "sourceId", "", "Specify an identifier that relates to the source type")
	cmd.Flags().StringVar(&sourceType, "sourceType", "", "Specify a source type."+
		"Valid values are: user, role, org")
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("notificationMessage")
	cmd.MarkFlagRequired("permissionName")
	cmd.MarkFlagRequired("sourceId")
	cmd.MarkFlagRequired("sourceType")
	return cmd
}

func addPermissions(id string, isAdminMode bool, notifyRecipients bool, notificationMessage string,
	permissionName string, sourceId string, sourceType string) {
	var permissionsResponse api.GetPermissions
	log := logging.GetConsoleLogger()
	requestBodySchema := api.ModifyPermissionsRequest{
		NotifyRecipients:    notifyRecipients,
		NotificationMessage: notificationMessage,
	}
	permissions := api.PermissionsDetail{
		PermissionName: permissionName,
		SourceType:     sourceType,
		SourceId:       sourceId,
		ContentId:      id,
	}
	requestBodySchema.ContentPermissionAssignments = append(requestBodySchema.ContentPermissionAssignments, permissions)
	requestBody, err := json.Marshal(requestBodySchema)
	if err != nil {
		log.Error().Err(err).Msg("error marshalling request body")
	}
	requestUrl := "v2/content/" + id + "/permissions/add"
	client, request := factory.NewHttpRequestWithBody("PUT", requestUrl, requestBody)
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
