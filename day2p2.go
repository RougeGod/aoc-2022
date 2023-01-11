package main

import "fmt"
import "os"
import "strings"


func main() {
filecontents, err := os.ReadFile("./" + os.Args[1])
if (err != nil) {
	fmt.Println("oh no it broke")
	return
}

file := string(filecontents)

process(file)


}

func process(file string) {
lines := strings.Split(file, "\n");
var totalPoints = 0
var count = 0
for (count < len(lines)) {
if ((len(lines[count])) == 0) {
	fmt.Println(count)
	fmt.Print("It's all over: ")
	fmt.Println(totalPoints)
	os.Exit(0)
	} 

//A lookup table because I was very tired and this was faster than figuring out the better way
//Also ran into syntax issues with nested switch statements
if(lines[count][0] == 65) {
if (lines[count][2] == 88) {
totalPoints += 3
}
if (lines[count][2] == 89) {
totalPoints += 4
}
if (lines[count][2] == 90) {
totalPoints += 8
}
}
if(lines[count][0] == 66) {
if (lines[count][2] == 88) {
totalPoints += 1
}
if (lines[count][2] == 89) {
totalPoints += 5
}
if (lines[count][2] == 90) {
totalPoints += 9
}
}
if(lines[count][0] == 67) {
if (lines[count][2] == 88) {
totalPoints += 2
}
if (lines[count][2] == 89) {
totalPoints += 6
}
if (lines[count][2] == 90) {
totalPoints += 7
}
}



count++
}

}




