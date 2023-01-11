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
var thirdHalf string
var totalPoints = 0
var count = 0
var fhc int //first half count, used to iterate over the first half of the string
var shc int
var thc int //i promise i'm not high on mistletoe
var found bool

lines := strings.Split(file, "\n");

for (count < len(lines)) {
firstHalf = lines[count]
secondHalf = lines[count + 1]
thirdHalf = lines[count + 2]
fhc = 0
found = false

for (fhc < len(firstHalf)) {
shc = 0
for (shc < len(secondHalf)) {
	if (found) {
	break
	} 
	if (secondHalf[shc] == firstHalf[fhc]) {
		thc = 0
		for (thc < len(thirdHalf)) {
			if ((thirdHalf[thc] == secondHalf[shc])&&(thirdHalf[thc] ==firstHalf[fhc])) {
				found = true
				totalPoints += addScore(thirdHalf[thc])
				break
				}
			thc++
			}
		}
	shc++
	}
fhc++
}



count += 3
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


