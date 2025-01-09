package day13

import (
	"bufio"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Tuple struct {
	X, Y uint64
}

type Machine struct {
	ID      int
	ButtonA Tuple
	ButtonB Tuple
	Prize   Tuple
}

type Result struct {
	PushA uint64
	PushB uint64
}

func (r Result) Tokens() uint64 {
	return r.PushA*3 + r.PushB*1
}

func (r Result) String() string {
	return fmt.Sprintf("(PushA:%d PushB:%d Tokens:%d)", r.PushA, r.PushB, r.Tokens())
}

var (
	ButtonRegex = regexp.MustCompile(`^Button (A|B): X\+(\d+), Y\+(\d+)$`)
	PrizeRegex  = regexp.MustCompile(`^Prize: X=(\d+), Y=(\d+)$`)
)

func Read(input string) (m []*Machine) {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	var ID int
	for scanner.Scan() {
		ID += 1
		// ButtonA
		line := scanner.Text()
		parsed := ButtonRegex.FindStringSubmatch(line)
		X, _ := strconv.ParseInt(parsed[2], 10, 64)
		Y, _ := strconv.ParseInt(parsed[3], 10, 64)
		ButtonA := Tuple{uint64(X), uint64(Y)}
		// ButtonB
		scanner.Scan()
		line = scanner.Text()
		parsed = ButtonRegex.FindStringSubmatch(line)
		X, _ = strconv.ParseInt(parsed[2], 10, 64)
		Y, _ = strconv.ParseInt(parsed[3], 10, 64)
		ButtonB := Tuple{uint64(X), uint64(Y)}
		// Prize
		scanner.Scan()
		line = scanner.Text()
		parsed = PrizeRegex.FindStringSubmatch(line)
		X, _ = strconv.ParseInt(parsed[1], 10, 64)
		Y, _ = strconv.ParseInt(parsed[2], 10, 64)
		Prize := Tuple{uint64(X), uint64(Y)}
		m = append(m, &Machine{ID, ButtonA, ButtonB, Prize})
		// Scan next empty line
		scanner.Scan()
	}
	return
}

func SolvePart1(machines []*Machine) (total uint64) {
	for _, m := range machines {
		var plays []Result
		// generate every possible combination
		for a := uint64(0); a < 100; a++ {
			for b := uint64(0); b < 100; b++ {
				X := a*m.ButtonA.X + b*m.ButtonB.X
				Y := a*m.ButtonA.Y + b*m.ButtonB.Y
				if X == m.Prize.X && Y == m.Prize.Y {
					plays = append(plays, Result{a, b})
				}
			}
		}
		if len(plays) > 0 {
			best := plays[0]
			for _, play := range plays[1:] {
				if play.Tokens() < best.Tokens() {
					best = play
				}
			}
			fmt.Printf("machine %d best play %v amongst %v\n", m.ID, best, plays)
			total += best.Tokens()
		} else {
			fmt.Printf("machine %d no winning play\n", m.ID)
		}
	}
	return
}

func SolvePart2(machines []*Machine) (total uint64) {
	for _, m := range machines {
		AX := float64(m.ButtonA.X)
		AY := float64(m.ButtonA.Y)
		BX := float64(m.ButtonB.X)
		BY := float64(m.ButtonB.Y)
		PX := float64(m.Prize.X)
		PY := float64(m.Prize.Y)
		// solution of the system of 2 equations with 2 unknowns a and b
		b := (PY - PX*AY/AX) / (BY - BX*AY/AX)
		a := (PX - b*BX) / AX
		// rounding because of precision
		a = math.Round(a*100) / 100
		b = math.Round(b*100) / 100
		// testing for integer results
		if a == math.Trunc(a) && b == math.Trunc(b) {
			best := Result{uint64(a), uint64(b)}
			fmt.Printf("machine %d best play %v\n", m.ID, best)
			total += best.Tokens()
		} else {
			fmt.Printf("machine %d no winning play\n", m.ID)
		}
	}
	return
}
