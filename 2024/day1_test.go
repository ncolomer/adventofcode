package adventofcode2024

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	//go:embed testdata/day1_sample.txt
	Day1Sample string
	//go:embed testdata/day1_puzzle.txt
	Day1Puzzle string
)

func TestDay1Part1Sample(t *testing.T) {
	// Given
	list1, list2, err := Day1Read(Day1Sample)
	require.NoError(t, err)
	// When
	totalDistance, err := SolveDay1Part1(list1, list2)
	require.NoError(t, err)
	// then
	assert.Equal(t, 11, totalDistance)
}

func TestDay1Part1Puzzle(t *testing.T) {
	// Given
	list1, list2, err := Day1Read(Day1Puzzle)
	require.NoError(t, err)
	// When
	totalDistance, err := SolveDay1Part1(list1, list2)
	require.NoError(t, err)
	// then
	assert.Equal(t, 2375403, totalDistance)
	t.Logf("day 1 part 1 result is: %d", totalDistance)
}

func TestDay1Part2Sample(t *testing.T) {
	// Given
	list1, list2, err := Day1Read(Day1Sample)
	require.NoError(t, err)
	// When
	similarityScore, err := SolveDay1Part2(list1, list2)
	require.NoError(t, err)
	// then
	assert.Equal(t, 31, similarityScore)
}

func TestDay1Part2Puzzle(t *testing.T) {
	// Given
	list1, list2, err := Day1Read(Day1Puzzle)
	require.NoError(t, err)
	// When
	totalDistance, err := SolveDay1Part2(list1, list2)
	require.NoError(t, err)
	// then
	assert.Equal(t, 23082277, totalDistance)
	t.Logf("day 1 part 2 result is: %d", totalDistance)
}
