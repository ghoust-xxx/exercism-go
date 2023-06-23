package matrix

import (
	"errors"
	"strconv"
	"strings"
)


type Matrix [][]int

func New(s string) (Matrix, error) {
	var m Matrix
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		cols := strings.Split(strings.Trim(line, " "), " ")
		var iline []int
		for _, val := range cols {
			num, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			iline = append(iline, num)
		}
		m = append(m, iline)
	}
	if m != nil {
		cols := len(m[0])
		for _, val := range m {
			if len(val) != cols {
				return nil, errors.New("Wrong size")
			}
		}
	}
	return m, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	if m == nil {
		return m
	}
	out := make([][]int, len(m[0]))
	for _, line := range m {
		for j, val := range line {
			out[j] = append(out[j], val)
		}
	}
	return out
}

func (m Matrix) Rows() [][]int {
	if m == nil {
		return m
	}
	out := make([][]int, len(m))
	for i, line := range m {
		for _, val := range line {
			out[i] = append(out[i], val)
		}
	}
	return out
}

func (m Matrix) Set(row, col, val int) bool {
	if m == nil {
		return false
	}
	rows := len(m)
	cols := len(m[0])
	if row >= rows || col >= cols || row < 0 || col < 0 {
		return false
	}
	m[row][col] = val
	return true
}
