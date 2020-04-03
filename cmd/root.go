package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

// Execute command.
func Execute() {
	rootCmd.AddCommand(cmdServer)
	rootCmd.AddCommand(cmdAdmin)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
