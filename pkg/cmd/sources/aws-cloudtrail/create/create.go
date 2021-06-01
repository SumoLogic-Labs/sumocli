package create

import (
	"github.com/spf13/cobra"
)

func NewCmdAWSCloudTrailSourceCreate() *cobra.Command {
	var ()

	cmd := &cobra.Command{
		Use: "create",
		Short: "AWS CloudTrail records API calls made to AWS. " +
			"This includes calls made using the AWS Management Console, AWS SDKs, command line tools, and higher-level AWS services. " +
			"Add an AWS CloudTrail Source to upload these messages to Sumo Logic. " +
			"The AWS CloudTrail Source automatically parses the logs prior to upload.",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	return cmd
}

func createAWSCloudTrailSource() {

}
