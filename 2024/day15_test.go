package adventofcode2024

import (
	_ "embed"
	"testing"

	"github.com/ncolomer/adventofcode/2024/day15"
	"github.com/stretchr/testify/assert"
)

var (
	//go:embed testdata/day15_sample1.txt
	Day15Sample1 string
	//go:embed testdata/day15_sample2.txt
	Day15Sample2 string
	//go:embed testdata/day15_sample3.txt
	Day15Sample3 string
	//go:embed testdata/day15_puzzle.txt
	Day15Puzzle string
)

func TestDay15Part1Sample2(t *testing.T) {
	// Given
	m, d := day15.Read(Day15Sample2)
	// When
	sum := day15.SolvePart1(m, d)
	// then
	assert.Equal(t, 2028, sum)
}

func TestDay15Part1Sample1(t *testing.T) {
	// Given
	m, d := day15.Read(Day15Sample1)
	// When
	sum := day15.SolvePart1(m, d)
	// then
	assert.Equal(t, 10092, sum)
}

func TestDay15Part1Puzzle(t *testing.T) {
	// Given
	m, d := day15.Read(Day15Puzzle)
	// When
	sum := day15.SolvePart1(m, d)
	// then
	assert.Equal(t, 1526673, sum)
	t.Logf("day15 part 1 result is: %d", sum)
}

func TestDay15Part2Sample3(t *testing.T) {
	// Given
	m, d := day15.Read(Day15Sample3)
	m = day15.Scale(m)
	// When
	sum := day15.SolvePart2(m, d)
	// then
	assert.Equal(t, 618, sum)
}

func TestDay15Part2Sample1(t *testing.T) {
	// Given
	m, d := day15.Read(Day15Sample1)
	m = day15.Scale(m)
	// When
	sum := day15.SolvePart2(m, d)
	// then
	assert.Equal(t, 9021, sum)
}

func TestDay15Part2Puzzle(t *testing.T) {
	// Given
	m, d := day15.Read(Day15Puzzle)
	m = day15.Scale(m)
	// When
	sum := day15.SolvePart2(m, d)
	// then
	assert.Equal(t, 1535509, sum)
	t.Logf("day15 part 2 result is: %d", sum)
}
