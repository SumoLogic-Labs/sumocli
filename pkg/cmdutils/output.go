package cmdutils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Output(apiResponse interface{}, httpResponse *http.Response, errorResponse error) {
	apiResponseJsonFormatted, err := json.MarshalIndent(apiResponse, "", "    ")
	if err != nil {
		fmt.Println("failed to marshal api response")
	}
	if httpResponse.StatusCode != 200 {
		fmt.Println(errorResponse)
	} else {
		fmt.Println(string(apiResponseJsonFormatted))
	}
}
