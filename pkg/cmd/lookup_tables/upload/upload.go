package upload

import (
	"fmt"
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip/types"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
	"os"
)

func NewCmdLookupTablesUpload(client *cip.APIClient) *cobra.Command {
	var (
		fileEncoding string
		fileLocation string
		id           string
		merge        bool
	)
	cmd := &cobra.Command{
		Use:   "upload",
		Short: "Create a request to populate a lookup table with a CSV file",
		Run: func(cmd *cobra.Command, args []string) {
			uploadData(fileEncoding, fileLocation, id, merge, client)
		},
	}
	cmd.Flags().StringVar(&fileEncoding, "fileEncoding", "UTF-8", "Encoding of the CSV file being uploaded.")
	cmd.Flags().StringVar(&fileLocation, "fileLocation", "", "File path of the CSV file.")
	cmd.Flags().StringVar(&id, "id", "", "Id of the lookup table to upload the file to.")
	cmd.Flags().BoolVar(&merge, "merge", false, "If this is set to true file contents will be merged with existing data in the lookup table. "+
		"By default the data with the same primary keys will be updated while the reset of the rows will be appended.")
	cmd.MarkFlagRequired("fileLocation")
	cmd.MarkFlagRequired("id")
	return cmd
}

func uploadData(fileEncoding string, fileLocation string, id string, merge bool, client *cip.APIClient) {
	var options types.LookupTableUploadFileOpts
	options.FileEncoding = optional.NewString(fileEncoding)
	options.Merge = optional.NewBool(merge)
	file, err := os.Open(fileLocation)
	if err != nil {
		fmt.Println("File location (" + fileLocation + ") couldn't be opened, please the file exists.")
		os.Exit(1)
	}
	data, response, err := client.UploadFile(file, id, &options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
