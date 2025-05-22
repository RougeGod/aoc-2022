package main

import "fmt"
import "os"
import "strings"
import "strconv"

type state struct {
	timeRemaining int
	ore           int
	oreBot        int
	clay          int
	clayBot       int
	obs           int
	obsBot        int
}

type blueprint struct {
	oreCost  int
	clayCost int
	obsCost  [2]int //{ore, clay}
	geoCost  [2]int //{ore, obs}
}

const INF int = 1000000

var solved map[state]int

func main() {
	filecontents, err := os.ReadFile("./" + os.Args[1])
	if err != nil {
		fmt.Println("oh no it broke")
		return
	}
	file := strings.Split(string(filecontents), "\n")
	var score int = 1
	for index, line := range file {
		solved = make(map[state]int) //clear solved
		score *= maximize(state{32, 0, 1, 0, 0, 0, 0}, toBlueprint(line))
		if index == 2 {
			break
		}
	}
	fmt.Println(score)
}

func toBlueprint(line string) blueprint {
	var output [6]int

	pos := 0
	for _, word := range strings.Split(line, " ") {
		num, err := strconv.Atoi(word)
		if err == nil {
			output[pos] = num
			pos++
		}
	}
	return blueprint{output[0], output[1], [2]int{output[2], output[3]}, [2]int{output[4], output[5]}}
}

//add(state, farm(state), {0, -1, -cost...})
func maximize(c state, b blueprint) int {
	if c.timeRemaining <= 1 {
		return 0
	}
	if sol, ok := solved[c]; ok {
		return sol
	}
	outcomes := make([]int, 0)
	var poss [4]int = canBuild(c, b)
	if poss[3] <= c.timeRemaining {
		outcomes = append(outcomes, maximize(add(farm(c, poss[3]), state{0, -b.geoCost[0], 0, 0, 0, -b.geoCost[1], 0}), b)+c.timeRemaining-poss[3])
	}
	if poss[2] <= c.timeRemaining && !enough(c, b)[2] {
		outcomes = append(outcomes, maximize(add(farm(c, poss[2]), state{0, -b.obsCost[0], 0, -b.obsCost[1], 0, 0, 1}), b))
	}
	if poss[1] <= c.timeRemaining && !enough(c, b)[1] {
		outcomes = append(outcomes, maximize(add(farm(c, poss[1]), state{0, -b.clayCost, 0, 0, 1, 0, 0}), b))
	}
	if poss[0] <= c.timeRemaining && !enough(c, b)[0] {
		outcomes = append(outcomes, maximize(add(farm(c, poss[0]), state{0, -b.oreCost, 1, 0, 0, 0, 0}), b))
	}

	solved[c] = max(outcomes...)
	return solved[c]
}

//return value:
//{we have more ore/oreBots than we need, clay/clayBots, obs/obsBots}
func enough(s state, b blueprint) [3]bool {
	var output [3]bool
	maxOreSpend := max(b.oreCost, b.clayCost, b.obsCost[0], b.geoCost[0])
	output[0] = s.ore >= maxOreSpend && s.ore+(s.timeRemaining*s.oreBot) >= maxOreSpend*s.timeRemaining
	output[1] = s.clay >= b.obsCost[1] && s.clay+(s.timeRemaining*s.clayBot) >= b.obsCost[1]*s.timeRemaining
	output[2] = s.obs >= b.geoCost[1] && s.obs+(s.obsBot*s.timeRemaining) >= b.geoCost[1]*s.timeRemaining
	return output
}

//return value:
//{number of minutes to build [ore bot, clay bot, obs bot, geode bot]}
//includes time to build the actual bot as well as gather req. resources
func canBuild(s state, b blueprint) [4]int {
	var output [4]int
	if s.obsBot == 0 {
		output[3] = INF
	} else {
		output[3] = max(ceilDiv(b.geoCost[1]-s.obs, s.obsBot), ceilDiv(b.geoCost[0]-s.ore, s.oreBot), 0) + 1
	}
	if s.clayBot == 0 {
		output[2] = INF
	} else {
		output[2] = max(ceilDiv(b.obsCost[1]-s.clay, s.clayBot), ceilDiv(b.obsCost[0]-s.ore, s.oreBot), 0) + 1
	}
	output[1] = max(ceilDiv(b.clayCost-s.ore, s.oreBot), 0) + 1
	output[0] = max(ceilDiv(b.oreCost-s.ore, s.oreBot), 0) + 1
	return output
}

//returns the new state if no bots were built for t minutes
func farm(s state, time int) state {
	return state{s.timeRemaining - time, s.ore + time*s.oreBot, s.oreBot, s.clay + time*s.clayBot, s.clayBot, s.obs + time*s.obsBot, s.obsBot}
}

func add(a state, b state) state {
	return state{a.timeRemaining + b.timeRemaining, a.ore + b.ore, a.oreBot + b.oreBot, a.clay + b.clay, a.clayBot + b.clayBot, a.obs + b.obs, a.obsBot + b.obsBot}
}

func ceilDiv(num, denom int) int {
	if num%denom == 0 {
		if num >= 0 {
			return num / denom
		} else {
			return (num / denom) - 1
		}
	} else {
		if num >= 0 {
			return (num / denom) + 1
		} else {
			return (num / denom)
		}
	}
}

func max(numbers ...int) int {
	var output int
	for _, num := range numbers {
		if num > output {
			output = num
		}
	}
	return output
}
