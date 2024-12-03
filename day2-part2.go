//go:build ignore

package main

import (
	"bufio"
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var count int
	for scanner.Scan() {
		rec := parseRecord(scanner.Text())
		recs := generateRemoveOne(rec)
		for _, rec := range recs {
			if validRecord(rec) {
				count++
				break
			}
		}
	}
	fmt.Println(count)
}

func parseRecord(line string) []int {
	fields := strings.Fields(line)
	rec := make([]int, 0, len(fields))
	for _, field := range fields {
		n, _ := strconv.Atoi(field)
		rec = append(rec, n)
	}
	return rec
}

func validRecord(rec []int) bool {
	if len(rec) < 2 {
		return false
	}

	diff := int(math.Abs(float64(rec[1] - rec[0])))
	dir := cmp.Compare(rec[0], rec[1])
	if dir == 0 {
		return false
	}
	if diff < 1 {
		return false
	}
	if diff > 3 {
		return false
	}
	for i := 2; i < len(rec); i++ {
		if cmp.Compare(rec[i-1], rec[i]) != dir {
			return false
		}
		diff := int(math.Abs(float64(rec[i] - rec[i-1])))
		if diff > 3 {
			return false
		}
		if diff < 1 {
			return false
		}
	}

	return true
}

func generateRemoveOne(rec []int) [][]int {
	recs := make([][]int, 0, len(rec))
	for i := 0; i < len(rec); i++ {
		cloned := slices.Clone(rec)
		recs = append(recs, slices.Delete(cloned, i, i+1))
	}
	return recs
}
