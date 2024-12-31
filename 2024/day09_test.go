package adventofcode2024

import (
	_ "embed"
	"fmt"
	"github.com/ncolomer/adventofcode/2024/day09"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	Day09Sample1 string = "2333133121414131402"
	Day09Sample2 string = "12345"
	Day09Sample3 string = "233313312141413140242"
	//go:embed testdata/day09_puzzle.txt
	Day09Puzzle string
)

func TestDecompress(t *testing.T) {
	testCases := []struct {
		sample   string
		expected string
	}{
		{Day09Sample2, "0..111....22222"},
		{Day09Sample3, "00...111...2...333.44.5555.6666.777.888899....1010"},
		{Day09Sample1, "00...111...2...333.44.5555.6666.777.888899"},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("sample%d", i+1), func(t *testing.T) {
			// Given
			dm := day09.NewDiskMap(tc.sample)
			// then
			assert.Equal(t, tc.expected, dm.String())
		})
	}
}

func TestCompact(t *testing.T) {
	testCases := []struct {
		sample   string
		expected string
	}{
		{Day09Sample2, "022111222......"},
		{Day09Sample3, "00101091119882887333744755556666.................."},
		{Day09Sample1, "0099811188827773336446555566.............."},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("sample%d", i+1), func(t *testing.T) {
			// Given
			dm := day09.NewDiskMap(tc.sample).Compact()
			// then
			assert.Equal(t, tc.expected, dm.String())
		})
	}
}

func TestDay09Part1Sample(t *testing.T) {
	testCases := []struct {
		sample   string
		expected int64
	}{
		{Day09Sample2, 60},
		{Day09Sample3, 2351},
		{Day09Sample1, 1928},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("sample%d", i+1), func(t *testing.T) {
			// Given
			dm := day09.NewDiskMap(tc.sample)
			// When
			dm.Compact()
			// then
			assert.Equal(t, tc.expected, dm.Checksum())
		})
	}
}

func TestDay09Part1Puzzle(t *testing.T) {
	// Given
	dm := day09.NewDiskMap(Day09Puzzle)
	// When
	sum := dm.Compact().Checksum()
	// then
	assert.Equal(t, int64(6334655979668), sum)
	t.Logf("Day09 part 1 result is: %d", sum)
}

func TestDay09Part2Sample(t *testing.T) {
	// Given
	dm := day09.NewDiskMap(Day09Sample1)
	// When
	compacted := dm.Compact2()
	sum := compacted.Checksum()
	// then
	assert.Equal(t, "00992111777.44.333....5555.6666.....8888..", compacted.String())
	assert.Equal(t, int64(2858), sum)
}

func TestDay09Part2Puzzle(t *testing.T) {
	// Given
	dm := day09.NewDiskMap(Day09Puzzle)
	// When
	sum := dm.Compact2().Checksum()
	// then
	assert.Equal(t, int64(6349492251099), sum)
	t.Logf("Day09 part 2 result is: %d", sum)
}
