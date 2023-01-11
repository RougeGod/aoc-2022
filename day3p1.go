package main

import "fmt"
import "io/ioutil"
import "os"
import "strings"


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
var firstHalf string
var secondHalf string
var totalPoints = 0
var count = 0
var fhc int //first half count, used to iterate over the first half of the string
var shc int
var found bool

lines := strings.Split(file, "\n");

for (count < len(lines)) {
firstHalf = lines[count][0:(len(lines[count])/2)]
secondHalf = lines[count][(len(lines[count])/2):len(lines[count])]
found = false

fhc = 0
for (fhc < len(firstHalf)) {
shc = 0
for (shc < len(secondHalf)) {
	if (found) {
	break
	}
	if (firstHalf[fhc] == secondHalf[shc]) {
		totalPoints += addScore(firstHalf[fhc])
		found = true
		}
		shc++
}
fhc++
}

count++
}
fmt.Println(totalPoints)
}

func addScore(item byte) int{
	var output int
	if (item < 123 && item > 96) {
	output = int(item - 96)
	}
	if (item < 91 && item > 64) {
	output = int(item - 38)
	}
return output

}


