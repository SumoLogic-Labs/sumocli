package cmd

import (
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Sets the scope to create operations",
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Sets the scope to list operations",
}

func init() {
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(listCmd)
}
