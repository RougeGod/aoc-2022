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

type coords struct {
	x int
	y int
}

//global variables?
var head coords
var tail coords
var visited = make([]coords, 20000)
var vcount int

func process(file string) {
var count int
var inner int
var command [2]string
var moveLen int

//keeps track of the end of the visited array. also the total score

lines := strings.Split(file, "\n");
visited[0].x = 0
visited[0].y = 0 
vcount = 1

for (count < len(lines)) {
command[0] = strings.Split(lines[count], " ")[0]
command[1] = strings.Split(lines[count], " ")[1]
moveLen, _ = strconv.Atoi(command[1])

inner = 0

if (command[0] == "U") {
	for (inner < moveLen) {
		head.y++
		moveTail()
		inner++
	}
}
if (command[0] == "D") {
	for (inner < moveLen) {
		head.y--
		moveTail()
		inner++
	}
}
if (command[0] == "R") {
	for (inner < moveLen) {
		head.x++
		moveTail()
		inner++
	}
}
if (command[0] == "L") {
	for (inner < moveLen) {
		head.x--
		moveTail()
		inner++
	}
}


count++	
fmt.Println(count, vcount)
}
}

func moveTail() {
	if ((abs(head.x - tail.x) <= 1)&&(abs(head.y - tail.y) <= 1)) {
		return
	}
	//move both x and y coords if not touching (if one is equal it won't move)
	
	if (head.x > tail.x) {
		tail.x++
	}
	if (head.x < tail.x) {
		tail.x--
	}
	if (head.y > tail.y) {
		tail.y++
	}
	if (head.y < tail.y) {
		tail.y--
	}
	var count int //initialized to 0
	var found bool //initialized to false
	count = 0
	found = false
	for (count < vcount) {
		if ((tail.x == visited[count].x)&&(tail.y == visited[count].y)) {
		found = true
		break
		}
	count++
	}
	if (!found) {
		visited[vcount] = tail
		vcount++
	}
}

//don't want to deal with conversions to use the built-in math.Abs (which expects floats)
func abs(x int) int{
	if (x >= 0) {
		return x
	} else {
		return 0-x
	}
}