/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ccbr",
	Short: "Use Rego Implement CIS-benchmark check",
	Long: `Use Rego Implement Kubernetes CIS-benchmark check`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


