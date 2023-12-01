package advent

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Day2P1(input string) int {
	tot := 0
	for _, l := range Lines(input) {
		fmt.Println(l)
	}
	return tot
}

func Day2P2(input string) int {
	return 0
}

func TestDay2_Ex1(t *testing.T) {
	assert.Equal(t, Day2P1(exD2), 42)
}

func TestDay2_Ex2(t *testing.T) {
	assert.Equal(t, Day2P2(exD2), 42)
}

func TestDay2_P1_Actual(t *testing.T) {
	fmt.Printf("Day 2 Part 1: %d\n", Day2P1(inputD2))
}

func TestDay2_P2_Actual(t *testing.T) {
	fmt.Printf("Day 2 Part 2: %d\n", Day2P2(inputD2))
}

const exD2 = ``

const inputD2 = ``
