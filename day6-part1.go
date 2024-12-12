//go:build ignore

package main

import (
	"fmt"
	"strings"
)

const (
	up = iota
	right
	down
	left
)

func main() {
	dir := up
	var cx, cy int
	lines := strings.Split(input, "\n")
	var maps [][]rune
	for _, line := range lines {
		maps = append(maps, []rune(line))
	}
	my := len(maps) - 1
	mx := len(maps[0]) - 1
	for i := range maps {
		for j := range maps[i] {
			if maps[i][j] == '^' {
				cx, cy = j, i
				maps[i][j] = 'X'
			}
		}
	}
	move := func() {
		switch dir {
		case up:
			cy--
		case right:
			cx++
		case down:
			cy++
		case left:
			cx--
		}
	}
	isBlocked := func() bool {
		switch dir {
		case up:
			return maps[cy-1][cx] == '#'
		case right:
			return maps[cy][cx+1] == '#'
		case down:
			return maps[cy+1][cx] == '#'
		case left:
			return maps[cy][cx-1] == '#'
		}
		return false

	}
	isFinish := func() bool {
		switch dir {
		case up:
			return cy-1 < 0
		case right:
			return cx+1 > mx
		case down:
			return cy+1 > my
		case left:
			return cx-1 < 0
		}
		return false
	}
	turn := func() {

		switch dir {
		case up:
			dir = right
		case right:
			dir = down
		case down:
			dir = left
		case left:
			dir = up
		}

	}
	for {
		if isFinish() {
			break
		}
		if isBlocked() {
			turn()
			continue
		}
		move()
		maps[cy][cx] = 'X'
	}
	var count int
	for _, m := range maps {
		for _, n := range m {
			if n == 'X' {
				count++
			}
		}
	}
	fmt.Println(count)
}

var sampleInput = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

