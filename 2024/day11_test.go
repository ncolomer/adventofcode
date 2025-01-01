package adventofcode2024

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/ncolomer/adventofcode/2024/day11"
	"github.com/stretchr/testify/assert"
)

var (
	Day11Puzzle = "112 1110 163902 0 7656027 83039 9 74"
)

func TestDay11Part1SampleArrangement(t *testing.T) {
	testCases := []struct {
		sample []string
	}{
		{[]string{
			"0 1 10 99 999",
			"1 2024 1 0 9 9 2021976",
		}},
		{[]string{
			"125 17",
			"253000 1 7",
			"253 0 2024 14168",
			"512072 1 20 24 28676032",
			"512 72 2024 2 0 2 4 2867 6032",
			"1036288 7 2 20 24 4048 1 4048 8096 28 67 60 32",
			"2097446912 14168 4048 2 0 2 4 40 48 2024 40 48 80 96 2 8 6 7 6 0 3 2",
		}},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("sample%d", i+1), func(t *testing.T) {
			// Given
			s := day11.NewStones(tc.sample[0])
			for j, expected := range tc.sample[1:] {
				// When
				s = s.Blink()
				// then
				assert.Equalf(t, expected, s.String(), "blink %d has not the expected value", j)
			}
		})
	}
}

func TestDay11Part1Sample(t *testing.T) {
	// Given
	s := day11.NewStones("125 17")
	// When
	//for range 25 {s = s.Blink()}
	//count := s.Count()
	count := s.BlinkFast(25)
	// Then
	assert.Equal(t, int64(55312), count)
}

func TestDay11Part1Puzzle(t *testing.T) {
	// Given
	s := day11.NewStones(Day11Puzzle)
	// When
	//for range 25 {s = s.Blink()}
	//count := s.Count()
	count := s.BlinkFast(25)
	// then
	assert.Equal(t, int64(183620), count)
	t.Logf("day11 part 1 result is: %d", count)
}

func TestDay11Part2Puzzle(t *testing.T) {
	// Given
	s := day11.NewStones(Day11Puzzle)
	// When
	count := s.BlinkFast(75)
	// then
	assert.Equal(t, int64(220377651399268), count)
	t.Logf("day11 part 2 result is: %d", count)
}
