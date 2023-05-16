/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := createNewScript("helloworld")
		openScriptInEditor(path, "code")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createNewScript(name string) (string, error) {
	var path = fmt.Sprintf("bin/%s.sh", name)

	err := os.WriteFile(path, []byte("#!/bin/bash\n\n"), 0755)

	if err != nil {
		fmt.Println("Error creating new script:", err)
		return "", err
	}

	fmt.Println("Created new script at", path)

	return path, nil
}

func openScriptInEditor(path string, editor string) {
	validEditors := []string{"code", "emacs", "gedit", "nano", "vi", "vim"} // List of allowed editors

	if editor == "" {
		editor = "vi" // Default editor
	} else {
		editor = strings.ToLower(editor)
		found := false
		for _, validEditor := range validEditors {
			if validEditor == editor {
				found = true
				break
			}
		}

		if !found {
			fmt.Println("Invalid editor. Using default editor (vi).")
			editor = "vi" // Fall back to default editor
		}
	}

	cmd := exec.Command(editor, path)
	err := cmd.Run()

	if err != nil {
		fmt.Println("Error opening script in editor:", err)
	}
}
