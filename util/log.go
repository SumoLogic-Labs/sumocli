package util

import "log"

/*
type ErrorResponse struct {
	id string
	errors []
}
*/

// LogError: Logs any errors received
func LogError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func HttpError(statusCode int, responseBody string) bool {
	if statusCode == 401 {
		log.Println("Unauthorized access please check the user exists,  are valid.")
		apiCallSuccess := false
		return apiCallSuccess
	} else if statusCode == 200 {
		apiCallSuccess := true
		return apiCallSuccess
	}
	apiCallSuccess := true
	return apiCallSuccess
}
