package adventofcode2024

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	//go:embed testdata/day2_sample.txt
	Day2Sample string
	//go:embed testdata/day2_puzzle.txt
	Day2Puzzle string
)

func TestDay2Part1Sample(t *testing.T) {
	// Given
	reports, err := Day2Read(Day2Sample)
	require.NoError(t, err)
	// When
	safeReportCount, err := SolveDay2Part1(reports)
	require.NoError(t, err)
	// then
	assert.Equal(t, 2, safeReportCount)
}

func TestDay2Part1Puzzle(t *testing.T) {
	// Given
	reports, err := Day2Read(Day2Puzzle)
	require.NoError(t, err)
	// When
	safeReportCount, err := SolveDay2Part1(reports)
	require.NoError(t, err)
	// then
	assert.Equal(t, 483, safeReportCount)
	t.Logf("day 2 part 1 result is: %d", safeReportCount)
}

func TestDay2Part2Puzzle(t *testing.T) {
	// Given
	reports, err := Day2Read(Day2Puzzle)
	require.NoError(t, err)
	// When
	safeReportCount, err := SolveDay2Part2(reports)
	require.NoError(t, err)
	// then
	assert.Equal(t, 528, safeReportCount)
	t.Logf("day 2 part 2 result is: %d", safeReportCount)
}
