package adventofcode2024

import (
	_ "embed"
	"github.com/ncolomer/adventofcode/2024/day7"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	//go:embed testdata/day7_sample.txt
	Day7Sample string
	//go:embed testdata/day7_puzzle.txt
	Day7Puzzle string
)

func TestDay7Part1Sample(t *testing.T) {
	// Given
	m, err := day7.Read(Day7Sample)
	require.NoError(t, err)
	// When
	sum, err := day7.SolvePart1(m)
	require.NoError(t, err)
	// then
	assert.Equal(t, int64(3749), sum)
}

func TestDay7Part1Puzzle(t *testing.T) {
	// Given
	m, err := day7.Read(Day7Puzzle)
	require.NoError(t, err)
	// When
	sum, err := day7.SolvePart1(m)
	require.NoError(t, err)
	// then
	assert.Equal(t, int64(945512582195), sum)
	t.Logf("day7 part 1 result is: %d", sum)
}

func TestDay7Part2Sample(t *testing.T) {
	// Given
	m, err := day7.Read(Day7Sample)
	require.NoError(t, err)
	// When
	sum, err := day7.SolvePart2(m)
	require.NoError(t, err)
	// then
	assert.Equal(t, int64(11387), sum)
}

func TestDay7Part2Puzzle(t *testing.T) {
	// Given
	m, err := day7.Read(Day7Puzzle)
	require.NoError(t, err)
	// When
	sum, err := day7.SolvePart2(m)
	require.NoError(t, err)
	// then
	assert.Equal(t, int64(271691107779347), sum)
	t.Logf("day7 part 2 result is: %d", sum)
}
