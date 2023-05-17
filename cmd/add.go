/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var (
	scriptName   = "helloworld"
	validEditors = []string{"code", "emacs", "gedit", "nano", "vi", "vim"}
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if scriptName == "" {
			scriptName, err = promptUserForScriptName()
		}

		if err != nil {
			fmt.Println(err)
			return
		}

		if Editor == "" {
			Editor = promptUserForEditor()
		}

		path, _ := createNewScript(scriptName)
		openScriptInEditor(path, Editor)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&scriptName, "script", "s", "", "script file name")
}

func promptUserForScriptName() (string, error) {
	var result string
	var err error

	prompt := promptui.Prompt{
		Label: "What would you like to name your new script?",
	}

	result, err = prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	if result == "" {
		return "", errors.New("New script name was not provided. Exiting")
	}

	return result, nil
}

func promptUserForEditor() string {
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label: "Which text editor would you like to use?",
			Items: validEditors,
		}

		index, result, err = prompt.Run()
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	fmt.Printf("You choose %s\n", result)
	return result
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

	fmt.Printf("Opening %s in %s\n", path, editor)

	cmd := exec.Command(editor, path)
	err := cmd.Run()

	if err != nil {
		fmt.Println("Error opening script in editor:", err)
	}
}
