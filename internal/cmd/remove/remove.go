/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package remove

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove called")
	},
}
