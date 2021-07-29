package cmdutils

import (
	"fmt"
	"net/http"
)

func OutputError(httpResponse *http.Response) {
	if httpResponse.StatusCode == 400 {
		fmt.Println("Bad request, this could mean the resource already exists or the request is incorrect.")
	} else if httpResponse.StatusCode == 401 {
		fmt.Println("Unauthorized access please check the user exists, and credentials are valid.")
	} else if httpResponse.StatusCode == 403 {
		fmt.Println("This operation is not allowed for your account type or the user doesn't have the role capability to perform this action.")
	} else if httpResponse.StatusCode == 404 {
		fmt.Println("Requested resource could not be found.")
	} else if httpResponse.StatusCode == 405 {
		fmt.Println("Unsupported method for URL.")
	} else if httpResponse.StatusCode == 415 {
		fmt.Println("Invalid content type.")
	} else if httpResponse.StatusCode == 429 {
		fmt.Println("The API request rate is higher than 4 request per second or inflight API requests are higher than 10 requests per second.")
	} else if httpResponse.StatusCode == 500 {
		fmt.Println("Internal server error.")
	} else if httpResponse.StatusCode == 503 {
		fmt.Println("Service is currently unavailable.")
	}
}
