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

	"github.com/spf13/cobra"
)

type commandHandler func(*cobra.Command, []string)

var root = "bin"

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Run: func(*cobra.Command, []string) {
		fmt.Println("run called")

		filepaths, err := getFilePaths(root)

		if err != nil {
			log.Panicln("Error looking up filepaths in root directory:", err)
		}

		fmt.Println(filepaths)

		// prompt to select script
	},
}

func runDynamicTask(p string) commandHandler {
	return func(cmd *cobra.Command, args []string) {
		err := runScript(p)

		if err != nil {
			log.Println("Error running script:", err)
		}
	}
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
