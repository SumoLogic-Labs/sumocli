package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/roles"
)

var (
	numberResults string
	filter        string
)

var rolesListCmd = &cobra.Command{
	Use:   "roles",
	Short: "Lists the roles in the Sumo Logic tenancy",
	Long: `Lists the roles in the Sumo Logic tenancy. 
This command supports setting the number of results and filter based on name.`,
	Run: func(cmd *cobra.Command, args []string) {
		roles.ListRoleIds(numberResults, filter)
	},
}

func init() {
	listCmd.AddCommand(rolesListCmd)

	rolesListCmd.PersistentFlags().StringVar(&numberResults, "results", "", "Specify the number of results, this is set to 100 by default.")
	rolesListCmd.PersistentFlags().StringVar(&filter, "filter", "", "Specify the name of the role you want to retrieve")
}
