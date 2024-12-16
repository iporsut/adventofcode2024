//go:build ignore

package main

import (
	"fmt"
	"strings"
)

type Command int

const (
	Up Command = iota
	Down
	Left
	Right
)

func (c Command) String() string {
	switch c {
	case Up:
		return "Up"
	case Down:
		return "Down"
	case Left:
		return "Left"
	case Right:
		return "Right"
	}
	return ""
}

func main() {
	inputs := strings.Split(input, "\n\n")
	inputWarehouse := inputs[0]
	inputCommands := inputs[1]
	warehouse, robotX, robotY := parseWareHouse(inputWarehouse)
	commands := parseCommands(inputCommands)
	printWarehouse(warehouse)

	for _, command := range commands {
		robotX, robotY = moveRobot(warehouse, robotX, robotY, command)
	}
	printWarehouse(warehouse)
	fmt.Println(calcGPS(warehouse))
}

func assertRobot(warehouse [][]rune) {
	var count int
	for _, row := range warehouse {
		for _, r := range row {
			if r == '@' {
				count++
			}
		}
	}
	if count != 1 {
		panic("robot count is not 1")
	}
}

func calcGPS(warehouse [][]rune) int {
	var sum int
	for y, row := range warehouse {
		for x, r := range row {
			if r == '[' {
				sum += (x + y*100)
			}
		}
	}
	return sum
}

func isMoveableUpOrDownBox(warehouse [][]rune, x1, x2, y int, c Command) bool {
	switch c {
	case Up:
		if warehouse[y-1][x1] == '.' && warehouse[y-1][x2] == '.' {
			return true
		} else if warehouse[y-1][x1] == '[' && warehouse[y-1][x2] == ']' {
			if isMoveableUpOrDownBox(warehouse, x1, x2, y-1, Up) {
				return true
			}
			return false
		} else if warehouse[y-1][x1] == '#' || warehouse[y-1][x2] == '#' {
			return false
		} else if warehouse[y-1][x1] == ']' && warehouse[y-1][x2] == '.' {
			if isMoveableUpOrDownBox(warehouse, x1-1, x2-1, y-1, Up) {
				return true
			}
			return false
		} else if warehouse[y-1][x1] == '.' && warehouse[y-1][x2] == '[' {
			if isMoveableUpOrDownBox(warehouse, x1+1, x2+1, y-1, Up) {
				return true
			}
			return false
		} else if warehouse[y-1][x1] == ']' && warehouse[y-1][x2] == '[' {
			if isMoveableUpOrDownBox(warehouse, x1-1, x2-1, y-1, Up) && isMoveableUpOrDownBox(warehouse, x1+1, x2+1, y-1, Up) {
				return true
			}
			return false
		}
	case Down:
		if warehouse[y+1][x1] == '.' && warehouse[y+1][x2] == '.' {
			return true
		} else if warehouse[y+1][x1] == '[' && warehouse[y+1][x2] == ']' {
			if isMoveableUpOrDownBox(warehouse, x1, x2, y+1, Down) {
				return true
			}
			return false
		} else if warehouse[y+1][x1] == '#' || warehouse[y+1][x2] == '#' {
			return false
		} else if warehouse[y+1][x1] == ']' && warehouse[y+1][x2] == '.' {
			if isMoveableUpOrDownBox(warehouse, x1-1, x2-1, y+1, Down) {
				return true
			}
			return false
		} else if warehouse[y+1][x1] == '.' && warehouse[y+1][x2] == '[' {
			if isMoveableUpOrDownBox(warehouse, x1+1, x2+1, y+1, Down) {
				return true
			}
			return false
		} else if warehouse[y+1][x1] == ']' && warehouse[y+1][x2] == '[' {
			if isMoveableUpOrDownBox(warehouse, x1-1, x2-1, y+1, Down) && isMoveableUpOrDownBox(warehouse, x1+1, x2+1, y+1, Down) {
				return true
			}
			return false
		}
	}

	return false
}

func moveUpOrDownBox(warehouse [][]rune, x1, x2, y int, c Command) bool {
	switch c {
	case Up:
		if warehouse[y-1][x1] == '.' && warehouse[y-1][x2] == '.' {
			warehouse[y-1][x1] = '['
			warehouse[y-1][x2] = ']'
			warehouse[y][x1] = '.'
			warehouse[y][x2] = '.'
			return true
		} else if warehouse[y-1][x1] == '[' && warehouse[y-1][x2] == ']' {
			if moveUpOrDownBox(warehouse, x1, x2, y-1, Up) {
				warehouse[y-1][x1] = '['
				warehouse[y-1][x2] = ']'
				warehouse[y][x1] = '.'
				warehouse[y][x2] = '.'
				return true
			}
			return false
		} else if warehouse[y-1][x1] == '#' || warehouse[y-1][x2] == '#' {
			return false
		} else if warehouse[y-1][x1] == ']' && warehouse[y-1][x2] == '.' {
			if moveUpOrDownBox(warehouse, x1-1, x2-1, y-1, Up) {
				warehouse[y-1][x1] = '['
				warehouse[y-1][x2] = ']'
				warehouse[y][x1] = '.'
				warehouse[y][x2] = '.'
				return true
			}
			return false
		} else if warehouse[y-1][x1] == '.' && warehouse[y-1][x2] == '[' {
			if moveUpOrDownBox(warehouse, x1+1, x2+1, y-1, Up) {
				warehouse[y-1][x1] = '['
				warehouse[y-1][x2] = ']'
				warehouse[y][x1] = '.'
				warehouse[y][x2] = '.'
				return true
			}
			return false
		} else if warehouse[y-1][x1] == ']' && warehouse[y-1][x2] == '[' {
			if moveUpOrDownBox(warehouse, x1-1, x2-1, y-1, Up) && moveUpOrDownBox(warehouse, x1+1, x2+1, y-1, Up) {
				warehouse[y-1][x1] = '['
				warehouse[y-1][x2] = ']'
				warehouse[y][x1] = '.'
				warehouse[y][x2] = '.'
				return true
			}
			return false
		}
	case Down:
		if warehouse[y+1][x1] == '.' && warehouse[y+1][x2] == '.' {
			warehouse[y+1][x1] = '['
			warehouse[y+1][x2] = ']'
			warehouse[y][x1] = '.'
			warehouse[y][x2] = '.'
			return true
		} else if warehouse[y+1][x1] == '[' && warehouse[y+1][x2] == ']' {
			if moveUpOrDownBox(warehouse, x1, x2, y+1, Down) {
				warehouse[y+1][x1] = '['
				warehouse[y+1][x2] = ']'
				warehouse[y][x1] = '.'
				warehouse[y][x2] = '.'
				return true
			}
			return false
		} else if warehouse[y+1][x1] == '#' || warehouse[y+1][x2] == '#' {
			return false
		} else if warehouse[y+1][x1] == ']' && warehouse[y+1][x2] == '.' {
			if moveUpOrDownBox(warehouse, x1-1, x2-1, y+1, Down) {
				warehouse[y+1][x1] = '['
				warehouse[y+1][x2] = ']'
				warehouse[y][x1] = '.'
				warehouse[y][x2] = '.'
				return true
			}
			return false
		} else if warehouse[y+1][x1] == '.' && warehouse[y+1][x2] == '[' {
			if moveUpOrDownBox(warehouse, x1+1, x2+1, y+1, Down) {
				warehouse[y+1][x1] = '['
				warehouse[y+1][x2] = ']'
				warehouse[y][x1] = '.'
				warehouse[y][x2] = '.'
				return true
			}
			return false
		} else if warehouse[y+1][x1] == ']' && warehouse[y+1][x2] == '[' {
			if moveUpOrDownBox(warehouse, x1-1, x2-1, y+1, Down) && moveUpOrDownBox(warehouse, x1+1, x2+1, y+1, Down) {
				warehouse[y+1][x1] = '['
				warehouse[y+1][x2] = ']'
				warehouse[y][x1] = '.'
				warehouse[y][x2] = '.'
				return true
			}
			return false
		}
	}

	return false
}

