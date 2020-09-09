package util

import "log"

/*
type ErrorResponse struct {
	id string
	errors [] // TODO: This will be an array of something objects probably type ErrorObjects struct {}
}
*/

// TODO: Implement a better logging solution than outputing to console
// LogError: Logs any errors received
func LogError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// TODO: Refactor the HttpError function
func HttpError(statusCode int, responseBody string) bool {
	if statusCode == 401 {
		log.Fatalln("Unauthorized access please check the user exists,  are valid.")
		apiCallSuccess := false
		return apiCallSuccess
	} else if statusCode == 400 {

	}
	apiCallSuccess := true
	return apiCallSuccess
}
