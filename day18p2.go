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

var squareMap map[[4]coords]int

func main() {
	filecontents, err := os.ReadFile("./" + os.Args[1])
	if err != nil {
		fmt.Println("oh no it broke")
		return
	}
	lines := strings.Split(string(filecontents), "\n")
	squareMap = make(map[[4]coords]int)
	var max int
	for _, line := range lines {
		for _, square := range getSquares(toCoords(line)) {
			if square[0].x > max {
				max = square[3].x
			}
			if square[0].y > max {
				max = square[3].y
			}
			if square[0].z > max {
				max = square[3].z
			}
			squareMap[square]++
		}
	}

	var scanX map[coords]bool = pass(max, 'X')
	scanY := pass(max, 'Y')
	scanZ := pass(max, 'Z')
	for k := range scanX {
		if scanY[k] && scanZ[k] {
			for _, square := range getSquares([3]int{k.x, k.y, k.z}) {
				squareMap[square]++
			}
		}
	}

	var score int
	for _, v := range squareMap {
		if v == 1 {
			score++
		}
	}
	fmt.Println(score)
}

func pass(max int, scanvar rune) map[coords]bool {
	var minFound bool
	var maxSoFar int
	anomalies := make(map[coords]bool)
	for i := 0; i < max; i++ {
		for j := 0; j < max; j++ {
			minFound = false
			maxSoFar = 0
			for k := 0; k < max; k++ {
				switch scanvar {
				case 'Z':
					if squareMap[[4]coords{{i, j, k}, {i, j + 1, k}, {i + 1, j, k}, {i + 1, j + 1, k}}] == 1 {
						if !minFound {
							minFound = true
						} else if maxSoFar == 0 {
							maxSoFar = k
						} else {
							for m := maxSoFar; m < k; m++ {
								anomalies[coords{i, j, m}] = true
							}
							maxSoFar = 0
						}
					}
				case 'Y':
					if squareMap[[4]coords{{i, k, j}, {i, k, j + 1}, {i + 1, k, j}, {i + 1, k, j + 1}}] == 1 {
						if !minFound {
							minFound = true
						} else if maxSoFar == 0 {
							maxSoFar = k
						} else {
							for m := maxSoFar; m < k; m++ {
								anomalies[coords{i, m, j}] = true
							}
							maxSoFar = 0
						}
					}
				case 'X':
					if squareMap[[4]coords{{k, i, j}, {k, i, j + 1}, {k, i + 1, j}, {k, i + 1, j + 1}}] == 1 {
						if !minFound {
							minFound = true
						} else if maxSoFar == 0 {
							maxSoFar = k
						} else {
							for m := maxSoFar; m < k; m++ {
								anomalies[coords{m, i, j}] = true
							}
							maxSoFar = 0
						}
					}
				}
			}
		}
	}
	return anomalies
}

func getSquares(pos [3]int) [][4]coords {
	var output [][4]coords
	output = make([][4]coords, 6)
	output[0] = [4]coords{{pos[0], pos[1], pos[2]}, {pos[0], pos[1] + 1, pos[2]}, {pos[0] + 1, pos[1], pos[2]}, {pos[0] + 1, pos[1] + 1, pos[2]}}                 //Z=0
	output[1] = [4]coords{{pos[0], pos[1], pos[2] + 1}, {pos[0], pos[1] + 1, pos[2] + 1}, {pos[0] + 1, pos[1], pos[2] + 1}, {pos[0] + 1, pos[1] + 1, pos[2] + 1}} //Z = 1
	output[2] = [4]coords{{pos[0], pos[1], pos[2]}, {pos[0], pos[1], pos[2] + 1}, {pos[0], pos[1] + 1, pos[2]}, {pos[0], pos[1] + 1, pos[2] + 1}}                 //X = 0
	output[3] = [4]coords{{pos[0] + 1, pos[1], pos[2]}, {pos[0] + 1, pos[1], pos[2] + 1}, {pos[0] + 1, pos[1] + 1, pos[2]}, {pos[0] + 1, pos[1] + 1, pos[2] + 1}} //X = 1
	output[4] = [4]coords{{pos[0], pos[1], pos[2]}, {pos[0], pos[1], pos[2] + 1}, {pos[0] + 1, pos[1], pos[2]}, {pos[0] + 1, pos[1], pos[2] + 1}}                 //Y = 0
	output[5] = [4]coords{{pos[0], pos[1] + 1, pos[2]}, {pos[0], pos[1] + 1, pos[2] + 1}, {pos[0] + 1, pos[1] + 1, pos[2]}, {pos[0] + 1, pos[1] + 1, pos[2] + 1}} //Y = 1
	return output
}

func toCoords(line string) [3]int {
	var pos [3]int
	for index, number := range strings.Split(line, ",") {
		pos[index], _ = strconv.Atoi(number)
	}
	return pos
}
