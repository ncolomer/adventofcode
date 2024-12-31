package adventofcode2024

import (
	_ "embed"
	"fmt"
	"github.com/ncolomer/adventofcode/2024/day10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	//go:embed testdata/day10_sample1.txt
	Day10Sample1 string
	//go:embed testdata/day10_sample2.txt
	Day10Sample2 string
	//go:embed testdata/day10_sample3.txt
	Day10Sample3 string
	//go:embed testdata/day10_sample4.txt
	Day10Sample4 string
	//go:embed testdata/day10_sample5.txt
	Day10Sample5 string
	//go:embed testdata/day10_sample6.txt
	Day10Sample6 string
	//go:embed testdata/day10_sample7.txt
	Day10Sample7 string
	//go:embed testdata/day10_sample8.txt
	Day10Sample8 string
	//go:embed testdata/day10_puzzle.txt
	Day10Puzzle string
)

func TestDay10Part1Sample(t *testing.T) {
	testCases := []struct {
		sample   string
		expected int
	}{
		{Day10Sample1, 36},
		{Day10Sample2, 1},
		{Day10Sample3, 2},
		{Day10Sample4, 4},
		{Day10Sample5, 3},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("sample%d", i+1), func(t *testing.T) {
			// Given
			m, err := day10.Read(tc.sample)
			require.NoError(t, err)
			// When
			sum, err := day10.SolvePart1(m)
			require.NoError(t, err)
			// then
			assert.Equal(t, tc.expected, sum)
		})
	}
}

func TestDay10Part1Puzzle(t *testing.T) {
	// Given
	m, err := day10.Read(Day10Puzzle)
	require.NoError(t, err)
	// When
	sum, err := day10.SolvePart1(m)
	require.NoError(t, err)
	// then
	assert.Equal(t, 582, sum)
	t.Logf("day10 part 1 result is: %d", sum)
}

func TestDay10Part2Sample(t *testing.T) {
	testCases := []struct {
		sample   string
		expected int
	}{
		{Day10Sample6, 3},
		{Day10Sample7, 13},
		{Day10Sample8, 227},
		{Day10Sample1, 81},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("sample%d", i+1), func(t *testing.T) {
			// Given
			m, err := day10.Read(tc.sample)
			require.NoError(t, err)
			// When
			sum, err := day10.SolvePart2(m)
			require.NoError(t, err)
			// then
			assert.Equal(t, tc.expected, sum)
		})
	}
}

func TestDay10Part2Puzzle(t *testing.T) {
	// Given
	m, err := day10.Read(Day10Puzzle)
	require.NoError(t, err)
	// When
	sum, err := day10.SolvePart2(m)
	require.NoError(t, err)
	// then
	assert.Equal(t, 1302, sum)
	t.Logf("day10 part 2 result is: %d", sum)
}
