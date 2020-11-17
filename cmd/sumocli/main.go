package main

import (
	"github.com/wizedkyle/sumocli/pkg/cmd/root"
)

func main() {
	rootCmd := root.NewCmdRoot()
	rootCmd.Execute()
}