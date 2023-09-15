package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidLecture(t *testing.T) {
	tt := []struct {
		Name    string
		Lecture string
	}{
		{
			Name:    "bare key mode",
			Lecture: `"mykey"`,
		},
	}

	for _, test := range tt {
		t.Run(test.Name, func(t *testing.T) {
			_, err := NewLectureWalker(test.Lecture)
			assert.NoError(t, err, `lecture: %s`, test.Lecture)
		})
	}
}

func TestInvalidLecture(t *testing.T) {
	tt := []struct {
		Name    string
		Lecture string
	}{
		{
			Name:    "unquoted string not a keyword",
			Lecture: `mykey`,
		},
	}

	for _, test := range tt {
		t.Run(test.Name, func(t *testing.T) {
			_, err := NewLectureWalker(test.Lecture)
			assert.Error(t, err, `lecture: %s`, test.Lecture)
		})
	}
}
