package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/paulwrubel/lecture-lang/internal"
	"github.com/spf13/cobra"
)

var (
	outputFilename string
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build FILENAME",
	Short: "Build golang scripts from a lecture file",
	Long: `Generate golang source files from a lecture file (.ltr)
	
This will generate a single *.go file (default "lecture.go") that can be
directly run using "go run [FILENAME]"`,
	Args: cobra.ExactArgs(1),
	Run:  runBuild,
}

func init() {
	rootCmd.AddCommand(buildCmd)

	buildCmd.Flags().StringVarP(&outputFilename, "output", "o", "-", "filename to write the generated Golang to. '-' writes to stdout.")
}

func runBuild(cmd *cobra.Command, args []string) {
	lectureBytes, err := readLectureBytes(args[0])
	if err != nil {
		log.Fatalf("error reading lecture bytes: %s\n", err.Error())
	}

	golangBytes, err := internal.TranspileToGolang(lectureBytes)
	if err != nil {
		log.Fatalf("error transpiling to Golang: %s\n", err.Error())
	}

	err = writeOutputBytes(outputFilename, golangBytes)
	if err != nil {
		log.Fatalf("error writing Golang bytes: %s\n", err.Error())
	}
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

func writeOutputBytes(filename string, bytes []byte) error {
	var writer io.Writer
	if filename == "-" {
		writer = os.Stdout
	} else {
		// create folder path if necessary
		lastSlashIndex := strings.LastIndex(filename, "/")
		if lastSlashIndex != -1 {
			folder := filename[:lastSlashIndex]
			_, err := os.Stat(folder)
			if os.IsNotExist(err) {
				os.MkdirAll(folder, 0755)
			} else if err != nil {
				log.Fatalf("error getting output folder info: %s\n", err.Error())
			}
		}
		file, err := os.Create(filename)
		if err != nil {
			return fmt.Errorf("error creating output file: %s", err.Error())
		}
		writer = file
	}

	_, err := writer.Write(bytes)
	if err != nil {
		return fmt.Errorf("error writing output: %s", err.Error())
	}
	return nil
}
