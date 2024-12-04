//go:build ignore

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	pattern := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	input := string(b)
	mulCmds := pattern.FindAllString(input, -1)
	extractDigitsPattern := regexp.MustCompile(`\d{1,3}`)
	var sum int
	for _, cmd := range mulCmds {
		numStrs := extractDigitsPattern.FindAllString(cmd, -1)
		var a, b int
		a, _ = strconv.Atoi(numStrs[0])
		b, _ = strconv.Atoi(numStrs[1])
		sum += a * b
	}
	fmt.Println(sum)
}
