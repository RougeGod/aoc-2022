package main

import "fmt"
import "os"
import "strings"

type blizzard struct {
	x   int
	y   int
	dir int
}

type state struct {
	x    int
	y    int
	time int
}

type coords struct {
	x int
	y int
}

var GRID_HEIGHT int
var GRID_WIDTH int
var blizzards map[state]bool //true iff we cannot be there
var walls map[coords]bool

func main() {
	filecontents, err := os.ReadFile("./" + os.Args[1])
	if err != nil {
		fmt.Println("oh no it broke")
		return
	}
	file := strings.Split(string(filecontents), "\n")
	blizzards = make(map[state]bool)
	walls = make(map[coords]bool)
	GRID_HEIGHT = len(file) - 3   //newline, top wall, bottom wall
	GRID_WIDTH = len(file[1]) - 2 //walls don't count
	CYCLE_PERIOD := (GRID_HEIGHT * GRID_WIDTH) / gcd(GRID_HEIGHT, GRID_WIDTH)
	endPoint := coords{GRID_WIDTH, GRID_HEIGHT + 1}
	findWallsBlizzards(file, CYCLE_PERIOD)
	fmt.Println(BFS(state{1, 0, BFS(state{endPoint.x, endPoint.y, BFS(state{1, 0, 0}, endPoint, CYCLE_PERIOD)}, coords{1, 0}, CYCLE_PERIOD)}, endPoint, CYCLE_PERIOD))

}

func findWallsBlizzards(file []string, cycle int) {
	walls[coords{GRID_WIDTH, GRID_HEIGHT + 2}] = true
	walls[coords{1, -1}] = true //prevent expedition from going up past start point
	for y, line := range file {
		for x, char := range line {
			switch char {
			case '^':
				for t := 0; t < cycle; t++ {
					blizzards[state{x, ((y - (t % GRID_HEIGHT) + GRID_HEIGHT - 1) % GRID_HEIGHT) + 1, t}] = true
				}
			case '>':
				for t := 0; t < cycle; t++ {
					blizzards[state{(x+t-1)%GRID_WIDTH + 1, y, t}] = true
				}
			case '<':
				for t := 0; t < cycle; t++ {
					blizzards[state{((x - (t % GRID_WIDTH) + GRID_WIDTH - 1) % GRID_WIDTH) + 1, y, t}] = true
				}
			case 'v':
				for t := 0; t < cycle; t++ {
					blizzards[state{x, (y+t-1)%GRID_HEIGHT + 1, t}] = true
				}
			case '#':
				walls[coords{x, y}] = true
			case '.':
			default:
				fmt.Println("Unrecognized character")
			}
		}
	}
}

func BFS(s state, end coords, cycle int) int {
	var q []state = make([]state, 1)
	q[0] = s
	visited := make(map[state]bool)
	var top state
	for len(q) != 0 {
		top = q[0]
		//fmt.Println(top)
		q = q[1:]
		for _, n := range checkAround(top, cycle) {
			if !visited[state{n.x, n.y, n.time % cycle}] {
				q = append(q, state{n.x, n.y, n.time})
				visited[state{n.x, n.y, n.time % cycle}] = true
			}
			if (coords{n.x, n.y} == end) {
				return n.time
			}
		}
	}
	return 0
}

func checkAround(s state, cycle int) []state {
	output := make([]state, 0)
	if !isBlizzard(state{s.x, s.y + 1, (s.time + 1) % cycle}) {
		output = append(output, state{s.x, s.y + 1, s.time + 1})
	}
	if !isBlizzard(state{s.x + 1, s.y, (s.time + 1) % cycle}) {
		output = append(output, state{s.x + 1, s.y, s.time + 1})
	}
	if !isBlizzard(state{s.x - 1, s.y, (s.time + 1) % cycle}) {
		output = append(output, state{s.x - 1, s.y, s.time + 1})
	}
	if !isBlizzard(state{s.x, s.y - 1, (s.time + 1) % cycle}) {
		output = append(output, state{s.x, s.y - 1, s.time + 1})
	}
	if !isBlizzard(state{s.x, s.y, (s.time + 1) % cycle}) {
		output = append(output, state{s.x, s.y, s.time + 1})
	}
	return output
}

func isBlizzard(s state) bool {
	return blizzards[s] || walls[coords{s.x, s.y}]
}

func gcd(n1 int, n2 int) int {
	if n2 > n1 {
		temp := n2
		n2 = n1
		n1 = temp

	}
	for n2 != 0 {
		temp := n2
		n2 = n1 - (n1/n2)*n2
		n1 = temp
	}
	return n1
}
