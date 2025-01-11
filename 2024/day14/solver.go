package day14

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ncolomer/adventofcode/2024/util"
)

type Robot struct {
	ID, X, Y, VX, VY int
}

func (r *Robot) String() string {
	return fmt.Sprintf("%d: p=%v v=%v", r.ID, util.Coordinate{r.X, r.Y}, util.Coordinate{r.VX, r.VY})
}

var (
	RobotRegex = regexp.MustCompile(`^p=(\d+),(\d+) v=(-?\d+),(-?\d+)$`)
)

func Read(input string) (m []*Robot) {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	var ID int
	for scanner.Scan() {
		ID += 1
		line := scanner.Text()
		parsed := RobotRegex.FindStringSubmatch(line)
		X, _ := strconv.Atoi(parsed[1])
		Y, _ := strconv.Atoi(parsed[2])
		VX, _ := strconv.Atoi(parsed[3])
		VY, _ := strconv.Atoi(parsed[4])
		m = append(m, &Robot{ID, X, Y, VX, VY})
	}
	return
}

func Print(robots []*Robot, mapWidth, mapHeight int, middles bool) {
	matrix := make([][]rune, mapHeight)
	// initialize
	for i, _ := range matrix {
		matrix[i] = []rune(strings.Repeat(".", mapWidth))
	}
	// index
	coords := make(map[util.Coordinate]int)
	for _, robot := range robots {
		coords[util.Coordinate{X: robot.X, Y: robot.Y}] += 1
	}
	// update
	for p, n := range coords {
		matrix[p.Y][p.X] = []rune(strconv.Itoa(n))[0]
	}
	// drop middles
	if middles {
		matrix[mapHeight/2] = []rune(strings.Repeat(" ", mapWidth))
		for i, _ := range matrix {
			matrix[i][mapWidth/2] = ' '
		}
	}
	// display
	fmt.Println()
	for _, row := range matrix {
		fmt.Println(string(row))
	}
}

func Tick(robots []*Robot, mapWidth, mapHeight int) {
	for _, robot := range robots {
		robot.X = (robot.X + robot.VX) % mapWidth
		if robot.X < 0 {
			robot.X += mapWidth
		}
		robot.Y = (robot.Y + robot.VY) % mapHeight
		if robot.Y < 0 {
			robot.Y += mapHeight
		}
	}
}

func SolvePart1(robots []*Robot, mapWidth, mapHeight int) int {
	var TopLeft, TopRight, BottomLeft, BottomRight int
	for _, robot := range robots {
		switch {
		case robot.Y < mapHeight/2 && robot.X < mapWidth/2:
			// top left quadrant
			TopLeft += 1
		case robot.Y < mapHeight/2 && robot.X > mapWidth/2:
			// top right quadrant
			TopRight += 1
		case robot.Y > mapHeight/2 && robot.X < mapWidth/2:
			// bottom left quadrant
			BottomLeft += 1
		case robot.Y > mapHeight/2 && robot.X > mapWidth/2:
			// bottom right quadrant
			BottomRight += 1
		}
	}
	safetyFactor := TopLeft * BottomLeft * BottomRight * TopRight
	return safetyFactor
}

func SolvePart2(robots []*Robot, MapWidth, MapHeight int) int {
loop:
	for i := range 10000 {
		Tick(robots, MapWidth, MapHeight)
		coords := make(map[util.Coordinate]int)
		for _, robot := range robots {
			if _, ok := coords[util.Coordinate{X: robot.X, Y: robot.Y}]; ok {
				// we already found a robot at this position
				continue loop
			}
			coords[util.Coordinate{X: robot.X, Y: robot.Y}] += 1
		}
		iteration := i + 1
		return iteration
	}
	panic("unexpected")
}
