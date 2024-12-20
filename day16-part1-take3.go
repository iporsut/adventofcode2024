//go:build ignore

package main

import (
	"fmt"
	"strings"
)

var globalMinCost = -1

type costOk struct {
	cost int
	ok   bool
}

type visitedNode struct {
	x, y int
	dir  int
}

// cache visited nodes cost
var visitedNodes = make(map[visitedNode]costOk)

func main() {
	m, sx, sy, ex, ey := parseMap(input)
	cost, ok := findPath(m, sx, sy, ex, ey, 1, 0, make(map[int]map[int]bool), make(map[int]bool))

	fmt.Println(cost, ok)
	for k, v := range visitedNodes {
		if v.ok {
			fmt.Println(k, v)
		}
	}
}

func cloneVisitedPath(m map[int]map[int]bool) map[int]map[int]bool {
	cloned := make(map[int]map[int]bool)
	for k, v := range m {
		cloned[k] = make(map[int]bool)
		for kk, vv := range v {
			cloned[k][kk] = vv
		}
	}

	return cloned
}

func cloneTurnedPath(m map[int]bool) map[int]bool {
	cloned := make(map[int]bool)
	for k, v := range m {
		cloned[k] = v
	}
	return cloned
}

func clonePaths(m []points) []points {
	var cloned []points
	for _, p := range m {
		cloned = append(cloned, p)
	}
	return cloned
}

type points struct {
	x, y int
}

