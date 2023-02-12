package main

import "fmt"
import "os"
import "strings"
import "strconv"

/*THOUGHTS: Make a structure out of each number, with one for the value, one for the initial position, and one for its current position. 
Maybe they can be in a map with initial position as the key and (value, currentposition) as the value

Acutally never mind. Current position is found as the array index and is kinda redundant, so make every number a map of initial position to value, and have an array of them. 
There are repeated numbers, some repeated multiple times. 
To move a number, find its destination (current position + value % len(list)). If it's moving forwards, set the currentposition of every number with a currentposition 

This is undoubtedly faster with a linkedlist, and perhaps I will rewrite it with one to learn how



18512 is TOO HIGH
*/
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
	numbers[count].iPos = count
	count++
}

count = 0
for (count < len(numbers)) {
	moveNumber(findNumberI(count))
	//fmt.Println(numbers)
	count++
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

