package adventofcode2024

import (
	_ "embed"
	"github.com/ncolomer/adventofcode/2024/day01"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	//go:embed testdata/day01_sample.txt
	Day01Sample string
	//go:embed testdata/day01_puzzle.txt
	Day01Puzzle string
)

func TestDay01Part1Sample(t *testing.T) {
	// Given
	list1, list2, err := day01.Day01Read(Day01Sample)
	require.NoError(t, err)
	// When
	totalDistance, err := day01.SolveDay01Part1(list1, list2)
	require.NoError(t, err)
	// then
	assert.Equal(t, 11, totalDistance)
}

func TestDay01Part1Puzzle(t *testing.T) {
	// Given
	list1, list2, err := day01.Day01Read(Day01Puzzle)
	require.NoError(t, err)
	// When
	totalDistance, err := day01.SolveDay01Part1(list1, list2)
	require.NoError(t, err)
	// then
	assert.Equal(t, 2375403, totalDistance)
	t.Logf("day 01 part 1 result is: %d", totalDistance)
}

func TestDay01Part2Sample(t *testing.T) {
	// Given
	list1, list2, err := day01.Day01Read(Day01Sample)
	require.NoError(t, err)
	// When
	similarityScore, err := day01.SolveDay01Part2(list1, list2)
	require.NoError(t, err)
	// then
	assert.Equal(t, 31, similarityScore)
}

func TestDay01Part2Puzzle(t *testing.T) {
	// Given
	list1, list2, err := day01.Day01Read(Day01Puzzle)
	require.NoError(t, err)
	// When
	totalDistance, err := day01.SolveDay01Part2(list1, list2)
	require.NoError(t, err)
	// then
	assert.Equal(t, 23082277, totalDistance)
	t.Logf("day 01 part 2 result is: %d", totalDistance)
}
