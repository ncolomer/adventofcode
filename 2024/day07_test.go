package adventofcode2024

import (
	_ "embed"
	"github.com/ncolomer/adventofcode/2024/day07"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	//go:embed testdata/day07_sample.txt
	Day07Sample string
	//go:embed testdata/day07_puzzle.txt
	Day07Puzzle string
)

func TestDay07Part1Sample(t *testing.T) {
	// Given
	m, err := day07.Read(Day07Sample)
	require.NoError(t, err)
	// When
	sum, err := day07.SolvePart1(m)
	require.NoError(t, err)
	// then
	assert.Equal(t, int64(3749), sum)
}

func TestDay07Part1Puzzle(t *testing.T) {
	// Given
	m, err := day07.Read(Day07Puzzle)
	require.NoError(t, err)
	// When
	sum, err := day07.SolvePart1(m)
	require.NoError(t, err)
	// then
	assert.Equal(t, int64(945512582195), sum)
	t.Logf("day07 part 1 result is: %d", sum)
}

func TestDay07Part2Sample(t *testing.T) {
	// Given
	m, err := day07.Read(Day07Sample)
	require.NoError(t, err)
	// When
	sum, err := day07.SolvePart2(m)
	require.NoError(t, err)
	// then
	assert.Equal(t, int64(11387), sum)
}

func TestDay07Part2Puzzle(t *testing.T) {
	// Given
	m, err := day07.Read(Day07Puzzle)
	require.NoError(t, err)
	// When
	sum, err := day07.SolvePart2(m)
	require.NoError(t, err)
	// then
	assert.Equal(t, int64(271691107779347), sum)
	t.Logf("day07 part 2 result is: %d", sum)
}
