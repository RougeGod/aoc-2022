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
	farLeft := 20000000
	farRight := 0
	for (count < 27) {
		if (findCorner(squares[count], "L").x < farLeft) {
			farLeft = findCorner(squares[count], "L").x
		}
		if (findCorner(squares[count], "R").x > farRight) {
			farRight = findCorner(squares[count], "R").x
		}
	count++
	}
	
	score := 0
	count = farLeft
	sensor.y = 2000000
	for (count < farRight) {
	sensor.x = count
		if (isInAnySquare(sensor)) {
			score++
		}
		count++
		if (count % 10000 == 0) {
			fmt.Println(count)
		}
	}
	fmt.Println(score) 
}

func md(one coords, two coords) int{
	return (abs(two.x - one.x)) + (abs(two.y - one.y))
}

func isInSquare(q square, point coords) bool{
	for _, v := range beakons { 
	//exclude locations where beakons are known to be
	//this is stupid slow and only matters in one location with a beakon
		if (v == point) {
			return false
		}
	}
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