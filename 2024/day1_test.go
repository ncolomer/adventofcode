package day1

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	//go:embed testdata/day1_sample.txt
	sample string
	//go:embed testdata/day1_puzzle.txt
	puzzle string
)

func TestPart1Sample(t *testing.T) {
	// Given
	list1, list2, err := Read(sample)
	require.NoError(t, err)
	// When
	totalDistance, err := SolvePart1(list1, list2)
	require.NoError(t, err)
	// then
	assert.Equal(t, 11, totalDistance)
}

func TestPart1Puzzle(t *testing.T) {
	// Given
	list1, list2, err := Read(puzzle)
	require.NoError(t, err)
	// When
	totalDistance, err := SolvePart1(list1, list2)
	require.NoError(t, err)
	// then
	t.Logf("part 1 result is: %d", totalDistance)
}

func TestPart2Sample(t *testing.T) {
	// Given
	list1, list2, err := Read(sample)
	require.NoError(t, err)
	// When
	similarityScore, err := SolvePart2(list1, list2)
	require.NoError(t, err)
	// then
	assert.Equal(t, 31, similarityScore)
}

func TestPart2Puzzle(t *testing.T) {
	// Given
	list1, list2, err := Read(puzzle)
	require.NoError(t, err)
	// When
	totalDistance, err := SolvePart2(list1, list2)
	require.NoError(t, err)
	// then
	t.Logf("part 2 result is: %d", totalDistance)
}
