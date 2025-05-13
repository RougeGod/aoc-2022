package main

import "fmt"
import "os"
import "strings"
import "strconv"
import "time"

type state struct {
	timeRemaining int
	pos           string
}

type twostate struct {
	hstate state
	estate state
	opts   uint
}

type graphNode struct {
	position   string //the valve's name
	neighbours []string
	flowRate   int
}

var graph []graphNode
var valveEnum map[string]int //our distance matrix works on int/int, this gets the string value, if necessary
var dists [][]int
var options []string
var optEnum map[string]uint
var solved map[twostate]int

func parseLine(line string, ord int) {
	remLine := strings.Replace(line, ",", "", -1)
	remLine = strings.Replace(remLine, "rate=", "", -1)
	remLine = strings.Replace(remLine, ";", "", -1)
	input := strings.Split(remLine, " ")
	rate, _ := strconv.Atoi(input[4])
	graph[ord] = graphNode{input[1], input[9:], rate}
	valveEnum[input[1]] = ord
	if rate > 0 {
		options = append(options, input[1])
		optEnum[input[1]] = 1 << len(options)
	}
}

func main() {
	filecontents, err := os.ReadFile("./" + os.Args[1])
	if err != nil {
		fmt.Println("oh no it broke")
		return
	}
	file := strings.Split(string(filecontents), "\n")
	graph = make([]graphNode, len(file))
	valveEnum = make(map[string]int)
	optEnum = make(map[string]uint)
	solved = make(map[twostate]int)
	for i := range file {
		parseLine(file[i], i)
	}
	FWA(len(graph))
	t := time.Now().UnixMilli()
	fmt.Println(score(state{26, "AA"}, state{26, "AA"}, options))
	fmt.Println(len(solved))
	fmt.Println(time.Now().UnixMilli() - t)
}

func score(hstate state, estate state, options []string) int {
	var redOptions []string
	if hstate.timeRemaining <= 0 && estate.timeRemaining <= 0 {
		return 0
	}
	if len(options) == 0 {
		return 0
	}
	var optNum uint = 0
	for _, valve := range options {
		optNum += optEnum[valve]
	}
	solution, ok := solved[twostate{hstate, estate, optNum}]
	if ok {
		return solution
	} else {
		//maybe we've seen this before but human and elephant in opposite states
		solution, ok = solved[twostate{estate, hstate, optNum}]
		if ok {
			return solution
		}
	}
	maxScore := 0
	for index, valve := range options {
		var h int
		var e int
		redOptions = append([]string(nil), options...)
		redOptions = append(redOptions[:index], redOptions[index+1:]...)
		if hstate.timeRemaining > dists[valveEnum[hstate.pos]][valveEnum[valve]] {
			h = score(state{hstate.timeRemaining - dists[valveEnum[hstate.pos]][valveEnum[valve]] - 1, valve}, state{estate.timeRemaining, estate.pos}, redOptions) + graph[valveEnum[valve]].flowRate*(hstate.timeRemaining-dists[valveEnum[hstate.pos]][valveEnum[valve]]-1)
		}
		if estate.timeRemaining > dists[valveEnum[estate.pos]][valveEnum[valve]] {
			e = score(state{hstate.timeRemaining, hstate.pos}, state{estate.timeRemaining - dists[valveEnum[estate.pos]][valveEnum[valve]] - 1, valve}, redOptions) + graph[valveEnum[valve]].flowRate*(estate.timeRemaining-dists[valveEnum[estate.pos]][valveEnum[valve]]-1)
		}
		if h > maxScore {
			maxScore = h
		}
		if e > maxScore {
			maxScore = e
		}
	}
	solved[twostate{hstate, estate, optNum}] = maxScore
	return maxScore
}

func FWA(numVerts int) {
	dists = make([][]int, numVerts, numVerts)
	for count := 0; count < numVerts; count++ {
		dists[count] = make([]int, numVerts, numVerts)
		for inner := 0; inner < numVerts; inner++ {
			if inner == count {
				dists[count][inner] = 0
			} else {
				dists[count][inner] = 9999
			}
			for _, n := range graph[count].neighbours {
				enum, OK := valveEnum[n]
				if OK {
					dists[count][enum] = 1
				}
			}
		}
	}
	for k := 0; k < numVerts; k++ {
		for i := 0; i < numVerts; i++ {
			for j := 0; j < numVerts; j++ {
				if dists[i][j] > dists[i][k]+dists[k][j] {
					dists[i][j] = dists[i][k] + dists[k][j]
				}
			}
		}
	}
}
