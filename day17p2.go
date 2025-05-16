package main

import "fmt"
import "os"

type coords struct {
	x int8
	y uint64
}

type result struct {
	rocks uint64
	y     uint64
}

type state struct {
	topography [350]bool
	windex     uint16
	rindex     uint8
}

var filled map[coords]bool
var seen map[state]result
var maxY uint64
var rocks [][]coords

const TRIL uint64 = 1000000000000

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
	} else {
		return coords{-1, 0}
	}
}

func allCoordsClear(rock []coords, rockpos coords) bool {
	var newpos coords
	for _, diff := range rock {
		newpos = add(diff, rockpos)
		if !coordClear(newpos) {
			return false
		}
	}
	return true
}

func coordClear(place coords) bool {
	return !((place.x >= 7) || (place.x < 0) || (place.y <= 0) || (filled[place]))
}

func fillRock(rock []coords, rockpos coords) {
	var newpos coords
	for _, diff := range rock {
		newpos = add(diff, rockpos)
		filled[newpos] = true
		if newpos.y > maxY {
			maxY = newpos.y
		}
		for i := int8(0); i < 7; i++ {
			delete(filled, coords{i, maxY - 50})
			delete(filled, coords{i, maxY - 51})
			delete(filled, coords{i, maxY - 52})
			delete(filled, coords{i, maxY - 53})
		}
	}
}

func getTopography() [350]bool {
	var output [350]bool
	for x := 0; x < 7; x++ {
		for y := 0; y < 50; y++ {
			if coordClear(coords{int8(x), maxY - uint64(y)}) {
				output[x+7*y] = true
			}
		}
	}
	return output
}

func transferMap(ydiff uint64) {
	var newFilled map[coords]bool = make(map[coords]bool)
	for k := range filled {
		newFilled[coords{k.x, k.y + ydiff}] = true
		delete(filled, k)
	}
	for k := range newFilled {
		filled[coords{k.x, k.y}] = true
	}
}

func simulate(sequence string) {
	var windCount uint16 = 0
	var pos coords = coords{2, maxY + 4}
	var currState state
	seen = make(map[state]result)
	for rockCount := uint64(0); rockCount < TRIL; rockCount++ {
		for {
			if allCoordsClear(rocks[rockCount%5], add(pos, dir(sequence[windCount]))) {
				pos = add(pos, dir(sequence[windCount]))
			}
			windCount++
			windCount %= uint16(len(sequence))
			if allCoordsClear(rocks[rockCount%5], coords{pos.x, pos.y - 1}) {
				pos = coords{pos.x, pos.y - 1}
			} else {
				fillRock(rocks[rockCount%5], pos)

				if maxY > 50 { //topography breaks if y <= 50, and floor interferes anyways
					currState = state{getTopography(), windCount, uint8(rockCount % 5)}
					if (seen[currState] != result{0, 0}) {
						yDiff := maxY - seen[currState].y
						rockDiff := rockCount - seen[currState].rocks
						var cyclesRemaining uint64 = (TRIL - rockCount) / rockDiff
						rockCount += cyclesRemaining * rockDiff
						maxY += cyclesRemaining * yDiff
						transferMap(yDiff * cyclesRemaining)
					} else {
						seen[currState] = result{rockCount, maxY}
					}
				}
				pos = coords{2, maxY + 4}
				break
			}
		}
	}
}
