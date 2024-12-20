//go:build ignore

package main

import (
	"fmt"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type pos struct{ x, y int }

func main() {
	lines := strings.Split(input, "\n")
	mx, my := len(lines[0]), len(lines)
	antennas := make(map[rune][]pos)
	var antinodes [][]rune
	for _, line := range lines {
		nodes := make([]rune, len(line))
		antinodes = append(antinodes, nodes)
	}
	for y, line := range lines {
		for x, r := range []rune(line) {
			if (r >= '0' && r <= '9') || (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'z') {
				p, ok := antennas[r]
				if !ok {
					p = make([]pos, 0)
				}
				p = append(p, pos{x, y})
				antennas[r] = p
			}
		}
	}

	for _, ps := range antennas {
		for i := 0; i < len(ps)-1; i++ {
			for j := i + 1; j < len(ps); j++ {
				dx := abs(ps[i].x - ps[j].x)
				dy := abs(ps[i].y - ps[j].y)
				antinodes[ps[i].y][ps[i].x] = '#'
				antinodes[ps[j].y][ps[j].x] = '#'

				if ps[i].x <= ps[j].x {
					var ax1, ay1, ax2, ay2 int
					for c := 1; ; c++ {
						ax1, ay1 = ps[i].x-dx*c, ps[i].y-dy*c
						if !(ax1 >= 0 && ax1 < mx && ay1 >= 0 && ay1 < my) {
							break
						}

						antinodes[ay1][ax1] = '#'
					}

					for c := 1; ; c++ {
						ax2, ay2 = ps[j].x+dx*c, ps[j].y+dy*c
						if !(ax2 >= 0 && ax2 < mx && ay2 >= 0 && ay2 < my) {
							break
						}

						antinodes[ay2][ax2] = '#'
					}

				} else {
					var ax1, ay1, ax2, ay2 int
					for c := 1; ; c++ {
						ax1, ay1 = ps[i].x+dx*c, ps[i].y-dy*c
						if !(ax1 >= 0 && ax1 < mx && ay1 >= 0 && ay1 < my) {
							break
						}

						antinodes[ay1][ax1] = '#'
					}

					for c := 1; ; c++ {
						ax2, ay2 = ps[j].x-dx*c, ps[j].y+dy*c
						if !(ax2 >= 0 && ax2 < mx && ay2 >= 0 && ay2 < my) {
							break
						}

						antinodes[ay2][ax2] = '#'
					}
				}
			}
		}
	}

	var count int
	for _, row := range antinodes {
		for _, e := range row {
			if e == '#' {
				count++
			}
		}
	}
	fmt.Println(count)
}

var sampleInput = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

var input = `..........................4..............7..q.....
..........G..42.f......K.........7................
D.t...S......A....................................
..K.................................I.............
G....D...f.tA..H.S..........o................N....
t....f..............4..A........B.........N.....q.
...b...k....f..h..........6.......................
..........b....m................7...............Q.
....h....G.2........K.i...........................
.F...2.....D....H..6..o........I..................
k.......b..................K......I.....e.....B...
.............Sp..o....n....R.............N........
F............d................2...................
.........i........................................
.....ma.....d......p.Q..n.....7....9..........N...
......m..H......S...8......n.....Q...e............
.i..............8......O.....I................c...
..d......k....R.....................9....z........
..p.......m......n...............P................
.......pLb...................W..j................q
.....C..1..........u.....c.....jO...Z..o.........V
..C.....i........X1......9......e....j.....B....c.
......................9...........Q..Z............
.d....h..L...............8........O...............
....C....r..L....R...............6................
...........h.............1.t......P.......V.......
.......L.1........................................
..................................................
X.......................................V.....W...
rx........a.X.......0....l..........6.........z...
..r........a.8.................................z..
................w.........l..............P....A...
..........E....s..w.j........l...............W....
...v...............c..............W..y...V.O......
.....X..g.Y...0w......l...................u.......
.C.......Y...0....................................
...g..UJ...0........v.............................
.U...aY...........................................
....5........Y....MUJ..........B..................
.......g...5M........J.......w.........u..Z.......
................TE................................
..U....r....5.................J..........Z........
.......5...3......s........T......................
.............E.T..............................u...
...........v........y.......................P.....
................s.................................
x............M3........e..........................
........3...v......MT.............................
.............x....................................
....x..........3............y.....................`
