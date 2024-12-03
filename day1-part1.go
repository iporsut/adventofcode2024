//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var left, right []int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		pair := strings.Fields(line)
		l, _ := strconv.Atoi(pair[0])
		r, _ := strconv.Atoi(pair[1])
		left = append(left, l)
		right = append(right, r)
	}
	slices.Sort(left)
	slices.Sort(right)
	var result int
	for i := range left {
		if left[i] > right[i] {
			result += left[i] - right[i]
		} else {
			result += right[i] - left[i]
		}
	}
	fmt.Println(result)
}
