package main

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmd/root"
)

func main() {
	rootCmd := root.NewCmdRoot()
	rootCmd.Execute()
}
