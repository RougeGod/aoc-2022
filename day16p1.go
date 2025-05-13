package main

import "fmt"
import "os"
import "strings"
import "strconv"

type state struct {
	timeRemaining int
	pos           string
	options       []string
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

//parse step to guild the graph
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
	}
}

func main() {
	//solved = make(map[state]int)

	filecontents, err := os.ReadFile("./" + os.Args[1])
	if err != nil {
		fmt.Println("oh no it broke")
		return
	}
	file := strings.Split(string(filecontents), "\n")
	graph = make([]graphNode, len(file))
	valveEnum = make(map[string]int)
	for i := range file {
		parseLine(file[i], i)
	}
	FWA(len(graph))
	fmt.Println(score(state{30, "AA", options}))
}

func score(st state) int {
	var redOptions []string
	if st.timeRemaining <= 0 {
		return 0
	}
	if len(options) == 0 {
		return 0
	}

	maxScore := 0
	for index, valve := range st.options {
		redOptions = append([]string(nil), st.options...)
		redOptions = append(redOptions[:index], redOptions[index+1:]...)
		s := score(state{st.timeRemaining - dists[valveEnum[st.pos]][valveEnum[valve]] - 1, valve, redOptions}) + graph[valveEnum[valve]].flowRate*(st.timeRemaining-dists[valveEnum[st.pos]][valveEnum[valve]]-1)
		if s > maxScore {
			maxScore = s
		}
	}
	return maxScore

}

func FWA(numVerts int) {
	fmt.Println(valveEnum)
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
