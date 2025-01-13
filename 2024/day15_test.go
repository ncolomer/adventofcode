package adventofcode2024

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/ncolomer/adventofcode/2024/day15"
	"github.com/stretchr/testify/assert"
)

var (
	//go:embed testdata/day15_sample1.txt
	Day15Sample1 string
	//go:embed testdata/day15_sample2.txt
	Day15Sample2 string
	//go:embed testdata/day15_puzzle.txt
	Day15Puzzle string
)

func TestDay15Part1Sample2(t *testing.T) {
	// Given
	m := day15.Read(Day15Sample2)
	testCases := []struct {
		sample   []*day15.Machine
		expected uint64
	}{
		{sample: m[0:1], expected: 280},
		{sample: m[1:2], expected: 0},
		{sample: m[2:3], expected: 200},
		{sample: m[3:4], expected: 0},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("sample%d", i+1), func(t *testing.T) {
			// When
			actual := day15.SolvePart1(tc.sample)
			// then
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestDay15Part1Sample(t *testing.T) {
	// Given
	m := day15.Read(Day15Sample1)
	// When
	tokens := day15.SolvePart1(m)
	// then
	assert.Equal(t, uint64(480), tokens)
}

func TestDay15Part1Puzzle(t *testing.T) {
	// Given
	m := day15.Read(Day15Puzzle)
	// When
	tokens := day15.SolvePart1(m)
	// then
	assert.Equal(t, uint64(31589), tokens)
	t.Logf("day15 part 1 result is: %d", tokens)
}

func TestDay15Part2Sample(t *testing.T) {
	// Given
	m := day15.Read(Day15Sample1)
	for _, machine := range m {
		machine.Prize.X += 10000000000000
		machine.Prize.Y += 10000000000000
	}
	// When
	tokens := day15.SolvePart2(m)
	// then
	assert.Equal(t, uint64(875318608908), tokens)
}

func TestDay15Part2Puzzle(t *testing.T) {
	// Given
	m := day15.Read(Day15Puzzle)
	for _, machine := range m {
		machine.Prize.X += 10000000000000
		machine.Prize.Y += 10000000000000
	}
	// When
	sum := day15.SolvePart2(m)
	// then
	assert.Equal(t, uint64(98080815200063), sum)
	t.Logf("day15 part 2 result is: %d", sum)
}