var input = `..#.....##......#............#..................................#.....#..............#................#.........................#.
..........................................................................................................#..........#...#.......#
..................#....................#.......................................#.....#....#............#.......#..................
......#...#..##....................................................................#...#.......#..........#..#.....#..............
..#..#.#.............#.#.........#.........##..................#..............#..........................#........................
#..........#.................#..........#...............................#..#.................#.......##..##..#.#..................
......................#......................................................#........................#......................#....
........................#.#.....................................#..............................#.......#.............#.....#......
#.................#...#......#..........#.......#.....#............................................#...........##.........#.....#.
...........................................#.....#............##...............#..............................#...................
.....#..............#......................................#..........#............#..................##...............#......###.
..............#..#.....................................##................#............................#............#......#.......
#........#...................................................#..........................................#.........................
.................................................................................#................................#....#..#..#....
.............................#.###.........#.#...............#..........................................##........................
......................................................................................................................#.........#.
...............#............#............#...#.........#............#.............................#.#........................#....
........................#...............#...............#........#.........#................#....#.#....................#.........
........#...............#...........................................................................................#.............
.................#..........................................................#.....................................................
........#.#.#...................#..#.......#...........#..#............#......#.....................#.....#.........#..#..........
..#.......#..................................................................................................#......#.............
................##...##......................#....................................................#..............................#
............#......................................#...........................#.#...#...............#.....................##.....
..........#...................#.....................#...........................................##................................
.#.......................#...............................................................#........................................
...................................................................#..#................#.....#....................#...............
...........#...........................................................#..........................................................
.....#....#..#..........................................................#..................#.....................##...............
...#................................#.....#...................#...............##...........#..............#.......................
.......##.....#.............#..#.......................................................................#......................#...
......................................#....#.............................................#...#....................#...............
.......#.......#...............#..#.....#..............................#.#......#...............#.................................
......................#.......#.......#..........................................................#................#...........#...
..............................#........................#.......#..#...............................#...............................
...........................#..........#..##.......................#................#..............................................
..............#............#...##........#.................................#................#....#........................#..#....
..............#.........#............................................#.............#......#......#.#....#.........................
.......................#.................................#..#..#..#..............................................#................
........................###..#...............................................................#.........#................#......#..
#........................................................................................#........................................
........................................#..............................##.....#.........................#...........#.............
.............#.........#...#................#..#..........#.....#...........................................................#.....
.......................................#......#.............................#.........#................#..........................
..........................................................#......................#...........#...#.......#........................
#.....................................#.#...#.............................................................#...........#........#..
............................#..................................................................................#............#.....
#.................#.................................................................#..............#...................#..........
........#...................................................................#...............................................#.....
.................#...#..................#.............................................................#.......#...................
...#.....#.................................................................................................#..........#...........
..........#...................................#....#................#......#........................................#.......#.....
.....................................#...............#................................................................#...........
.#............#................................................................#.............................................#....
...........#.........#.....................#...........#^.......................#.....................#..........#................
.....#........#.................................................................................................#.......#.........
..........#..........................#.......#......................#..............................#....#......................#..
...#.........................#.................#....#........#....................................................................
...............#...#........................................#................................#....#........................#.#....
...........#.......................................#......#....#.......................................................#.......#..
#.........#...............#..#..#..........#..............#...........................#...........#.....................#.........
...................#.......................................................#......................................................
.............#..............................................#.......................................#.........#...........#......#
....#......#...#.................................#.....#......#.................................#.................................
...............#...............#......................................................................#...........................
.......................#.........................................................................................................#
.............................#............................................................................#.......................
...................#..#.........#...........................#......#...............#......#.......................................
...#....................................#...................#..#...........................#............#...............#........#
...................#................................................................#............##...............................
........................#.#...................#.....................................................................#...#.........
...................#......#.......#.#......................................................#...............#......................
....#...........................#................#.............................#..#....................#........##.....#..........
.........................#......................................................#.#......................................#........
.........#.............#.................#........................................................................................
.#.#.......#..#....#.......#...#..#......................................................................#.#......................
..#.................#............................................................................#................................
.#................................#...............................................................................................
..........#............#..............................................#...#...............##...............#..................#...
................................#.........................#.......................................................................
........................................................##.......#........................#.....................#.................
......#......................................................#......#.........#........................................#.......#..
...........................................................#............#....#......#.............................................
..................................#...................................................#.................................#...#.....
.......#...#..........#............#.......#........................................#..#.................#........................
....................#........................................#........#...........................................................
.................................##..#............#.................................#.............#..................#....#.......
.#.................#..........#.....#...........#.......................................................#...................#.....
........................#.........................................................................................................
...#..........................................#...........................#.................................#.....................
............#......................................................#.................................#.........#..............#...
.............#........##.........#..................#...#.............#......#....##................#................#............
..................#...........................................................................................#.................#.
.....................#..#..#....................#........................#.........#................................#.............
..............................................#...................#............................#..............................#...
.......................#.#...................................#.............................#...................................#..
................................#...........#.............................#.............#.........#..............................#
...........................#................................#...................................................#.....#...........
..........#..........#.....#..........#..#............#..........#.................#.......#.......#........................#....#
...#.......#...#..............................................................................................#......#......#.....
...........#......................#..#.................#......#.........#.....#.............#............##................#....#.
.......#...........................................#..................#...............#.............................#.............
....#.......#.................................#..................................................................#................
..#..................................#...#............................#.........#....#..........#............................#....
....................................................#.................#...........................................#...............
............#......................................#.........#..................#.....#.#.#.#..#..................................
#..............#.......#...#.#..............#.....#.#.............................................................................
...........#.#.................................#...#........................................#................#.........#....#.....
......#.........................................#.........................#...........................................#...........
...#......#.................................#......#............................................................................#.
...#.................#....##................#...........................................#.........#....#..........................
...........#...#...#..................#...................................#......................................#......#.....#.#.
...................................#....................#................#..................................................#.....
................#.....................................#...........#......#........................................................
..............................#......#..#..............................................##........#......#......................#.#
#...........................#....#....##.....#...........................##..#...................#.....#........#.................
..........#....#.#..................#..........#..#................#.................#..................................##.......#
...................#.....................#......#...........................................................#..........#..........
........#.....#.............................#..............................................#......................................
......#............#.#.................................................................................................#..#.......
.#.........#..................................................................................#....#...........#.................#
........#.....#.......................#.........................................#...........................................#.....
...................#.......#..#......#.................##......................................#...................#..............
.............................#.................................#........#.......#..................................#.........#....
....#....#..........#....###..#...................................................#...#................................#.......#..
...................#...................................................#.......#....#........................#....................
...#................#...........#.......................#.....................................#...#......#......#.................
.......#........................#.........#.............................................................#...#...#.................
....................#.........#..............................#..........................#....................#..#.................
..#..........................#.#...................................#..................#..........#................#...............`