func moveRobot(warehouse [][]rune, x, y int, command Command) (int, int) {
	switch command {
	case Up:
		if warehouse[y-1][x] == '#' {
			return x, y
		}
		if warehouse[y-1][x] == '.' {
			warehouse[y][x] = '.'
			warehouse[y-1][x] = '@'
			return x, y - 1
		}
		if warehouse[y-1][x] == '[' {
			// try to move the box up
			if isMoveableUpOrDownBox(warehouse, x, x+1, y-1, Up) && moveUpOrDownBox(warehouse, x, x+1, y-1, Up) {
				warehouse[y][x] = '.'
				warehouse[y-1][x] = '@'
				return x, y - 1
			}

			return x, y
		}
		if warehouse[y-1][x] == ']' {
			// try to move the box up
			if isMoveableUpOrDownBox(warehouse, x-1, x, y-1, Up) && moveUpOrDownBox(warehouse, x-1, x, y-1, Up) {
				warehouse[y][x] = '.'
				warehouse[y-1][x] = '@'
				return x, y - 1
			}
		}
	case Down:
		if warehouse[y+1][x] == '#' {
			return x, y
		}
		if warehouse[y+1][x] == '.' {
			warehouse[y][x] = '.'
			warehouse[y+1][x] = '@'
			return x, y + 1
		}
		if warehouse[y+1][x] == '[' {
			// try to move the box down
			if isMoveableUpOrDownBox(warehouse, x, x+1, y+1, Down) && moveUpOrDownBox(warehouse, x, x+1, y+1, Down) {
				warehouse[y][x] = '.'
				warehouse[y+1][x] = '@'
				return x, y + 1
			}
			return x, y
		}
		if warehouse[y+1][x] == ']' {
			// try to move the box down
			if isMoveableUpOrDownBox(warehouse, x-1, x, y+1, Down) && moveUpOrDownBox(warehouse, x-1, x, y+1, Down) {
				warehouse[y][x] = '.'
				warehouse[y+1][x] = '@'
				return x, y + 1
			}
			return x, y
		}
	case Left:
		if warehouse[y][x-1] == '#' {
			return x, y
		}
		if warehouse[y][x-1] == '.' {
			warehouse[y][x] = '.'
			warehouse[y][x-1] = '@'
			return x - 1, y
		}
		if warehouse[y][x-1] == ']' {
			// find the next open space
			for i := x - 1; i >= 0; i-- {
				if warehouse[y][i] == '.' {
					// shift the box to the left
					for j := i; j < x; j++ {
						warehouse[y][j] = warehouse[y][j+1]
					}
					warehouse[y][x] = '.'
					warehouse[y][x-1] = '@'
					return x - 1, y
				} else if warehouse[y][i] == '#' {
					return x, y
				}
			}
			return x, y
		}
	case Right:
		if warehouse[y][x+1] == '#' {
			return x, y
		}
		if warehouse[y][x+1] == '.' {
			warehouse[y][x] = '.'
			warehouse[y][x+1] = '@'
			return x + 1, y
		}
		if warehouse[y][x+1] == '[' {
			// find the next open space
			for i := x + 1; i < len(warehouse[y]); i++ {
				if warehouse[y][i] == '.' {
					// shift the box to the right
					for j := i; j > x; j-- {
						warehouse[y][j] = warehouse[y][j-1]
					}
					warehouse[y][x] = '.'
					warehouse[y][x+1] = '@'
					return x + 1, y
				} else if warehouse[y][i] == '#' {
					return x, y
				}
			}
			return x, y
		}
	}

	return x, y
}

func parseWareHouse(input string) ([][]rune, int, int) {
	var warehouse [][]rune
	var robotX, robotY int
	for y, line := range strings.Split(input, "\n") {
		var row []rune
		var i int
		for _, r := range line {
			if r == '@' {
				robotX = i
				robotY = y
				row = append(row, '@')
				row = append(row, '.')
				i += 2
			}
			if r == 'O' {
				row = append(row, '[')
				row = append(row, ']')
				i += 2
			}
			if r == '#' || r == '.' {
				row = append(row, r)
				row = append(row, r)
				i += 2
			}
		}
		warehouse = append(warehouse, row)
	}
	return warehouse, robotX, robotY
}

func printWarehouse(warehouse [][]rune) {
	for _, row := range warehouse {
		fmt.Println(string(row))
	}
	fmt.Println()
}

func parseCommands(input string) []Command {
	var commands []Command
	for _, r := range input {
		switch r {
		case '^':
			commands = append(commands, Up)
		case 'v':
			commands = append(commands, Down)
		case '<':
			commands = append(commands, Left)
		case '>':
			commands = append(commands, Right)
		}
	}
	return commands
}

var ex = `#######
#...#.#
#.....#
#..OO@#
#..O..#
#.....#
#######

<vv<<^^<<^^`

var ex2 = `##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`

