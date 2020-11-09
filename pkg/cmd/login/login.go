package login

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

func NewCmdLogin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			configFile := configPath()
			fmt.Println("Sumocli requires an access key and secret key.")
			fmt.Println("Sumocli will store the access key and secret key in plain text in" +
				" the following file for use by subsequent commands:")
			fmt.Printf(configFile)
			confirmation := userConfirmation()
			if confirmation == true {
				// TODO: Access Key and Secret
				// TODO: write to disk with Viper
			} else {
				os.Exit(1)
			}

			return nil
		},
	}

	return cmd
}

func configPath() string {
	var filePath string = ".sumocli/credentials/creds.json"
	homeDirectory, _ := os.UserHomeDir()
	configFile := filepath.Join(homeDirectory, filePath)
	fmt.Println(configFile)
	return configFile
}

func userConfirmation() bool {
	prompt := promptui.Prompt{
		Label: "Do you want to proceed?",
	}

	result, err := prompt.Run()
	resultLower := strings.ToLower(result)

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	if resultLower == "yes" {
		return true
	} else {
		fmt.Println("Error: Login cancelled")
		return false
	}
}

func ReadLoginFile() {

}

func writeLoginFile() {
	// https://github.com/spf13/viper
	// https://medium.com/@jomzsg/the-easy-way-to-handle-configuration-file-in-golang-using-viper-6b3c88d2ee79
}
