//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lineNumber := 1
	fmt.Printf("INSERT INTO records (ID, Val) VALUES\n")
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		fields := strings.Fields(line)

		for _, f := range fields {
			fmt.Printf("(%d, %s),\n", lineNumber, f)
		}

		lineNumber++
	}
}