func findPath(m [][]rune,
	sx, sy,
	ex, ey int,
	currentDir int,
	cost int,
	visitedPath map[int]map[int]bool,
	turnedPath map[int]bool,
) (int, bool) {
	if globalMinCost != -1 && cost > globalMinCost {
		fmt.Println("cost", globalMinCost, cost)
		return 0, false
	}

	if visitedPath[sy] == nil {
		visitedPath[sy] = make(map[int]bool)
	}
	visitedPath[sy][sx] = true

	if sx == ex && sy == ey {
		if globalMinCost == -1 || cost < globalMinCost {
			globalMinCost = cost
			fmt.Println("cost-done", cost)
		}
		return cost, true
	}

	var costs []costOk

	// up
	if currentDir == 0 && m[sy-1][sx] != '#' && !visitedPath[sy-1][sx] {
		if cc, ok := visitedNodes[visitedNode{sx, sy - 1, 0}]; ok {
			costs = append(costs, cc)
		} else {
			clonedVisitedPath := cloneVisitedPath(visitedPath)
			newTurnedPath := make(map[int]bool)
			c, ok := findPath(m, sx, sy-1, ex, ey, 0, cost+1, clonedVisitedPath, newTurnedPath)
			costs = append(costs, costOk{c, ok})
			if ok {
				visitedNodes[visitedNode{sx, sy - 1, 0}] = costOk{c, ok}
			}
		}
	}

	// right
	if currentDir == 1 && m[sy][sx+1] != '#' && !visitedPath[sy][sx+1] {
		if cc, ok := visitedNodes[visitedNode{sx + 1, sy, 1}]; ok {
			costs = append(costs, cc)
		} else {
			clonedVisitedPath := cloneVisitedPath(visitedPath)
			newTurnedPath := make(map[int]bool)
			c, ok := findPath(m, sx+1, sy, ex, ey, 1, cost+1, clonedVisitedPath, newTurnedPath)
			costs = append(costs, costOk{c, ok})
			if ok {
				visitedNodes[visitedNode{sx + 1, sy, 1}] = costOk{c, ok}
			}
		}
	}

	// down
	if currentDir == 2 && m[sy+1][sx] != '#' && !visitedPath[sy+1][sx] {
		if cc, ok := visitedNodes[visitedNode{sx, sy + 1, 2}]; ok {
			costs = append(costs, cc)
		} else {
			clonedVisitedPath := cloneVisitedPath(visitedPath)
			newTurnedPath := make(map[int]bool)
			c, ok := findPath(m, sx, sy+1, ex, ey, 2, cost+1, clonedVisitedPath, newTurnedPath)
			costs = append(costs, costOk{c, ok})
			if ok {
				visitedNodes[visitedNode{sx, sy + 1, 2}] = costOk{c, ok}
			}
		}
	}

	// left
	if currentDir == 3 && m[sy][sx-1] != '#' && !visitedPath[sy][sx-1] {
		if cc, ok := visitedNodes[visitedNode{sx - 1, sy, 3}]; ok {
			costs = append(costs, cc)
		} else {
			clonedVisitedPath := cloneVisitedPath(visitedPath)
			newTurnedPath := make(map[int]bool)
			c, ok := findPath(m, sx-1, sy, ex, ey, 3, cost+1, clonedVisitedPath, newTurnedPath)
			costs = append(costs, costOk{c, ok})
			if ok {
				visitedNodes[visitedNode{sx - 1, sy, 3}] = costOk{c, ok}
			}
		}
	}

	// turn to up
	if (currentDir == 1 || currentDir == 3) && m[sy-1][sx] != '#' && !visitedPath[sy-1][sx] && !turnedPath[0] {
		if cc, ok := visitedNodes[visitedNode{sx, sy, 0}]; ok {
			costs = append(costs, cc)
		} else {
			clonedVisitedPath := cloneVisitedPath(visitedPath)
			turnedPath[0] = true
			c, ok := findPath(m, sx, sy, ex, ey, 0, cost+1000, clonedVisitedPath, turnedPath)
			costs = append(costs, costOk{c, ok})
			if ok {
				visitedNodes[visitedNode{sx, sy, 0}] = costOk{c, ok}
			}
		}
	}

	// turn to right
	if (currentDir == 0 || currentDir == 2) && m[sy][sx+1] != '#' && !visitedPath[sy][sx+1] && !turnedPath[1] {
		if cc, ok := visitedNodes[visitedNode{sx, sy, 1}]; ok {
			costs = append(costs, cc)
		} else {
			clonedVisitedPath := cloneVisitedPath(visitedPath)
			turnedPath[1] = true
			c, ok := findPath(m, sx, sy, ex, ey, 1, cost+1000, clonedVisitedPath, turnedPath)
			costs = append(costs, costOk{c, ok})
			if ok {
				visitedNodes[visitedNode{sx, sy, 1}] = costOk{c, ok}
			}
		}
	}

	// turn to down
	if (currentDir == 1 || currentDir == 3) && m[sy+1][sx] != '#' && !visitedPath[sy+1][sx] && !turnedPath[2] {
		if cc, ok := visitedNodes[visitedNode{sx, sy, 2}]; ok {
			costs = append(costs, cc)
		} else {
			clonedVisitedPath := cloneVisitedPath(visitedPath)
			turnedPath[2] = true
			c, ok := findPath(m, sx, sy, ex, ey, 2, cost+1000, clonedVisitedPath, turnedPath)
			costs = append(costs, costOk{c, ok})
			if ok {
				visitedNodes[visitedNode{sx, sy, 2}] = costOk{c, ok}
			}
		}
	}

	// turn to left
	if (currentDir == 0 || currentDir == 2) && m[sy][sx-1] != '#' && !visitedPath[sy][sx-1] && !turnedPath[3] {
		if cc, ok := visitedNodes[visitedNode{sx, sy, 3}]; ok {
			costs = append(costs, cc)
		} else {
			clonedVisitedPath := cloneVisitedPath(visitedPath)
			turnedPath[3] = true
			c, ok := findPath(m, sx, sy, ex, ey, 3, cost+1000, clonedVisitedPath, turnedPath)
			costs = append(costs, costOk{c, ok})
			if ok {
				visitedNodes[visitedNode{sx, sy, 3}] = costOk{c, ok}
			}
		}
	}

	if len(costs) == 0 {
		return 0, false
	}

	var minCost costOk
	for _, c := range costs {
		if !minCost.ok && c.ok {
			minCost = c
			continue
		}

		if c.cost < minCost.cost && c.ok {
			minCost = c
		}
	}

	if minCost.ok && minCost.cost < globalMinCost {
		globalMinCost = minCost.cost
		return minCost.cost, minCost.ok
	}

	return 0, false
}

func parseMap(input string) ([][]rune, int, int, int, int) {
	var m [][]rune
	var sx, sy, ex, ey int
	for y, l := range strings.Split(input, "\n") {
		m = append(m, []rune(l))
		for x, r := range []rune(l) {
			if r == 'S' {
				sx, sy = x, y
			}
			if r == 'E' {
				ex, ey = x, y
			}
		}
	}
	return m, sx, sy, ex, ey
}

var ex1 = `###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`

var ex2 = `#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################`

