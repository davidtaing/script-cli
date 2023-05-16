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

var script string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := createNewScript("helloworld")
		openScriptInEditor(path, Editor)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&script, "script", "s", "", "script file name")
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
		editor = "gedit" // Default editor
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
			fmt.Println("Invalid editor. Using default editor (gedit).")
			editor = "gedit" // Fall back to default editor
		}
	}

	cmd := exec.Command(editor, path)
	err := cmd.Run()

	if err != nil {
		fmt.Println("Error opening script in editor:", err)
	}
}
