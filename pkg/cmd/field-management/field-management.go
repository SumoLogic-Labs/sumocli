package field_management

import (
	"github.com/spf13/cobra"
	NewCmdFieldManagementCreateField "github.com/wizedkyle/sumocli/pkg/cmd/field-management/create-field"
	NewCmdFieldManagementDeleteField "github.com/wizedkyle/sumocli/pkg/cmd/field-management/delete-field"
	NewCmdFieldManagementDisableCustomField "github.com/wizedkyle/sumocli/pkg/cmd/field-management/disable-custom-field"
	NewCmdFieldManagementEnableCustomField "github.com/wizedkyle/sumocli/pkg/cmd/field-management/enable-custom-field"
	NewCmdFieldManagementGetBuiltinField "github.com/wizedkyle/sumocli/pkg/cmd/field-management/get-builtin-field"
	NewCmdFieldManagementGetCapacityInfo "github.com/wizedkyle/sumocli/pkg/cmd/field-management/get-capacity-info"
	NewCmdFieldManagementGetCustomField "github.com/wizedkyle/sumocli/pkg/cmd/field-management/get-custom-field"
	NewCmdFieldManagementListBuiltinFields "github.com/wizedkyle/sumocli/pkg/cmd/field-management/list-builtin-fields"
	NewCmdFieldManagementListCustomFields "github.com/wizedkyle/sumocli/pkg/cmd/field-management/list-custom-fields"
	NewCmdFieldManagementListDroppedFields "github.com/wizedkyle/sumocli/pkg/cmd/field-management/list-dropped-fields"
)

func NewCmdFieldManagement() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "field-management",
		Short: "Manage fields",
		Long: "Fields allow you to reference log data based on meaningful associations. " +
			"They act as metadata tags that are assigned to your logs so you can search with them. " +
			"Each field contains a key-value pair, where the field name is the key. Fields may be referred to as Log Metadata Fields.",
	}
	cmd.AddCommand(NewCmdFieldManagementCreateField.NewCmdFieldManagementCreateField())
	cmd.AddCommand(NewCmdFieldManagementDeleteField.NewCmdFieldManagementDeleteField())
	cmd.AddCommand(NewCmdFieldManagementDisableCustomField.NewCmdFieldManagementDisableCustomField())
	cmd.AddCommand(NewCmdFieldManagementEnableCustomField.NewCmdFieldManagementEnableCustomField())
	cmd.AddCommand(NewCmdFieldManagementGetBuiltinField.NewCmdFieldManagementGetBuiltinField())
	cmd.AddCommand(NewCmdFieldManagementGetCapacityInfo.NewCmdFieldManagementGetCapacityInfo())
	cmd.AddCommand(NewCmdFieldManagementGetCustomField.NewCmdFieldManagementGetCustomField())
	cmd.AddCommand(NewCmdFieldManagementListBuiltinFields.NewCmdFieldManagementListBuiltinFields())
	cmd.AddCommand(NewCmdFieldManagementListCustomFields.NewCmdFieldManagementListCustomFields())
	cmd.AddCommand(NewCmdFieldManagementListDroppedFields.NewCmdFieldManagementListDroppedFields())
	return cmd
}
