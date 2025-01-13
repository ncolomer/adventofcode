package day14

import (
	"bufio"
	"strings"

	"github.com/ncolomer/adventofcode/2024/util"
)

const (
	Wall  = '#'
	Robot = '@'
	Box   = 'O'
)

func Read(input string) (m util.Map[rune], d []util.Direction) {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	var runes [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		runes = append(runes, []rune(line))
	}
	m = util.NewMap[rune](runes)
	for scanner.Scan() {
		for _, r := range scanner.Text() {
			d = append(d, util.DirectionFrom(r))
		}
	}
	return
}

func Move(m util.Map[rune], robot util.Coordinate, move util.Direction) {
}

func findRobot(m util.Map[rune]) util.Coordinate {
	for c, v := range m.Iterate() {
		if v == Robot {
			return c
		}
	}
	panic("robot not found")
}

func SolvePart1(m util.Map[rune], moves util.Direction) int {
	robot := findRobot(m)
	for move := range moves {
		var position util.Coordinate
		switch move {
		case util.DirectionUp:
			position = robot.MoveY(-1)
		case util.DirectionDown:
			position = robot.MoveY(1)
		case util.DirectionLeft:
			position = robot.MoveX(-1)
		case util.DirectionRight:
			position = robot.MoveX(1)
		}
		if !m.IsValid(position) {
			panic("unexpected")
		}
		switch m.At(position) {
		case Wall:
		case Box:
		}
	}
	return 0
}

func SolvePart2(m util.Map[rune], move util.Direction) int {
	panic("unimplemented")
}
