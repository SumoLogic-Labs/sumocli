package cmd

import (
	"encoding/base64"
	"fmt"
	"github.com/spf13/cobra"
)

var sumoApiEndpoints = map[string]string{
	"au":  "https://api.au.sumologic.com/api/",
	"ca":  "https://api.ca.sumologic.com/api/",
	"de":  "https://api.de.sumologic.com/api/",
	"in":  "https://api.in.sumologic.com/api/",
	"jp":  "https://api.jp.sumologic.com/api/",
	"us1": "https://api.sumologic.com/api/",
	"us2": "https://api.us2.sumologic.com/api/",
}

var (
	accessId    string
	accessKey   string
	apiEndpoint string
)

var rootCmd = &cobra.Command{
	Use:   "sumocli",
	Short: "Manages and automates Sumo Logic using the Sumo Logic API",
	Long:  `Sumocli is a CLI tool that allows you to manage Sumo Logic based on the capabilities in the public API.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&accessId, "id", "", "Provide a Sumo Logic access ID")
	rootCmd.PersistentFlags().StringVar(&accessKey, "key", "", "Provide a Sumo Logic access key")
	rootCmd.PersistentFlags().StringVar(&apiEndpoint, "endpoint", "", "Provide the deployment region code")
}

func Execute() error {
	return rootCmd.Execute()
}

func GetApiEndpoint(apiEndpoint string) string {
	apiEndpointUrl := sumoApiEndpoints[apiEndpoint]
	return apiEndpointUrl
}

func GetApiCredentials(accessId string, accessKey string) string {
	accessCredentials := accessId + ":" + accessKey
	accessCredentialsEnc := base64.StdEncoding.EncodeToString([]byte(accessCredentials))
	fmt.Println(accessCredentials)
	fmt.Println(accessCredentialsEnc)
	return accessCredentialsEnc
}
