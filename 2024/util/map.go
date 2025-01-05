package util

import (
	"fmt"
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

func (m Map[T]) IsValid(c Coordinate) bool {
	return c.X >= 0 && c.X < m.Width() &&
		c.Y >= 0 && c.Y < m.Height()
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

func (m Map[T]) String() (s string) {
	for _, row := range m.matrix {
		for _, cell := range row {
			s += fmt.Sprintf("%v", cell)
		}
	}
	return
}