var input = `#############################################################################################################################################
#.#...............#.........#.....#...........#.............#.............#.#.......#.#.........#...........#.......#...........#.........#E#
#.#.#.###########.#.#.#####.#.#.#.###.#####.#.#.#########.#.#.###.###.###.#.#.###.#.#.#.#.###.#.#.#######.#.#####.#.#.#.#.#.###.#.#.#####.#.#
#.#.#.....#.....#...#.....#.#...............#.#...#...#.....#.....#...#.#.#.#.#...#...........#.#.#.....#.#.#...#.#.#.#.#...#...#.#.#...#...#
#.#.#####.#.#######.#.###.#.#.#.#######.###.#.#.#.#.#.#.###########.###.#.#.#.#.###.#####.###.###.#.###.#.#.#.#.#.#.#.#.#.#.#.###.#.#.#####.#
#...#...#.#.....#...#...#.#.....#.....#.#.#...#.#.#.#.#.#.....#.....#.......#.#...#.#...#...#.....#...#.#.#...#.#.#...#.#.#.#.#...#.#.......#
#.###.#.#.#.###.#.#.#####.###.#####.#.#.#.#.###.#.#.#.#.###.#.#.#########.###.###.#.#.#.###.###########.#.#####.#.#######.#.#.#####.#.#######
#...#.#.#.#...#.#.........#.#.......#.#.#.......#.#.#.......................#.#.#.#.#.#...#...#.......#.#.#.....#.......#.#.#...#...#.#.....#
###.#.#.#.###.#.###.#######.#.#.###.###.#.#.#####.#####.#.###.#.#########.#.#.#.#.#.#.###.#.#.#.#####.#.#.#.#####.#####.#.#.###.#.###.#.###.#
#...#.#.#.#.#.#.....#.......#.#.#...#...#.#.....#.....#...#.#.#...........#.....#.#...#.#.#.....#.......#.#.#...#.....#...#.#.#.#...#...#.#.#
#.###.#.#.#.#.###.#.#.#.#####.#.#####.###.#.###.#####.#####.#.###.#######.#######.#####.#.#.#############.#.###.#####.#####.#.#.###.###.#.#.#
#.#...#...#.........................#...#.#...#.#.....#.....#...#.#...........................#...#.................#...#...#.#...#.......#.#
#.#.###.#######.#.#.#######.#.#####.###.###.#.#.#.#####.#######.###.#.###.#.###.#.#.###.#######.#.#.###.#.###.###.#.###.#.#.#.###.#########.#
#.#.#...#.....#.#.........#.....#.#...#...#...#...#...#.......#...#.#.......#.#.#.#...#.........#...#.....#.#.............#.......#.....#...#
#.#.#####.###.#.###.#####.#.#.#.#.###.###.#.#########.#.#####.###.#.#####.###.#.#.###.#.#############.#####.###.#.###.#.#.###.#####.###.#.###
#.........#.#.#...#...#...#.........#...#...#.#.......#.....#...#.#.#...#...#...#.......#.........#...#...#.....#...#...#...#...#...#...#...#
#.#########.#.###.###.#.#####.#######.#.#####.#.###.#######.#.###.#.#.###.#.#.#######.#.#.#######.#####.#.###.###########.#.###.#.###.###.#.#
#.#.#.....#.#.#.......#...#.........#.#.........#...#.....#.#...#...#...#.#.#.#.....#.#.........#.......#...#...........#.#.#.....#.#.#.#.#.#
#.#.#.#.#.#.#.#.###########.#####.#.#.###########.###.###.#.###.#####.#.#.#.#.#.#.#.#.###########.#########.###########.#.#.#.#####.#.#.#.#.#
#.#...#.#...#.#.#.........#.#.....#.#...#.....#...#.....#.#...#...#...#.#.#...#.#.#.#...#.......#.#...#...#.#...#.....#.#.#.#.......#.#...#.#
#.#####.###.#.###.#######.#.#.###.#.#####.###.#.#.#.#####.###.###.#.###.#.#######.#.#####.#####.###.#.#.#.#.#.#.#.#.#.#.###.#.###.###.#####.#
#.......#.#.#...........#...#.#...#.....#...#...#.#...#...#...#.#.#...#...#.....#.#.#.....#.....#...#.#...#.#.#...#...#.#...#...#...#.....#.#
#########.#.#############.###.#######.#.#.#.#########.#.###.###.#.###.#######.#.#.#.#.#####.#.###.###.#.#.#.#.#####.#.#.#.#######.#.#####.#.#
#.#.....#...#...........#.....#.....#.#.#.#.........#.#.....#...#.#...#.....#.#...#.#.....#...#.#...#.#.#...#.#.....#.#.#.......#.#.....#...#
#.#.#.###.###.#.#######.#######.###.###.#.#.#.#####.#.#######.#.#.#.###.###.#.#####.#.#.###.###.#.#.#.#.#####.#.#######.#######.#.###.#####.#
#...#...#.....#.....#.#.........#.#.#...#.....#...#.........#.#.#.......#.#...#...#.#.#...#.#.....#.#...#...#.#.......#.......#.#.#.#.......#
#.#####.#.###.#####.#.###########.#.#.###.#####.#.#########.#.#.#######.#.#####.#.#.#.###.#.#.#####.###.#.#.#.#######.#######.#.#.#.#########
#.#...#.#.......#...#.....#...#.......#.....#...#...#...#...#.#.........#.#.....#...#...#.#.#.....#...#...#.....#...#...#...#...#.........#.#
#.#.#.#.#.#####.#.#######.#.#.#.#############.###.#.###.#.#####.###.#####.#.###########.#.#.#########.#.#######.#.#.###.#.#.#############.#.#
#.#.#...#.....#.#...#...#...#...#.....#.......#...#...#.......#.#.#.#.#...#.#.......#...#.#...#...#...#.#.....#.#.#...#...#.......#.....#...#
###.#####.###.#.###.#.#.#.#######.###.#.#######.#####.#######.#.#.#.#.#.#.#.#.#####.#.###.###.#.#.#.#####.###.###.###.#####.#####.#.#.#####.#
#...#.....#.#...#.#...#.#...#.....#.#...#.......#...#...#...#.#.....#.......#...#...#...#.....#.#.#.#.......#...#...#.....#.........#.......#
#.#.#.###.#.#####.#####.###.#.#####.#######.#####.#.#.#.#.#.#.#####.#.#########.#.###.#.###.#.#.#.#.#.#####.#.#.###.#####.#######.#.#.#######
#.#.#.....#.#.........#.#.#.#...#.#.....#...#.....#.#.#.#.#.#.......#.#.........#...#.#...#.....#.#...#...#...#.........#.......#.#.........#
#.#######.#.#.#######.#.#.#.###.#.#.#.#.#.#.#.#####.#.#.###.#####.###.#######.#####.#####.#######.#####.#.###############.#####.#.#######.#.#
#.#.......#...#...#.....#...#...#...#.#...#.#...#.#.#.#...#.........#.......#...#.#.......#.....#.......#...........#.......#...#.........#.#
#.#.#.###.#.#####.#.#.#.#.###.#######.#####.###.#.#.#####.#####.#.#.#######.###.#.#########.#######################.#.#.###.#.###.#######.#.#
#...#...#.#.......#.#.#.#...#.......#...#...#...#.#.#...#.....#.#.#.......#...#.................#.........#.......#...#...#.#...#.......#.#.#
#.###.#.#.#######.#.###.###########.#.#.#.###.###.#.#.#.#.###.#.#.#.#####.###.#.#.###########.#.#.###.###.#.#####.#######.###.#.###.###.#.#.#
#.....#.#.........#...#.....#.....#.#.#.#.#...#...#...#.#.#...#...#.....#.#.#...#.#...........#...#...#.....#.#.........#.....#...#...#.....#
#####.#.#.#########.#.#####.#.###.#.#.#.###.###.#.#####.###.#####.#####.#.#.#######.#########.#####.###.#####.#.#####.#####.#.###.#.#####.###
#.....#.#...#.#...#.......#.....#.#.#.#.#...#...#.....#.....#...#.#.#...#...#.......#.....#.........#.#.#...#...#...#.#...#.#...........#...#
#.###.#.#.#.#.#.#.###.#.#.#######.#.###.#.###.#.###.#########.###.#.#.#####.#.#####.#####.#.#########.#.#.#.#####.#.#.#.#.#.###.#.#.#.#.#.#.#
#.#...#...#...#.#...#.#.#.....#...#...#.#...#.#...........#.....#.#.#.#.......#.........#.....#.....#...#.#.......#.#.#.#.....#.#.#.#.#...#.#
#.###.###.#####.###.#.#.#####.#.#####.#.###.#.#.#######.###.###.#.#.#.#########.#.#.###.#######.###.#.###.#########.###.#####.#.#.###.###.#.#
#...#...#.#.....#...#...#...#.#.....#...#...#.....#.......#.#.#.#...#.#.......#.#...#.#...........#.#...#.#.......#.....#.#.....#...#.#...#.#
#.#.#.###.#.#####.#####.#.#.#######.#.###.#######.#.#.###.#.#.#.###.#.#.#####.#.#####.###.#########.#####.#######.#######.#.#######.#.#.###.#
#.#.......#.....#.#.....#.#.....#.#.#.#...#.....#...#.#.#.#...#.....#...#.#...#...#...#.......#.....#...#.........#.......#...#...#.#.#...#.#
#.###.#####.###.#.###.###.###.#.#.#.#.#.###.###.###.#.#.#.###.###########.#.#.###.#.###.#####.#.###.#.#.#########.#.#####.###.#.#.#.#.###.#.#
#...#.....#.#...#...#.....#...#...#...#...#.#.#.#...#...#...................#...#.#.#...#.#...#...#.#.#.....#...#.#...#.....#...#.#.........#
###.#.###.###.#####.#.#####.#.###########.#.#.#.#.#####.###########.#.#####.###.#.#.#.###.#.#.###.#.#.#####.#.#.#.###.#.#.#######.#.#.#.#.###
#...#...#.....#...#.#.#.....#.#.......#...#.....#...#.......#...#...#.#.....#...#.#.#...#.#.#.#...#.#...#...#.#.#...#...#.......#.#.....#...#
#.###.#.#######.###.###.#######.#####.#.###.###.#.#.#.#####.#.#.#.#####.###.#.###.#.###.#.#.###.#######.#.###.#.###.#.#######.#.#.###.#####.#
#.#...#.#.....#.....#...#.......#.....#.#...#...#.#...#...#...#...#...#.#.#.#.....#.....#.#...#.........#.....#...#.#.....#...#...#.......#.#
#.#####.#####.#.#####.###.#.#.#####.###.#.###.###.#####.#.#########.#.#.#.#.#.#####.#####.#.#.#####.#.###########.#.#####.#.#.#####.#.###.#.#
#.....#.#...#.......#.#...#...#...#.#...#.#.#...#.#.....#...........#...#.#.#.#...#.#.......#...#...#.....#.....#...#...#.....#.....#...#.#.#
#####.#.#.#.#.#####.#.#.#######.#.#.###.#.#.###.#.#.###.#########.#######.#.#.#.#.#.#######.#.###.###.###.#.#.#.#######.#####.#.###.#.#.###.#
#.....#.#.#.#.......#.#.....#...#.#...#.#.#.#...#.#.#...#...#...#.#...#.....#.#.#...........#.#...#.......#.#.#.............#.#.#.#...#.#...#
#####.#.#.#.#.###.###.#####.#.#.#.###.#.#.#.#.###.#.#.###.#.#.#.###.#.#######.#############.#.#.###.#######.#.###########.###.#.#.#####.#.###
#...#.#...#.#.#...........#.....#.#...#.......#.#.#.#.#...#...#.#...#.......#...........#...#.#.....#...#.#.#...#.........#...#.....#...#...#
#.#.#.#####.#.#####.#.#######.###.#.#########.#.#.#.###.###.###.#.#########.###########.#####.#######.#.#.#.###.#######.###.###.#####.#####.#
#.#.#.....#.#.....#.#.#.....#.#...#...........#...#.......#.#...#.....#.......#.......#.........#.........#...#.......#.#...#...#...#.......#
#.#.#.###.#.###.#.###.#.###.###.#####.#####.###.#########.#.#.#######.#.###.###.#####.#########.###.#########.#.#.###.#.#.###.###.#.#.#.#.#.#
#.#.#.#.....#.........#.#.#...#.....#...#.....#.#.....#.#...#.......#.#.#...#...#...#.......#.#.....#.........#...#.#.#.#.#...#...#...#.#.#.#
#.#.#.#.#############.#.#.###.#####.#.#.#.###.#.#.#.#.#.###.#######.#.#.###.#.###.#.###.###.#.#.#.###.#########.#.#.#.#.#.#####.###.###.#.#.#
#.#.#.....#.........#.#.#...#.#.....#.#...#...#...#.#.#...........#...#.......#.#.#...#.#.....#.#.#.......#...#.....#.#.#...#...#...#...#...#
#.#.#.#.#.#.#######.#.#.#.###.#.#####.###########.#.#.###.###########.#####.###.#.#.#.#.#######.#.#.###.#.###.#######.#.###.#.#######.#####.#
#.#...#.#.#.#...#.#...#.#...#...#...#...#.......#.#.#...#.#.....#...#.........#.#.#...#.....#...#.......#.........#.....#.....#.....#...#...#
#.###.#.#.#.#.#.#.#.#.#.#.#.#####.#.#.#.#.###.###.#.###.#.#.###.#.#.#########.#.#.#.#.#####.#.###.###.###########.#.#####.#####.###.###.#.###
#.#...#...#.#.#...#.#.#.#.#.#.....#...#.....#.#...#...#.#.#.#.#...#.#.....#.#.#...#.#.....#...#.#.....#.........#.#.......#...#.#.......#.#.#
#.#.#.#####.#.#####.#.#.###.#.#############.#.#.#####.#.#.#.#.#####.#.###.#.#.#.###.###########.#######.#########.#########.#.#.#.#######.#.#
#.#.#.....#.#...#...#.#...#.#...#...#...#...#...#.....#.#.#.....#.#...#.#.....#...#.........#.....#...............#.#.......#.#.#.#...#...#.#
#.#.#.###.#.###.#.###.###.#.###.#.#.#.#.#####.###.#####.#######.#.#####.#.#.#.###.#########.#.###.#.#########.#.#.#.#.#######.#.###.#.#.###.#
#.#...#...#.#...#.#...#...#.....#...#.#.....#.#.#.#...#.#.....#...#...#.#.#.#...............#...#...#.......#.....#.#.#.......#.....#.#.#...#
#.###.#.###.#.#.#.#####.###.#######.#.#####.#.#.#.###.#.#.###.###.#.###.#.#.#.#########.#######.#####.#####.###.###.#.#.#.#####.#####.#.###.#
#.#.........#.#.#.....#.#.....#.....#.#.....#...#.....#.#.#...#.#.#...#.....#.#.......#.#.....#...#...#...#...#.....#.#...............#...#.#
#.###########.#.#####.#.###.#.#.#####.#.#.#######.#####.#.#.#.#.#.###.#####.#.#.#####.#.#.###.###.#.###.#.###.#####.#.#.#######.#########.#.#
#.#...#...#...#.#.....#...#.#...#...#.#.#.........#...#...#.#...#.....#.#...#.....#.#.#...#...#...#.#...#.........#.#...#.......#.........#.#
#.#.#.#.#.###.###.#######.#.###.#.#.#.#.###########.#.#####.#########.#.#.#######.#.#.###.#.###.###.###.#########.#.#.###.#####.#.#######.#.#
#.#.#.#.#...#...................#.#.#.#.........#.#.#...#...#.........#...#.......#.....#.#.....#.#...#.....#...#.#.......#.....#.#.....#...#
#.#.#.#.###.###########.#.#.#####.#.#.###.#####.#.#.###.#.#.#.#########.###.#.#########.#####.#.#.###.#######.#.#.#####.#####.#.###.###.###.#
#.#.#.#.#...#...#.....#...#.....#.#...#...#...#...#.#.#.....#...........#.#.#.......#.#...#...#...#.#.....#...#.#.............#.....#.#...#.#
#.#.#.#.#.###.#.###.#######.#####.#####.#.#.#.###.#.#.#####.#.###.#######.#.#######.#.###.#.#####.#.#####.#.###.#.###.#.#####.#.#####.###.###
#.#.#...#.....#...#.........#...#.#.....#...#...#.#.#.......#...#.#.....#...#.....#.#.......#...#...#...#...#...#...#.....#.#.#.....#...#...#
#.#.#########.###.#.#########.#.#.#.#.#.#.###.###.#.###########.#.#.###.#.###.###.#.###.#####.#.###.#.#.#####.#####.#####.#.#.#.###.#.#.###.#
#.#...#.......#...#.....#...#.#...#.#...#.........#.........#...#.#...#.#...#.#...#...#.#.....#.......#.#...#.#...#.#.....#...#.#.#...#.#...#
#.###.#.#.#.#.#.###.#####.#.#.#######.#######.#########.###.#.###.#.#.#.###.#.#.#####.###.###########.###.#.#.#.#.#.#.#.#.#.###.#.#####.#.#.#
#...#.#.#.....#.#...#.....#...#.....#.......#.#...#.......#...#.......#...#.#.#.....#.........#.#...#.#...#.#.#.#.....#...#.#...#.......#.#.#
#.###.#.###.#.#.#.###.#########.#.#.#######.#.#.#.#.#.#########.###.#####.#.###.#.###########.#.#.#.###.###.#.#.#####.#####.###.#.#######.#.#
#.............#.#...#.......#...#.#.....#...#...#.#.#.........#.....#.#...#...#.#.#.........#...#.#...#...#.#.#.#.....#...#...#.#.#...#.....#
#.#####.###.#.#.###.###.#.#.###.#.#####.#.#######.#.###.###.###.###.#.#.#####.###.#.###.###.#####.###.#.#.#.#.#.#.#####.#####.#.#.#.#.#.###.#
#...#...#...#...#.#.#.#.#.#.....#...#...#.......#...#...#...#...#...#.#.....#.....#.....#.......#.#...#...#...#.#.#...#.....#...#.#.#.#...#.#
#####.#.###.#.###.#.#.#.#.#######.#.#.###.#.###.#####.###.###.###.###.#####.#####.#####.#####.#.#.#.###.#######.#.###.#.#.#####.#.#.#.###.#.#
#...#.#.......#.......#...#.......#.#...#.....#.........#.#...#.....#.....#.....#.....#...#...#...#.#.........#.#...#.#.#.........#.#.#...#.#
#.#.#.#.###.###.###########.#######.###.#.#############.#.#.#######.###.#.#####.#####.###.#.#######.###.#####.#.###.#.#.#######.###.#.#.###.#
#.#...........#.#.................#.#...#.#.........#...#.#.........#...#...#.#.....#...#.#.......#...#.#.....#...#.#.#.#.......#...#...#.#.#
#.#####.###.#.###.#######.#######.###.#####.###.#.#.#####.###.#######.#####.#.#####.#####.#.###.#.#.#.#.#.#####.###.#.#.###.###.#.#######.#.#
#.#.....#...#.......#...#.#...#...#...#.....#.#...#...#.....#.#.....#.....#.#.....#.#.....#...#.#.....#.#.#...#.....#.#.#...#...#.#...#.....#
#.#.###.#.#.###.#####.#.###.#.#.#.#.#####.###.#######.#.#####.#.###.#####.#.#.###.#.#.#######.#.###.#.###.#.###.###.#.#.#.###.###.#.#.#.#####
#...#...#.............#.....#...#.#.#.....#.......#...#.......#.#.#...#...#...#.#.#.#...#...#.#...#.#.....#.....#.....#.#...#...#...#.#.....#
#.###.#.#.#.#.###.###############.#.###.###.#######.###.#####.#.#.###.#.#######.#.#.###.#.###.###.#########.#####.#.###.###.#######.#.#####.#
#...#.#...#.#.#...#.....#.....#...#.....#.........#.#...#...#...#.....#.......#...#.#...#.#...#.#.......#...#.....#...#...#.#.....#.#.....#.#
###.#######.#.#.###.#####.#.###.#################.#.#.###.#.###.#.#######.###.#####.#.###.#.###.#######.#.###.#######.###.#.#.#.###.#.#.###.#
#.#.........#.#.#...#...#.#.#...#...#.......#...#.#.......#.#...#...#.....#...#...#.#.#.......#.......#.#...#...#...#...#.#...#...#.#.#.#...#
#.###########.#.#.#.#.#.#.#.#.###.#.#.###.#.#.#.#.#########.#.#####.#.###.#.#.#.#.#.#.#######.###.#.###.###.#.###.#.#.#.#.###.###.#.#.#.#.###
#...#.........#.#.#.#.#...#.......#.#.#.#.#...#.#.....#...#.#...#.#...#...#.#.#.............#...#.#...#...#.#.#...#...#.#...#...#.#.#.#...#.#
#.#.#.#.#######.###.#.#.###########.#.#.#.#####.#.#.#.#.#.#.###.#.#####.###.###.###########.###.#####.###.#.#.#.###.#.#####.#####.#.###.#.#.#
#.....#.......#...#...........#.....#...#.#.....#...#.#.#.#.....#.#.....#...................#.#...#...#.#...#.#.....#.#...#.#.....#...#...#.#
#.#.###.#####.###.#####.#.#.#.#####.###.#.#.#########.#.#.#######.#.#####.###################.###.#.#.#.#####.#####.#.#.#.#.#.#######.###.#.#
#...#.........#.#...#.....#.#.....#.#...#.#.#...#.....#.#.......#...#.........#.......#.........#.#.#.#...#...#.#...#.#.#...#.......#...#...#
###.#.#.#####.#.###.#####.#.#####.#.###.#.#.#.#.#.###.#.###.###.#.#############.#####.#.#.#.#####.#.#.#.###.###.#.###.#.#########.#.###.#.#.#
#.................#...#...#.#.....#...#.#.#.#.#...#.#.#...#...#.#.#...........#.#.....#.#.........#.#.#...#...#.#...#.#.....#.....#.......#.#
#.#.###.###.#####.###.#####.#.#######.###.#.#.#####.#.###.###.#.#.###.#######.#.###.###.###########.#.#.#.###.#.###.#######.#.#.#######.#.#.#
#.#.....#...#...#...#...#...#.......#.....#.#.#.......#...#.#.#.......#.......#...#...#...#.....#...#.#.#...#.#...#.....#.#.#.#.#.....#.#.#.#
#.#.###.#.#.#.#.#.#####.#.#######.#.#######.#.###.#####.###.#.#########.#########.###.###.#.###.###.#.###.#.#.#.###.#.#.#.#.#.#.#.###.#.#.#.#
#.#...#...#...#.#.#.....#...#...#.#.#.....#.#...#...#.......#.........#...#.....#.#.#...#...#.#...#.#...#.#.#...#...#.#...#...#.#...#.#.#.#.#
#.#.#.#.###.###.#.#.#######.#.#.###.#.#####.#.#.#####.###.#########.#.###.#.###.#.#.###.#####.###.#.###.#.#.###.#.#.#.#########.###.#.#.#.#.#
#...#.....#.#.....#.#.......#.#.....#.....#.#.#.......#...#.......#.#...#.#...#...#...#...#...#.#.#.#.#.#.#.....#.#.#...#...#...#.#.#.#.#.#.#
#.#####.#.#.#.###.#.#.#######.#######.###.#.#######.###.#.#.#####.#####.#.###.#####.#.###.#.#.#.#.#.#.#.#.#######.#.###.#.#.#.###.#.#.###.#.#
#.#.....#...#.....#...#...#...#.....#.#...#.#.....#.#...#...#...........#.....#.#...#.#...#.#...#...#.#.#.....#...#...#...#...#...#.#.....#.#
#.#.#.#####.#####.#####.#.#.###.###.#.#.#.#.#.###.###.#######.#################.#.#.###.###.###.#####.#.#######.#####.#.#########.#.#####.#.#
#...#...........#.#...#.#.#...#.#...#.#.#.#...#.#.#...#.......#...#...........#...#...#...#...#.....#.#.........#.....#...#.......#.....#...#
#.#########.#.###.#.###.#.###.#.#####.#.#######.#.#.###.#######.#.#.#######.###.###.#.###.#.#.#####.#.###########.#######.#.###########.#.#.#
#.....#.......#...#.....#...#.#.......#.#.......#.#...#.#...#...#...#.....#...#.....#.#...#.#...#.........#.#.....#.#.......#.....#...#.#...#
#######.#.#.###.###########.#.#.#######.#.###.#.#.###.#.#.#.#.#######.###.###.#####.#.#.###.###.#########.#.#.#####.#.#####.#.###.#.#.#.###.#
#.......#.#.....#.......#.....#.#.....#...#...#.....#.#.#.#...#.#.....#.....#.#.....#...#...#.#...#...#.....#.....#.....#.......#...#.#.#...#
#.#.#####.#.#####.#####.#.###.#.###.#####.#.#######.#.#.#.#####.#.#####.###.#.#.#.#.#.###.#.#.###.#.#.###.#######.#####.#############.#.#.###
#.#.#.....................#...#...#.......#...........#.#...#.....#.....#.#.#...#.#.#.........#.#...#...#.#.....#.....#...............#.#.#.#
#.#.#.###.#.#.#.###.#.#####.#####.#####.#.#########.#####.#.###.#.#.#####.#.#####.#.#.#######.#.#######.###.#.#.#####.#############.###.#.#.#
#.#.#.#.#...#...#...#.........................#...................#.#...#.#.#.....#.........#...#.#.....#...#.#...#...#...........#.#...#.#.#
#.###.#.#.#######.###.###.#.#.#.#####.#.###.#.#.#.#.#.###.###.###.#.#.#.#.#.#.###.###.#####.###.#.#.#####.###.#.###.#########.###.#.#.###.#.#
#.#...#...........#.#.#...#.#.#.#...#.........................#...#.....#...............................#...#.#.#...#.......#...#...#.#.#...#
#.#.###.#.###.#####.#.#.#####.#.#.###.###.#.#####.#####.#.#######.#####.#.#.#.#.#.###.#.#.#.###.#.#.#.#.#.###.#.#.###.#####.###.#####.#.###.#
#.............................#.#...#.....#...#...#.....#...............#.#.#.#...#.....#.#...#.........#.#...#.#...#...#.#...#.#.....#.....#
#.#######.#####.###.###########.#.#.#.#######.#.###.#####.###.#.#.#######.#.#.###.#.#.###.###.###.#####.#.#.#######.###.#.###.#.#.#####.#####
#.#.......#.....#...#.#...#.....#.#...#.....#.#...#.........#.#.........#...#...#...#...#.#.#.#.#.....#.#.#.........#.......................#
#.###.###.#.#####.###.#.#.#.#########.#.###.#.###.###.#.###.###########.#.#####.#####.#.#.#.#.#.#####.#.#.###########.###.#.#########.###.#.#
#...........................#.......#.........#.#...#.#.#...#...#.....#.#.#...#.....#.#.....#.#.#.....#.#...#.....#...#.#...#...#.....#.#.#.#
###.#.#.#.#####.###.#.#.#####.#####.#####.#.#.#.#.###.#.#.#.#.#.#.###.#.#.#.#.#####.#########.#.#.#####.###.#.###.#.###.#.###.#.#.#####.#.#.#
#S....#.......#.......#...........#.......#.....#.......#.....#...#.....#...#.....#...........#.......#.......#.....#.........#...#.........#
#############################################################################################################################################`
