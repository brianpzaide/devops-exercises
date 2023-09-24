package main

import (
	"bears_in_the_forest/internal/algorithms"
	"bears_in_the_forest/internal/ds"
	"errors"
	"fmt"
	"strings"
)

const (
	BEAR         = -1
	NO_BEAR      = 0
	PATH_ELEMENT = 1
)

type PathElement struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

func solveMaze(grid [][]int, row_start, col_start, row_end, col_end int) (string, error) {
	if grid[row_start][col_start] == BEAR || grid[row_end][col_end] == BEAR{
		return "no path", nil
	}
	Size := len(grid)

	start := Size*row_start + col_start
	end := Size*row_end + col_end

	var i, j int

	mst := constructMSTfromMatrix(grid)

	dsp := algorithms.ShortestPath(mst, mst.GetVertex(start))
	if _, ok := dsp[end]; !ok {
		return "no path", nil
	}

	path := make([]PathElement, 0)
	path = append(path, PathElement{Row: row_end, Col: col_end})

	k, ok := dsp[end]

	for k != -1 && ok {
		i = mst.GetVertex(k).Data.Row
		j = mst.GetVertex(k).Data.Col
		path = append(path, PathElement{Row: i, Col: j})
		k, ok = dsp[k]

	}

	return getPath(path), nil
}

func getPath(path []PathElement) string {
	pathElements := ""
	for i := len(path) - 1; i >= 0; i-- {
		pathElements += fmt.Sprintf("(%d, %d) ", path[i].Row, path[i].Col)
	}

	return pathElements[:len(pathElements)-1]
}

func constructMatrix(mazeSize int, mazeString string) ([][]int, error) {

	grid := make([][]int, 0)
	rows := strings.Split(mazeString, "\n")
	fmt.Println(len(rows))
	if mazeSize != len(rows) {
		return nil, errors.New("ill formed data, ForestSize and length of ForestString are not compatible")
	}
        for i := 0; i<mazeSize;i++{
		if mazeSize != len(rows[i]) {
			return nil, errors.New("ill formed data, ForestString must be a square grid")
		}
	}

	for _, row := range rows {

		sl := make([]int, 0)

		for j := 0; j < mazeSize; j++ {
			if string(row[j]) == "." {
				sl = append(sl, NO_BEAR)
			} else if string(row[j]) == "B" {
				sl = append(sl, BEAR)
			} else {
				return nil, errors.New("forest string should only contain dots or letter B")
			}
		}
		grid = append(grid, sl)
	}
	return grid, nil
}

func constructMSTfromMatrix(grid [][]int) *(ds.Graph) {

	Size := len(grid)

	mstGraph := &(ds.Graph{
		AllEdges:    make([]*(ds.Edge), 0),
		AllVertices: make(map[int]*(ds.Vertex)),
	})

	for i := 0; i < Size; i++ {

		for j := 0; j < Size; j++ {

			if j < Size-1 {

				if grid[i][j] == 0 && grid[i][j+1] == 0 {

					mstGraph.AddEdge(Size*i+j, Size*i+(j+1), 1)
				}
			}

			if i < Size-1 {

				if grid[i][j] == 0 && grid[i+1][j] == 0 {

					mstGraph.AddEdge(Size*i+j, Size*(i+1)+j, 1)
				}
			}
		}
	}

	// fmt.Println(mstGraph.String())

	for i := 0; i < Size; i++ {

		for j := 0; j < Size; j++ {

			if v := mstGraph.GetVertex(Size*i + j); v != nil {
				v.Data = ds.Coordinates{Row: i, Col: j}
			}
		}
	}

	return mstGraph
}
