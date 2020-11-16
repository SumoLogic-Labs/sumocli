package list

import (
	"github.com/spf13/cobra"
)

func NewCmdUserList() *cobra.Command {
	var (
		email           string
		numberOfResults string
		sortBy          string
		output          string
	)

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists Sumo Logic users",
		Long: `The following fields can be exported using the --output command:
firstName
lastName
email
roleIds
id
isActive
isLocked
isMfaEnabled
lastLoginTimestamp
`,
	}

	cmd.Flags().StringVar(&email, "email", "", "Specify the email address of the user")
	cmd.Flags().StringVar(&numberOfResults, "results", "", "Specify the number of results, this is set to 100 by default.")
	cmd.Flags().StringVar(&sortBy, "sort", "", "Sort the results by firstName, lastName or email")
	cmd.Flags().StringVar(&output, "output", "", "Specify the field to export the value from")

	return cmd
}

func listUsers() {

}

func validateSort() {

}
