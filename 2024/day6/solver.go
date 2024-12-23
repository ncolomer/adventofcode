package day6

import (
	"bufio"
	"strings"
)

func Read(input string) (m Map, err error) {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	var matrix [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []rune(line))
	}
	m = NewMap(matrix)
	return
}

func SolvePart1(m Map) (res int, err error) {
	m.ComputeGuardRoute()
	res = m.TotalVisited()
	return
}

func SolvePart2(m Map) (res int, err error) {
	original := m.Copy()
	m.ComputeGuardRoute()
	for y, row := range m.matrix {
		for x, cell := range row {
			coordinates := Coordinates{x, y}
			if cell == Visited && coordinates != original.guard {
				s := original.Copy()
				s.Set(coordinates, WallAlt)
				if s.ComputeGuardRoute() == StatusLooping {
					res += 1
				}
			}
		}
	}
	return
}
