package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lecture",
	Short: "Build and run lecture source files",
	Long: `lecture is used to transpile code from lecture files (.ltr)
	
Currently, it is designed to transpile into Golang, so the go binary is
required to correctly use this tool`,
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

}

func readLectureBytes(filename string) ([]byte, error) {
	var reader io.Reader
	if filename == "-" {
		reader = os.Stdin
	} else {
		file, err := os.Open(filename)
		if err != nil {
			return nil, fmt.Errorf("error opening lecture file: %s", err.Error())
		}
		reader = file
	}

	lectureBytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("error reading lecture: %s", err.Error())
	}

	return lectureBytes, nil
}
