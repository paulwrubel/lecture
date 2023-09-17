package internal

import (
	"fmt"
	"io"
)

func TranspileToGolang(lectureBytes []byte) ([]byte, error) {
	walker, err := NewLectureWalker(string(lectureBytes))
	if err != nil {
		return nil, fmt.Errorf("error creating lecture walker: %s", err.Error())
	}
	golangListener := &GolangLectureListener{}
	walker.Walk(golangListener)
	if len(golangListener.Errors) > 0 {
		return nil, golangListener.Errors[0]
	}
	golangBytes, err := io.ReadAll(golangListener.OutputBytes)
	if err != nil {
		return nil, fmt.Errorf("error reading transpiling bytes: %s", err.Error())
	}
	return golangBytes, nil
}
