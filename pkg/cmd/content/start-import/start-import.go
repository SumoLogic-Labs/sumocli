package start_import

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"os"
)

func NewCmdStartImport() *cobra.Command {
	var (
		file        string
		folderId    string
		isAdminMode bool
		overwrite   bool
	)

	cmd := &cobra.Command{
		Use:   "start-import",
		Short: "Schedule an asynchronous import of content inside an existing folder with the given identifier. The start-import command can be used to create or update content within a folder.",
		Run: func(cmd *cobra.Command, args []string) {
			startImport(file, folderId, isAdminMode, overwrite)
		},
	}
	cmd.Flags().StringVar(&file, "file", "", "File path that contains Sumo Logic content in JSON format")
	cmd.Flags().StringVar(&folderId, "folderId", "", "Specify the folder ID to import into must be in hexadecimal format. Use sumocli content get-path to get the ID of a folder")
	cmd.Flags().BoolVar(&isAdminMode, "isAdminMode", false, "Set to true if you want to perform the request as a content administrator")
	cmd.Flags().BoolVar(&overwrite, "overwrite", false, "Set to true if you want to overwrite existing content with the same name")
	cmd.MarkFlagRequired("folderId")
	return cmd
}

func startImport(file string, folderId string, isAdminMode bool, overwrite bool) {
	var responseType api.ResponseType
	log := logging.GetConsoleLogger()
	fileData, err := os.ReadFile(file)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to read file")
	}
	err = json.Unmarshal(fileData, responseType)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal file data")
	}

	// Read file
	// Unmarshall type and have if statements check which type it is
	// create requestBodySchema

	//requestUrl := "v2/content/folders/" + folderId + "/import"
}
