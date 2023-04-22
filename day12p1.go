package main

import "fmt"
import "os"
import "strings"

type coords struct {
	x int
	y int
}

var visited map[coords]coords //key is place that we are, value is place that we came from
var queue []coords
var field [][]int //only holds elevation data. the indices are the coordinates
var START_POINT coords
var END_POINT coords

func main() {
	filecontents, err := os.ReadFile("./" + os.Args[1])
	if err != nil {
		fmt.Println("oh no it broke")
		return
	}
	file := string(filecontents)
	visited = make(map[coords]coords)
	queue = make([]coords, 0)
	field = make([][]int, 0) //makes a 2-D slice of ints, with zero rows. Each row must be initialized and appended later.
	process(file)
}

func process(file string) {
	lines := strings.Split(file, "\n")
	for count, line := range lines {
		field = append(field, make([]int, len(line)))
		row := []rune(line)
		for inner, r := range row {
			if r == 69 { //E
				field[count][inner] = 26
				END_POINT = coords{inner, count}
			} else if r == 83 { //S
				field[count][inner] = 1
				START_POINT = coords{inner, count}
			} else {
				field[count][inner] = int(r - 96) //subtract the starting offset of lowercase letters

			}
		}
	}
	BFS()
}

func BFS() {
	queue = append(queue, START_POINT)
	count := 0
	for len(queue) > 0 {
		for _, spot := range searchAround(queue[count]) {
			visited[spot] = queue[count]
			queue = append(queue, spot)
			if spot == END_POINT {
				traverse()
			}
		}
		count++
	}
}

func traverse() {
	count := 0
	currentLocation := END_POINT
	for currentLocation != START_POINT {
		currentLocation = visited[currentLocation]
		count++
	}
	fmt.Println("Final Score: ", count)
	os.Exit(0)
}

func searchAround(centre coords) []coords {
	output := make([]coords, 0)
	if centre.x > 0 { //search to the left but only if not on the left edge
		_, found := visited[coords{centre.x - 1, centre.y}]
		if (!found) && (field[centre.y][centre.x-1] <= field[centre.y][centre.x]+1) {
			output = append(output, coords{centre.x - 1, centre.y})
		}
	}
	if centre.x < len(field[0])-1 {
		_, found := visited[coords{centre.x + 1, centre.y}]
		if (!found) && (field[centre.y][centre.x+1] <= field[centre.y][centre.x]+1) {
			output = append(output, coords{centre.x + 1, centre.y})
		}
	}
	if centre.y > 0 {
		_, found := visited[coords{centre.x, centre.y - 1}]
		if (!found) && (field[centre.y-1][centre.x] <= field[centre.y][centre.x]+1) {
			output = append(output, coords{centre.x, centre.y - 1})
		}
	}
	if centre.y < len(field)-1 {
		_, found := visited[coords{centre.x, centre.y + 1}]
		fmt.Println(coords{centre.x, centre.y})
		if (!found) && (field[centre.y+1][centre.x] <= field[centre.y][centre.x]+1) {
			output = append(output, coords{centre.x, centre.y + 1})
		}
	}
	return output
}
