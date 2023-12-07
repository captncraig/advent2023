package advent

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Day6(input string) (int, int) {
	tot := 1
	times := []int{}
	dists := []int{}
	ls := Lines(input)
	for _, n := range strings.Split(ls[0], " ") {
		nn, _ := strconv.Atoi(n)
		times = append(times, nn)
	}
	for _, n := range strings.Split(ls[1], " ") {
		nn, _ := strconv.Atoi(n)
		dists = append(dists, nn)
	}
	fmt.Println(times, dists)
	waysToWin := func(t int, d int) int {
		sum := 0
		for j := 0; j < t; j++ {
			dist := (t - j) * j
			if dist > d {
				sum++
			}
		}
		return sum
	}
	for i, t := range times {
		d := dists[i]
		sum := waysToWin(t, d)
		tot *= sum
	}
	// part 2
	l0 := strings.ReplaceAll(ls[0], " ", "")
	l1 := strings.ReplaceAll(ls[1], " ", "")
	t2, _ := strconv.Atoi(l0)
	d2, _ := strconv.Atoi(l1)
	return tot, waysToWin(t2, d2)
}

func TestDay6_Ex1(t *testing.T) {
	a, b := Day6(exD6)
	assert.Equal(t, 288, a)
	assert.Equal(t, 71503, b)
}

func TestDay6_Actual(t *testing.T) {
	a, b := Day6(inputD6)
	fmt.Printf("Day 6 Part 1: %d\n", a)
	fmt.Printf("Day 6 Part 2: %d\n", b)
}

const exD6 = `7 15 30
9 40 200`

const inputD6 = `61 70 90 66
643 1184 1362 1041`
