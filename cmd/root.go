package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/util"
)

var rootCmd = &cobra.Command{
	Use:   "sumocli",
	Short: "Manages and automates Sumo Logic using the Sumo Logic API",
	Long:  `Sumocli is a CLI tool that allows you to manage Sumo Logic based on the capabilities in the public API.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&util.AccessId, "id", "", "Provide a Sumo Logic access ID")
	rootCmd.PersistentFlags().StringVar(&util.AccessKey, "key", "", "Provide a Sumo Logic access key")
	rootCmd.PersistentFlags().StringVar(&util.ApiEndpoint, "endpoint", "", "Provide the deployment region code")
}

func Execute() error {
	return rootCmd.Execute()
}
