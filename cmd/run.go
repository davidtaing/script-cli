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

var root = "bin"

// runCmd represents the task command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

var runScriptCmd = &cobra.Command{
	Use:   "script",
	Short: "Execute a bash script.",
	Long:  "Execute a bash script.",
}

func generateDynamicCommand(name string) *cobra.Command {
	var d = fmt.Sprintf("Execute %s.", name)

	var cmd = &cobra.Command{
		Use:   name,
		Short: d,
		Long:  d,
		Run:   runDynamicTask(name),
	}

	return cmd
}

func runTask(cmd *cobra.Command, args []string) {
	p := "bin/test.sh"

	err := runScript(p)

	if err != nil {
		log.Println("Error running script:", err)
	}
}

func runDynamicTask(p string) func(cmd *cobra.Command, args []string) {
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
	filepaths, err := getFilePaths(root)

	if err != nil {
		log.Panicln("Error looking up filepaths in root directory:", err)
	}

	for _, p := range filepaths {
		runScriptCmd.AddCommand(generateDynamicCommand(p))
	}

	runCmd.AddCommand(runScriptCmd)
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// taskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// taskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
