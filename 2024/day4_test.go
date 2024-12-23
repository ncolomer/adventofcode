package adventofcode2024

import (
	_ "embed"
	"github.com/ncolomer/adventofcode/2024/day4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	//go:embed testdata/day4_sample.txt
	Day4Sample string
	//go:embed testdata/day4_puzzle.txt
	Day4Puzzle string
)

func TestDay4Part1Sample(t *testing.T) {
	// Given
	matrix, err := day4.Day4Read(Day4Sample)
	require.NoError(t, err)
	// When
	occurrences, err := day4.SolveDay4Part1(matrix)
	require.NoError(t, err)
	// then
	assert.Equal(t, 18, occurrences)
}

func TestDay4Part1Puzzle(t *testing.T) {
	// Given
	matrix, err := day4.Day4Read(Day4Puzzle)
	require.NoError(t, err)
	// When
	occurrences, err := day4.SolveDay4Part1(matrix)
	require.NoError(t, err)
	// then
	assert.Equal(t, 2545, occurrences)
	t.Logf("day 4 part 1 result is: %d", occurrences)
}

func TestDay4Part2Sample(t *testing.T) {
	// Given
	matrix, err := day4.Day4Read(Day4Sample)
	require.NoError(t, err)
	// When
	occurrences, err := day4.SolveDay4Part2(matrix)
	require.NoError(t, err)
	// then
	assert.Equal(t, 9, occurrences)
}

func TestDay4Part2Puzzle(t *testing.T) {
	// Given
	matrix, err := day4.Day4Read(Day4Puzzle)
	require.NoError(t, err)
	// When
	occurrences, err := day4.SolveDay4Part2(matrix)
	require.NoError(t, err)
	// then
	assert.Equal(t, 1886, occurrences)
	t.Logf("day 4 part 2 result is: %d", occurrences)
}
