package util

import (
	"fmt"
	"iter"
	"slices"
)

type Coordinate struct {
	X int
	Y int
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%d, %d)", c.X, c.Y)
}

func (c Coordinate) MoveX(dx int) Coordinate {
	return Coordinate{c.X + dx, c.Y}
}

func (c Coordinate) MoveY(dy int) Coordinate {
	return Coordinate{c.X, c.Y + dy}
}

func (c Coordinate) Next(d Direction) Coordinate {
	switch d {
	case DirectionUp:
		return c.MoveY(-1)
	case DirectionDown:
		return c.MoveY(1)
	case DirectionLeft:
		return c.MoveX(-1)
	case DirectionRight:
		return c.MoveX(1)
	}
	panic("unexpected")
}

type Direction int

const (
	DirectionUp Direction = iota
	DirectionDown
	DirectionLeft
	DirectionRight
)

func (d Direction) rune() rune {
	return map[Direction]rune{
		DirectionUp:    '^',
		DirectionDown:  'v',
		DirectionLeft:  '<',
		DirectionRight: '>',
	}[d]
}

func DirectionFrom(r rune) Direction {
	return map[rune]Direction{
		'^': DirectionUp,
		'v': DirectionDown,
		'<': DirectionLeft,
		'>': DirectionRight,
	}[r]
}

type Map[T any] struct {
	matrix [][]T
}

func NewMap[T any](matrix [][]T) Map[T] {
	return Map[T]{matrix}
}

func (m Map[T]) Height() int {
	return len(m.matrix)
}

func (m Map[T]) Width() int {
	return len(m.matrix[0])
}

func (m Map[T]) At(c Coordinate) T {
	return m.matrix[c.Y][c.X]
}

func (m Map[T]) Set(c Coordinate, v T) {
	if m.IsValid(c) {
		m.matrix[c.Y][c.X] = v
	}
}

func (m Map[T]) IsValid(c Coordinate) bool {
	return c.X >= 0 && c.X < m.Width() &&
		c.Y >= 0 && c.Y < m.Height()
}

func (m Map[T]) Swap(c1, c2 Coordinate) {
	if m.IsValid(c1) && m.IsValid(c2) {
		v1, v2 := m.At(c1), m.At(c2)
		m.Set(c1, v2)
		m.Set(c2, v1)
	}
}

func (m Map[T]) Sides(c Coordinate) []Coordinate {
	return slices.DeleteFunc([]Coordinate{
		{c.X + 1, c.Y},
		{c.X - 1, c.Y},
		{c.X, c.Y + 1},
		{c.X, c.Y - 1},
	}, func(c Coordinate) bool {
		return !m.IsValid(c)
	})
}

func (m Map[T]) Diagonals(c Coordinate) []Coordinate {
	return slices.DeleteFunc([]Coordinate{
		{c.X + 1, c.Y + 1},
		{c.X + 1, c.Y - 1},
		{c.X - 1, c.Y + 1},
		{c.X - 1, c.Y - 1},
	}, func(c Coordinate) bool {
		return !m.IsValid(c)
	})
}

func (m Map[T]) Print(f func(T) string) {
	for _, row := range m.matrix {
		for _, cell := range row {
			fmt.Print(f(cell))

		}
		fmt.Println()
	}
	return
}

func (m Map[T]) Iterate() iter.Seq2[Coordinate, T] {
	return func(yield func(Coordinate, T) bool) {
		for y, row := range m.matrix {
			for x, cell := range row {
				if !yield(Coordinate{x, y}, cell) {
					return
				}
			}
		}
	}
}
