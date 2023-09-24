package main

import (
	"fmt"
	"net/http"
	"strings"
)

func (app *application) solve(w http.ResponseWriter, r *http.Request) {

	var input struct {
		ForestSize int    `json:"forestSize"`
		Forest     string `json:"forest"`
		Start      []int  `json:"start"`
		End        []int  `json:"end"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {

		app.clientError(w, err.Error(), http.StatusBadRequest)
		return
	}

	gridMatrix, err := constructMatrix(input.ForestSize, input.Forest)
	if err != nil {
		app.clientError(w, err.Error(), http.StatusBadRequest)
		return
	}

	path, err := solveMaze(gridMatrix, input.Start[0], input.Start[1], input.End[0], input.End[1])
	if err != nil {
		app.serverError(w, err)
		return
	}
	if !strings.Contains(path, fmt.Sprintf("(%d, %d)", input.Start[0], input.Start[1])) {
		path = "no path"
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"path": path}, nil)
	if err != nil {
		app.serverError(w, err)
		return
	}
}
