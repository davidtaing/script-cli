/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package update

import (
	"fmt"

	"github.com/spf13/cobra"
)

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
	},
}
