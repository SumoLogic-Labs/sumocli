package field_management

import (
	NewCmdFieldManagementCreateField "github.com/SumoLogic-Labs/sumocli/pkg/cmd/field_management/create_field"
	NewCmdFieldManagementDeleteField "github.com/SumoLogic-Labs/sumocli/pkg/cmd/field_management/delete_field"
	NewCmdFieldManagementDisableCustomField "github.com/SumoLogic-Labs/sumocli/pkg/cmd/field_management/disable_custom_field"
	NewCmdFieldManagementEnableCustomField "github.com/SumoLogic-Labs/sumocli/pkg/cmd/field_management/enable_custom_field"
	NewCmdFieldManagementGetBuiltinField "github.com/SumoLogic-Labs/sumocli/pkg/cmd/field_management/get_builtin_field"
	NewCmdFieldManagementGetCapacityInfo "github.com/SumoLogic-Labs/sumocli/pkg/cmd/field_management/get_capacity_info"
	NewCmdFieldManagementGetCustomField "github.com/SumoLogic-Labs/sumocli/pkg/cmd/field_management/get_custom_field"
	NewCmdFieldManagementListBuiltinFields "github.com/SumoLogic-Labs/sumocli/pkg/cmd/field_management/list_builtin_fields"
	NewCmdFieldManagementListCustomFields "github.com/SumoLogic-Labs/sumocli/pkg/cmd/field_management/list_custom_fields"
	NewCmdFieldManagementListDroppedFields "github.com/SumoLogic-Labs/sumocli/pkg/cmd/field_management/list_dropped_fields"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/cobra"
)

func NewCmdFieldManagement(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "field-management",
		Short: "Manage fields",
		Long: "Fields allow you to reference log data based on meaningful associations. " +
			"They act as metadata tags that are assigned to your logs so you can search with them. " +
			"Each field contains a key-value pair, where the field name is the key. Fields may be referred to as Log Metadata Fields.",
	}
	cmd.AddCommand(NewCmdFieldManagementCreateField.NewCmdFieldManagementCreateField(client))
	cmd.AddCommand(NewCmdFieldManagementDeleteField.NewCmdFieldManagementDeleteField(client))
	cmd.AddCommand(NewCmdFieldManagementDisableCustomField.NewCmdFieldManagementDisableCustomField(client))
	cmd.AddCommand(NewCmdFieldManagementEnableCustomField.NewCmdFieldManagementEnableCustomField(client))
	cmd.AddCommand(NewCmdFieldManagementGetBuiltinField.NewCmdFieldManagementGetBuiltinField(client))
	cmd.AddCommand(NewCmdFieldManagementGetCapacityInfo.NewCmdFieldManagementGetCapacityInfo(client))
	cmd.AddCommand(NewCmdFieldManagementGetCustomField.NewCmdFieldManagementGetCustomField(client))
	cmd.AddCommand(NewCmdFieldManagementListBuiltinFields.NewCmdFieldManagementListBuiltinFields(client))
	cmd.AddCommand(NewCmdFieldManagementListCustomFields.NewCmdFieldManagementListCustomFields(client))
	cmd.AddCommand(NewCmdFieldManagementListDroppedFields.NewCmdFieldManagementListDroppedFields(client))
	return cmd
}
