package day10

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
	for _, paths := range m.GetTrails() {
		uniq := make(map[Coordinates]struct{})
		for _, path := range paths {
			uniq[path.Last()] = struct{}{}
		}
		score := len(uniq)
		res += score
	}
	return
}

func SolvePart2(m Map) (res int, err error) {
	for _, paths := range m.GetTrails() {
		rating := len(paths)
		res += rating
	}
	return
}
