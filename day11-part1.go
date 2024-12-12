//go:build ignore

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func blink(n string) []string {
	if n == "0" {
		return []string{"1"}
	}
	if len(n)%2 == 0 {
		v, _ := strconv.Atoi(n[len(n)/2:])
		return []string{n[:len(n)/2], fmt.Sprint(v)}
	}

	v, _ := strconv.Atoi(n)
	return []string{fmt.Sprint(v * 2024)}
}

func main() {
	vs := strings.Fields(input)

	for i := 0; i < 25; i++ {
		var vss []string
		for _, v := range vs {
			vss = append(vss, blink(v)...)
		}

		fmt.Println(i, len(vss))
		vs = vss
	}
	fmt.Println(len(vs))
}

var ex = `125 17`

var input = `3028 78 973951 5146801 5 0 23533 857`

// var input = `0`
