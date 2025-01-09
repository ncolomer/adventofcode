package adventofcode2024

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/ncolomer/adventofcode/2024/day13"
	"github.com/stretchr/testify/assert"
)

var (
	//go:embed testdata/day13_sample.txt
	Day13Sample string
	//go:embed testdata/day13_puzzle.txt
	Day13Puzzle string
)

func TestDay13Part1SampleMachines(t *testing.T) {
	// Given
	m := day13.Read(Day13Sample)
	testCases := []struct {
		sample   []*day13.Machine
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
			actual := day13.SolvePart1(tc.sample)
			// then
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestDay13Part1Sample(t *testing.T) {
	// Given
	m := day13.Read(Day13Sample)
	// When
	tokens := day13.SolvePart1(m)
	// then
	assert.Equal(t, uint64(480), tokens)
}

func TestDay13Part1Puzzle(t *testing.T) {
	// Given
	m := day13.Read(Day13Puzzle)
	// When
	tokens := day13.SolvePart1(m)
	// then
	assert.Equal(t, uint64(31589), tokens)
	t.Logf("day13 part 1 result is: %d", tokens)
}

func TestDay13Part2Sample(t *testing.T) {
	// Given
	m := day13.Read(Day13Sample)
	for _, machine := range m {
		machine.Prize.X += 10000000000000
		machine.Prize.Y += 10000000000000
	}
	// When
	tokens := day13.SolvePart2(m)
	// then
	assert.Equal(t, uint64(875318608908), tokens)
}

func TestDay13Part2Puzzle(t *testing.T) {
	// Given
	m := day13.Read(Day13Puzzle)
	for _, machine := range m {
		machine.Prize.X += 10000000000000
		machine.Prize.Y += 10000000000000
	}
	// When
	sum := day13.SolvePart2(m)
	// then
	assert.Equal(t, uint64(98080815200063), sum)
	t.Logf("day13 part 2 result is: %d", sum)
}
