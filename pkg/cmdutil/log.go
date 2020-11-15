package cmdutil

import "log"

// LogError: Logs any errors received
func LogError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
