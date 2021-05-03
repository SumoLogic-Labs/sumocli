package upload

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"os"
)

func NewCmdLookupTablesUpload() *cobra.Command {
	var (
		file         string
		fileEncoding string
		id           string
		merge        bool
	)

	cmd := &cobra.Command{
		Use:   "upload",
		Short: "Create a request to populate a lookup table with a CSV file.",
		Run: func(cmd *cobra.Command, args []string) {
			UploadCSV(file, fileEncoding, id, merge)
		},
	}
	cmd.Flags().StringVar(&file, "file", "", "Specify file path of the CSV file to upload")
	cmd.Flags().StringVar(&fileEncoding, "fileEncoding", "UTF-8", "Specify encoding of the CSV file")
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the lookup table to upload the data to")
	cmd.Flags().BoolVar(&merge, "merge", false, "This indicates whether the file contents will be merged with existing data in the lookup table or not. "+
		"If this is true then data with the same primary keys will be updated while the rest of the rows will be appended.")
	cmd.MarkFlagRequired("file")
	cmd.MarkFlagRequired("id")
	return cmd
}

func UploadCSV(file string, fileEncoding string, id string, merge bool) {
	//var requestStatus api.LookupTableRequestId
	log := logging.GetConsoleLogger()
	//requestUrl := "v1/lookupTables/" + id + "/upload"
	fileData, err := os.ReadFile(file)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to read CSV file")
	}
	fmt.Println(fileData)
	fileDataString := string(fileData)
	fmt.Println(fileDataString)
}
