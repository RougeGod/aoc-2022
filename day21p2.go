package main

import "fmt"
import "os"
import "strings"
import "strconv"


var scores map[string]int //for known scores
var monkeys map[string]string //for still unknown scores

func main() {
	scores = make(map[string]int)
	monkeys = make(map[string]string)
filecontents, err := os.ReadFile("./" + os.Args[1])
if (err != nil) {
	fmt.Println("oh no it broke")
	return
}

file := string(filecontents)
createmaps(file)
//I had a really stupid way of finding the answer and it wasn't really in the code when I got it so it has been removed

}


func createmaps(file string) {
	lines := strings.Split(file, "\n")
	count := 0
	for (count < len(lines)) {
		if (len(strings.Split(lines[count], " ")) == 2) { //the monkey just has a number
			num, _ := strconv.Atoi(strings.Split(lines[count], ": ")[1])
			scores[strings.Split(lines[count], ": ")[0]] = num
		} else {
			monkeys[strings.SplitN(lines[count], ": ", 2)[0]] = strings.SplitN(lines[count], ": ", 2)[1]
		}
	count++
	}
}

func getMonkeyScore(m string) int{
	score, found := scores[m]
	if (found) { //base case
		return score
	} else {
		command := strings.Split(monkeys[m], " ")
		switch (command[1]) {
			case "+": {
				return (getMonkeyScore(command[0]) + getMonkeyScore(command[2]))
			}
			case "-": {
				return (getMonkeyScore(command[0]) - getMonkeyScore(command[2]))
			}
			case "*": {
				return (getMonkeyScore(command[0]) * getMonkeyScore(command[2]))
			}
			case "/": {
				return (getMonkeyScore(command[0]) / getMonkeyScore(command[2]))
			} 
			default: {
				fmt.Println("oops")
				return 0
			}
			
		}

	}
	
}
