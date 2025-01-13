package day15

import (
	"bufio"
	"iter"
	"strings"

	"github.com/ncolomer/adventofcode/2024/util"
)

const (
	Wall         = '#'
	Robot        = '@'
	Box          = 'O'
	WideBoxLeft  = '['
	WideBoxRight = ']'
	Empty        = '.'
)

func Read(input string) (m util.Map[rune], d []util.Direction) {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	var runes [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		runes = append(runes, []rune(line))
	}
	m = util.NewMap[rune](runes)
	for scanner.Scan() {
		for _, r := range scanner.Text() {
			d = append(d, util.DirectionFrom(r))
		}
	}
	return
}

func Scale(m util.Map[rune]) util.Map[rune] {
	var scaled [][]rune
	next, _ := iter.Pull2(m.Iterate())
	for range m.Height() {
		var row []rune
		for range m.Width() {
			_, r, _ := next()
			switch r {
			case Wall:
				row = append(row, Wall, Wall)
			case Box:
				row = append(row, WideBoxLeft, WideBoxRight)
			case Empty:
				row = append(row, Empty, Empty)
			case Robot:
				row = append(row, Robot, Empty)
			}
		}
		scaled = append(scaled, row)
	}
	return util.NewMap(scaled)
}

func findRobot(m util.Map[rune]) util.Coordinate {
	for c, v := range m.Iterate() {
		if v == Robot {
			return c
		}
	}
	panic("robot not found")
}

func SolvePart1(m util.Map[rune], moves []util.Direction) (res int) {
	var push func(util.Coordinate, util.Direction) bool
	push = func(c util.Coordinate, d util.Direction) bool {
		next := c.Next(d)
		if !m.IsValid(next) {
			return false
		}
		switch m.At(next) {
		case Empty:
			m.Swap(c, next)
			return true
		case Box:
			if push(next, d) {
				m.Swap(c, next)
				return true
			}
		}
		return false
	}
	robot := findRobot(m)
	// apply moves
	for _, move := range moves {
		nextPosition := robot.Next(move)
		switch m.At(nextPosition) {
		case Wall:
			continue // facing wall, don't move
		case Box:
			if !push(nextPosition, move) {
				continue // facing a box that can't move
			}
		}
		m.Swap(robot, nextPosition)
		robot = nextPosition
	}
	// compute boxes' GPS coordinates
	for c, v := range m.Iterate() {
		if v == Box || v == WideBoxLeft {
			res += 100*c.Y + c.X
		}
	}
	return
}

func SolvePart2(m util.Map[rune], moves []util.Direction) (res int) {
	var push func(util.Coordinate, util.Direction, bool) bool
	push = func(c util.Coordinate, d util.Direction, apply bool) bool {
		var selfLeft, selfRight util.Coordinate
		switch m.At(c) {
		case WideBoxLeft:
			selfLeft = c
			selfRight = c.MoveX(1)
		case WideBoxRight:
			selfLeft = c.MoveX(-1)
			selfRight = c
		default:
			panic("unexpected")
		}
		switch d {
		case util.DirectionUp, util.DirectionDown:
			nextLeft, nextRight := selfLeft.Next(d), selfRight.Next(d)
			// []
			// --
			if !m.IsValid(nextLeft) || !m.IsValid(nextRight) {
				return false
			}
			// []
			// XX
			if m.At(nextLeft) == Wall || m.At(nextRight) == Wall {
				return false
			}
			// []
			// ..
			if m.At(nextLeft) == Empty && m.At(nextRight) == Empty {
				if apply {
					m.Swap(selfLeft, nextLeft)
					m.Swap(selfRight, nextRight)
				}
				return true
			}
			// []
			// []
			if m.At(nextLeft) == WideBoxLeft && m.At(nextRight) == WideBoxRight {
				if push(nextLeft, d, false) {
					if apply {
						push(nextLeft, d, true)
						m.Swap(selfLeft, nextLeft)
						m.Swap(selfRight, nextRight)
					}
					return true
				}
				return false
			}
			// []
			//[][]
			if m.At(nextLeft) == WideBoxRight && m.At(nextRight) == WideBoxLeft {
				if push(nextLeft, d, false) && push(nextRight, d, false) {
					if apply {
						push(nextLeft, d, true)
						push(nextRight, d, true)
						m.Swap(selfLeft, nextLeft)
						m.Swap(selfRight, nextRight)
					}
					return true
				}
				return false
			}
			// []
			//[].
			if m.At(nextLeft) == WideBoxRight {
				if push(nextLeft, d, false) {
					if apply {
						push(nextLeft, d, true)
						m.Swap(selfLeft, nextLeft)
						m.Swap(selfRight, nextRight)
					}
					return true
				}
				return false
			}
			// []
			// .[]
			if m.At(nextRight) == WideBoxLeft {
				if push(nextRight, d, false) {
					if apply {
						push(nextRight, d, true)
						m.Swap(selfLeft, nextLeft)
						m.Swap(selfRight, nextRight)
					}
					return true
				}
				return false
			}
		case util.DirectionLeft:
			next := selfLeft.Next(d)
			// |[] + X[]
			if !m.IsValid(next) || m.At(next) == Wall {
				return false
			}
			// .[]
			if m.At(next) == Empty {
				if apply {
					m.Swap(selfLeft, next)
					m.Swap(selfRight, selfLeft)
				}
				return true
			}
			// ][]
			if m.At(next) == WideBoxRight {
				if push(next, d, false) {
					if apply {
						push(next, d, true)
						m.Swap(selfLeft, next)
						m.Swap(selfRight, selfLeft)
					}
					return true
				}
				return false
			}
		case util.DirectionRight:
			next := selfRight.Next(d)
			// []| + []X
			if !m.IsValid(next) || m.At(next) == Wall {
				return false
			}
			// [].
			if m.At(next) == Empty {
				if apply {
					m.Swap(selfRight, next)
					m.Swap(selfLeft, selfRight)
				}
				return true
			}
			// [][
			if m.At(next) == WideBoxLeft {
				if push(next, d, false) {
					if apply {
						push(next, d, true)
						m.Swap(selfRight, next)
						m.Swap(selfLeft, selfRight)
					}
					return true
				}
				return false
			}
		default:
			panic("unexpected")
		}
		return false
	}
	robot := findRobot(m)
	// apply moves
	for _, move := range moves {
		nextPosition := robot.Next(move)
		switch m.At(nextPosition) {
		case Wall:
			continue // facing wall, don't move
		case WideBoxLeft, WideBoxRight:
			if !push(nextPosition, move, true) {
				continue // facing a box that can't move
			}
		}
		m.Swap(robot, nextPosition)
		robot = nextPosition
	}
	// compute boxes' GPS coordinates
	for c, v := range m.Iterate() {
		if v == Box || v == WideBoxLeft {
			res += 100*c.Y + c.X
		}
	}
	return
}
