package day6

import (
	"fmt"
	"slices"
)

const (
	Wall    = '#'
	WallAlt = 'O'
	Visited = 'X'
)

type Direction int

const (
	DirectionUp Direction = iota
	DirectionDown
	DirectionLeft
	DirectionRight
)

func (d Direction) rune() rune {
	return [...]rune{'^', 'v', '<', '>'}[d]
}

func DirectionFrom(r rune) Direction {
	for i, d := range [...]rune{'^', 'v', '<', '>'} {
		if r == d {
			return Direction(i)
		}
	}
	return Direction(-1)
}

type Status int

const (
	StatusExitedMap Status = iota
	StatusLooping
)

type Coordinates struct {
	x int
	y int
}

func (c Coordinates) WithX(x int) Coordinates {
	return Coordinates{x, c.y}
}

func (c Coordinates) WithY(y int) Coordinates {
	return Coordinates{c.x, y}
}

type Bump struct {
	wall Coordinates
	from Coordinates
}

type Map struct {
	matrix [][]rune
	guard  Coordinates
	bumps  []Bump
}

func NewMap(matrix [][]rune) Map {
	var guard Coordinates
LocateGuardLoop:
	for y, row := range matrix {
		for x, cell := range row {
			switch cell {
			case DirectionUp.rune(), DirectionDown.rune(), DirectionLeft.rune(), DirectionRight.rune():
				guard = Coordinates{x, y}
				break LocateGuardLoop
			}
		}
	}
	return Map{
		matrix: matrix,
		guard:  guard,
		bumps:  nil,
	}
}

func (m *Map) Copy() Map {
	copied := Map{
		matrix: nil,
		guard:  m.guard,
		bumps:  nil,
	}
	// Copy array
	for _, row := range m.matrix {
		var cp []rune
		for _, cell := range row {
			cp = append(cp, cell)
		}
		copied.matrix = append(copied.matrix, cp)
	}
	for _, c := range m.bumps {
		copied.bumps = append(copied.bumps, c)
	}
	return copied
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

func (m *Map) Get(c Coordinates) rune {
	return m.matrix[c.y][c.x]
}

func (m *Map) Set(c Coordinates, value rune) {
	m.matrix[c.y][c.x] = value
}

func (m *Map) ComputeGuardRoute() Status {
	for {
		guardCoordinates := m.guard
		guardDirection := DirectionFrom(m.Get(m.guard))
		// Compute next movement
		guardNewCoordinates := guardCoordinates
		guardNewDirection := guardDirection
		switch guardDirection {
		case DirectionUp:
			guardNewCoordinates.y -= 1
		case DirectionDown:
			guardNewCoordinates.y += 1
		case DirectionRight:
			guardNewCoordinates.x += 1
		case DirectionLeft:
			guardNewCoordinates.x -= 1
		}
		for {
			// Check guard is still in map
			if !m.IsWithinMap(guardNewCoordinates) {
				m.Set(guardCoordinates, Visited)
				return StatusExitedMap
			}
			// Is new position a wall?
			position := m.Get(guardNewCoordinates)
			if position == Wall || position == WallAlt {
				// Wall bumped already?
				bump := Bump{guardNewCoordinates, guardCoordinates}
				if slices.Contains(m.bumps, bump) {
					m.Set(guardCoordinates, Visited)
					return StatusLooping
				}
				// Register bumped wall
				m.bumps = append(m.bumps, bump)
				// Fix position
				switch guardNewDirection {
				case DirectionUp:
					guardNewDirection = DirectionRight
					guardNewCoordinates.y += 1
					guardNewCoordinates.x += 1
				case DirectionDown:
					guardNewDirection = DirectionLeft
					guardNewCoordinates.y -= 1
					guardNewCoordinates.x -= 1
				case DirectionRight:
					guardNewDirection = DirectionDown
					guardNewCoordinates.x -= 1
					guardNewCoordinates.y += 1
				case DirectionLeft:
					guardNewDirection = DirectionUp
					guardNewCoordinates.x += 1
					guardNewCoordinates.y -= 1
				}
			} else {
				break
			}
		}
		// Commit movement
		m.Set(guardCoordinates, Visited)
		m.Set(guardNewCoordinates, guardNewDirection.rune())
		m.guard = guardNewCoordinates
	}
}

func (m *Map) TotalVisited() (res int) {
	for _, row := range m.matrix {
		for _, cell := range row {
			if cell == Visited {
				res += 1
			}
		}
	}
	return
}

func (m *Map) Print() {
	fmt.Println()
	for _, row := range m.matrix {
		for _, cell := range row {
			fmt.Print(string(cell))
		}
		fmt.Println()
	}
	return
}
