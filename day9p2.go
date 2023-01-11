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
var knots [10]coords //element 0 is the head
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
		knots[0].y++
		moveTail()
		inner++
	}
}
if (command[0] == "D") {
	for (inner < moveLen) {
		knots[0].y--
		moveTail()
		inner++
	}
}
if (command[0] == "R") {
	for (inner < moveLen) {
		knots[0].x++
		moveTail()
		inner++
	}
}
if (command[0] == "L") {
	for (inner < moveLen) {
		knots[0].x--
		moveTail()
		inner++
	}
}


count++	
fmt.Println(count, vcount)
}
}

func moveTail() {
	leader := 0
	trailer := 1
	for (trailer < 10) {
	if ((abs(knots[leader].x - knots[trailer].x) <= 1)&&(abs(knots[leader].y - knots[trailer].y) <= 1)) {
		return
	}
	//move both x and y coords if not touching (if one is equal it won't move)
	
	if (knots[leader].x > knots[trailer].x) {
		knots[trailer].x++
	}
	if (knots[leader].x < knots[trailer].x) {
		knots[trailer].x--
	}
	if (knots[leader].y > knots[trailer].y) {
		knots[trailer].y++
	}
	if (knots[leader].y < knots[trailer].y) {
		knots[trailer].y--
	}
	if (trailer == 9) {
	count := 0
	found := false
	for (count < vcount) {
		if ((knots[9].x == visited[count].x)&&(knots[9].y == visited[count].y)) {
		found = true
		break
		}
	count++
	}
	if (!found) {
		visited[vcount] = knots[9]
		vcount++
	}
	}
	leader++
	trailer++
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