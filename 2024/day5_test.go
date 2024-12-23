package adventofcode2024

import (
	_ "embed"
	"github.com/ncolomer/adventofcode/2024/day5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	//go:embed testdata/day5_sample.txt
	Day5Sample string
	//go:embed testdata/day5_puzzle.txt
	Day5Puzzle string
)

func TestDay5Part1Sample(t *testing.T) {
	// Given
	ordering, updates, err := day5.Day5Read(Day5Sample)
	require.NoError(t, err)
	// When
	sum, err := day5.SolveDay5Part1(ordering, updates)
	require.NoError(t, err)
	// then
	assert.Equal(t, 143, sum)
}

func TestDay5Part1Puzzle(t *testing.T) {
	// Given
	ordering, updates, err := day5.Day5Read(Day5Puzzle)
	require.NoError(t, err)
	// When
	sum, err := day5.SolveDay5Part1(ordering, updates)
	require.NoError(t, err)
	// then
	assert.Equal(t, 4872, sum)
	t.Logf("day 5 part 1 result is: %d", sum)
}

func TestDay5Part2Sample(t *testing.T) {
	// Given
	ordering, updates, err := day5.Day5Read(Day5Sample)
	require.NoError(t, err)
	// When
	sum, err := day5.SolveDay5Part2(ordering, updates)
	require.NoError(t, err)
	// then
	assert.Equal(t, 123, sum)
}

func TestDay5Part2Puzzle(t *testing.T) {
	// Given
	ordering, updates, err := day5.Day5Read(Day5Puzzle)
	require.NoError(t, err)
	// When
	sum, err := day5.SolveDay5Part2(ordering, updates)
	require.NoError(t, err)
	// then
	assert.Equal(t, 5564, sum)
	t.Logf("day 5 part 2 result is: %d", sum)
}
