package guard

import (
	"fmt"
	"os"
)

func MustNotBeEmpty(value string, msg string) {
	if value == "" {
		fmt.Println(msg)
		os.Exit(0)
	}
}