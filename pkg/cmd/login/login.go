package login

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdRoleList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Sumocli requires an access key and secret key.")
			fmt.Println("Sumocli will store the access key and secret key in plain text in" +
				"in the following file for use by subsequent commands:")
			// TODO: Add the file path here
			// TODO: Conformation to proceed
			// TODO: Access Key and Secret
			// TODO: write to disk with Viper
			return nil
		},
	}

	return cmd
}
