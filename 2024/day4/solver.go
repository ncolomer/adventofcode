package day4

import (
	"bufio"
	"regexp"
	"strings"
)

func Day4Read(input string) (matrix [][]rune, err error) {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []rune(line))
	}
	return
}

func SolveDay4Part1(matrix [][]rune) (res int, err error) {
	height := len(matrix)
	width := len(matrix[0])
	var strs []string
	// accumulating horizontally
	for y := 0; y < height; y++ {
		s := ""
		r := ""
		for x := 0; x < width; x++ {
			c := string(matrix[y][x])
			s = s + c
			r = c + r
		}
		strs = append(strs, s, r)
	}
	// accumulating vertically
	for x := 0; x < width; x++ {
		s := ""
		r := ""
		for y := 0; y < height; y++ {
			c := string(matrix[y][x])
			s = s + c
			r = c + r
		}
		strs = append(strs, s, r)
	}
	// accumulating diagonal1
	for x := 0; x < width; x++ {
		s := ""
		r := ""
		for y := 0; y < width-x; y++ {
			c := string(matrix[y][y+x])
			s = s + c
			r = c + r
		}
		strs = append(strs, s, r)
	}
	for x := 0; x < width-1; x++ {
		s := ""
		r := ""
		for y := 1; y < height-x; y++ {
			c := string(matrix[y+x][y-1])
			s = s + c
			r = c + r
		}
		strs = append(strs, s, r)
	}
	// accumulating diagonal2
	for x := 0; x < width; x++ {
		s := ""
		r := ""
		for y := 0; y < width-x; y++ {
			c := string(matrix[height-1-y][y+x])
			s = s + c
			r = c + r
		}
		strs = append(strs, s, r)
	}
	for x := 0; x < width-1; x++ {
		s := ""
		r := ""
		for y := 1; y < height-x; y++ {
			c := string(matrix[height-1-y-x][y-1])
			s = s + c
			r = c + r
		}
		strs = append(strs, s, r)
	}

	re := regexp.MustCompile(`XMAS`)
	for _, str := range strs {
		res += len(re.FindAllString(str, -1))
	}

	return
}

func SolveDay4Part2(matrix [][]rune) (res int, err error) {
	height := len(matrix)
	width := len(matrix[0])
	patterns := [][][]rune{
		{
			{'M', '.', 'M'},
			{'.', 'A', '.'},
			{'S', '.', 'S'},
		}, {
			{'M', '.', 'S'},
			{'.', 'A', '.'},
			{'M', '.', 'S'},
		}, {
			{'S', '.', 'S'},
			{'.', 'A', '.'},
			{'M', '.', 'M'},
		}, {
			{'S', '.', 'M'},
			{'.', 'A', '.'},
			{'S', '.', 'M'},
		},
	}
	match := func(pattern [][]rune, x [][]rune) bool {
		for i, _ := range pattern {
			for j, r := range pattern[i] {
				if r != '.' && x[i][j] != r {
					return false
				}
			}
		}
		return true
	}
	for x := 0; x <= width-3; x++ {
		for y := 0; y <= height-3; y++ {
			window := [][]rune{
				matrix[y+0][x : x+3],
				matrix[y+1][x : x+3],
				matrix[y+2][x : x+3],
			}
			for _, pattern := range patterns {
				if match(pattern, window) {
					res += 1
				}
			}
		}
	}
	return
}
