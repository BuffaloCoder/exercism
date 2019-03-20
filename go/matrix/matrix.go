package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix defines the structure of our matrix (rows holding the data in rows
// and cols being the number of columns)
type Matrix struct {
	rows [][]int
	cols int
}

// New creates a new Matrix from a given string
func New(data string) (*Matrix, error) {
	rows := strings.Split(data, "\n")
	var colCount int
	rowData := make([][]int, len(rows))
	for rIndex, row := range rows {
		cols := strings.Split(strings.Trim(row, " "), " ")
		if colCount == 0 {
			colCount = len(cols)
		} else if len(cols) != colCount {
			return nil, errors.New("all rows should be the same length")
		}
		colData := make([]int, colCount)
		for cIndex, col := range cols {
			if col == "" {
				continue
			}
			val, err := strconv.Atoi(col)
			if err != nil {
				return nil, err
			}
			colData[cIndex] = val
		}
		rowData[rIndex] = colData
	}

	m := Matrix{rows: rowData, cols: colCount}
	return &m, nil
}

// Rows returns a copy of the rows of the matrix
func (m *Matrix) Rows() [][]int {
	rCopy := make([][]int, len(m.rows))
	for rIndex, row := range m.rows {
		rCopy[rIndex] = make([]int, len(row))
		copy(rCopy[rIndex], row)
	}
	return rCopy
}

// Cols returns a copy of the columns of the matrix
func (m *Matrix) Cols() [][]int {
	cCopy := make([][]int, m.cols)
	for _, row := range m.rows {
		for cIndex, col := range row {
			cCopy[cIndex] = append(cCopy[cIndex], col)
		}
	}
	return cCopy
}

// Set updates the internal matrix
func (m *Matrix) Set(row, col, val int) bool {
	if row >= len(m.rows) || row < 0 {
		return false
	}
	if col >= len(m.rows[row]) || col < 0 {
		return false
	}
	m.rows[row][col] = val
	return true
}
