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
		cycle++ //all commands add at least one cycle
		if ((cycle % 40 == 20)&&(cycle < 221)) {
			score += (cycle * x)
		}
	if (command == "addx") {
		regChange, _ = strconv.Atoi(strings.Split(lines[count], " ")[1])
		cycle++ //addx adds another cycle, then the register gets changed
		if ((cycle % 40 == 20)&&(cycle < 221)) {
			score += (cycle * x) 
			
		}
		x += regChange
	}	
	count++
}
fmt.Println(score)
}
