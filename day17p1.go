package main

import "fmt"
import "os"

type coords struct {
	x int
	y int
}

var filled map[coords]bool
var maxY int

var rocks [][]coords

func main() {
	filecontents, err := os.ReadFile("./" + os.Args[1])
	if err != nil {
		fmt.Println("oh no it broke")
		return
	}
	rocks = make([][]coords, 5)
	rocks[0] = []coords{{0, 0}, {1, 0}, {2, 0}, {3, 0}}
	rocks[1] = []coords{{1, 0}, {0, 1}, {1, 1}, {1, 2}, {2, 1}}
	rocks[2] = []coords{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}}
	rocks[3] = []coords{{0, 0}, {0, 1}, {0, 2}, {0, 3}}
	rocks[4] = []coords{{0, 0}, {1, 0}, {1, 1}, {0, 1}}
	filled = make(map[coords]bool)
	file := string(filecontents)
	simulate(file)
	fmt.Println(maxY)
}

func add(c1 coords, c2 coords) coords {
	return coords{c1.x + c2.x, c1.y + c2.y}
}

func dir(d byte) coords {
	if d == '>' {
		return coords{1, 0}
	} else if d == '<' {
		return coords{-1, 0}
	} else {
		fmt.Println("oups")
		return coords{0, 0}

	}
}

func allCoordsClear(rock []coords, rockpos coords) bool {
	var newpos coords
	for _, diff := range rock {
		newpos = add(diff, rockpos)
		if (newpos.x >= 7) || (newpos.x < 0) || (newpos.y <= 0) || (filled[newpos]) {
			return false
		}
	}
	return true
}

func fillRock(rock []coords, rockpos coords) {
	var newpos coords
	for _, diff := range rock {
		newpos = add(diff, rockpos)
		filled[newpos] = true
		if newpos.y > maxY {
			maxY = newpos.y
		}
	}
}

func simulate(sequence string) {
	var windCount = 0
	var pos coords = coords{2, maxY + 4}
	for rockCount := 0; rockCount < 2022; rockCount++ {
		for {
			if allCoordsClear(rocks[rockCount%5], add(pos, dir(sequence[windCount]))) {
				pos = add(pos, dir(sequence[windCount]))
			}
			windCount++
			windCount %= len(sequence)
			if allCoordsClear(rocks[rockCount%5], add(pos, coords{0, -1})) {
				pos = add(pos, coords{0, -1})
			} else {
				fillRock(rocks[rockCount%5], pos)
				pos = coords{2, maxY + 4}
				break
			}
		}
	}
}
