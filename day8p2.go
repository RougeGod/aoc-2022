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

type tree struct {
	height int
	vdLeft int
	vdRight int
	vdUp int
	vdDown int
}

func process(file string) {
lines := strings.Split(file, "\n");
var currentTree string
var treeHeight int //so that multi=value atoi works correctly
var count int
var inner int
var finalScore int
forest := make([]tree, 9801)
//there are 99 rows and 99 cols so 9801 total trees
for (count < 99) {
	inner = 0
	for (inner < 99) {
		currentTree = strings.Split(lines[count], "")[inner]
		treeHeight, _ = strconv.Atoi(currentTree)
		forest[count*99 + inner].height = treeHeight
		//all ints are automatically initialized as zero
		inner++
	}
	count++
}

count = 0
//for each tree, check if it's blocked on the left
for (count < len(forest)) {
	inner = count
	for (inner % 99 != 0) {
		inner-- //no need to compare to itself but a need to compare it to the edge tree
		if ((forest[inner].height >= forest[count].height)||(inner % 99 == 0)) {
			forest[count].vdLeft = (count - inner)
			break
		}
	}
	count++
}

count = 0


for (count < len(forest)) {
		inner = count
		for (inner % 99 != 98) {
			inner++
			if ((forest[inner].height >= forest[count].height)||(inner % 99 == 98)) {
				forest[count].vdRight = (inner - count)
				break
				}	
			}
	count++
}
count = 0

//blocked top
for (count < len(forest)) {
	inner = count - 99
	for (inner >= 0) {
		if ((forest[inner].height >= forest[count].height)||(inner < 99)) {
			forest[count].vdUp = ((count - inner)/99)
			break
		}
		inner = inner - 99
	}
	
	inner = count + 99
	for (inner < len(forest)) {
		if ((forest[inner].height >= forest[count].height)||(inner > len(forest) - 99)) {
			forest[count].vdDown = ((inner - count) / 99)
			break
		}
		inner = inner + 99		
	}
	count++
}

count = 0
finalScore = 0 //serving as the score counter


for (count < len(forest)) {
inner = (forest[count].vdUp * forest[count].vdDown * forest[count].vdLeft * forest[count].vdRight)
fmt.Println("Count: ", count, " Down: ", forest[count].vdDown, " Up: ", forest[count].vdUp, " Left: ", forest[count].vdLeft, " Right: ", forest[count].vdRight, " Inner: ", inner)
if  (inner > finalScore) {
	finalScore = inner
}

count++
	
}

fmt.Println(finalScore)
}