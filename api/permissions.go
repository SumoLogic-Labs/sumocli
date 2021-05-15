package api

type GetPermissions struct {
	ExplicitPermissions []PermissionsDetail `json:"explicitPermissions"`
	ImplicitPermissions []PermissionsDetail `json:"implicitPermissions"`
}

type ModifyPermissionsRequest struct {
	ContentPermissionAssignments []PermissionsDetail `json:"contentPermissionAssignments"`
	NotifyRecipients             bool                `json:"notifyRecipients"`
	NotificationMessage          string              `json:"notificationMessage"`
}

type PermissionsDetail struct {
	PermissionName string `json:"permissionName"`
	SourceType     string `json:"sourceType"`
	SourceId       string `json:"sourceId"`
	ContentId      string `json:"contentId"`
}
