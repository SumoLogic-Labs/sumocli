package login

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type SumoApiEndpoint struct {
	Name     string
	Code     string
	Endpoint string
}

type SumoAuth struct {
	AccessId  string `json:"accessid"`
	AccessKey string `json:"accesskey"`
	Region    string `json:"region"`
	Endpoint  string `json:"endpoint"`
}

func NewCmdLogin() *cobra.Command {
	cmd := &cobra.Command{
		Use: "login",
		RunE: func(cmd *cobra.Command, args []string) error {
			configFile := configPath()
			fmt.Println("Sumocli requires an access id and access key.")
			fmt.Println("Sumocli will store the access id and access key in plain text in" +
				" the following file for use by subsequent commands:")
			fmt.Printf(configFile)
			confirmation := userConfirmation()
			if confirmation == true {
				getCredentials()
			} else {
				os.Exit(1)
			}

			ReadCredentials()
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

func getCredentials() {
	var credentials SumoAuth

	sumoApiEndpoints := []SumoApiEndpoint{
		{Name: "Australia", Code: "au", Endpoint: "https://api.au.sumologic.com/api/"},
		{Name: "Canada", Code: "ca", Endpoint: "https://api.ca.sumologic.com/api/"},
		{Name: "Germany", Code: "de", Endpoint: "https://api.de.sumologic.com/api/"},
		{Name: "Ireland", Code: "eu", Endpoint: "https://api.eu.sumologic.com/api/"},
		{Name: "India", Code: "in", Endpoint: "https://api.in.sumologic.com/api/"},
		{Name: "Japan", Code: "jp", Endpoint: "https://api.jp.sumologic.com/api/"},
		{Name: "USA1", Code: "us1", Endpoint: "https://api.sumologic.com/api/"},
		{Name: "USA2", Code: "us2", Endpoint: "https://api.us2.sumologic.com/api/"},
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
	credentials.AccessId = accessIdResult
	credentials.AccessKey = accessKeyResult
	credentials.Region = sumoApiEndpoints[regionResultIndex].Code
	credentials.Endpoint = sumoApiEndpoints[regionResultIndex].Endpoint

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	configFilePath := filepath.Dir(configPath())
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		err := os.MkdirAll(configFilePath, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
	credentialFile, _ := json.MarshalIndent(credentials, "", "  ")
	err = ioutil.WriteFile(configPath(), credentialFile, 0644)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Credentials file saved to: " + configPath())
	}

	return
}

func ReadCredentials() (string, string) {
	viper.SetConfigName("creds")
	viper.AddConfigPath(filepath.Dir(configPath()))
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Credentials file not found at: " + configPath())
			fmt.Println("Please run sumocli login to continue...")
			os.Exit(1)
		} else {
			log.Fatal(err)
		}
	}
	accessid := viper.GetString("accessid")
	accesskey := viper.GetString("accesskey")
	endpoint := viper.GetString("endpoint")

	accessCredentials := accessid + ":" + accesskey
	accessCredentialsEnc := base64.StdEncoding.EncodeToString([]byte(accessCredentials))
	accessCredentialsComplete := "Basic " + accessCredentialsEnc
	return accessCredentialsComplete, endpoint
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
