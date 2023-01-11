package main

import "fmt"
import "os"
import "strings"
import "strconv"

func main() {
filecontents, err := os.ReadFile("./" + os.Args[1])
if (err != nil) {
	fmt.Println("oh no it broke")
	return
}

file := string(filecontents)
makemaps(file)
}

var beakons map[coords]coords
var squares []square

type coords struct {
 x int
 y int
}

type square struct {
	centre coords
	radius int
}

func makemaps(file string) {
	beakons = make(map[coords]coords)
	squares = make([]square, 27)
	lines := strings.Split(file, "\n")
	count := 0
	var sensor coords
	var beakon coords
	var ln []string
	for (count < len(lines)) {
		lines[count] = strings.TrimPrefix(lines[count], "Sensor at x=")
		ln = strings.Split(lines[count], " ")
		ln[0] = strings.TrimSuffix(ln[0], ",")
		ln[1] = strings.TrimSuffix((strings.TrimPrefix(ln[1], "y=")), ":")
		ln[6] = strings.TrimSuffix((strings.TrimPrefix(ln[6], "x=")), ",")
		ln[7] = strings.TrimPrefix(ln[7], "y=")
		sensor.x, _ = strconv.Atoi(ln[0])
		sensor.y, _ = strconv.Atoi(ln[1])
		beakon.x, _ = strconv.Atoi(ln[6])
		beakon.y, _ = strconv.Atoi(ln[7])
		beakons[sensor] = beakon
		count++
	}
	count = 0
	for k, v := range beakons {
		squares[count].centre = k
		squares[count].radius = md(k, v)
		count++
	}
	count = 0
	for (count < len(squares)) {
		searchAroundSquare(squares[count])
		//fmt.Println(count)
		count++
	}

}
	//if only one unfound spot, all of its neighbours are inside squares, and as such it is one spot
	//outside of four seperate squares
	
/*To search around a square: Go to its top, go up one square, go down-right radius + 1 spots,
go down-left r+1 spots, go up-left r+1 spots, go up-right r spots (or r+1, double search here or there is fine) */
func searchAroundSquare(q square) {
	var searcher coords
	inner := 0
		searcher = findCorner(q, "U")
		searcher.y--
		inner = 0
		for (inner <= q.radius) {
			searcher.x++
			searcher.y++
			if ((!isInAnySquare(searcher))&&(isValid(searcher))) {
				fmt.Println(searcher.x*4000000 + searcher.y)
				os.Exit(0)
			}
			inner++
		}
		inner = 0
		for (inner <= q.radius) {
			searcher.x--
			searcher.y++
			if ((!isInAnySquare(searcher))&&(isValid(searcher))) {
				fmt.Println(searcher.x*4000000 + searcher.y)
				os.Exit(0)
			}
			inner++
		}
		inner = 0
		for (inner <= q.radius) {
			searcher.x--
			searcher.y--
			if ((!isInAnySquare(searcher))&&(isValid(searcher))) {
				fmt.Println(searcher.x*4000000 + searcher.y)
				os.Exit(0)
			}
			inner++
		}
		inner = 0
		for (inner <= q.radius) {
			searcher.x++
			searcher.y--
			if ((!isInAnySquare(searcher))&&isValid(searcher)) {
				fmt.Println(searcher.x*4000000 + searcher.y)
				os.Exit(0)
			}
			inner++
		}
}

func isValid(cand coords) bool{
	if (cand.x > 4000000)||(cand.x < 0) {
		return false
	}
	if (cand.y > 4000000)||(cand.y < 0) {
		return false
	}
	return true
}
func md(one coords, two coords) int{
	return (abs(two.x - one.x)) + (abs(two.y - one.y))
}

func isInSquare(q square, point coords) bool{ //beakons no longer need to be excluded
	return (md(q.centre, point) <= q.radius)
}

func isInAnySquare(point coords) bool{
	count := 0
	for (count < len(squares)) {
		if (isInSquare(squares[count], point)) {
			return true
		}
	count++
	}
	return false
}

func findCorner(q square, dir string) coords{
	var output coords
	output = q.centre
	switch (dir) {
		case "L": output.x -= q.radius
		break
		case "R": output.x += q.radius
		break
		case "U": output.y -= q.radius
		break
		case "D": output.y += q.radius
		break
		default: fmt.Println("something has gone terribly wrong")
	}
	return output
}

func abs(s int) int{
	if (s >= 0) {
		return s 
	}
	return s*-1
}