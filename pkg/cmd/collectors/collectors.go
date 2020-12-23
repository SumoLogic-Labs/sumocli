package collectors

import (
	"github.com/spf13/cobra"
	cmdCollectorCreate "github.com/wizedkyle/sumocli/pkg/cmd/collectors/create"
	cmdCollectorGet "github.com/wizedkyle/sumocli/pkg/cmd/collectors/get"
	cmdCollectorList "github.com/wizedkyle/sumocli/pkg/cmd/collectors/list"
)

func NewCmdCollectors() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "collectors <command>",
		Short: "Managed collectors",
	}

	cmd.AddCommand(cmdCollectorCreate.NewCmdCollectorCreate())
	cmd.AddCommand(cmdCollectorGet.NewCmdCollectorGet())
	cmd.AddCommand(cmdCollectorList.NewCmdCollectorList())
	return cmd
}
