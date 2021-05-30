package partitions

import (
	"github.com/spf13/cobra"
	NewCmdPartitionsCancelRetentionUpdate "github.com/wizedkyle/sumocli/pkg/cmd/partitions/cancel-retention-update"
	NewCmdPartitionsCreate "github.com/wizedkyle/sumocli/pkg/cmd/partitions/create"
	NewCmdPartitionsDecommission "github.com/wizedkyle/sumocli/pkg/cmd/partitions/decommission"
	NewCmdPartitionsGet "github.com/wizedkyle/sumocli/pkg/cmd/partitions/get"
	NewCmdPartitionsList "github.com/wizedkyle/sumocli/pkg/cmd/partitions/list"
	NewCmdPartitionsUpdate "github.com/wizedkyle/sumocli/pkg/cmd/partitions/update"
)

func NewCmdPartitions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "partitions",
		Short: "Manage partitions",
		Long:  "Creating a Partition allows you to improve search performance by searching over a smaller number of messages.",
	}
	cmd.AddCommand(NewCmdPartitionsCancelRetentionUpdate.NewCmdPartitionsCancelRetentionUpdate())
	cmd.AddCommand(NewCmdPartitionsCreate.NewCmdPartitionCreate())
	cmd.AddCommand(NewCmdPartitionsDecommission.NewCmdPartitionsDecommission())
	cmd.AddCommand(NewCmdPartitionsGet.NewCmdPartitionsGet())
	cmd.AddCommand(NewCmdPartitionsList.NewCmdPartitionsList())
	cmd.AddCommand(NewCmdPartitionsUpdate.NewCmdPartitionUpdate())
	return cmd
}
