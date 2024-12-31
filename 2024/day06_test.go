package adventofcode2024

import (
	_ "embed"
	"github.com/ncolomer/adventofcode/2024/day06"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	//go:embed testdata/day06_sample.txt
	Day06Sample string
	//go:embed testdata/day06_puzzle.txt
	Day06Puzzle string
)

func TestDay06Part1Sample(t *testing.T) {
	// Given
	m, err := day06.Read(Day06Sample)
	require.NoError(t, err)
	// When
	sum, err := day06.SolvePart1(m)
	require.NoError(t, err)
	// then
	assert.Equal(t, 41, sum)
}

func TestDay06Part1Puzzle(t *testing.T) {
	// Given
	m, err := day06.Read(Day06Puzzle)
	require.NoError(t, err)
	// When
	sum, err := day06.SolvePart1(m)
	require.NoError(t, err)
	// then
	assert.Equal(t, 5208, sum)
	t.Logf("day06 part 1 result is: %d", sum)
}

func TestDay06Part2Sample(t *testing.T) {
	// Given
	m, err := day06.Read(Day06Sample)
	require.NoError(t, err)
	// When
	sum, err := day06.SolvePart2(m)
	require.NoError(t, err)
	// then
	assert.Equal(t, 6, sum)
}

func TestDay06Part2Puzzle(t *testing.T) {
	// Given
	m, err := day06.Read(Day06Puzzle)
	require.NoError(t, err)
	// When
	sum, err := day06.SolvePart2(m)
	require.NoError(t, err)
	// then
	assert.Equal(t, 1972, sum)
	t.Logf("day06 part 2 result is: %d", sum)
}
