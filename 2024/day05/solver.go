package day05

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

var (
	OrderingRuleRegex = regexp.MustCompile(`^(\d+)\|(\d+)$`)
	UpdateRegex       = regexp.MustCompile(`\d+`)
)

type Page struct {
	// could be optimized using equivalent of set instead of array
	After []int
}

func (p *Page) IsBefore(page int) bool {
	for _, a := range p.After {
		if a == page {
			return true
		}
	}
	return false
}

type PageOrdering map[int]*Page

func (po PageOrdering) Insert(before int, after int) {
	if ptr, ok := po[before]; ok {
		ptr.After = append(ptr.After, after)
	} else {
		po[before] = &Page{After: []int{after}}
	}
	// ensure page after is created
	if _, ok := po[after]; !ok {
		po[after] = &Page{}
	}
}

func Day05Read(input string) (pageOrdering PageOrdering, updates [][]int, err error) {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	pageOrdering = make(PageOrdering)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		switch {
		case OrderingRuleRegex.MatchString(line):
			s := OrderingRuleRegex.FindAllStringSubmatch(line, -1)
			before, _ := strconv.Atoi(s[0][1])
			after, _ := strconv.Atoi(s[0][2])
			pageOrdering.Insert(before, after)
		default:
			u := UpdateRegex.FindAllString(line, -1)
			if u == nil {
				continue
			}
			var pages []int
			for _, page := range u {
				p, _ := strconv.Atoi(page)
				pages = append(pages, p)
			}
			updates = append(updates, pages)
		}
	}
	return
}

func SolveDay05Part1(pageOrdering PageOrdering, updates [][]int) (res int, err error) {
	valid, _ := CheckUpdates(pageOrdering, updates)
	// compute checksum
	for _, update := range valid {
		res += update[len(update)/2]
	}
	return
}

func SolveDay05Part2(pageOrdering PageOrdering, updates [][]int) (res int, err error) {
	_, invalid := CheckUpdates(pageOrdering, updates)
	var corrected [][]int
	for _, update := range invalid {
		for i := 0; i < len(update)-1; i++ {
			if !pageOrdering[update[i]].IsBefore(update[i+1]) {
				// swap
				update[i], update[i+1] = update[i+1], update[i]
				// go fix backward
				for j := i - 1; j >= 0; j-- {
					if !pageOrdering[update[j]].IsBefore(update[j+1]) {
						// swap
						update[j], update[j+1] = update[j+1], update[j]
					} else {
						break
					}
				}

			}
		}
		corrected = append(corrected, update)
	}
	// compute checksum
	for _, update := range corrected {
		res += update[len(update)/2]
	}
	return
}

func CheckUpdates(pageOrdering PageOrdering, updates [][]int) (valid [][]int, invalid [][]int) {
updateLoop:
	for _, update := range updates {
		for i := 0; i < len(update)-1; i++ {
			before := update[i]
			after := update[i+1]
			if !pageOrdering[before].IsBefore(after) {
				invalid = append(invalid, update)
				continue updateLoop
			}
		}
		valid = append(valid, update)
	}
	return
}
