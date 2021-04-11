package get

import (
	"github.com/spf13/cobra"
)

func NewCmdLookupTablesGet() *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Gets a Sumo Logic lookup table based on the given identifier",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the lookup table you want to retrieve")
	cmd.MarkFlagRequired("id")
	return cmd
}

func getLookupTable(id string) {

}
