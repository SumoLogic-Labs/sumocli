package configure

import (
	"encoding/json"
	"fmt"
	"github.com/SumoLogic-Labs/sumocli/api"
	"github.com/SumoLogic-Labs/sumocli/internal/authentication"
	"github.com/SumoLogic-Labs/sumocli/internal/encryption"
	"github.com/SumoLogic-Labs/sumocli/pkg/logging"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

func NewCmdConfigure() *cobra.Command {
	var showAccessId bool

	cmd := &cobra.Command{
		Use:   "configure",
		Short: "Sets Sumo Logic credentials",
		Long:  "Interactively sets the Sumo Logic Access Id, Access Key and API endpoint.",
		Run: func(cmd *cobra.Command, args []string) {
			if showAccessId == true {
				accessId := authentication.ReadAccessId()
				if accessId != "" {
					fmt.Println("The access id currently configured for authentication is: " + accessId)
					os.Exit(1)
				} else {
					os.Exit(1)
				}
			}
			configFile := authentication.ConfigPath()
			fmt.Println("Sumocli requires an access id and access key.")
			fmt.Println("Sumocli will encrypt and store the access id and access key in" +
				" the following file for use by subsequent commands: " + configFile)
			confirmation := userConfirmation()
			if confirmation == true {
				getCredentials()
			} else {
				os.Exit(1)
			}
			return
		},
	}
	cmd.Flags().BoolVar(&showAccessId, "showAccessId", false, "Shows the plain text access key. This command "+
		"is useful for identify which access key is being used. If this flag is set you cannot set login credentials.")
	return cmd
}

func getCredentials() {
	var credentials api.SumoAuth
	log := logging.GetConsoleLogger()
	sumoApiEndpoints := []api.SumoApiEndpoint{
		{Name: "Australia", Code: "au", Endpoint: "https://api.au.sumologic.com/api"},
		{Name: "Canada", Code: "ca", Endpoint: "https://api.ca.sumologic.com/api"},
		{Name: "Germany", Code: "de", Endpoint: "https://api.de.sumologic.com/api"},
		{Name: "Ireland", Code: "eu", Endpoint: "https://api.eu.sumologic.com/api"},
		{Name: "India", Code: "in", Endpoint: "https://api.in.sumologic.com/api"},
		{Name: "Japan", Code: "jp", Endpoint: "https://api.jp.sumologic.com/api"},
		{Name: "USA1", Code: "us1", Endpoint: "https://api.sumologic.com/api"},
		{Name: "USA2", Code: "us2", Endpoint: "https://api.us2.sumologic.com/api"},
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "{{ .Name }}",
		Inactive: "{{ .Name }}",
		Selected: "{{ .Name }}",
		Details: `
------- Sumo Logic API Endpoints -------
{{ "Name:" }} {{ .Name }}
{{ "Code:" }} {{ .Code }}
{{ "Endpoint" }} {{ .Endpoint }}`,
	}

	searcher := func(input string, index int) bool {
		endpoint := sumoApiEndpoints[index]
		name := strings.Replace(strings.ToLower(endpoint.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	promptRegion := promptui.Select{
		Label:     "Please select your Sumo Logic tenancy endpoint",
		Items:     sumoApiEndpoints,
		Templates: templates,
		Size:      8,
		Searcher:  searcher,
	}

	promptAccessId := promptui.Prompt{
		Label: "Please enter your Sumo Logic Access Id",
		Mask:  '*',
	}

	promptAccessKey := promptui.Prompt{
		Label: "Please enter your Sumo Logic Access Key",
		Mask:  '*',
	}

	accessIdResult, err := promptAccessId.Run()
	accessKeyResult, err := promptAccessKey.Run()
	regionResultIndex, _, err := promptRegion.Run()
	credentials.Version = "v1"
	credentials.AccessId = encryption.EncryptData(accessIdResult)
	credentials.AccessKey = encryption.EncryptData(accessKeyResult)
	credentials.Region = sumoApiEndpoints[regionResultIndex].Code
	credentials.Endpoint = sumoApiEndpoints[regionResultIndex].Endpoint

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	configFilePath := filepath.Dir(authentication.ConfigPath())
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		err := os.MkdirAll(configFilePath, 0755)
		if err != nil {
			log.Fatal().Err(err)
		}
	}
	credentialFile, _ := json.MarshalIndent(credentials, "", "  ")
	err = os.WriteFile(authentication.ConfigPath(), credentialFile, 0644)
	if err != nil {
		log.Fatal().Err(err)
	} else {
		fmt.Println("Credentials file saved to: " + authentication.ConfigPath())
	}

	return
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
		fmt.Println("Login cancelled")
		return false
	}
}
