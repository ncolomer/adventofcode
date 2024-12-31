package day03

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

type Instruction struct {
	enabled  bool
	operandA int
	operandB int
}

func Day03Read(input string) (instructions []Instruction, err error) {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	re := regexp.MustCompile(`(do\(\)|don't\(\)|mul\((\d{1,3}),(\d{1,3})\))`)
	enabled := true
	for scanner.Scan() {
		line := scanner.Text()
		for _, match := range re.FindAllStringSubmatch(line, -1) {
			switch {
			case match[1] == "do()":
				enabled = true
			case match[1] == "don't()":
				enabled = false
			default:
				operandA, _ := strconv.Atoi(match[2])
				operandB, _ := strconv.Atoi(match[3])
				instructions = append(instructions, Instruction{enabled, operandA, operandB})
			}
		}
	}
	return
}

// SolveDay03Part1 compute sum of instructions' result
func SolveDay03Part1(instructions []Instruction) (res int, err error) {
	for _, i := range instructions {
		res += i.operandA * i.operandB
	}
	return
}

// SolveDay03Part2 compute sum of instructions' result if enabled
func SolveDay03Part2(instructions []Instruction) (res int, err error) {
	for _, i := range instructions {
		if i.enabled {
			res += i.operandA * i.operandB
		}
	}
	return
}
