package remove

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip/types"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
)

func NewCmdPermissionsRemove(client *cip.APIClient) *cobra.Command {
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
		Use:   "remove",
		Short: "Remove permissions from a content item with the given identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			removePermissions(id, isAdminMode, notifyRecipients, notificationMessage, permissionName, sourceId, sourceType,
				client)
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

func removePermissions(id string, isAdminMode bool, notifyRecipients bool, notificationMessage string,
	permissionName string, sourceId string, sourceType string, client *cip.APIClient) {
	var (
		contentPermission types.ContentPermissionAssignment
		options           *types.ContentPermissionsOpts
		updateRequest     types.ContentPermissionUpdateRequest
	)
	contentPermission.ContentId = id
	contentPermission.PermissionName = permissionName
	contentPermission.SourceId = sourceId
	contentPermission.SourceType = sourceType
	updateRequest.ContentPermissionAssignments = append(updateRequest.ContentPermissionAssignments, contentPermission)
	updateRequest.NotifyRecipients = notifyRecipients
	updateRequest.NotificationMessage = notificationMessage
	options.IsAdminMode = optional.NewString(cmdutils.AdminMode(isAdminMode))
	data, response, err := client.RemoveContentPermissions(updateRequest,
		id,
		options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
