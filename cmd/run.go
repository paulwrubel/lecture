/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/paulwrubel/lecture-lang/internal"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run builds and runs lecture",
	Long: `Run lecture files (.ltr) or input directly.
	
This will run a single lecture file.`,
	Args: cobra.ExactArgs(1),
	Run:  runRun,
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func runRun(cmd *cobra.Command, args []string) {
	// read input
	lectureBytes, err := readLectureBytes(args[0])
	if err != nil {
		log.Fatalf("error reading lecture bytes: %s\n", err.Error())
	}

	// transpile to golang
	golangBytes, err := internal.TranspileToGolang(lectureBytes)
	if err != nil {
		log.Fatalf("error transpiling to Golang: %s\n", err.Error())
	}

	// create temporary output file for golang code
	tempFile, err := os.CreateTemp("", "compiled_lecture_*.go")
	if err != nil {
		log.Fatalf("error creating temporary file: %s\n", err.Error())
	}
	defer func(f *os.File) {
		_ = f.Close()
		err := os.Remove(f.Name())
		if err != nil {
			log.Fatalf("error removing temporary file: %s\n", err.Error())
		}
	}(tempFile)

	_, err = tempFile.Write(golangBytes)
	if err != nil {
		log.Fatalf("error writing Golang bytes to temporary file: %s\n", err.Error())
	}

	// prepare the `go run` command to run the file we just made
	goRunCommand := exec.Command("go", "run", tempFile.Name())
	goRunCommand.Env = os.Environ()

	//
	stdout, err := goRunCommand.StdoutPipe()
	if err != nil {
		log.Fatalf("error getting stdout of lecture subprocess: %s\n", err.Error())
	}

	stderr, err := goRunCommand.StderrPipe()
	if err != nil {
		log.Fatalf("error getting stderr of lecture subprocess: %s\n", err.Error())
	}

	err = goRunCommand.Start()
	if err != nil {
		log.Fatalf("error starting lecture subprocess: %s\n", err.Error())
	}

	group := &errgroup.Group{}
	scanFunc := func(r io.Reader) error {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			outputText := scanner.Text()
			fmt.Println(outputText)
		}
		return scanner.Err()
	}
	group.TryGo(func() error {
		err := scanFunc(stdout)
		if err != nil {
			return fmt.Errorf("error scanning stdout: %s", err.Error())
		}
		return nil
	})
	group.TryGo(func() error {
		err := scanFunc(stderr)
		if err != nil {
			return fmt.Errorf("error scanning stderr: %s", err.Error())
		}
		return nil
	})

	err = group.Wait()
	if err != nil {
		log.Fatalf("error scanning lecture subprocess output: %s\n", err.Error())
	}

	err = goRunCommand.Wait()
	if err != nil {
		log.Fatalf("error waiting for lecture subprocess: %s\n", err.Error())
	}
}
