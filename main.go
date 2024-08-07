package main

import (
	"github.com/sirusxcoder/studybuddy/cmd"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "studybuddy"}

func main() {
	cmd.Execute()
}

func init() {
	rootCmd.AddCommand(cmd.serveCmd)
}
