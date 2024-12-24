package adventofcode2024

import (
	_ "embed"
	"fmt"
	"github.com/ncolomer/adventofcode/2024/day8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	//go:embed testdata/day8_sample1.txt
	Day8Sample1 string
	//go:embed testdata/day8_sample2.txt
	Day8Sample2 string
	//go:embed testdata/day8_sample3.txt
	Day8Sample3 string
	//go:embed testdata/day8_sample4.txt
	Day8Sample4 string
	//go:embed testdata/day8_sample5.txt
	Day8Sample5 string
	//go:embed testdata/day8_puzzle.txt
	Day8Puzzle string
)

func TestDay8Part1Samples(t *testing.T) {
	testCases := []struct {
		sample   string
		expected int
	}{
		{Day8Sample2, 2},
		{Day8Sample3, 4},
		{Day8Sample4, 4},
		{Day8Sample1, 14},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("sample%d", i+1), func(t *testing.T) {
			// Given
			m, err := day8.Read(tc.sample)
			require.NoError(t, err)
			// When
			sum, err := day8.SolvePart1(m)
			require.NoError(t, err)
			// then
			assert.Equal(t, tc.expected, sum)
		})
	}
}

func TestDay8Part1Puzzle(t *testing.T) {
	// Given
	m, err := day8.Read(Day8Puzzle)
	require.NoError(t, err)
	// When
	sum, err := day8.SolvePart1(m)
	require.NoError(t, err)
	// then
	assert.Equal(t, 413, sum)
	t.Logf("day8 part 1 result is: %d", sum)
}

func TestDay8Part2Sample(t *testing.T) {
	testCases := []struct {
		sample   string
		expected int
	}{
		{Day8Sample5, 9},
		{Day8Sample1, 34},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("sample%d", i+1), func(t *testing.T) {
			// Given
			m, err := day8.Read(tc.sample)
			require.NoError(t, err)
			// When
			sum, err := day8.SolvePart2(m)
			require.NoError(t, err)
			// then
			assert.Equal(t, tc.expected, sum)
		})
	}

}

func TestDay8Part2Puzzle(t *testing.T) {
	// Given
	m, err := day8.Read(Day8Puzzle)
	require.NoError(t, err)
	// When
	sum, err := day8.SolvePart2(m)
	require.NoError(t, err)
	// then
	assert.Equal(t, 1417, sum)
	t.Logf("day8 part 2 result is: %d", sum)
}
