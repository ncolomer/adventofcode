package day8

import (
	"fmt"
	"slices"
)

type Coordinate struct {
	x int
	y int
}

func (c Coordinate) Move(xd int, yd int) Coordinate {
	return Coordinate{c.x + xd, c.y + yd}
}

type Map struct {
	matrix   [][]rune
	antennas map[rune][]Coordinate
}

func NewMap(matrix [][]rune) Map {
	antennas := make(map[rune][]Coordinate)
	for y, row := range matrix {
		for x, cell := range row {
			if cell != '.' {
				antennas[cell] = append(antennas[cell], Coordinate{x, y})
			}
		}
	}
	return Map{matrix, antennas}
}

func (m Map) Height() int {
	return len(m.matrix)
}

func (m Map) Width() int {
	return len(m.matrix[0])
}

func (m Map) IsValid(c Coordinate) bool {
	return c.x >= 0 && c.x < m.Width() && c.y >= 0 && c.y < m.Height()
}

func (m Map) Print(antinodes []Coordinate) {
	fmt.Println()
	for y, row := range m.matrix {
		for x, cell := range row {
			if slices.Contains(antinodes, Coordinate{x, y}) {
				fmt.Print("#")
			} else {
				fmt.Print(string(cell))
			}
		}
		fmt.Println()
	}
	return
}
