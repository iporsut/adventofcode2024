//go:build ignore

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type freq struct {
	val   string
	count int
}

func blink(f freq) []freq {
	if f.val == "0" {
		return []freq{{val: "1", count: f.count}}
	}
	if len(f.val)%2 == 0 {
		v, _ := strconv.Atoi(f.val[len(f.val)/2:])
		return []freq{{val: f.val[:len(f.val)/2], count: f.count}, {val: fmt.Sprint(v), count: f.count}}
	}

	v, _ := strconv.Atoi(f.val)
	return []freq{{val: fmt.Sprint(v * 2024), count: f.count}}
}

func main() {
	var freqs []freq

	vs := strings.Fields(input)
	for _, v := range vs {
		freqs = append(freqs, freq{val: v, count: 1})
	}

	m := make(map[string]int)
	for _, f := range freqs {
		if _, ok := m[f.val]; ok {
			m[f.val] += f.count
		} else {
			m[f.val] = f.count
		}
	}
	freqs = make([]freq, 0, len(m))
	for val, count := range m {
		freqs = append(freqs, freq{val: val, count: count})
	}

	for i := 0; i < 75; i++ {
		var vss []freq
		for _, f := range freqs {
			vss = append(vss, blink(f)...)
		}
		m := make(map[string]int)
		for _, f := range vss {
			if _, ok := m[f.val]; ok {
				m[f.val] += f.count
			} else {
				m[f.val] = f.count
			}
		}
		freqs = make([]freq, 0, len(m))
		for val, count := range m {
			freqs = append(freqs, freq{val: val, count: count})
		}
	}

	var count int
	for _, f := range freqs {
		count += f.count
	}
	fmt.Println(count)
}

var ex = `125 17`

var input = `3028 78 973951 5146801 5 0 23533 857`

// var input = `0`
