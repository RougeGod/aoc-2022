package main

import "fmt"
import "strings"
import "strconv"
import "os"

type coords struct {
	x int
	y int
}

var filled map[coords]bool
var maxY int

func main() {
	filecontents, err := os.ReadFile("./" + os.Args[1])
	if err != nil {
		fmt.Println("oh no it broke")
		return
	}
	filled = make(map[coords]bool)
	file := strings.Split(string(filecontents), "\n")
	for _, line := range file {
		parseLine(line)
	}
	fmt.Println(simulate())
}

func parseLine(line string) {
	noarrows := strings.ReplaceAll(line, " -> ", ",")
	spots := strings.Split(noarrows, ",")
	for count := 0; count < len(spots)-2; count += 2 {
		fillSpaces(space2coord(spots[count], spots[count+1]), space2coord(spots[count+2], spots[count+3]))
	}
}

func simulate() int {
	score := 0
	ended := false
	var pos coords = coords{500, 0}
	for !ended {
		if pos.y > maxY {
			ended = true
		} else if !filled[coords{pos.x, pos.y + 1}] {
			pos.y++
		} else if !filled[coords{pos.x - 1, pos.y + 1}] {
			pos.x--
			pos.y++
		} else if !filled[coords{pos.x + 1, pos.y + 1}] {
			pos.x++
			pos.y++
		} else {
			filled[pos] = true
			pos = coords{500, 0}
			score++
		}
	}
	return score
}

func space2coord(x string, y string) coords {
	xpos, _ := strconv.Atoi(x)
	ypos, _ := strconv.Atoi(y)
	maxY = max(maxY, ypos)
	return coords{xpos, ypos}
}

func fillSpaces(c1 coords, c2 coords) {
	if c1.x == c2.x {
		for y := min(c1.y, c2.y); y <= max(c1.y, c2.y); y++ {
			filled[coords{c1.x, y}] = true
		}
	} else if c1.y == c2.y {
		for x := min(c1.x, c2.x); x <= max(c1.x, c2.x); x++ {
			filled[coords{x, c1.y}] = true
		}
	} else {
		fmt.Println("something is terribly wrong")
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
