package adventofcode2024

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/ncolomer/adventofcode/2024/day12"
	"github.com/stretchr/testify/assert"
)

var (
	//go:embed testdata/day12_sample1.txt
	Day12Sample1 string
	//go:embed testdata/day12_sample2.txt
	Day12Sample2 string
	//go:embed testdata/day12_sample3.txt
	Day12Sample3 string
	//go:embed testdata/day12_sample4.txt
	Day12Sample4 string
	//go:embed testdata/day12_sample5.txt
	Day12Sample5 string
	//go:embed testdata/day12_puzzle.txt
	Day12Puzzle string
)

func TestDay12Part1Sample(t *testing.T) {
	testCases := []struct {
		sample   string
		expected int
	}{
		{sample: Day12Sample1, expected: 1930},
		{sample: Day12Sample2, expected: 140},
		{sample: Day12Sample3, expected: 772},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("sample%d", i+1), func(t *testing.T) {
			// Given
			m := day12.Read(tc.sample)
			// When
			actual := day12.SolvePart1(m)
			// then
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestDay12Part1Puzzle(t *testing.T) {
	// Given
	m := day12.Read(Day12Puzzle)
	// When
	price := day12.SolvePart1(m)
	// then
	assert.Equal(t, 1457298, price)
	t.Logf("day12 part 1 result is: %d", price)
}

func TestDay12Part2Sample(t *testing.T) {
	testCases := []struct {
		sample   string
		expected int
	}{
		{sample: Day12Sample1, expected: 1206},
		{sample: Day12Sample2, expected: 80},
		{sample: Day12Sample3, expected: 436},
		{sample: Day12Sample4, expected: 236},
		{sample: Day12Sample5, expected: 368},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("sample%d", i+1), func(t *testing.T) {
			// Given
			m := day12.Read(tc.sample)
			// When
			actual := day12.SolvePart2(m)
			// then
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestDay12Part2Puzzle(t *testing.T) {
	// Given
	m := day12.Read(Day12Puzzle)
	// When
	sum := day12.SolvePart2(m)
	// then
	assert.Equal(t, 921636, sum)
	t.Logf("day12 part 2 result is: %d", sum)
}
