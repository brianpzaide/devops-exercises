package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {

	type input struct {
		ForestSize int    `json:"forestSize"`
		Forest     string `json:"forest"`
		Start      []int  `json:"start"`
		End        []int  `json:"end"`
	}
	testData := input{
		ForestSize: 5,
		Forest:     "....\nB...\n..BB\n....",
		Start:      []int{0, 0},
		End:        []int{3, 3},
	}

	gridMatrix, err := constructMatrix(testData.ForestSize, testData.Forest)
	if err != nil {
		t.Fatal(err.Error())
	}

	path, err := solveMaze(gridMatrix, testData.Start[0], testData.Start[1], testData.End[0], testData.End[1])
	if err != nil {
		t.Fatal(err.Error())
	}

	if !strings.Contains(path, fmt.Sprintf("(%d, %d)", testData.Start[0], testData.Start[1])) {
		fmt.Println("no path")
	}

	fmt.Println(path)

}
