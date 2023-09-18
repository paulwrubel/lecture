package cmd

import (
	"io"
	"log"
	"os"
	"strings"

	"github.com/paulwrubel/lecture/internal"
	"github.com/spf13/cobra"
)

var (
	outputFilename string
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build FILENAME",
	Short: "build will generate a golang script from a lecture file",
	Long: `Generate golang source files from a lecture file (.ltr)
	
This will generate a single *.go file (prints to stdout by default) that can be
directly run using "go run [FILENAME]"

If the first argument is '-', it will read from stdin for the lecture text.`,
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

	writer, err := getOutputWriter(outputFilename)
	if err != nil {
		log.Fatalf("error getting output writer: %s\n", err.Error())
	}
	defer writer.Close()

	golangBytes, err := internal.TranspileToGolang(lectureBytes)
	if err != nil {
		log.Fatalf("error transpiling to Golang: %s\n", err.Error())
	}

	_, err = writer.Write(golangBytes)
	if err != nil {
		log.Fatalf("error writing Golang bytes: %s\n", err.Error())
	}
}

func getOutputWriter(filename string) (io.WriteCloser, error) {
	if filename == "-" {
		return internal.NopWriteCloser{
			Writer: os.Stdout,
		}, nil
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
		return os.Create(filename)
	}
}
