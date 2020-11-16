package unlock

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	util2 "github.com/wizedkyle/sumocli/pkg/cmdutil"
	"io/ioutil"
	"os"
)

func NewCmdUnlockUser() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "unlock",
		Short: "Unlocks a Sumo Logic user account",
		Run: func(cmd *cobra.Command, args []string) {
			unlockUser(id)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the user account to unlock")

	return cmd
}

func unlockUser(id string) {
	if id == "" {
		fmt.Println("--id field needs to be specified.")
		os.Exit(0)
	}

	requestUrl := "v1/users/" + id + "/unlock"
	client, request := factory.NewHttpRequest("POST", requestUrl)
	response, err := client.Do(request)
	util2.LogError(err)

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	util2.LogError(err)

	if response.StatusCode != 204 {
		var responseError api.ResponseError
		jsonErr := json.Unmarshal(responseBody, &responseError)
		util2.LogError(jsonErr)
	} else {
		fmt.Println("User account was unlocked.")
	}
}
