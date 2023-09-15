package main

import (
	"fmt"
	"os"
)

func main() {
	lectureText := `
		okay, hear me out. 
			let's say one is literally 1. 
			let's say two is literally 2. 
			let's say three is one plus two.
			let's say four is two plus two.
			let's say fifteen is one plus two plus three plus four plus literally 5.
			then we have one.
			then we have two.
			then we have three.
			then we have four.
			then we have fifteen.
			then we have literally 50 plus literally 5 plus four plus literally 2 plus two.
		i rest my case.
	`
	lectureWalker, err := NewLectureWalker(lectureText)
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
