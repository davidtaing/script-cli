/*
Copyright © 2023 NAME HERE adavidtaing@gmail.com

*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"

	directory "github.com/davidtaing/scriptcli/internal/dir"
	"github.com/davidtaing/scriptcli/internal/promptutil"
)

var root = "bin"

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Run: func(*cobra.Command, []string) {
		fp, err := directory.GetFilePaths(root)

		if err != nil {
			log.Panicln("Error looking up filepaths in root directory:", err)
		}

		s, err := promptutil.PromptSelectItems(fp, "Select script to run")

		if err != nil {
			fmt.Println("An invalid script was selected, exiting Run command")
		}

		err = runScript(s)

		if err != nil {
			fmt.Println("Failed to run script: " + s)
		}
	},
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
