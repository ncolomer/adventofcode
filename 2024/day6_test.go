package adventofcode2024

import (
	_ "embed"
	"github.com/ncolomer/adventofcode/2024/day6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	//go:embed testdata/day6_sample.txt
	Day6Sample string
	//go:embed testdata/day6_puzzle.txt
	Day6Puzzle string
)

func TestDay6Part1Sample(t *testing.T) {
	// Given
	m, err := day6.Read(Day6Sample)
	require.NoError(t, err)
	// When
	sum, err := day6.SolvePart1(m)
	require.NoError(t, err)
	// then
	assert.Equal(t, 41, sum)
}

func TestDay6Part1Puzzle(t *testing.T) {
	// Given
	m, err := day6.Read(Day6Puzzle)
	require.NoError(t, err)
	// When
	sum, err := day6.SolvePart1(m)
	require.NoError(t, err)
	// then
	assert.Equal(t, 5208, sum)
	t.Logf("day6 part 1 result is: %d", sum)
}

func TestDay6Part2Sample(t *testing.T) {
	// Given
	m, err := day6.Read(Day6Sample)
	require.NoError(t, err)
	// When
	sum, err := day6.SolvePart2(m)
	require.NoError(t, err)
	// then
	assert.Equal(t, 6, sum)
}

func TestDay6Part2Puzzle(t *testing.T) {
	// Given
	m, err := day6.Read(Day6Puzzle)
	require.NoError(t, err)
	// When
	sum, err := day6.SolvePart2(m)
	require.NoError(t, err)
	// then
	assert.Equal(t, 1972, sum)
	t.Logf("day6 part 2 result is: %d", sum)
}
