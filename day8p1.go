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
	blockedLeft bool
	blockedRight bool
	blockedUp bool
	blockedDown bool
}

func process(file string) {
lines := strings.Split(file, "\n");
var currentTree string
var treeHeight int //so that multi=value atoi works correctly
var count int
var inner int
forest := make([]tree, 9801)
//there are 99 rows and 99 cols so 9801 total trees
for (count < 99) {
	inner = 0
	for (inner < 99) {
		currentTree = strings.Split(lines[count], "")[inner]
		treeHeight, _ = strconv.Atoi(currentTree)
		forest[count*99 + inner].height = treeHeight
		//all booleans are automatically initialized as false
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
		if (forest[inner].height >= forest[count].height) {
			forest[count].blockedLeft = true
		}
	}
	count++
}

count = 0


for (count < len(forest)) {
		if (forest[count].blockedLeft) {
		inner = count
		for (inner % 99 != 98) {
			inner++
			if (forest[inner].height >= forest[count].height) {
				forest[count].blockedRight = true
				}	
			}
		}
	count++
}
count = 0

//blocked top
for (count < len(forest)) {
	inner = count - 99
	for (inner >= 0) {
		if (forest[inner].height >= forest[count].height) {
			forest[count].blockedUp = true
		}
		inner = inner - 99
	}
	inner = count + 99
	for (inner < len(forest)) {
		if (forest[inner].height >= forest[count].height) {
			forest[count].blockedDown = true
		}
		inner = inner + 99		
	}
	count++
}

count = 0
inner = 0 //serving as the anti-score counter
for (count < len(forest)) {
	if (forest[count].blockedLeft) {
		fmt.Println(count, " Blocked on the left")
	}
	if (forest[count].blockedRight) {
		fmt.Println(count, " Blocked on the right")
	}
	if (forest[count].blockedUp) {
		fmt.Println(count, "Blocked on the top")
	}
	if (forest[count].blockedDown) {
		fmt.Println("Blocked on the bottom")
	}
	
	if ((forest[count].blockedUp && forest[count].blockedDown)&&((forest[count].blockedLeft)&&(forest[count].blockedRight))) {
		fmt.Println(count, " is fully blocked")
	}
	
	
	if (!((forest[count].blockedUp && forest[count].blockedDown)&&((forest[count].blockedLeft)&&(forest[count].blockedRight)))) {
		inner++
		fmt.Println("Tree ", count, " is visible")
	}
	count++
	
}

fmt.Println(inner)
}