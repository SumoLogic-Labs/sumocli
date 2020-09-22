package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/version"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Sets the scope to create operations",
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Sets the scope to list operations",
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays the version of sumocli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.AppName + " " + version.Version + " " + version.BuildVersion)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(versionCmd)
}
