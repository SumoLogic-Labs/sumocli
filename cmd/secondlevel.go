package cmd

import (
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Sets the scope to create operations",
}

func init() {
	rootCmd.AddCommand(createCmd)
}
