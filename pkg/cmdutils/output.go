package cmdutils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Output(data interface{}, response *http.Response, err error, message string) {
	dataJsonFormatted, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println("failed to marshal api response")
	}
	if data == nil {
		if response.StatusCode != 204 && response.StatusCode != 200 {
			fmt.Println(err)
		} else {
			fmt.Println(message)
		}
	} else {
		if response.StatusCode != 200 && response.StatusCode != 201 {
			fmt.Println(err)
		} else {
			fmt.Println(string(dataJsonFormatted))
		}
	}
}
