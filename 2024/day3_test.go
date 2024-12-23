package adventofcode2024

import (
	_ "embed"
	"github.com/ncolomer/adventofcode/2024/day3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	//go:embed testdata/day3_sample1.txt
	Day3Sample1 string
	//go:embed testdata/day3_sample2.txt
	Day3Sample2 string
	//go:embed testdata/day3_puzzle.txt
	Day3Puzzle string
)

func TestDay3Part1Sample(t *testing.T) {
	// Given
	instructions, err := day3.Day3Read(Day3Sample1)
	require.NoError(t, err)
	// When
	sum, err := day3.SolveDay3Part1(instructions)
	require.NoError(t, err)
	// then
	assert.Equal(t, 161, sum)
}

func TestDay3Part1Puzzle(t *testing.T) {
	// Given
	instructions, err := day3.Day3Read(Day3Puzzle)
	require.NoError(t, err)
	// When
	sum, err := day3.SolveDay3Part1(instructions)
	require.NoError(t, err)
	// then
	assert.Equal(t, 183788984, sum)
	t.Logf("day 3 part 1 result is: %d", sum)
}

func TestDay3Part2Sample(t *testing.T) {
	// Given
	instructions, err := day3.Day3Read(Day3Sample2)
	require.NoError(t, err)
	// When
	sum, err := day3.SolveDay3Part2(instructions)
	require.NoError(t, err)
	// then
	assert.Equal(t, 48, sum)
}

func TestDay3Part2Puzzle(t *testing.T) {
	// Given
	instructions, err := day3.Day3Read(Day3Puzzle)
	require.NoError(t, err)
	// When
	sum, err := day3.SolveDay3Part2(instructions)
	require.NoError(t, err)
	// then
	assert.Equal(t, 62098619, sum)
	t.Logf("day 3 part 2 result is: %d", sum)
}
