package adventofcode2024

import (
	_ "embed"
	"testing"

	"github.com/ncolomer/adventofcode/2024/day14"
	"github.com/ncolomer/adventofcode/2024/util"
	"github.com/stretchr/testify/assert"
)

var (
	//go:embed testdata/day14_sample.txt
	Day14Sample string
	//go:embed testdata/day14_puzzle.txt
	Day14Puzzle string
)

func TestDay14Part1Single(t *testing.T) {
	// Given
	robot := &day14.Robot{ID: 11, X: 2, Y: 4, VX: 2, VY: -3}
	MapWidth, MapHeight := 11, 7
	expected := []util.Coordinate{
		{4, 1},
		{6, 5},
		{8, 2},
		{10, 6},
		{1, 3},
	}
	// When
	for i := range 5 {
		day14.Tick([]*day14.Robot{robot}, MapWidth, MapHeight)
		assert.Equal(t, expected[i].X, robot.X)
		assert.Equal(t, expected[i].Y, robot.Y)
	}
}

func TestDay14Part1Sample(t *testing.T) {
	// Given
	robots := day14.Read(Day14Sample)
	MapWidth, MapHeight := 11, 7
	// When
	for range 100 {
		day14.Tick(robots, MapWidth, MapHeight)
	}
	safetyFactor := day14.SolvePart1(robots, MapWidth, MapHeight)
	// then
	assert.Equal(t, 12, safetyFactor)
}

func TestDay14Part1Puzzle(t *testing.T) {
	// Given
	robots := day14.Read(Day14Puzzle)
	MapWidth, MapHeight := 101, 103
	// When
	for range 100 {
		day14.Tick(robots, MapWidth, MapHeight)
	}
	safetyFactor := day14.SolvePart1(robots, MapWidth, MapHeight)
	// Then
	assert.Equal(t, 215987200, safetyFactor)
	t.Logf("day14 part 1 result is: %d", safetyFactor)
}

func TestDay14Part2Puzzle(t *testing.T) {
	// Given
	robots := day14.Read(Day14Puzzle)
	MapWidth, MapHeight := 101, 103
	// When
	iteration := day14.SolvePart2(robots, MapWidth, MapHeight)
	// Then
	assert.Equal(t, 8050, iteration)
	t.Logf("day14 part 2 result is: %d", iteration)
}
