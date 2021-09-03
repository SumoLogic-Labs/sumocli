package main

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmd/root"
)

func main() {
	rootCmd := root.NewCmdRoot()
	rootCmd.Execute()
}
