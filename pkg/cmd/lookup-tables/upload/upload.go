package upload

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
	"mime/multipart"
	"os"
	"path"
)

type uploadRequest struct {
	File string `json:"file"`
}

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
	log := logging.GetConsoleLogger()
	/*
		var requestStatus api.LookupTableRequestId
		//requestUrl := "v1/lookupTables/" + id + "/upload"
		fileData, err := os.ReadFile(file)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to read CSV file")
		}
		fmt.Println(fileData)
		fileDataString := string(fileData)
		fmt.Println(fileDataString)

	*/

	fileDir, _ := os.Getwd()
	fileName := "test.csv"
	filePath := path.Join(fileDir, fileName)

	fileData, _ := os.Open(filePath)
	fileInfo, _ := fileData.Stat()
	defer fileData.Close()

	fileContents, _ := io.ReadAll(fileData)
	fmt.Println(fileContents)
	binString := ""
	for _, n := range fileContents {
		binString = fmt.Sprintf("%s%.8b", binString, n)
	}
	fmt.Println(binString)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", fileInfo.Name())
	if err != nil {
		log.Error().Err(err).Msg("failed to create form file")
	}

	part.Write([]byte(binString))
	writer.Close()

	fmt.Println(body)

	client, request := factory.NewHttpUploadRequest("POST", "v1/lookupTables/00000000007CE861/upload", body)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to upload file")
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read response body")
	}
	fmt.Println(string(responseBody))
}
