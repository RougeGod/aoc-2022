package main

import "fmt"
import "os"
import "strings"
import "regexp"
import "strconv"

type command struct {
	dist      int
	turnRight bool
}

type coords struct {
	x int
	y int
}

var rowEnds [][2]int
var colEnds [][2]int
var colStarted []bool
var walls map[coords]bool //true iff there is a wall tile there

func main() {
	filecontents, err := os.ReadFile("./" + os.Args[1])
	if err != nil {
		fmt.Println("oh no it broke")
		return
	}
	file := strings.Split(string(filecontents), "\n")
	rowEnds = make([][2]int, len(file)-3)
	colEnds = make([][2]int, 0)
	colStarted = make([]bool, 0)
	walls = make(map[coords]bool)
	var commands []command
	for index, line := range file {
		if index < len(file)-3 {
			process(line, index)
		}
		if index == len(file)-2 {
			commands = parseCommandString(line)
		}
	}
	var facing int = 0
	var pos coords
	pos.y = 0
	pos.x = rowEnds[0][0]
	//for k, _ := range(walls) {
	//   fmt.Printf("(%d, %d)\n", k.x, k.y)
	//}
	for ; walls[coords{0, pos.x}]; pos.x++ {
	}
	for _, comm := range commands {
		pos = move(pos, comm.dist, facing)
		if comm.turnRight {
			facing++
		} else {
			facing += 3
		}
		facing %= 4
	}
	facing = (facing + 1) % 4
	fmt.Println(1000*(pos.y+1) + 4*(pos.x+1) + facing)
}

func move(pos coords, dist int, facing int) coords {
	for count := 0; count < dist; count++ {
		newPos := getNewPos(pos, facing)
		if walls[newPos] {
			return pos
		} else {
			pos = newPos
		}
	}
	return pos
}

func getNewPos(pos coords, facing int) coords {
	switch facing {
	case 0:
		//fmt.Println(pos.x, rowEnds[pos.y])
		if pos.x == rowEnds[pos.y][1] {
			pos.x = rowEnds[pos.y][0]
		} else {
			pos.x++
		}
	case 1:
		if pos.y == colEnds[pos.x][1] {
			pos.y = colEnds[pos.x][0]
		} else {
			pos.y++
		}
	case 2:
		if pos.x == rowEnds[pos.y][0] {
			pos.x = rowEnds[pos.y][1]
		} else {
			pos.x--
		}
	case 3:
		if pos.y == colEnds[pos.x][0] {
			pos.y = colEnds[pos.x][1]
		} else {
			pos.y--
		}
	default:
		panic("Facing out of range!")
	}
	return pos
}

func parseCommandString(line string) []command {
	var rightTurn bool
	var dist int
	var output []command = make([]command, 0)
	line += "L"
	commands := regexp.MustCompile("[0-9]{1,2}[L|R]").FindAllString(line, -1)
	for _, com := range commands {
		switch com[len(com)-1] {
		case 'R':
			rightTurn = true
		case 'L':
			rightTurn = false
		default:
			panic("Invalid Turn")
		}
		dist, _ = strconv.Atoi(com[:len(com)-1])
		output = append(output, command{dist, rightTurn})
	}
	return output
}

func getRowStart(line string) int {
	firstPound := strings.Index(line, "#")
	firstDot := strings.Index(line, ".")
	if firstPound == -1 {
		return firstDot
	} else if firstDot == -1 {
		return firstPound
	} else if firstDot > firstPound {
		return firstPound
	} else {
		return firstDot
	}
}

func process(line string, row int) {
	rowEnd := len(line) - 1
	rowStart := getRowStart(line)
	rowEnds[row] = [2]int{rowStart, rowEnd}
	if rowEnd > len(colEnds) {
		//		oldColCount := len(colEnds)
		colEnds = append(colEnds, make([][2]int, rowEnd-len(colEnds)+1)...)
		colStarted = append(colStarted, make([]bool, rowEnd-len(colStarted)+1)...)
		for count := rowStart; count < len(colEnds); count++ {
			if !colStarted[count] {
				colEnds[count][0] = row
				colStarted[count] = true
			}
		}
	}
	if rowEnd < len(colEnds) {
		for count := rowEnd; count < len(colEnds); count++ {
			if colEnds[count][1] == 0 {
				colEnds[count][1] = row - 1
			}
		}
	}
	for col, char := range line {
		if char == ' ' {
			if colStarted[col] && colEnds[col][1] == 0 {
				colEnds[col][1] = row - 1
			}
		} else {
			if !colStarted[col] {
				colStarted[col] = true
				colEnds[col][0] = row
			}
			colEnds[col][1] = row
			if char == '#' {
				walls[coords{col, row}] = true
			}
		}
	}
}
