package main

import "fmt"
import "os"
import "strings"
import "strconv"

type coords struct {
	x int
	y int
	z int
}

func main() {
	filecontents, err := os.ReadFile("./" + os.Args[1])
	if err != nil {
		fmt.Println("oh no it broke")
		return
	}
	var score int
	lines := strings.Split(string(filecontents), "\n")
	squareMap := make(map[[4]coords]int)
	for _, line := range lines {
		for _, square := range getSquares(line) {
			squareMap[square] += 1
			if squareMap[square] == 1 {
				score++
			} else if squareMap[square] == 2 {
				score--
			}
		}
	}
	fmt.Println(score)
}

func getSquares(line string) [][4]coords {
	var pos [3]int
	var output [][4]coords
	output = make([][4]coords, 6)
	for index, number := range strings.Split(line, ",") {
		pos[index], _ = strconv.Atoi(number)
	}
	output[0] = [4]coords{{pos[0], pos[1], pos[2]}, {pos[0], pos[1] + 1, pos[2]}, {pos[0] + 1, pos[1], pos[2]}, {pos[0] + 1, pos[1] + 1, pos[2]}} //Z=0
	output[1] = [4]coords{{pos[0], pos[1], pos[2] + 1}, {pos[0], pos[1] + 1, pos[2] + 1}, {pos[0] + 1, pos[1], pos[2] + 1}, {pos[0] + 1, pos[1] + 1, pos[2] + 1}} //Z = 1
	output[2] = [4]coords{{pos[0], pos[1], pos[2]}, {pos[0], pos[1], pos[2] + 1}, {pos[0], pos[1] + 1, pos[2]}, {pos[0], pos[1] + 1, pos[2] + 1}} //X = 0
	output[3] = [4]coords{{pos[0] + 1, pos[1], pos[2]}, {pos[0] + 1, pos[1], pos[2] + 1}, {pos[0] + 1, pos[1] + 1, pos[2]}, {pos[0] + 1, pos[1] + 1, pos[2] + 1}} //X = 1
	output[4] = [4]coords{{pos[0], pos[1], pos[2]}, {pos[0], pos[1], pos[2] + 1}, {pos[0] + 1, pos[1], pos[2]}, {pos[0] + 1, pos[1], pos[2] + 1}} //Y = 0
	output[5] = [4]coords{{pos[0], pos[1] + 1, pos[2]}, {pos[0], pos[1] + 1, pos[2] + 1}, {pos[0] + 1, pos[1] + 1, pos[2]}, {pos[0] + 1, pos[1] + 1, pos[2] + 1}} //Y = 1
	return output
}