var input = `##################################################
#.#O.O..O...........O#.......O.O...O..O....OOOOO.#
#.#.OO....O#OO......#O...#...O....#......#.OOOO.##
#OO..#..O........O..O.O..O....O...#..O..#.#.#...O#
#.O..#O....O.....O.O..O..OO...O...O..O#...#O...#O#
#O..#.O...#.#O.O.#.....O.OOO.#OO..#....O.OOO....O#
#.#..O.O....O..O...O...O.O.......#.O.O..O..O..#..#
#.O.......O.....O...O.O....O..OOOOO.#.....O..#...#
#OO.....O#...O......O...O..##O.#..O........O.##..#
##.OO..O...OO...O#..OOOO......O..OO..OO.O.#O.O#..#
#........O......O...O....O.........OO.OO...O.....#
#..O..OO..O.....OO......O.#O..#O......O.O#..O..OO#
#.O......O..#O.OOO...OO.O....#...O......O..OO.O..#
#..#O.##..OOO..O.O....O..O..O.OO#...O....O..#OO.O#
#O..O........O...##O..O....OO.O.O..O...#..O......#
#....O.OOO...#O.O......O.#OO....O.....#O.O..O..O.#
#O......#.#O.##.........O.#OO...O.O.OO.......O.O.#
#..OO..#...O.#..#.....O.....O.....O#O.##...O...OO#
#..OO......O..........O...O.O..O......O..O.....OO#
#...O..O..O......#..OOOO#O....O....OO....O.......#
#.OOO....OO.#.O.OO......O.....O..#..O...#....O.OO#
#.O#...O..OOO.O.#..#O..#O.O...O.O#O.#.O......O..O#
#O.OOO.O..O......O......##O.OO...O..O#.....O...O.#
#........OO...#..O...O..O#O.O.O...OO....O....O..##
#...#.O...O.OO....OOO..#@O....OO.OO.OO.......O...#
#.....OO........OOO..OOOO.OO..OO.O....O...#OO....#
##....OO#..OO.#......O......O#.O....O.O..........#
#......O..#O.O.O.OOO....OO.....#...O..#O.O..#..O.#
#O.OO...O.#.O.O....O#.OO..O....O.O....O...O......#
##..O..O.....O............#O........O......OOO.O.#
##..........#..O...O..O..O...OO....OO....O.O.....#
#.O..O.O.#O....#OO...#O.O.OOOOO.#...O.....O..O..##
#.......OO..O...O.O..O.....O#OO.OOO.........O....#
#.........#..OO..O....O..#.O.O.O...O.OO......O...#
#...........OO...O##..O#....OO.O...OO.O.....O.O..#
#.....OO...O#.#.#OO.....#..OO........O....#...O#.#
#O.O.O......#....O...OO........O.#O...........O.O#
#.OO.........O..O..#.#O#....O.OO#O...OO.OO...OO.##
#..OOO.O...#.....O..O.......O.........O#...OOO...#
#.O..O#.....#...OOO.#...O.##O...O...O..OO..OO.O..#
#.O..O...O.....O...O#O.....O......O.#.O.O....OOO.#
#OO.O..O.O..O.#.O.....O.##...O........#..O..O.#O.#
#...O.....O#......#OO..#........OOO.#..OO.OO.....#
#O.....#.#O.O.OO..#......OOO.......O..........#.##
##.......O..O...O..#O#......O#..O....O........O.O#
#..#.....#...O...O.....OO...OO.#O....OO#O........#
#..O..O.O..OOO...O....O..OO..O.....O..OOOO.#OOO..#
##O.O..O#.OO.O.OO.O..O.#..O...O.O.OO...#......O..#
#....O.O..##.#O..#.O...O.O.#O.O.#.......OO..O....#
##################################################

><v<^>>>^vv^^>><<v^><>^<v<<vv^v>^v^>><v>^>^<<v<<^v<^v>^<v>>^<v>vv^v<<^^<><<<>^^^v<><>v<>vv><>v<v^><^>v>vv<<<v^^v<^^v<<v^vvv<>><^>v>v^>^<<<<v<<>>^<>v>><>>^<^>>v<v>^><v<^><vv><<>v>v>v^>vv^vv^v>><>>v><<>>v<vv^^v<>v^><v^v<v<><v^><^v<<>>^<^>v<><^<^^^v<>^<>^<>v^v^>>^^v<<<>^><v><<<^>>v^<>vv<^^<v^>>><v^^v>><><^>>^<<>^<^<<^v^^<>><v>>>^vv<^<<>v>^>v^<v><>>>^<^>><<>^^^><^vvv<>^<<v><^v<<^v<<v>vvvv^<>vv<<>>v><v^>^<v<<^^<>v<><<vv<^<<<^>><^>^^>v<^>>>>v^v>>v^<v>v^v^<<v<<>vv>^<<<^<<^<<^<^^^<v<v^><<vv>>><>^<^>^>^^<>v>^>>>v^<^><<<<v<v<>v>^<^v<^v<<^^>>^>^>vv><^<v>vv^^><<<<>v<v>^v>^<<>v<>^^<<^><^v^v<v>^^^<vvv<>><^>vvv>>^<>^v^>^v>vvv^^^v<^<v<v<<^^^><<v<^<>^^<>>><<><^v><><<v<v<^<<^^v^^v<^^><v<vv>^<<>^^<<<v^vv>^^v<>^v^<<^^v<><>^<v^^^<<v>^><>><^>>^<^vv>v>>v<<^^^v<v><^<>><^<^<v<vvv^v<<<^^^^>v>v<<><<v>v^<>v<^vv^>v<>vvv<vv<<^>>><v<>^<^^^><^v<>>^^vv^>^vv<^><><^vv^v>>v>>><v>v<>vv<^>^>^<<v<<><>^v^<><v><^v^>><<v<<vv<v<>^v<<^v<^>v^><>vv<vv^vvvv^>>>>^>>^^>v^^^^<<vvv>^>v^<v>^v<^<v^>^<<<><v<v<>v>>^^^<v>>>>v^^><>>vv<^>vv<>
^<vv^>^><^<>^>><<vvv>vv<^v^>>^^^vvvv<<<<>>>v>>^>^><^>^v<v>>v<><v^^^>^v<v^^^^^<>><^vvvv^^>^^^<>vv>v>^<>>>v^^<v^>v<v<^v<<>>vvv>v^v><>>^v><^^><^<v>>>v^v^>>^>v>>^vv<^^<^<v>^>v><^^v<<v^<^><<^v^>^><v>v<vvv><><v^<^vvvv^^^v>v^>>^^>v>v>^v^^><^v^><v>><>v>><<v<^<>^>v^vv>>>>v<^^>v>^v<v<><<^^<v<>><vvv>^^^v^>v^^<>v^>>^vv><^v<^^<<vv<><>>vvv<>>v^v<^>>^>v<>^^>><>v^^>v>v>><v<<>^v>^<^<v>^^>>vv>^v<>^v^v^<^v^^><><<<^^>^^>v>^>>^^^v><v^>v<vv<>^^<<^><v<^^<^v<v<<>vv><<v<>><><<<><><>v^v>^vv<^<><^<>>^^><>>v>^^v<^<>><vv<v<v><><>^^^<^>>^><vv<^<>><<<^><v^><>^v<^<v><<v>^vv^>^<>^v><^<<<><>^^^^v><^<^v>^^<>^><>v<<^>^vv<^^<<>^v>vvv<v>>^>^>>>vv^^^<<<^^>v><>v>^<>>>>v<<<^>>^^>^<>^>v>v><<^<<><v^vvv>^^^^^^>vv^vv<v<v^^<>v^^^<>>^^<v><<<<^<^<v^>^>^><><<<<v^v>>><^<><^<>v<>v^<^<v<vv<^<><^^^v^v^><v>v>>vvvv^<v<<^>>^>^^vv<>vv>^^^<^><v>^<>><>^v>v>vv^^v>>>><v^<><v^^><<>>vv>><vv^<^v^<vv<>>>><<v<<v>><v>v^<<v^>^<^^^^<vv^<v^^^<^><^^><v^v>^^<^>v>>^><vvv>v>v<<v^^>v^v>>>vvv<>v<v<vvv>>><v<^><^^v>^<v<^vv>>><v<>v^>v>v<vv<><v>^v>^>^><v<<^>v^<^v>
>>>^<><<>v>vvv^<><>^>^v^v>><v^><<<><v>>^>>vv^><<>vvvv^v><vv<v<<<>^<<v<<<^>v^>vv><^>^><^^v^<^>>^^v<<v>v<v^>>>><>vvv>^>^^>>v<^>^<v>^>>^^<^<<v><^<<^v>v<v<<vv^<<v>>^>><vv^><<^^<v^^><^<^><<><v>vv<><<^<^vv>><<>>>>v>><<>>^v^<<v>>>^<v^v<>^<>>v^<v<^^^^<^>>^v<v^vv^^<vv^>^^v<>^<<vv^^v>>vvvv^>v^^v^^>v<v>>vv^v<vvvvv<v><<>^^^v^<<>>vv^>>><>^><^>>v<>>v<^><>^<^>>v^>vv^v>v<<<v^v^v^^^<^^v<v^<v^>v<vv^v^<v<^^v^vvv<vv<^<v<^^v>^v^v<<><>^^v>><<v^v>^v<vv<<><^>^^<v<vvv<^v>^^>v^v><v<v<v^^>>^<<<<^v^^v<<<^>v<v<v<v>^^^>>vv<^v<>>>^^v^<v<v><^^<>>^v<>^^vv<vv<>v^vvvv^<^v^^v<^><^>><vv^>vv<<<^><v<>>vv^<^v<v^v<<^>>v<<>><<<>><^<<<v>>v^v><><v^<<^^<v<^>^<>v<>v>^^>v>^<vv^>>^^><<v<>^v<vv^^v><^>^>v<^<>^>>^<^vv^v^v^v^^<<<v>>v^v>vv>>>>v<v><v<>v<>v^<^vv^vvv^<<>v<vv>v>>vv>^>>^^v>vv><>>^>>v<v<>><>><<v<<>>^^v<>^>vv<vv<^><>vvv<^vv>^><>^^v^<^>^^<^>>>vv^^v^>^v^vv>v<>^^^^v>>v^v^<<^vv>v^^><><v^<^v<v>^v^>^^<>^<^^^>^vv<^>>><><<><>^<v^<>v>v<^vvv<v><<^v^^>v>^^><^v^<<v>>>vv<v^>>v^^^v<^^<>^>^>>^^<^^<^<^<><v>v^<<>>v><^<^>v<^^>^>^<v^<^>vv<>^>^^^v
^^>>><>^>vv>>^>^^<^>^<vvv^v^^v<><><>>^^^<<vv^>>>^^<v>v>><^<^><<^<>^^vv<v>>vv^<<>^><v^v>>^><>>v<<>><<><v<><><v>><<>^^^^^<<>>v>^v<>^>v^<^>><^^<v<><^<vv<v>^<^^^>vv^^>^^>v>><v^<><<v><<>^<v>^^>><v<>>^^^v^>^vvv<vv<^vv^^><<v<>>>^^^vvv<<<<vv^^v<<^vv>v^v><><<>^^<<<^v^^^<v>vv<^^>^v<v<<^v<<<>vv>^v^v^^<^v^v<^vv^<^>^<^<<<>v<^^<<><><^><vv>^^>v>><^v>vv><^>^>^>^vv<v^<<>^<^>^v><<^v^^^vv^><vv><^<^^<^^^v<v^vvv^<<^<v>^v^<v>^><^>^><>v><v>v>^^>>>^^>v>^^>^<v<v>v<v>^<><><v>v><^>><^><^>^vv<^<>v^<^>><v<vv^v^<vv^<<^<^v>^>><<^>>^^>v^^>>>^<v><<^<^v^>^vv><v<vv>^vvv<^<^<<vv<>v>v>^v^v>^v<v>v^^<v><^^<^>v^^>>^<>><<vv><^>>^^v^<v>v<>>^^^vvv^>>>v^>^>^>v>>^>^><<><vv>vv<v^>^<v>v^<v>^^<v^<<^><^<<<>>^v>^><^<v>v<^^v<vv^^<vvvv<><<>>^^<<^>^^>>>v^^<^v^v<<^><<^^<vv<<<vv<>v>v>v>^<^<^<<>^^^^^vv<>>>^vv^^^<^^^v>>^<vvv>><>^<>^^^v^vv^<<<v^^v<>v^>^^><^<>^>>^v>v^v>v<>vv<^>^^<vv>>v<>v<v^><v><>v<^vv^>vv<^<>v<><v>vv>>>>^<<<<v<<vv>^^^^>v>>^<<>^>^v<^>vv>>^<<v>^v<>>^>v<v>><^<>^^><<<^<^><^>^<>><vv^^<^>v^>v^<^^>^>><<v>vv<v^<>>^v>vvv>>vv^^>v>><><>
vv<v^<<<v><vv>>^>^>>><<^^>vv^v<><^>>v^^<<>><><><<vv<v<>>^v<^^><^><^^vvv^>><>>>v^^<^>><>><<>vv<<<v^>v>v^<<v^v^vvv<>vv^>><><>vv^>^^>vv>^><<>>v<>^^v><><v<^<<v<<vv<>v>>><vv^^>^>v<<>>>^^>v^<v<v<^<^v>v^v<>><<^^<^>^v^vv<<><<^vv>>vvv>^v<^v^>>vvvvv^^^<^<v><^v><>v^v<v>><>>>>vvvvv^<^><><^>>><><vvvv>v><<>^^>v^v^>vv><>>^<^vvv^vv^vv^^><>><^<<>^vv>^v^v>>><^v^<<v^^>><^<v<vvvv<v<v^vv>v>>>><>vv<<^vv^>>>v^>^<<v^v<<vv<v><<><<v^^<>>vv>v<v<<<>>^^v<v^<^v^v<<<v<v^^^>v<v>>>>>>v<<<>v><><>>^>^>vvv<><^>>>v<<vv^<^<v<v><^<>v>v>><<v^<^vv>v>v<vv^^><v^<><v^v<<<^v<v^^>><^vv^<<>^><<<^v^^><vvv<>^<>^^^<>v^^>>><^>^<^<v^<vv><>^v>^>>>>v<v>^>v<<><>^<^<v<^<><<>v^v<<^v<^><><<<^<><<^^<><^>><<^v<><v><>v<v>v^vv<<^v<v><<<vv>>>^>v>v^>>^vv<<v^v^<<^>^^<v>^v^v^^<><vv^><v^>vv^<v<>v^v^<^>>vv^v^^><<^<>v<^^^v<<^>>>^<v<<<^>>>>>v^<^v^>v<>^v^<><<^<^^<v^^v>v<v^<vv^<>^<^><<^v^v^vv><^^^>>vv><<vv>>v<v>v<^<v><v<v<>^<<^vv^v<<><<^<vv<v<<v<>>^^v>>^>v<><>^<<><><<v^<<v>^^v<<>>^>^<<^<^v>^v>>^>^^>>v<v><v<<^<>^v^^>v<vv><>^v^<>^<>>vv<<<<><>v><^v<<>v<v><>v<
<v<vvv^<^<v<^^^v>>><^^^^<vv>>^^<^>^v<v<vv>vvv^>^>^<><<^^>v>^<^vv>^><>v<>v^><>^v^><>^v><>>v>^v<<v>v^v^><v<>>^^>v^v^^vv^>>><<<^><^><^vv><<^>vvv>^^v><<<>><<v><>^><v^v<>><^>>v>>^<<>>v<<>^>>>^v<<v<<^><><<vvv<v^>><>v<<^>^>v<v^<^v><>v>^<>>^^v^>>^v><vv<v<^><<^>>><^^>^v<<vvv^^>>v<>^<v>^<<>>vvvvvv>v<<<^<v^v<><v>>>v^>v<>^<<v<v>v<>v^^v><<>^^vv>vv<v^>^<<^^^v<v>v<^^>><^<>^<v^^>>v^v><<>v>><><vv>><^v^v^^>>>v<^^vv<^<<>><<>>v<^<^<^<^^v<v>vvv<>>v^^><^>>^v^^^v>v<^vvv^^^<v^^<v>^<<>^>^>vv>vv<<v>^><vv^<>v^<>^v^<^v<^><>vv<^<vv<v<>v<<^>^v><>^>>^<><<><vv<>v^^<><><<>^>^^v<>^<>>>v^>v<>vv^vv<^v<v<v>^^v^v<<>^^^<v>v<<>v^vv^v<vv^^<<vv>v<v>^<v<<^>vv<<>vvvvv<v<<^vv^<v<>>^v<<>vv^v<vv>^^v^^v^v<^vv>v>><v><^^>><^<v>v^vv<>^v<>><>v^v><>>^>v<<v^>^vvvvv<<^^^<vv<>vv<>^<<<<vv^>v^^^^^^vv^v^^v>v<^<^^^>^v^<<^<v^^<<<>>>v>v^^v>^^^<<<vv<<<v>><><^><>^<><><v^><^><^^^^vv^>^^^vv><>v^<v^<><v><>^<>>^^vvvv>>^v^v>^>><v>v>^><^<>v<><^<>v<<v^vv><>v^^v<v<><<>^<v<<<>>>^><<v<v<^v^>^^^v<^^^<>v<<^^v<^vv<<>>^v<v^v<<<>v<>><>^>v^^<<v^<>v^vvvvv^^><><<^>v
<><^>^^<^vv<^>^>v>^<^><<^><^>^vv>>>>^>><<>^<^>^^<v>>^^v^>^v<vv^>^v<v^<>>^<><>>>>^v^vv<vv<<^v^<v>^><v^v^vv^<<>>>>^<>vv><^^^v^^<<v<v>^>v>v>><<v><>v<>^><>>v>^vv^^^>>^v<>^^>>>^v<<vvv<>><<><<^^<^><^^v^>v<^>><v^>>v^vv^>v^><v^><^>>>><v<<vvv><v^>vvv<<v<vv<<^^^v>^<v>^<>>><><v^>>>>><>v>vvv^v<vvvv><>^v<><vv^<<^v<<<^>><<>>>^^><><v<^^^^<v<v>vvvv<v<^<>^<v<>>>v<^v^^v>v>v^<><^><v>^vvv^^>vv^>>^<^v><^><><<>>^^^^^v<^^v^v<v<><<>>>^<v><v<>^>vv>^>><^<vv>^>><<><>vv<vvv>>vv^>>>vvv^^v<>^<^v<^v^^v<<^vv<^<^<^v>^^><v><vv^<^>^^>>^vv^<><^^v>>v<<^v>^^v^^<vv>>^<<^v^v^^><^<>>^>v^<v>v>>^v>v>>^<vv<><>v^^><^^>><^^<<v<vvv>v^v>v^<>>v^>vv^>>vv^>vv^^>>>>^v<>v>v>^^^>vv>v^<>^^>vv><<>v^^^^v^>^^^v<<>>^<<<>v^^<^>v^^v>>^vvv<v<<v>^v>>^>^<vv<v^><^^>^vv>v^^^>v<>^^<>>><>><><>v<>v>><>v>^^>^<^>^<<^>><<>^^^^v<<<>>v^^>^<<v<vv<>>^<>>>^v^^^<^vv^<>^>>^vv^^vv>^<^>^>^>>>v>>v^>^<vv><^>>v<><><><<^><>>v<v<><>v^>^>^>^<<vv>v<<vvvvv^<^<^<v^^v<^v^<v^<vv<v>^>>^<^v>^>vv>^v>v><v^><<><><><><>v<><>^v^>>vv^v>>v^^^v^>v>>^v<>^vv<<^^^>v<><v^>^^>v><<^^v^v>^^>^
v><v<^v<><>^^>>>>v^><>v<^>v>^^>v^<<<>>>v<<>>>^vvv<>>^>>^^v<^^>^^v<>^^<^v<<v<^><^<>^<>^>^v>><<<^^<>^v><^v^<^<>^^<vv>vv>>>^^vvv<v<v>>vv^v<<>^<^<v^<<vv<v>>v^^v<^^^^^v><><v><><>v^>v<><<<^><<^^vv>>^<^<<>v<>><v>^<<v>^v><v<>><^>vvv^>v^<><^>>>>^^>>vv^v><^>^>v>^>v<<<^>v<<^^<>v^^^><><v<^^>^>v>^v<<><^vv^><>v>^vvv^vv>^^>vvv<v^>^<vv>v<^<^v<^<>^vv^^v^v^^>^^^vv^^>v<v>v^^<><v>^<^vvv<<<^^^^<v><>>>>>>^<^<>^><^^v^^^^<<^>^><^<v<^>vvv>>v^>^vv<vv>^v>><<^<^^>^^>^^<>>v<^v>v>>^><vv>^v^^v<>^v>vv<^^<>>v^^<<^>^^><v><><vv^<^vvv><^^>>>><<vvv>^><<>^^^^>v^^<<>^vv><>v^^>v>^v<<v^^^^v^^v^<<vv<<v>vvv^>v^<^>>v^>^>^^^v<^v^<v<<>><<>>>^v^>><^>^>>v^<<^<>^vv>>^^v<<<^vvv^<>^vv<>v^v^>^<<v^^v<>^v<vvv<<v<^<<v><<v>>v<^v^><<vv>v>^>v^>^><>>v^^<^^<v>^^>^v><<^>^><v><<v<>vv^<^^<^v^vv<v<<<v^><^>>>>^><<^^^^><^v<><<<<v<^^^>>v<^>^v^<^<^^<>v^>^v^<v<^>><><<>^v<^vv<^^v>vv^^^v^><>>^^vv<^<>^>><^v<><>>^^^v>v<<^>v>>^><^v^<<<<^^>v^>v>v^v<v<><^v^>><v><<^<>vv<<<v^^v<^>>v<^^^>>>v>>vv^<>>^<vv^<<v>^<<vv><<v^^^<v^>^>^<v<>>^v<v>v<^^^vvv>>v>v><<vvv<<^^>>>>
v><><v>v^^^>^v^^v>>>><><^v><v<>>^<<><^>^v^v^<v<^<<>^^<^v^>^>>v<<<>>v<v>><vv>^<^^<<<v^<<>vvvvvv>>^><v<^<v><vv^><>>>vv<^^><>v>^<>v^^v<v>^<<^<<^^<v<<^v>^<<>v><v><<<<>vv><v<>v^^v^vv>^><^<^<<^vv><vv^<<v><v^<><<>^v<>^vv>>>^<v<><><v<>^^^v<v<v><v^vv>v^>v<vvvv>>vv>v>>^>vv^>vv<v<v<<^v>>v<^^>vvv>^^>>><v>v>v>vv^>v^<<^>><>^^><v<v<<v><v^<^<^v><>^>v<<v^^^v<^^<><>>vv>v^<vv><><v^<v^vv>vv^v^<>><>><vv>>>^^>>>^<>>^>>>v>><>>^^^v^>^><v^>^^>^<<^<^v^^>^^>^^<>v<>><<v<^>^>^v<<>v^v^v^<^^<>v<<^><<v>^<<v^>>vvvv^<^v>v<^>><<<^^^>^v^<v<<v>><><^^<>v>v>^<v<<v><<<<vv^^<>>>>^<<<v^<vv^><<^v<^vv^v><^^<v>v<<^^>^>><<>vv^v><>^<<>^^<>>^v^v<v<<^v>>>vv><^>^>^v^^v^<vv<^v<>>^vv^^^^>^<>vv<v<<v>^^v>^v>v<<vvvv<>^^><<vv>^v^^><<^^><>^>v<>^v^>^v^v<>^<v<v<^^<^^>vv><v<>><>v<^><^v^>>>vv<^>v<<<<^<^v>><>^><^^><^v<<^v^<v<v^v<<>><<<v><^><>>>^^>v>vvv<>^v>vv>>vvv<^v<v>^v^v<v<^v^><v^>^><>^^<v^><>v<<<v<>^^v<<v>v>v<<v>^v<><>v<vv<>><^^^<>^v>v>v^<>>^>v<^><vv^>v>>><><>^^>v>>^^<^>v^v<><^<v<v><<><<v<>^>v<<>^>^>>v>>^vvvv>^<><^<v<>v<>^<^<v^>>v<v<^><><^<<<
v<^^vv<<<>^v<v^><^^^v<<<<v>^<^<<^v^v<<>vv>^vv>^<^v^><><<>^>>^>v<^>v<<<vv<>vv<<v^^v^v<<^>>><^>^><^>^>v^^>v>><v>^v>^v^<^^^^<^v<vv^^<<<v^vv>>v<>v>^v^^vvv<>>v><>>><^^^vv^v^^><vv<>>v<>^vv^<^>><^v^vv^>>vv<v<>^>v><^v^>^>>^>^>>^<><>^^^^vvv^><v<v<^vv<<v<<v>^^>^^><<^^>>v>v^^>^<<><^v>vvv>>v<<><vv><vv^>v<^v<v<>^>v^^><v^>^^v<>^<v><>v^<^>v^>^<<v>v^^>^<>v<v<>v>v^^vv<<><v>^v^v<>v<v><<^v<<<^<><<>vv^>^><^^<v<^vvv>><<vv^vv^><^>vvv^<<v<v^<<>v>>v><vv<<<><>v<^<<v<<^^^^v^vv><v>^<vv<^vv^<v^<^v<v>^v^^>^>><<^^^>v<>^>^>v<>v<<><v>vv<^>vv<v>^<><^<^^><<^>v<<<>><>v>>v><>><>>v>>v^^vv<^>>>^>>>>vv>v<^vv>v>>^>>^>>v<>v<<v^v<v<>v>vvv<>vvvvv^^<v>><^^>v>v<v>^<><<<>>v><>v>v<^<v>>^>^vvv>v<><v>v^<<^>>vv^<^v<<<^<^^<>>vv^v><<<v^v<^^><v<^vv><<^<v>^^vvv<><^<>v>>v>v<^>vvv<v^v^<<>><v<>^<v>^vv>vv<^><>vv^^^<v<<><^>>^v^^v>^^^vv<^><^>^v><vv<v^>>v<^v><<^<<^^v<>v>vvv<^v<^v><><v<><>>^v>>^<^>>vvvv>>^>><^v<<^<^>^<^^^v><^^<vv^^<<vv>vv<>^<v<^<v>v^><^<<<<^>v><^v<v^^<v^<<><^>>v^>^>v>>v^>><vv^^<^<<^<>v><<^<>>v<>><><^^v<<^<^^><v>><<v>v>^<><^<<<^^>
^v<>>>^v<^^^^v<>vvv^>v<^^vvv><v>>>>>v^^^^><^>^>v<v>^^<><^<^<>><^<v>^^>v>^>^v<vv^>^<^>^^<<<>^<<v<<^^^>^^v<v<^>v^>vv^vv^^<^^>><<v><<>^>vvv>^><^v^^>v^^v><^<v>^^^<<>><^v>>vv<^^^v<><^>^<>>>>>^^<<<<vvv>^^v<v<>v<<^^v<>>>v<vv<<^<<v^^^v>^>v^><^<v>v^^^>>><>v^^v^vvv<v>v>^v^>^<<^vvv^vv<>>>>>^<v>^v>>v^v<>^^><><^><v>>>^>v>>>vv^v>^>>v<vvv>^vv^<v<<v<>>v>>>v<>>>^><><v>>><>v>>v^<>vvv<v>>>>v>^>>vvv^^<v>^>>>>v^^v>>vvv<>>><v^v<^>v^v<^<^^<v<<vv>v^>^v>>vv>vvv^>>>>><><^><<<^vv^><v<<^>>v<>^>^^v>^<v><<<>^vv^<<^v<v<vvv>vv<vvvv^><<^v>vv^v^^vv^^v<<>^vv^^<v^v><<<>vv^vv^v>^<>>v<v^>><^>^<^<<<^><<<^v^v<<<<v<v<><>>>v^><v^v>^^v<^<<<^v><^v^vv^v>v<^^vv>>^>>>v^^^v^v>vv><><<>><^^^^v^<v><vvv^>>v^vv^>>^v<<<>>>^^^^>^v>><>>>v>^<>>^v><v<<<^v^v>^<<^^v<<<>^<<<<<^v^^^v>>^<><v^<^^vv^v^^v<^<>><^>v<v<<^>>>^>vvv<<^<vv>>^><^v<>^><^>v<<<<^<><<>^<^v^>><<<^><>><>^v>v<>^^vv^v<v<v<<^v<v>>>v<^^<^v>^>v^^>><vv<<>>^^>vv^<vv^<^<><<vvvv<><<^^>>>>v^<^^>><<<v<>^<><vv>><v^><>>^><<v<>vv<><>>^vvv<v^><^^^>vv>><>>v>><vvv<>^><v>><v<><<v>>^vv^v^vvv^v<v<v^v
v^>^<v<<v>v^^vv>vv<^><<<v>v>^><<^>>^>^v<v<>>^vvvv^<v^v^^<^<vv<vvvv^v^^^^<v>><^v>^<^<>vv<v<>^>>>><^v>^vv^><^v^>^v>^v<>^v<>>>^<^vv<v><>vv<^^<>^v^><vv>>v^^<^>v<>v>>><>v^vv>^>v>vv^><<<^>v^<^^^<<<>v>><<v^^<<>v>><^^v<vv<>>>><>>^<<v>^><^v>^^>^>>v^>>>^>>>^>^^<^^<><^<<^^<>vv^><^v>v<>v<vv^>>^<<<<<^^v><^>v>v^^<^<vv><>^v<^^vvv>>^>>>><v>>v>><^<^v<^>^<v>>^<>vvvv>vv>>>>>^<^v>>v>v<<<v^vv<v>^<v<>v<>^>v^vvv>^><v^v<>v<^^v^>>>>^v<>v<<<<^vv>^<>vvv>^<^^<<>vv^>v^^vv<^>vvv^vv>vv^<<v>^^v<<^>v>^vvv^<>><<^<>>>v^<vvv^>^vv^<^<v<v<v^<v^<<^vvv^>^>>^<vv^^<^<<^v<<<v>vv>vv<<<>^v<<vv>v^v<v^^^^<^vv^<vv<^<^^^v<>vv^<>^^^<^>vv<<v<<>v>v^^^v<^<^^<>><v<^>^>vv^>vvv>>><^>>>>>v>>^v<>>>^^>>>v^<^v>v<^>^>v<^^<v^^<^v<<^^<v^>>>>>>^>>v>vv<>^>^v>^>^><v^v<>^>^^<vv><v>^v><v>^>^v<>^^v^v^<v<<^<v^v^v<^>>v><vv<<vv<v^>v>^<^v<<><^^^>^v<v>^<^^<v<^<<>v<><^^vvv<v<v>v>^v<<v^>><>>v>^^v<<><<v^>^>><^^^>^>vvv^v^<^>^v^v>><^>^^<v<v><<><vv<<<^>v><><^>>>vvv><>v>><><>v^^<><<<<^v<><<^<^v^>v<vv^vv>>>^v^><v^^v>>v<v^<>v^>><^<^v^^<>^<^><>>><vv^>>^<<>v<<<>v^vvv^^
v<^>vv>^>v><v^<>>>>^<v^>><v<<v>^^>^<^v>>>^>v>vv^vv^>>^><vv><>><vv<^^>v^^>><v<<<<<v^^><<^<v<v>vv<><>vvvvvv>^^^>v^>v><v^<>><vv<><^><^v>v^^v<v<v<>>>^^>^v^^vvvv>^>>^<>v>^v>>><^v<><>>>^^>v^v^<<<^v>v>>>^>v^^v>^v>v<^<v<vv<><>^^>^>><v>>v^v<^<^v^<v<<^vv^<<>v^><^^v^v<<><^>^^>v^>v<^><^v<>><>>>>>^^<v<>v>v<>^><>><^vvv<>^^>><><>^<v^v<><>v^^>^v^<v^vv<^<<^<<^<>v<^<vv<>>^<>^^^^<<^<^><>><<<<v>>>>vv<^<<>>^^v>>^^v<v<<v^>>>v<><<<v<^v<^>>>>>>>^<v<<vv<v>vv^<>vvv^<v>v><v^^^^^^v<^>v>^^vv^^><v^<^v<^^^^<>v<<v><<v<v<v^>>vv<^<^v><^>vv<^<v>v<v^<<><>^<^v>>>>^<>>v^vvvvv^v<<^>^<<^>vv^>><<^<v^^<v^<>^>v>^><vvv><v^v<>>v^v>^^v^^vvv>v>>^vvv^<>>><v>>v>^>^<v<^v^v^>^<v<^v<<<v><>>^<>>v^^>^<<<>^v><<>>^<>v^^v>^>^v<>v^<^>vvv>><><<^<<vv^^<vv<<<<v^^<^>>v>^><^<v^<^<<>v<<^<<^>^^vv<v^^^^<^<<<<<<<^^<>v<v^>^>v>^vv<^^<^<vvvv<<v^^<><v^>><<v><^^^>^>v<<^^^>>><>>v>v<v>^>^<<vv^><v^^v^<v^v^^^<v^>>^<vv<<vv<>>^<^^v^vv>^>^><>>vv<^^<v>^>>>^<<v^<<^^^v<<<vvv^>>v<^v>^<vv<><><^v^^^vv^>^v<<^^^<>>^<vv<>^><v<<^^<<>><<>^v><>^<><<>v<<><><<>>^><^<^^<><v^>v>
^^^>v<^<><vvvv<vv<v^^^><^v>^>vv^v>><>^>v>v>v^^>>^<>v^v<>v><<v<v>^<>v^>v^<vvv<^<><<<^>^v><<^<vvvv>vv>>><vv<<v^<<vv>^>v<<v<<<<<v><><^><v^v>v<v<><^v^>v^><><^<^v^v^vv<^<>v><>^>><^><<^<v>vv<<v^<^>v^>><<<><>vv>>vv<v<^>^>>>>v<vv<^>v<<><^>v><>v>>v<v><^v>v><^>v<<v<>v<<>^v>>vv><^^v^v^>>^^^^^><><<v^^<>><v><vv^v>>>^>^^^^v><v><>>^>>>^^^>vv>v>v<^<v><^<v^v>^>v^^v^v<v>^v>v^v>v>><v^>^^v><^>>>^v>>v<v>^>^>v>v^^^v^^^<>^v><^v><>>v>v^>vv<>v<>>^v<>vv>><>v>v>vv^<^^<<<v^^v^^vvv^v>vv<<^>^v>>>v<vv^<>>^vv<<<>>>><v<>v^^^vvv<^^^v<>v^^><<>^<^vv<^^^v^<^vv><>>^<vv<><^v^><^^v><><>vv<vv<vv^>vv<><^^v^v^vv<<^v<>vv^>^>^<vv>^<<^>vv^>><^vv^v<vv<<<>>^v^><v^>><^>^v<vv^<<v<v^v^><<>>>^>>>>v^^v>^<^<>>^<v^^v^^<<>>^v>>^v<v<>>v^>v^>><^>vv<<<v^>><><<<>>v^v><^>v><<>^^v<<<<>^<v^>^<v^><>^vv^><v^<v^<>>><>v><><^>^v>>v<>^^>>^>><v^^^<<>><vv^v>v^><>v^^^<>>v^v>v<v^^^><vv>^^>><>^><<v^>^v<^^>^^<^^^^v<>>^^^<>>v^^vv>><>^<v<>v^<^^^^^^><vv^<^>>vvvv^v<v<^>^^<>v<><^><v>>>>v^>><v<v>^><>>v^>v^><>^<<<^^>>^<<<<^^vv>>^^>^<<v>^<v<^>><^<v>^><v^<v<><>v>v<vv<
^<^>>>v^v<<<^>>>v<<>>^<>^>^v>^v<v^>>^^vv>^^^^v><^^^<>vv>v>>^^<v<^<<><<v>><^^v><^^>^<>>^vvv>>>v^><^<v>v<<<v^v<vv<<<<<>vv<><<<>v^>vv^^v>>vv^v^^<vv^^>v^^v<>^<<^^>>><v<^^<<>^>^^v<^^vvv<>v<>><<v^<<<>v<^><vvv<<>^^v>>>v<^^>v>^>><<v<<<v<^>>>^^^<^^v<^^v<v>>><^>^^>><v<v^<<v<>>v><^v>^<vv>>v^vvv^>^<v<^^<<v<v><<<><^<<v>v^^<>^>>><^v^<><^^>^><vv^v<>>>^><^v<^^<<>vv><<vv<<v<v><><>v^<v>^>>>v<<v^<>v>vv<v>>vv^<>v<>v<><<v>>^v^v>v<>>^<>^vv^^>vv>v>^^^v>>>v<<^<v<<<<>^>vvv>><v>v<<>v><^^^>>vv>><^v<v><vv><<><><^^>^<<^^v^^<vv<vvv<v<>>^<^v>v<<<^^>>^^^>>>v<>^v>>v^>><<>>^^vv<>v<<v^^><v<>v<v>v<>>>>><<><v<vv<><^v^vv^v>><<<^<<>><>vv>v>>^^<>><v<<><^><^^^v<^v><>^vv^v><>^vv<>v^>^<>v>>v^>^><^><v^^^^v>^^>vv><v<>v<>v><><v><>>v<^<>>>v^<<>>>^<v>v<>v<^^<>^>^v>>v>^^>>v<<v<<v>^>v>>^v<^<^v>>>v<v><vv<vv>>v>v>>^>vv^<>v>v^<<><vv^vv><<^v^<>>>^v<vv>^<v^^>>^^<v>>v><>v^^>v>vvv^<v>>><<<v><<^><<<vv<^v^v>^>^><>^vv>vv<><<vvv^vv^<<<<<^<<>>^>v<v^^<>v>v^vv>><<>vv><<><>>v^^>^^^v>vv^v^^v^<^v><>v<><^>^<<<^v<<<^^v^><>>^^<>v><>>vv>^<^>^>>v>><<v^>v<>
<^<v^<>v<^<^>>v<>><<^vv^<<>v^v^v^><>v<v<<v><vv>v^^>^^^^<v<<^<<>><<^v>^<<<>vv<>^>><<<v>v<^>>^^^v^v>^vv^v>v^<vv^^<v<<<^^^>>v><>^>^>^>>>vv>><><><<<^<<^vv<>>><v<>>^^<<vv<><v<vv>><>v^<vv<<v>^<^^>>>>><><<^<<>^^^v>^<>v^<^<^^>^<v^><^v>^>^>><v><v<v<>^^v>^<v<v>^>^v<>vvv<<>>>v^v><^^vv<>^^>^<^vvv^^<<>><v^>>>vvv<>v<><>v^><<^v>^v^<<>v^^^^v<>^><^><<^vvv<^<^^>^^>v><>>^><>>>v>^<^^<>v^<v^>^v^v>v^vv>vv><>^vv<<><<^>v>>>>>^v><v^<>^v<^><^v<v<>^vv^^^v>^><^v<^^^<><>vvv^^>>>>>>>^v<<v^^>><v>>>^v<<v^>>>^^<^<v>>^>>><<^v>>^v>^<>^<><<><v<^<vv>^>>>v>v^<vv^>><v^^v^>^^<^^v<<v><vv>>v<<>v<v<v>>vv>v<<v<<<<v>^^<v<^>>><^^>vv<vv^<^><^>v<><^^v<>^>>v>^v^v<>v>v>v>^vv<^^^><^<>^>vv^<<v^<<<<<<v<v<>^<>>>v>^v>><vv^^v<>^>vv^><v<v^v<<v<^>v>>^^>>vv<vv><>^<vv<v<v>>v<^v<<<>>><><><^^><v^v>^^v>vvv>v><^v^<<>vvvv><^^^vv>>v^<v^^^^<<^^><<^v>>^^>>>><^<^>>>vv>>>vv>^^>^^<^v^^<<^v^>><^^^>><v<>^v<vv>>v>^v^>><>^v^>^><<vvv^^v^^^v>vv<v>>>v>v^v<>>^vv^>v<^v^><>>^v<v^<><<>vvvv>^>^v^v<vv^><v^v>>>v>v^v<>vvv><<vvv>>v>^vv^vvvv^<^>^<<^v<<<<>><^^<>v<^^<>v^v>>
<vv^<^^^^<<>^^<<v<^^>>>^^^^<<^^>>^^v>^vv>>v><>^<^v^^>v<<vvv>^><v<>^<vv^^<^^>v>^><^^<<><>^v^><>>v<>v<>v<^v^>>vvv^^<>^><<v^^^^vvv>>^>^^<<>^>v><v^>>>v^<vv^v>v^<^><<<<v^^><vvv<>^>>v<<>^^v><^^<<<<^>^>v^v><<>^><^^>>v<<<^<>v>>vv>v^>v^>^<>>^vv^v^v>vv^<<^>v^v<<<v>v<v<^v^>>^^><vvvvvvv<^v^>^v<^^^vv^^<^>>^^>v<v>>>><>v<>^><v<>v>^>v<<><^v^^^^><>>>v<>><>^v<^<vv^>>>v^^v><v^^<v^>^^<<<<>>><<^<<vvv^^^<<><<<<<<v^<><vvvvvv>v^<<v<^<>>v^v<vvv>^^^v><<<>>^^v>><<^<<v<v^<v^><><v^v^^<<>><v^^<vv^>^v<><<>v<v^^<^<vv^><^v<><^v><v^^<<>>^>>v>^><>v^<<^<v<>^<>^<^^^>v<^<v<v<<<<v<^<><^v>v<v<v>^<<<v>>>vv<<><<>v><<v<vvvv>><>^>^vv^v<<>^^<>>^^><<<vv<^>>vvv^<^><>^^>^>>^v<>v<v<>^>>^^>>v^vv^^<<v^^<v^v><<v<<v<v>>>>>><><<v^>>v>>>^>^v><^^<<<^^<<<^<vv^>^^<v<v<vvv^vv^^v<>v<^^^^<^v><>>vv<<>v<v^>><><v<v>^><^><>^<><vv<^><>vv^v^<v>^><<<>>^v^>v^^>^^<>^v^>>v<>>>^^v<^^v^>>>>><>^v>^^>v>^^><^<<^>vv^>><vvvv><>vv><<<vv<>vv^v<>>><^^^^vv^><^v>>^>^>>v<^^v^v^>v^>>^<v<>^^><^<v^v>v>^vvv>^^><^v<><^v<^>v^<>^v<^vv<><<<><^><><<v^vv<^>vv<v><v^>>v>^^^<^><<<
>^<v^>>v^v>v><>^><^^>^v<^vv><>^v>>v<v>>>v^<^^><><vv<^>>vv^^^vv>>>v><>v><vvv><v^v^<>^<v><>^<><^>v<<><^<>^>^^^<v><>>v^<><v^vvv<^v^<<^>^<^v>v>vv<>v>^v><vv^<><v>^v<>vv^v^v<>vv<<^vvvvv^^>>>>^v>>^>^v<><v^<^v>>>>^<^>>v^vv^^>><>^><>><><v<^^v><<v<><>^<<v<v><v<<^>^>><<^^vvv<vv^>vv>>v>^><v<^<>^>vv>>>>>>v<^>>v><v>v>v^v><>^>>vv<vv<^>>v<><>>^v><^<<>>vv<<v^^>>>^^<<<>>^><^^vv^v^<vvv>v^><>^v^>>^^^<v^v^>v<>^^^vvv^>^>^v<<v>v^>^>v>^<><^^>>><vv<vvvv>><<vv>^<^>^vv><>^<<^<vvvv^^^>^^vv<vv<<v^v^<>^>v<<><<<>><^^>><><>^^^>^>>v<>>^>^>^v^<>><>^^^^vv<^<>>>>>>^v^<^>>>v<<>v<<>v>^vvv>>v<^^>v<^<vv<>v^v>^v^<^^v<v><v^vv^>v<<<<^vv>>>><<v>><^>><>><>>v<>v^<<^^v<v^><>vv><^>v><vvv>><>vv><vv<<>^v^<^<^<<^v>^vv><><v<<vv^v^v<v^^>>v^^<v<><^v>^^<<^v<<^vv<v>v^>^<<<vv<v><v^^>v>>vv^^v^<><>^vv<v>>>>^v>>>^v<<<><<<v>^>v><<><v^^^v^<^^<^v<v<v^^v^^>v>>^^vvv>>>v>^^v>^v>>^<<^><><^^<<^>><^<>vv<vv^v<v^v^^^><<<v<<v<v^v<<<v<>><>^<>><<<><^v<vv^v^><>v^vvvvv^^><^^^v<^<<^>vv<^<^v<^>^<<v>^v>>>v<>>^>>><^^>vv^v<v<^^<v><v^<<<><<^<><^v>v^>>^^<v>>><>v^<v<<
v<v>v<>>^<^vv^>>^^<<^>>vv<>^^>^v>vv>^>vv^><<v<^v<>^<v>v^>^>vv^<><><v^<<>vv^<>^>^<^<>v<v<vv>v><<>v<v^v^<v<^^^><>vv>v><>^^>^>>><vv<v>><vv<<<^<<>^>v<v><<v>v<<v^>vvv>^^vv^<v^^<^>vv^^v^><^<<^^v>v<<<^><<>>v<v<>vv^vv<v>^<<v^v>^^vv>v<<<<vv^vv>>v<<<^><>v^^v><>>v>>^^^^<v>>>^>>>^vv>>^v<>>v>v^<v<^vvv<v<v<>v<v<<<>>^>>v^>><<v<^>vv^v>><^^^<<>^^>vv<<^^<v>><<<^^<^><>^>>>>v>><>^><<<<v<<<^<>>vv>^>v<^<^v^^v^>^^>v<><><><<v^^v<v><v<<<v^^v^^<>^><v<vv>v^>^^v^v^v<v<>>^<<v<vv>>^^<^^<<v^v<<>><^>v<><>v^^v^>>^vv><^v>v>vv^><v<>vv>v><^^^<<v>^^<v>v<^v<v><<<>vv>v><<<<^^<<><<>v><vv><<^v>vv^>v<<^<vv^^>^vv<^v<>^>><^^<^<vvv<^^^v^vvvv>^v>vvv<<><>><vv><v><<v^^v>^v^v<<<<v><vv><^v<><<<>>^v>^^v><>v<>><v<^<v^<>v>vv<>v<v^<v^^^vv><^>v<<><<v<v<><><^>>v>^^<<v<<^vvvv>v>>v<^>>^v^><vvv><>^><^^v<<>v<>v^^v<<>>><<<<v<vv><vv^^^><><>^>>v><^^^^v>^^><<^<^>v<^<<>^^vv^v><^<v<vv<>v^>^<>v>>v>v><<<<<>v<^<>v<>^^>^^>>>vv<<<^^^v>v<^<v>^>v^<>>v<v>v<<>v^<^>^v<v<><>^^<>>><^><v^^^>^^vv>><<><v<^<>v>>vv<^v<^^>>^<>v^^^<^>^^^^<<v^vvv><<>^>^^v^v<v^<>><^<<<><
v^^>>v^>^>v><>>^<<v^^vv^>>>>>>>>^^<^^>^v>^^^v<><<^vv<>^>><<<<>v<^<v^^>><>vv<><>^vv^>>>vvv>^>v>>^v<v^v^vv^>v<^^^^>vv<v<<^<<>v<<v<v^<<<^vv^v<^>>vv>>><>v^^v>>^><><^v><><vv>^v>>v<<>>>>vv^v^<v>v>vvv><^^^v<<<^v<^v>v^<^>^v>v>^><v>v<<><v<v<>v^v^vv^^<>v^v>^>v^>^^v><vvv^<v<^^^<>><>^v>^^<v<>v>vv^^>>^^v<<v<>><v<v<v>>^^<^<<^><<<^vv<>>><^^>^<>^v^^<>^<^<<^><>>^v^><>^v<^^v>>>v^<><v^>^<><<v>>^<vv><<<>^^vv^<>^^<^^v>^>^>vvv<v>vv<v<^>><><^^v>>^><>v><^><<vv^v^^^vvv^>>v<<<<><<<^<><<>v<>^<^>v>^v^>><<>^v>v^^>^^<>^^^<>v<>vv<>v^><<<<><v><v<<>^^^<<>v^v^v<^<><v^v>^^>v>>^>>^>^><v>v^<^^>v<^><v>>vv^v^^>^<^>v^v<>v^><^<^v^<<<><><>>>^vv^><v<^><>><^^>>vv><><^<v^<^v<^<vvv<>^^>v^v^>>>>>v^^<>v^^>vv<<<^v><vv^><>^<^<^v^vvvv<><>^><v^^<^><v>><<v<v^><>^<^<<^^vvv><<vv^^^<<<<<>v>v<><^v<v<v^<v^>><^>^^<><>^^vv^><vv>vv<^v<<><<^v^^>><^^<^><vv><^<<v>>^<vv>>^<^^<>^^<^>^<>v^<^v><>>>v^>^><<>^^>^<>>vv<^^><<><><v>v^^<<v^v>>vv<^^<<^<v^^^^<v^>^^^v><^v<vvv><<^^<<>v^v>v<<^>><><<^v^<vv^^^>><<^^>><>>v<^><vv>^^^^<v<vv<v><vv>^v<>>v>^^<>v>><>^v<v<<`
