package day7

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

const (
	OperatorAdd Operator = iota
	OperatorMultiply
)

type Operator int

type Equation struct {
	TestValue int64
	Numbers   []int64
}

var (
	TestValueRegex = regexp.MustCompile(`^(\d+):(.*)$`)
	NumberRegex    = regexp.MustCompile(`\d+`)
)

func Read(input string) (e []Equation, err error) {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		parsed := TestValueRegex.FindStringSubmatch(line)
		testValue, _ := strconv.ParseInt(parsed[1], 10, 64)
		var numbers []int64
		for _, match := range NumberRegex.FindAllString(parsed[2], -1) {
			number, _ := strconv.ParseInt(match, 10, 64)
			numbers = append(numbers, number)
		}
		e = append(e, Equation{testValue, numbers})
	}
	return
}

func SolvePart1(eqs []Equation) (res int64, err error) {
	var IsPossiblyTrue func(test int64, acc int64, numbers []int64) bool
	IsPossiblyTrue = func(test int64, acc int64, numbers []int64) bool {
		if len(numbers) == 0 {
			return test == acc
		}
		return IsPossiblyTrue(test, acc+numbers[0], numbers[1:]) || IsPossiblyTrue(test, acc*numbers[0], numbers[1:])
	}
	for _, eq := range eqs {
		if IsPossiblyTrue(eq.TestValue, 0, eq.Numbers) {
			res += eq.TestValue
		}
	}
	return
}

func SolvePart2(eqs []Equation) (res int64, err error) {
	var IsPossiblyTrue func(test int64, acc int64, numbers []int64) bool
	IsPossiblyTrue = func(test int64, acc int64, numbers []int64) bool {
		if len(numbers) == 0 {
			return test == acc
		}
		concatenated, _ := strconv.ParseInt(strconv.FormatInt(acc, 10)+strconv.FormatInt(numbers[0], 10), 10, 64)
		return IsPossiblyTrue(test, acc+numbers[0], numbers[1:]) || IsPossiblyTrue(test, acc*numbers[0], numbers[1:]) || IsPossiblyTrue(test, concatenated, numbers[1:])
	}
	for _, eq := range eqs {
		if IsPossiblyTrue(eq.TestValue, 0, eq.Numbers) {
			res += eq.TestValue
		}
	}
	return
}
