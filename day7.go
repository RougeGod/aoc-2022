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
process(file)
fmt.Println(getFinalAnswer())//for both parts! the returned int is for part 1 nd the returned string is for part 2
}

//make a map of all directories (key) to their sizes (value) 
var sizes map[string]int

func process(file string) {
	sizes = make(map[string]int)
	fullPath := "root" //because having a slash there breaks my parsing i think
	commands := strings.Split(file, "$ ")	
	commands = commands[1:]
	//the 0th element is blank due to strange behaviour of the Split() function which leaves element 0 blank when 
	//the string starts with the seperator, as it does here
	var output []string
	for _, command := range commands { //runs for every command
		output = strings.Fields(command)//Fields returns a slice with all "words", split at whitespace characters
		if (output[0] == "cd") {
			if (output[1] == "..") {
				fullPath = fullPath[0:strings.LastIndex(fullPath,"/")] //removes the last / and anything following from fullPath
			} else {
				fullPath += "/" + output[1]
			}
		} else if (output[0] == "ls") {
			sumSizes(output, fullPath)
		} 
	} 
}

func sumSizes(sizeList []string, path string) {
	dirSize := 0
	for _, word := range sizeList {
		fileSize, _ := strconv.Atoi(word)
		dirSize += fileSize
	}
	addToPath(dirSize, path)
}

func addToPath(dirSize int, path string) {
	splitPath := strings.Split(path, "/")
	var summedPath string
	for count, _ := range splitPath {
		if (count != 0) {
			summedPath += "/"
		}
		summedPath += splitPath[count]
		sizes[summedPath] += dirSize
	}
}

func getFinalAnswer() (int, int) {
	score := 0	
	MIN_DELETED_SIZE := sizes["root"] - 40000000
	contender := "root"
	for dir, dirSize := range sizes {
		if (dirSize < 100000) {
			score += dirSize //to solve part 1's problem
		}
		if ((dirSize > MIN_DELETED_SIZE)&&(sizes[dir] < sizes[contender])) {
			contender = dir //since we already have the map set up, why not solve part 2 while we're here?
		}
	}
	return score, sizes[contender]
}
