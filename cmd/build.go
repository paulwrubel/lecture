package cmd

import (
	"fmt"
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

	buildCmd.Flags().StringVarP(&outputFilename, "output", "o", "lecture.go", "golang file to generate")
}

func runBuild(cmd *cobra.Command, args []string) {
	lectureBytes, err := os.ReadFile(args[0])
	if err != nil {
		log.Fatalf("error opening specified lecture source file: %s\n", err.Error())
	}

	lectureWalker, err := internal.NewLectureWalker(string(lectureBytes))
	if err != nil {
		log.Fatalf("error parsing lecture source file: %s\n", err.Error())
	}
	lastSlashIndex := strings.LastIndex(outputFilename, "/")
	if lastSlashIndex != -1 {
		outputFolder := outputFilename[:lastSlashIndex]
		_, err = os.Stat(outputFolder)
		if os.IsNotExist(err) {
			os.MkdirAll(outputFolder, 0755)
		} else if err != nil {
			log.Fatalf("error getting output folder info: %s\n", err.Error())
		}
	}
	outputFile, err := os.Create(outputFilename)
	if err != nil {
		log.Fatalf("error creating output file: %s\n", err.Error())
	}
	golangListener := &internal.GolangLectureListener{
		OutputFile: outputFile,
	}
	lectureWalker.Walk(golangListener)
	if len(golangListener.Errors) > 0 {
		for _, err := range golangListener.Errors {
			fmt.Printf("error in listener: %s\n", err.Error())
		}
		return
	}
}
