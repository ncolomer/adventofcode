package adventofcode2024

import (
	_ "embed"
	"github.com/ncolomer/adventofcode/2024/day04"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	//go:embed testdata/day04_sample.txt
	Day04Sample string
	//go:embed testdata/day04_puzzle.txt
	Day04Puzzle string
)

func TestDay04Part1Sample(t *testing.T) {
	// Given
	matrix, err := day04.Day04Read(Day04Sample)
	require.NoError(t, err)
	// When
	occurrences, err := day04.SolveDay04Part1(matrix)
	require.NoError(t, err)
	// then
	assert.Equal(t, 18, occurrences)
}

func TestDay04Part1Puzzle(t *testing.T) {
	// Given
	matrix, err := day04.Day04Read(Day04Puzzle)
	require.NoError(t, err)
	// When
	occurrences, err := day04.SolveDay04Part1(matrix)
	require.NoError(t, err)
	// then
	assert.Equal(t, 2545, occurrences)
	t.Logf("day04 part 1 result is: %d", occurrences)
}

func TestDay04Part2Sample(t *testing.T) {
	// Given
	matrix, err := day04.Day04Read(Day04Sample)
	require.NoError(t, err)
	// When
	occurrences, err := day04.SolveDay04Part2(matrix)
	require.NoError(t, err)
	// then
	assert.Equal(t, 9, occurrences)
}

func TestDay04Part2Puzzle(t *testing.T) {
	// Given
	matrix, err := day04.Day04Read(Day04Puzzle)
	require.NoError(t, err)
	// When
	occurrences, err := day04.SolveDay04Part2(matrix)
	require.NoError(t, err)
	// then
	assert.Equal(t, 1886, occurrences)
	t.Logf("day04 part 2 result is: %d", occurrences)
}
