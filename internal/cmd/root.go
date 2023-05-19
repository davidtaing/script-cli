/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/davidtaing/scriptcli/internal/cmd/add"
	"github.com/davidtaing/scriptcli/internal/cmd/remove"
	"github.com/davidtaing/scriptcli/internal/cmd/run"
	"github.com/davidtaing/scriptcli/internal/cmd/update"
)

var Editor string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "script-cli",
	Short: "Automate Bash Script Execution with Ease",
	Long:  `Automate Bash Script Execution with Ease.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&Editor, "editor", "e", "", "The editor to use when creating a new script.")

	rootCmd.AddCommand(add.AddCmd)
	rootCmd.AddCommand(remove.RemoveCmd)
	rootCmd.AddCommand(run.RunCmd)
	rootCmd.AddCommand(update.UpdateCmd)
}
