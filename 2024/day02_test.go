package adventofcode2024

import (
	_ "embed"
	"github.com/ncolomer/adventofcode/2024/day02"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	//go:embed testdata/day02_sample.txt
	Day02Sample string
	//go:embed testdata/day02_puzzle.txt
	Day02Puzzle string
)

func TestDay02Part1Sample(t *testing.T) {
	// Given
	reports, err := day02.Day02Read(Day02Sample)
	require.NoError(t, err)
	// When
	safeReportCount, err := day02.SolveDay02Part1(reports)
	require.NoError(t, err)
	// then
	assert.Equal(t, 2, safeReportCount)
}

func TestDay02Part1Puzzle(t *testing.T) {
	// Given
	reports, err := day02.Day02Read(Day02Puzzle)
	require.NoError(t, err)
	// When
	safeReportCount, err := day02.SolveDay02Part1(reports)
	require.NoError(t, err)
	// then
	assert.Equal(t, 483, safeReportCount)
	t.Logf("day02 part 1 result is: %d", safeReportCount)
}

func TestDay02Part2Puzzle(t *testing.T) {
	// Given
	reports, err := day02.Day02Read(Day02Puzzle)
	require.NoError(t, err)
	// When
	safeReportCount, err := day02.SolveDay02Part2(reports)
	require.NoError(t, err)
	// then
	assert.Equal(t, 528, safeReportCount)
	t.Logf("day02 part 2 result is: %d", safeReportCount)
}
