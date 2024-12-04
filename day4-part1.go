//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := parseInput()
	var x, y int
	var count int
	for y = range input {
		for x = range input[y] {
			if checkForward(input, x, y) {
				count++
			}
			if checkBackward(input, x, y) {
				count++
			}
			if checkUp(input, x, y) {
				count++
			}
			if checkDown(input, x, y) {
				count++
			}
			if checkForwardUp(input, x, y) {
				count++
			}
			if checkForwardDown(input, x, y) {
				count++
			}
			if checkBackwardUp(input, x, y) {
				count++
			}
			if checkBackwardDown(input, x, y) {
				count++
			}
		}
	}
	fmt.Println(count)
}

func parseInput() []string {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func checkForward(input []string, x, y int) bool {
	var sb strings.Builder
	if x+3 >= len(input[y]) {
		return false
	}
	sb.WriteByte(input[y][x])
	sb.WriteByte(input[y][x+1])
	sb.WriteByte(input[y][x+2])
	sb.WriteByte(input[y][x+3])
	return sb.String() == "XMAS"
}

func checkBackward(input []string, x, y int) bool {
	var sb strings.Builder
	if x-3 < 0 {
		return false
	}
	sb.WriteByte(input[y][x])
	sb.WriteByte(input[y][x-1])
	sb.WriteByte(input[y][x-2])
	sb.WriteByte(input[y][x-3])
	return sb.String() == "XMAS"
}

func checkUp(input []string, x, y int) bool {
	var sb strings.Builder
	if y-3 < 0 {
		return false
	}
	sb.WriteByte(input[y][x])
	sb.WriteByte(input[y-1][x])
	sb.WriteByte(input[y-2][x])
	sb.WriteByte(input[y-3][x])
	return sb.String() == "XMAS"
}

func checkDown(input []string, x, y int) bool {
	var sb strings.Builder
	if y+3 >= len(input) {
		return false
	}
	sb.WriteByte(input[y][x])
	sb.WriteByte(input[y+1][x])
	sb.WriteByte(input[y+2][x])
	sb.WriteByte(input[y+3][x])
	return sb.String() == "XMAS"
}

func checkForwardUp(input []string, x, y int) bool {
	var sb strings.Builder
	if x+3 >= len(input[y]) || y-3 < 0 {
		return false
	}
	sb.WriteByte(input[y][x])
	sb.WriteByte(input[y-1][x+1])
	sb.WriteByte(input[y-2][x+2])
	sb.WriteByte(input[y-3][x+3])
	return sb.String() == "XMAS"
}

func checkForwardDown(input []string, x, y int) bool {
	var sb strings.Builder
	if x+3 >= len(input[y]) || y+3 >= len(input) {
		return false
	}
	sb.WriteByte(input[y][x])
	sb.WriteByte(input[y+1][x+1])
	sb.WriteByte(input[y+2][x+2])
	sb.WriteByte(input[y+3][x+3])
	return sb.String() == "XMAS"
}

func checkBackwardUp(input []string, x, y int) bool {
	var sb strings.Builder
	if x-3 < 0 || y-3 < 0 {
		return false
	}
	sb.WriteByte(input[y][x])
	sb.WriteByte(input[y-1][x-1])
	sb.WriteByte(input[y-2][x-2])
	sb.WriteByte(input[y-3][x-3])
	return sb.String() == "XMAS"
}

func checkBackwardDown(input []string, x, y int) bool {
	var sb strings.Builder
	if x-3 < 0 || y+3 >= len(input) {
		return false
	}
	sb.WriteByte(input[y][x])
	sb.WriteByte(input[y+1][x-1])
	sb.WriteByte(input[y+2][x-2])
	sb.WriteByte(input[y+3][x-3])
	return sb.String() == "XMAS"
}
