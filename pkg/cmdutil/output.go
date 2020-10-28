package cmdutil

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func OutputToFile(data []byte) {

	now := time.Now().Format("2006-01-02-15-04-05")
	path, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	filePath := path + "/" + now + "-output.json"
	log.Println("creating output file at: " + filePath)
	_, err = os.Create(filePath)
	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println()
}
