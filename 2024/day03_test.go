package adventofcode2024

import (
	_ "embed"
	"github.com/ncolomer/adventofcode/2024/day03"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	//go:embed testdata/day03_sample1.txt
	Day03Sample1 string
	//go:embed testdata/day03_sample2.txt
	Day03Sample2 string
	//go:embed testdata/day03_puzzle.txt
	Day03Puzzle string
)

func TestDay03Part1Sample(t *testing.T) {
	// Given
	instructions, err := day03.Day03Read(Day03Sample1)
	require.NoError(t, err)
	// When
	sum, err := day03.SolveDay03Part1(instructions)
	require.NoError(t, err)
	// then
	assert.Equal(t, 161, sum)
}

func TestDay03Part1Puzzle(t *testing.T) {
	// Given
	instructions, err := day03.Day03Read(Day03Puzzle)
	require.NoError(t, err)
	// When
	sum, err := day03.SolveDay03Part1(instructions)
	require.NoError(t, err)
	// then
	assert.Equal(t, 183788984, sum)
	t.Logf("day 03 part 1 result is: %d", sum)
}

func TestDay03Part2Sample(t *testing.T) {
	// Given
	instructions, err := day03.Day03Read(Day03Sample2)
	require.NoError(t, err)
	// When
	sum, err := day03.SolveDay03Part2(instructions)
	require.NoError(t, err)
	// then
	assert.Equal(t, 48, sum)
}

func TestDay03Part2Puzzle(t *testing.T) {
	// Given
	instructions, err := day03.Day03Read(Day03Puzzle)
	require.NoError(t, err)
	// When
	sum, err := day03.SolveDay03Part2(instructions)
	require.NoError(t, err)
	// then
	assert.Equal(t, 62098619, sum)
	t.Logf("day 03 part 2 result is: %d", sum)
}
