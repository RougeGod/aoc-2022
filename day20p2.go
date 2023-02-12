package main

import "fmt"
import "os"
import "strings"
import "strconv"


func main() {
filecontents, err := os.ReadFile("./" + os.Args[1])
if (err != nil) {
	fmt.Println("oh no it broke")
	return
}

file := string(filecontents)

process(file)
}

var numbers []numberInfo //yay global variables!
const DEC_KEY = 811589153

type numberInfo struct {
	value int
	iPos int
}


func process(file string) {
lines := strings.Split(file, "\n");
numbers = make([]numberInfo, len(lines))
var count int
for (count < len(lines)) {
	numbers[count].value, _ = strconv.Atoi(lines[count])
	numbers[count].value *= DEC_KEY
	numbers[count].iPos = count
	count++
}

mixCount := 0
for (mixCount < 10) {
count = 0
for (count < len(numbers)) {
	moveNumber(findNumberI(count))
	//fmt.Println(numbers)
	count++
}
mixCount++
}
printFinalAnswer()
}

func moveNumber(startingIndex int) {
	var count int
    temp := numbers[startingIndex]
	destination := (startingIndex + (numbers[startingIndex].value % (len(numbers)-1))) //Because the number is removed from the list while in limbo
	//the below two statements handle wraparounds
	if (destination <= 0) {
		destination-- 
		//if it needs to wrap around the left side, add an additional leftward move, because the move from 0 to end doesn't displace anything
	}
	if (destination >= len(numbers)) {
		destination++ //same wrapping thing as above
	}
	destination %= len(numbers)
	for (destination < 0) {
		destination += len(numbers) //positive modulo, essentially
	}
	if (destination < startingIndex) {//moving left
		count = startingIndex
		for (count > destination) {
		numbers[count] = numbers[count - 1]	//shift everything to the right, starting with the moved number being replaced
		count--
		}
	}
	if (destination > startingIndex) {//moving right
		count = startingIndex
		for (count < destination) {
		numbers[count] = numbers[count + 1]	//shift everything to the right, starting with the moved number being replaced
		count++
		}
	}
	numbers[destination] = temp
	//if not moving, neither condition hits. no special case needed
	
	
}

func findNumberI(targetIPos int) int{
	count := 0
	for (count < len(numbers)) {
		if (numbers[count].iPos == targetIPos) {
			return count
		}
	count++
	}
	fmt.Println("something is terribly wrong")
	return -1
}


func printFinalAnswer() {
	count := 0
	score := 0
	for (count < len(numbers)) {
		if (numbers[count].value == 0) {
			score += numbers[(count + 1000) % len(numbers)].value
			score += numbers[(count + 2000) % len(numbers)].value
			score += numbers[(count + 3000) % len(numbers)].value
			break
		}
		count++
	}
	fmt.Println(score)
}

