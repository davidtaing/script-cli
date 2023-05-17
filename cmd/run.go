/*
Copyright © 2023 NAME HERE adavidtaing@gmail.com

*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type commandHandler func(*cobra.Command, []string)

var root = "bin"

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Run: func(*cobra.Command, []string) {
		fp, err := getFilePaths(root)

		if err != nil {
			log.Panicln("Error looking up filepaths in root directory:", err)
		}

		s, err := promptSelectScript(fp)

		if err != nil {
			fmt.Println("An invalid script was selected, exiting Run command")
		}

		err = runScript(s)

		if err != nil {
			fmt.Println("Failed to run script: " + s)
		}
	},
}

func promptSelectScript(scripts []string) (string, error) {
	p := promptui.Select{
		Label: "Select script to run",
		Items: scripts,
	}

	i, _, err := p.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	return scripts[i], nil
}

func getFilePaths(root string) ([]string, error) {
	var filepaths []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			filepaths = append(filepaths, path)
			if err != nil {
				log.Println("Error opening file:", err)
			}
		}

		return nil
	})

	if err != nil {
		log.Println("Error:", err)
	}

	return filepaths, err
}

func runScript(filepath string) error {
	fmt.Printf("Excecuting script: %v\n\n", filepath)

	cmd := exec.Command("./" + filepath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(runCmd)
}
