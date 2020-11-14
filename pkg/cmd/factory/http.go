package factory

import (
	"fmt"
)

func NewHttpClient() {

}

func HttpError(statusCode int) bool {
	if statusCode == 401 {
		fmt.Println("Unauthorized access please check the user exists,  are valid.")
		apiCallSuccess := false
		return apiCallSuccess
	} else if statusCode == 403 {
		fmt.Println("This operation is not allowed for your account type or the user doesn't have the role capability to perform this action.")
		apiCallSuccess := false
		return apiCallSuccess
	} else if statusCode == 404 {
		fmt.Println("Requested resource could not be found.")
		apiCallSuccess := false
		return apiCallSuccess
	} else if statusCode == 405 {
		fmt.Println("Unsupported method for URL.")
		apiCallSuccess := false
		return apiCallSuccess
	} else if statusCode == 415 {
		fmt.Println("Invalid content type.")
		apiCallSuccess := false
		return apiCallSuccess
	} else if statusCode == 429 {
		fmt.Println("The API request rate is higher than 4 request per second or inflight API requests are higher than 10 requests per second.")
		apiCallSuccess := false
		return apiCallSuccess
	} else if statusCode == 500 {
		fmt.Println("Internal server error.")
		apiCallSuccess := false
		return apiCallSuccess
	} else if statusCode == 503 {
		fmt.Println("Service is currently unavailable.")
		apiCallSuccess := false
		return apiCallSuccess
	} else if statusCode == 200 {
		apiCallSuccess := true
		return apiCallSuccess
	}
	apiCallSuccess := true
	return apiCallSuccess
}
