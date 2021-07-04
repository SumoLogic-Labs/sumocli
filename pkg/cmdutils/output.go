package cmdutils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Output(apiResponse interface{}, httpResponse *http.Response, errorResponse error, message string) {
	apiResponseJsonFormatted, err := json.MarshalIndent(apiResponse, "", "    ")
	if err != nil {
		fmt.Println("failed to marshal api response")
	}
	if apiResponse == nil {
		if httpResponse.StatusCode != 204 {
			fmt.Println(errorResponse)
		} else {
			fmt.Println(message)
		}
	} else {
		if httpResponse.StatusCode != 200 {
			fmt.Println(errorResponse)
		} else {
			fmt.Println(string(apiResponseJsonFormatted))
		}
	}
}
