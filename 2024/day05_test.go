package adventofcode2024

import (
	_ "embed"
	"github.com/ncolomer/adventofcode/2024/day05"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	//go:embed testdata/day05_sample.txt
	Day05Sample string
	//go:embed testdata/day05_puzzle.txt
	Day05Puzzle string
)

func TestDay05Part1Sample(t *testing.T) {
	// Given
	ordering, updates, err := day05.Day05Read(Day05Sample)
	require.NoError(t, err)
	// When
	sum, err := day05.SolveDay05Part1(ordering, updates)
	require.NoError(t, err)
	// then
	assert.Equal(t, 143, sum)
}

func TestDay05Part1Puzzle(t *testing.T) {
	// Given
	ordering, updates, err := day05.Day05Read(Day05Puzzle)
	require.NoError(t, err)
	// When
	sum, err := day05.SolveDay05Part1(ordering, updates)
	require.NoError(t, err)
	// then
	assert.Equal(t, 4872, sum)
	t.Logf("day05 part 1 result is: %d", sum)
}

func TestDay05Part2Sample(t *testing.T) {
	// Given
	ordering, updates, err := day05.Day05Read(Day05Sample)
	require.NoError(t, err)
	// When
	sum, err := day05.SolveDay05Part2(ordering, updates)
	require.NoError(t, err)
	// then
	assert.Equal(t, 123, sum)
}

func TestDay05Part2Puzzle(t *testing.T) {
	// Given
	ordering, updates, err := day05.Day05Read(Day05Puzzle)
	require.NoError(t, err)
	// When
	sum, err := day05.SolveDay05Part2(ordering, updates)
	require.NoError(t, err)
	// then
	assert.Equal(t, 5564, sum)
	t.Logf("day05 part 2 result is: %d", sum)
}
