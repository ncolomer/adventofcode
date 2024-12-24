package day8

import (
	"bufio"
	"slices"
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
	var combine func(antenna Coordinate, others []Coordinate) []Coordinate
	combine = func(antenna Coordinate, others []Coordinate) (antinodes []Coordinate) {
		if len(others) == 0 {
			return nil
		}
		for i, other := range others {
			if antenna == other {
				continue
			}
			xd := other.x - antenna.x
			yd := other.y - antenna.y
			antinodes = append(antinodes,
				antenna.Move(-xd, -yd),
				antenna.Move(xd*2, yd*2),
			)
			concat := slices.Concat(others[:i], others[i+1:])
			antinodes = append(antinodes, combine(other, concat)...)
		}
		return
	}
	antinodes := make(map[Coordinate]struct{})
	for _, coordinates := range m.antennas {
		if len(coordinates) < 2 {
			// no antinode possible
			continue
		}
		for _, candidate := range combine(coordinates[0], coordinates[1:]) {
			if m.IsValid(candidate) {
				antinodes[candidate] = struct{}{}
			}
		}
	}
	//m.Print(slices.Collect(maps.Keys(antinodes)))
	return len(antinodes), nil
}

func SolvePart2(m Map) (res int, err error) {
	var combine func(antenna Coordinate, others []Coordinate) []Coordinate
	combine = func(antenna Coordinate, others []Coordinate) (antinodes []Coordinate) {
		if len(others) == 0 {
			return nil
		}
		for i, other := range others {
			if antenna == other {
				continue
			}
			xd := other.x - antenna.x
			yd := other.y - antenna.y
			for c := antenna; m.IsValid(c); c = c.Move(-xd, -yd) {
				antinodes = append(antinodes, c)
			}
			for c := antenna.Move(xd, yd); m.IsValid(c); c = c.Move(xd, yd) {
				antinodes = append(antinodes, c)
			}
			concat := slices.Concat(others[:i], others[i+1:])
			antinodes = append(antinodes, combine(other, concat)...)
		}
		return
	}
	antinodes := make(map[Coordinate]struct{})
	for _, coordinates := range m.antennas {
		if len(coordinates) < 2 {
			// no antinode possible
			continue
		}
		for _, candidate := range combine(coordinates[0], coordinates[1:]) {
			if m.IsValid(candidate) {
				antinodes[candidate] = struct{}{}
			}
		}
	}
	//m.Print(slices.Collect(maps.Keys(antinodes)))
	return len(antinodes), nil
}
