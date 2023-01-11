package main

import "fmt"
import "io/ioutil"
import "os"
import "strings"
import "strconv"


func main() {
filecontents, err := ioutil.ReadFile("./" + os.Args[1])
if (err != nil) {
	fmt.Println("oh no it broke")
	return
}

file := string(filecontents)

process(file)
}

var crt [240]bool //true if pixel is lit, false otherwise

func process(file string) {
lines := strings.Split(file, "\n");
count := 0 
cycle := 0
score := 0
x := 1
var command string
var regChange int

for (count < len(lines)) {
	command = strings.Split(lines[count], " ")[0]
		if (abs(x - (cycle % 40)) <= 1) {
			crt[cycle] = true
		}
		cycle++
	if (command == "addx") {
		regChange, _ = strconv.Atoi(strings.Split(lines[count], " ")[1])
		if (abs(x - (cycle % 40)) <= 1) {
			crt[cycle] = true
		}
		cycle++ //addx adds another cycle, then the register gets changed
		x += regChange
	}	
	count++
}
fmt.Println(score)
drawScreen()
}

func drawScreen() {
	row := 0
	col := 0
	for (row < 6) {
		col = 0
		for (col < 40) {
			if (crt[row*40 + col]) {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		col++
		}
		fmt.Println("")
		row++
	}
}

func abs(x int) int{
	if (x >= 0) {
		return x
	} else {
		return 0-x
	}
}