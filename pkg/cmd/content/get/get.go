package get

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdGet(client *cip.APIClient) *cobra.Command {
	var (
		contentId string
		path      string
	)
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Gets a content item corresponding to the provided path or via the identifier of the content.",
		Run: func(cmd *cobra.Command, args []string) {
			getContent(contentId, path, client)
		},
	}
	cmd.Flags().StringVar(&contentId, "contentId", "", "Specify the id of the content")
	cmd.Flags().StringVar(&path, "path", "", "Specify the path of the content you want to retrieve (e.g. /Library/Users/user@demo.com/SampleFolder)")
	return cmd
}

func getContent(contentId string, path string, client *cip.APIClient) {
	if contentId != "" {
		apiResponse, httpResponse, errorResponse := client.GetPathById(contentId)
		if errorResponse != nil {
			cmdutils.OutputError(httpResponse, errorResponse)
		} else {
			cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
		}
	} else if path != "" {
		apiResponse, httpResponse, errorResponse := client.GetItemByPath(path)
		if errorResponse != nil {
			cmdutils.OutputError(httpResponse, errorResponse)
		} else {
			cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
		}
	} else if contentId != "" && path != "" {
		fmt.Println("Please specify only contentId or path.")
	}
}
