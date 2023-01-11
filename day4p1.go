package main

import "fmt"
import "io/ioutil"
import "os"
import "strings"
import "regexp"
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
expression := regexp.MustCompile(`[^\d]`)

lines := strings.Split(file, "\n");
count :=0
overlaps :=0
var c4 int //count 4, for the ranges
var ranges []string
var numberRange [4]int
var swap int

for (count < len(lines)) {
c4 = 0
ranges = expression.Split(lines[count], -1)
for (c4 < 4) {
numberRange[c4], _ = strconv.Atoi(ranges[c4]) //underscore because we're sending the error value out the window
c4++
}

if ((numberRange[2] < numberRange[0])||((numberRange[0] == numberRange[2])&&(numberRange[1] < numberRange[3]))) { 
//make sure that the lower of the ranges is the first one listed
swap = numberRange[2]
numberRange[2] = numberRange[0]
numberRange[0] = swap
swap = numberRange[3]
numberRange[3] = numberRange[1]
numberRange[1] = swap
}

fmt.Println(numberRange)
if (numberRange[1] >= numberRange[3]) {
fmt.Println("overlap")
overlaps++
}

fmt.Println(overlaps)
count++
}

}