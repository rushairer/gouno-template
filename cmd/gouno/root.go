package gouno

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gouno",
	Short: "gouno is a tool to generate go code",
	Long: `gouno is a tool to generate go code.
It can generate go code from proto file.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing root command: %v", err)
		os.Exit(1)
	}
}
