package main

import (
	"fmt"
	"os"
)

func main() {
	lectureWalker, err := NewLectureWalker("okay, hear me out. let's say one is 1. then we have one. i rest my case.")
	if err != nil {
		panic(err)
	}
	outputFolder := "output"
	outputFilename := "lecture.go"
	_, err = os.Stat(outputFolder)
	if os.IsNotExist(err) {
		os.MkdirAll(outputFolder, 0755)
	} else if err != nil {
		panic(err)
	}
	outputFile, err := os.Create(fmt.Sprintf("%s/%s", outputFolder, outputFilename))
	if err != nil {
		panic(err)
	}
	golangListener := &GolangLectureListener{
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
