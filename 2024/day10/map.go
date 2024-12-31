package day10

import (
	"fmt"
	"slices"
	"strconv"
)

type Coordinates struct {
	x int
	y int
}

func (c Coordinates) AddX(dx int) Coordinates {
	return Coordinates{c.x + dx, c.y}
}
func (c Coordinates) AddY(dy int) Coordinates {
	return Coordinates{c.x, c.y + dy}
}

func (c Coordinates) String() string {
	return fmt.Sprintf("(%d, %d)", c.x, c.y)
}

type Path struct {
	steps []Coordinates
}

func (p *Path) First() Coordinates {
	return p.steps[0]
}

func (p *Path) Last() Coordinates {
	return p.steps[len(p.steps)-1]
}

func (p *Path) Put(c Coordinates) Path {
	var steps []Coordinates
	for _, step := range p.steps {
		steps = append(steps, step)
	}
	return Path{append(steps, c)}
}

func (p *Path) Contains(c Coordinates) bool {
	return slices.Contains(p.steps, c)
}

type Map struct {
	matrix     [][]int
	trailheads []Coordinates
}

func NewMap(input [][]rune) Map {
	var matrix [][]int
	var trailheads []Coordinates
	for y, row := range input {
		var acc []int
		for x, cell := range row {
			digit, err := strconv.Atoi(string(cell))
			if err == nil {
				acc = append(acc, digit)
				if digit == 0 {
					c := Coordinates{x, y}
					trailheads = append(trailheads, c)
				}
			} else {
				acc = append(acc, -1)
			}
		}
		matrix = append(matrix, acc)
	}
	return Map{matrix, trailheads}
}

func (m *Map) Height() int {
	return len(m.matrix)
}

func (m *Map) Width() int {
	return len(m.matrix[0])
}

func (m *Map) IsWithinMap(c Coordinates) bool {
	if c.x >= 0 && c.x < m.Width() && c.y >= 0 && c.y < m.Height() {
		return true
	}
	return false
}

func (m *Map) GetHeight(c Coordinates) int {
	return m.matrix[c.y][c.x]
}

func (m *Map) Next(pos Coordinates) (res []Coordinates) {
	return slices.DeleteFunc([]Coordinates{
		pos.AddX(1),
		pos.AddX(-1),
		pos.AddY(1),
		pos.AddY(-1),
	}, func(x Coordinates) bool {
		return !m.IsWithinMap(x)
	})
}

func (m *Map) GetTrails() map[Coordinates][]Path {
	var walk func(Path) []Path
	walk = func(p Path) (sp []Path) {
		last := p.Last()
		height := m.GetHeight(last)
		if height == 9 {
			sp = append(sp, p)
			return
		}
		for _, nc := range m.Next(last) {
			if m.GetHeight(nc) == height+1 {
				paths := walk(p.Put(nc))
				sp = append(sp, paths...)
			}
		}
		return
	}
	res := make(map[Coordinates][]Path)
	for _, trailhead := range m.trailheads {
		res[trailhead] = walk(Path{[]Coordinates{trailhead}})
	}
	return res
}

func (m *Map) Print() {
	fmt.Println()
	for _, row := range m.matrix {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}
	return
}

func (m *Map) PrintWithPath(p Path) {
	fmt.Println()
	for y, row := range m.matrix {
		for x, cell := range row {
			c := Coordinates{x, y}
			if p.Contains(c) {
				fmt.Print(cell)
			} else {
				fmt.Print("░")
			}
		}
		fmt.Println()
	}
	return
}
