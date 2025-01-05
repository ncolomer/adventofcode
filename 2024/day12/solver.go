package day12

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/ncolomer/adventofcode/2024/util"
)

func Read(input string) util.Map[rune] {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	var runes [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		runes = append(runes, []rune(line))
	}
	return util.NewMap[rune](runes)
}

type Region struct {
	plant     string
	plots     []util.Coordinate
	area      int
	perimeter int
	sides     int
}

func (r Region) String() string {
	return fmt.Sprintf("[%v] area:%2d perimeter:%2d sides:%2d", r.plant, r.area, r.perimeter, r.sides)
}

func CountCorners(c util.Coordinate, m util.Map[rune]) (corner int) {
	corners := [][]util.Coordinate{
		// up left corner
		{c.MoveX(-1).MoveY(0), c.MoveX(-1).MoveY(-1), c.MoveX(0).MoveY(-1)},
		// up right corner
		{c.MoveX(0).MoveY(-1), c.MoveX(1).MoveY(-1), c.MoveX(1).MoveY(0)},
		// down right corner
		{c.MoveX(1).MoveY(0), c.MoveX(1).MoveY(1), c.MoveX(0).MoveY(1)},
		// down left corner
		{c.MoveX(0).MoveY(1), c.MoveX(-1).MoveY(1), c.MoveX(-1).MoveY(0)},
	}
	for _, pos := range corners {
		if (!m.IsValid(pos[0]) || m.At(pos[0]) != m.At(c)) &&
			(!m.IsValid(pos[2]) || m.At(pos[2]) != m.At(c)) {
			corner += 1
		}
		if (!m.IsValid(pos[1]) || m.At(pos[1]) != m.At(c)) &&
			(m.IsValid(pos[0]) && m.At(pos[0]) == m.At(c)) &&
			(m.IsValid(pos[2]) && m.At(pos[2]) == m.At(c)) {
			corner += 1
		}
	}
	return
}

func FindRegions(m util.Map[rune]) []*Region {
	visited := make(map[util.Coordinate]struct{})

	var visit func(util.Coordinate) *Region
	visit = func(current util.Coordinate) *Region {
		var r Region
		visited[current] = struct{}{}
		r.plant = string(m.At(current))
		r.plots = append(r.plots, current)
		r.area += 1
		vBorder := current.X == 0 || current.X == m.Width()-1
		hBorder := current.Y == 0 || current.Y == m.Height()-1
		if vBorder {
			r.perimeter += 1
		}
		if hBorder {
			r.perimeter += 1
		}
		for _, next := range m.Sides(current) {
			if m.At(current) != m.At(next) {
				r.perimeter += 1
				continue
			}
			if _, ok := visited[next]; ok {
				continue
			}
			rr := visit(next)
			r.plots = append(r.plots, rr.plots...)
			r.area += rr.area
			r.perimeter += rr.perimeter
			r.sides += rr.sides
		}
		corners := CountCorners(current, m)
		r.sides += corners
		return &r
	}

	var regions []*Region
	for x := range m.Width() {
		for y := range m.Height() {
			c := util.Coordinate{x, y}
			if _, ok := visited[c]; !ok {
				region := visit(c)
				regions = append(regions, region)
			}
		}
	}
	return regions
}

func SolvePart1(m util.Map[rune]) (res int) {
	for _, region := range FindRegions(m) {
		res += region.area * region.perimeter
	}
	return
}

func SolvePart2(m util.Map[rune]) (res int) {
	for _, region := range FindRegions(m) {
		res += region.area * region.sides
	}
	return
}
