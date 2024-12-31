package adventofcode2024

import (
	_ "embed"
	"fmt"
	"github.com/ncolomer/adventofcode/2024/day08"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	//go:embed testdata/day08_sample1.txt
	Day08Sample1 string
	//go:embed testdata/day08_sample2.txt
	Day08Sample2 string
	//go:embed testdata/day08_sample3.txt
	Day08Sample3 string
	//go:embed testdata/day08_sample4.txt
	Day08Sample4 string
	//go:embed testdata/day08_sample5.txt
	Day08Sample5 string
	//go:embed testdata/day08_puzzle.txt
	Day08Puzzle string
)

func TestDay08Part1Samples(t *testing.T) {
	testCases := []struct {
		sample   string
		expected int
	}{
		{Day08Sample2, 2},
		{Day08Sample3, 4},
		{Day08Sample4, 4},
		{Day08Sample1, 14},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("sample%d", i+1), func(t *testing.T) {
			// Given
			m, err := day08.Read(tc.sample)
			require.NoError(t, err)
			// When
			sum, err := day08.SolvePart1(m)
			require.NoError(t, err)
			// then
			assert.Equal(t, tc.expected, sum)
		})
	}
}

func TestDay08Part1Puzzle(t *testing.T) {
	// Given
	m, err := day08.Read(Day08Puzzle)
	require.NoError(t, err)
	// When
	sum, err := day08.SolvePart1(m)
	require.NoError(t, err)
	// then
	assert.Equal(t, 413, sum)
	t.Logf("day08 part 1 result is: %d", sum)
}

func TestDay08Part2Sample(t *testing.T) {
	testCases := []struct {
		sample   string
		expected int
	}{
		{Day08Sample5, 9},
		{Day08Sample1, 34},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("sample%d", i+1), func(t *testing.T) {
			// Given
			m, err := day08.Read(tc.sample)
			require.NoError(t, err)
			// When
			sum, err := day08.SolvePart2(m)
			require.NoError(t, err)
			// then
			assert.Equal(t, tc.expected, sum)
		})
	}

}

func TestDay08Part2Puzzle(t *testing.T) {
	// Given
	m, err := day08.Read(Day08Puzzle)
	require.NoError(t, err)
	// When
	sum, err := day08.SolvePart2(m)
	require.NoError(t, err)
	// then
	assert.Equal(t, 1417, sum)
	t.Logf("day08 part 2 result is: %d", sum)
}
