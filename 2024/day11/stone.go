package day11

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	NumberRegex = regexp.MustCompile(`\d+`)
)

type Stones []int64

func NewStones(input string) (stones Stones) {
	for _, match := range NumberRegex.FindAllString(input, -1) {
		number, _ := strconv.ParseInt(match, 10, 64)
		stones = append(stones, number)
	}
	return stones
}

func blinkOnce(nbr int64) Stones {
	str := strconv.FormatInt(int64(nbr), 10)
	switch {
	case nbr == 0:
		return []int64{1}
	case len(str)%2 == 0:
		half := len(str) / 2
		left, _ := strconv.ParseInt(str[:half], 10, 64)
		right, _ := strconv.ParseInt(str[half:], 10, 64)
		return []int64{left, right}
	default:
		return []int64{nbr * 2024}
	}
}

func (s *Stones) Blink() (stones Stones) {
	for _, stone := range *s {
		stones = append(stones, blinkOnce(stone)...)
	}
	return stones
}

func (s *Stones) Count() int64 {
	return int64(len(*s))
}

func (s *Stones) String() string {
	var str []string
	for _, nbr := range *s {
		str = append(str, strconv.FormatInt(nbr, 10))
	}
	return strings.Join(str, " ")
}

type Key struct {
	Stone int64
	Depth int
}

type Index map[Key]int64

func (s *Stones) BlinkFast(times int) int64 {
	index := make(Index)
	var recur func(int64, int) int64
	recur = func(stone int64, times int) int64 {
		if times == 0 {
			return 1
		}
		key := Key{stone, times}
		sum, found := index[key]
		if !found {
			for _, forked := range blinkOnce(stone) {
				sum += recur(forked, times-1)
			}
			index[key] = sum
		}
		return sum
	}
	var sum int64
	for _, nbr := range *s {
		sum += recur(nbr, times)
	}
	return sum
}
