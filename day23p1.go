package main

import "fmt"
import "os"
import "strings"
import "sort"

//YOOOOOOOOOOOOOO FIRST TRY BABYYYYYY!
//Print statements used for debugging haave been left in here, commented out.

func main() {
filecontents, err := os.ReadFile("./" + os.Args[1])
if (err != nil) {
	fmt.Println("oh no it broke")
	return
}
file := string(filecontents)
process(file)
}

type elf struct {
	currentPos  int
	proposedPos int
}

var grid [10000]bool //a rather hacky way to make coordinates. it's a 100x100 grid but i don't want to deal with 2-D arrays
					//with only 10 rounds of moves, no way that any elf leaves the 100x100 grid (from 70x70). Top corner of the file
					//is at (11, 11) or array position 1111, starting from (0,0)
var elves [2458]elf //number of elves, counted in notepad++. it could be calculated in code but this is faster


func process(file string) { //initializes grid and elf positions
lines := strings.Split(file, "\n")
var line []string
count := 0
inner := 0
elfcount := 0 //to assign to the elves array
for (count < len(lines)) {
	line = strings.Split(lines[count], "")
	for (inner < len(line)) {
		if (line[inner] == "#") {
			grid[(count*100) + inner + 1111] = true //1111 is an adjustment to kick the grid out of the top left corner
			elves[elfcount].currentPos = count*100 + inner + 1111
			//fmt.Println(elves[elfcount])
			elfcount++
		}
		inner++
	}
	inner = 0
	count++
}
proposeMoves(0)
}


func proposeMoves(moveNum int) {
	var elfCheck [4]bool
	var count int
	for elf, _ := range elves {
		//fmt.Println(elves[elf])
		elfCheck = checkAround(elves[elf].currentPos)
		count = moveNum
		elves[elf].proposedPos = -1 //flag value. if this remains afterwards then that elf stays still
		for (count < 4 + moveNum) { //run count 4 times
			if (elfCheck[count % 4]) {
				switch (count % 4) {
					case 0: elves[elf].proposedPos = elves[elf].currentPos - 100;
					case 1: elves[elf].proposedPos = elves[elf].currentPos + 100;  
					case 2: elves[elf].proposedPos = elves[elf].currentPos - 1;
					case 3: elves[elf].proposedPos = elves[elf].currentPos + 1;  
					default: fmt.Println("all is lost")
				}
			if (elves[elf].currentPos != -1) { //if proposal has been made, don't check further
				break
			}
			}
			count++
		}
    //fmt.Println(elves[elf].proposedPos)	
	}
	moveElves(moveNum)
}


//updates the grid array and each elf's current position. there are no ways that a new elf moves into the 
//spot of an elf that just moved, so it's safe to clear that space. Elf does not move if its proposed
//position is in the conflicts array. 
func moveElves(moveNum int) {
	conflicts := checkConflicts()
	//fmt.Println(conflicts)
	for count, _ := range elves {
		if (elves[count].proposedPos > conflicts[len(conflicts) - 1]) { //special case to not break go's searchInt function
			grid[elves[count].currentPos] = false
			grid[elves[count].proposedPos] = true
			elves[count].currentPos = elves[count].proposedPos
			//fmt.Println("no conflict, moved to ", elves[count].currentPos)
		} else if (!(conflicts[sort.SearchInts(conflicts, elves[count].proposedPos)] == elves[count].proposedPos)) {
			grid[elves[count].currentPos] = false
			grid[elves[count].proposedPos] = true
			elves[count].currentPos = elves[count].proposedPos
			//fmt.Println("no conflict, moved to ", elves[count].currentPos)
		} 
	}
	if (moveNum < 9) { //run for 10 moves (the last move number is 9 since it starts at zero)
	proposeMoves(moveNum + 1)
	} else {
		printGrid()
		fmt.Println(findFinalAnswer())
	}
}


//runs through every elf's proposed position and returns a slice of ints with all of the positions 
//that have a proposal conflict and thus can't be moved to in the move portion of the turn... 
//return by reference or by value doesn't matter
func checkConflicts() []int{
	positions := make([]int, 2458) 			//holds all elves' proposed positions
	conflicts := make([]int, 1, 1230) //slice capacity is the number of elves divided by 2 (11 in the example) + 1 for the special "no move" position of -1
									//there cannot be more than (elves/2) conflicts because each conflict needs two elves
	conflicts[0] = -1
	for count, elf := range elves {
		positions[count] = elf.proposedPos
	}
	//checks through all positions, sorting them for easier check of two being the same. Note that this
	//also makes the final conflicts array a sorted array. 
	sort.Ints(positions)
	count := 1
	for (count < len(positions)) {
		if ((positions[count - 1] == positions[count])&&(positions[count] != -1)) {
			conflicts = append(conflicts, positions[count])
		}
		count++
	}
	return conflicts
	}

//output[0] is "free to the north", [1] is south, [2], is west, and [3] is east. if all are true, returns all false because that elf won't move
func checkAround(location int) [4]bool{
	var output [4]bool
	if ((location % 100 == 0)||(location % 100 == 99)||(location / 100 == 0)||(location / 100 == 99)) { //if on grid edge, something is terribly wrong
		//fmt.Println("something is terribly wrong")
		//output = {false, false, false, false}
		return output //nothing has been set to true yet so it's all false 
	}
	if (!((grid[location-101])||(grid[location-100])||(grid[location-99]))) {
		output[0] = true //indicate that north direction is free
	}
	if (!((grid[location+1])||(grid[location+101])||(grid[location-99]))) {
		output[3] = true //indicate that east direction is free
	}
	if (!((grid[location+101])||(grid[location+100])||(grid[location+99]))) {
		output[1] = true //indicate that south direction is free
	}
	if (!((grid[location-101])||(grid[location-1])||(grid[location+99]))) {
		output[2] = true //indicate that west direction is free
	}
	if (output[0]&&output[1]&&output[2]&&output[3]) { //if all directions are valid (ie elf has no neighbours)
		var blank [4]bool //why can't i set the whole array?
		return blank
	}
	//fmt.Println(output)
	return output
}


func printGrid() { //very useful for debugging. unnecessary in getting the final answer
	count := 0
	for (count < len(grid)) {
		if (grid[count]) {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		if (count % 100 == 99) {
			fmt.Print("\n")
		}
	count++
	}
}

func anythingInRow(row int) bool{
	count := row*100
	for (count < ((row+1) * 100)) {
		if (grid[count]) {
			return true
		}
		count++
	}
	return false
}

func anythingInCol(col int) bool{
	count := col
	for (count < len(grid)) {
		if (grid[count]) {
			return true
		}
		count += 100
	}
	return false
}

func findFinalAnswer() int{
	var left, right, top, bottom int
	count := 0
	for (!anythingInCol(count)) {
		count++
	}
	left = count
	count = 99
	for (!anythingInCol(count)) {
		count--
	}
	right = count
	count = 0 
	for (!anythingInRow(count)) {
		count++
	}
	top = count
	count = 99
	for (!anythingInRow(count)) {
		count--
	}
	bottom = count
	return ((bottom-top+1)*(right-left+1)) - len(elves)//added 1 because bounds are inclusive for answer and exclusive in calculation
}