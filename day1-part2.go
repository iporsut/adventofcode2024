//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var right []int
	leftExists := make(map[int]bool)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		pair := strings.Fields(line)
		l, _ := strconv.Atoi(pair[0])
		r, _ := strconv.Atoi(pair[1])
		leftExists[l] = true
		right = append(right, r)
	}
	var result int
	for _, r := range right {
		if leftExists[r] {
			result += r
		}
	}
	fmt.Println(result)
}
