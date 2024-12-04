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
			if isXMAS(input, x, y) {
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

func isXMAS(input []string, x, y int) bool {
	if x+2 >= len(input[y]) || y+2 >= len(input) {
		return false
	}

	return isSliceXMAS(sliceXMAS(input, x, y))
}

func sliceXMAS(input []string, x, y int) []string {
	var slice []string
	slice = append(slice, input[y][x:x+3])
	slice = append(slice, input[y+1][x:x+3])
	slice = append(slice, input[y+2][x:x+3])
	return slice
}

func isSliceXMAS(slice []string) bool {
	var sb1 strings.Builder
	var sb2 strings.Builder
	sb1.WriteByte(slice[0][0])
	sb1.WriteByte(slice[1][1])
	sb1.WriteByte(slice[2][2])
	sb2.WriteByte(slice[0][2])
	sb2.WriteByte(slice[1][1])
	sb2.WriteByte(slice[2][0])

	return (sb1.String() == "MAS" || sb1.String() == "SAM") && (sb2.String() == "MAS" || sb2.String() == "SAM")
}
