package cmd

import (
	"github.com/spf13/cobra"
)

var (
	firstName    string
	lastName     string
	emailAddress string
	roleIds      []string
)

var usersCreateCmd = &cobra.Command{
	Use:   "user",
	Short: "Creates a Sumo Logic user",
	Long:  `Creates a Sumo Logic user by specifying the first name, last name, email and roleIds`,
	Run: func(cmd *cobra.Command, args []string) {
		GetApiCredentials(accessId, accessKey)

	},
}

func init() {
	createCmd.AddCommand(usersCreateCmd)

	usersCreateCmd.PersistentFlags().StringVar(&firstName, "fn", "", "First name of the user")
	usersCreateCmd.PersistentFlags().StringVar(&lastName, "ln", "", "Last name of the user")
	usersCreateCmd.PersistentFlags().StringVar(&emailAddress, "ea", "", "Email address for the user")
	usersCreateCmd.PersistentFlags().StringArrayVar(&roleIds, "roleids", []string{}, "Role Ids for user comma deliminated")

}
