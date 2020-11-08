package login

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func NewCmdLogin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Sumocli requires an access key and secret key.")
			fmt.Println("Sumocli will store the access key and secret key in plain text in" +
				"in the following file for use by subsequent commands:")
			userConfirmation()
			// TODO: Add the file path here
			// TODO: Access Key and Secret
			// TODO: write to disk with Viper
			return nil
		},
	}

	return cmd
}

func userConfirmation() {
	validate := func(input string) error {
		if input != "yes" {
			fmt.Println("Error: Login Cancelled")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Do you want to proceed?",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	fmt.Print(result)

}

func ReadLoginFile() {

}

func writeLoginFile() {

}
