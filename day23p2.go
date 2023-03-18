package main

import "fmt"
import "os"
import "strings"
import "sort"

func main() {
filecontents, err := os.ReadFile("./" + os.Args[1])
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

type elf struct {
	currentPos  coords
	proposedPos coords
}

var elves []elf 


func process(file string) { //initializes grid and elf positions
lines := strings.Split(file, "\n")
var line []string
count := 0
inner := 0
var currElf elf
for (count < len(lines)) {
	line = strings.Split(lines[count], "")
	for (inner < len(line)) {
		if (line[inner] == "#") {
			currElf = elf{coords{inner, count}, coords{inner, count}}
			elves = append(elves, currElf)
		}
		inner++
	}
	inner = 0
	count++
}
/*TESTING CODE GOES HERE*/
count = 0
for (true) {
	proposeMoves(count)
	moveElves(count)
	fmt.Println(count)
	count++
}

}


func proposeMoves(moveNum int) {
	var elfCheck [4]bool
	var count int
	winFlag := true
	for elf, _ := range elves {
		elfCheck = checkAround(elves[elf].currentPos)
		count = moveNum
		elves[elf].proposedPos = elves[elf].currentPos //flag value. if this remains afterwards then that elf stays still
		for (count < 4 + moveNum) { //run count 4 times
			if (elfCheck[count % 4]) {
				switch (count % 4) {
					case 0: elves[elf].proposedPos.y = elves[elf].currentPos.y - 1
					case 1: elves[elf].proposedPos.y = elves[elf].currentPos.y + 1;  
					case 2: elves[elf].proposedPos.x = elves[elf].currentPos.x - 1;
					case 3: elves[elf].proposedPos.x = elves[elf].currentPos.x + 1;  
					default: fmt.Println("all is lost")
				}
			if (elves[elf].currentPos != elves[elf].proposedPos) { 
			//if proposal has been made, don't check further and leave the inner for-loop (the outer one continues)
				winFlag = false
				break
			}
			}
			count++
		}
	}
	if (winFlag) {
		fmt.Println("No more movement on move ", moveNum + 1)//because I'm starting at zero
		os.Exit(0)
	}
}


//updates the grid array and each elf's current position. there are no ways that a new elf moves into the 
//spot of an elf that just moved, so it's safe to clear that space. Elf does not move if its proposed
//position is in the conflicts array. 
func moveElves(moveNum int) {
	conflicts := checkProposalConflicts()
	var conflictLocation int
	for elf, elfPos := range elves {
		if (elfPos.currentPos != elfPos.proposedPos) { //if the elf wants to move
			conflictLocation = sort.Search(len(conflicts), func(conIndex int) bool { //find where in the conflict list that the movement would be
															if (conflicts[conIndex].x == elfPos.proposedPos.x) {
																return (conflicts[conIndex].y >= elfPos.proposedPos.y)			
															} else {
																return (conflicts[conIndex].x >= elfPos.proposedPos.x)
															}
														    })
		if (conflictLocation == len(conflicts)||(conflicts[conflictLocation] != elfPos.proposedPos)) { //no conflict, moving is fine
			elves[elf].currentPos = elves[elf].proposedPos
		} else { //it can't move due to a conflict. Reset proposedPos
			elves[elf].proposedPos = elves[elf].currentPos
		}
		}
	}
	proposeMoves(moveNum + 1)
}


//runs through every elf's proposed position and returns a slice of ints with all of the positions
//that have a proposal conflict and thus can't be moved to in the move portion of the turn... 
//return by reference or by value doesn't matter
func checkProposalConflicts() []coords{
	var moves []coords
	var conflicts []coords
	sortElves(true)//sort elves by proposals so that the conflict slice is fully sorted and duplicates can be eliminated
	for _, elfPos := range elves {
		if (elfPos.proposedPos != elfPos.currentPos) {
			moves = append(moves, elfPos.proposedPos)//adds all positions that an elf is moving to to the moves array (elves that stay still cannot have a conflict)
		}
	} 
	count := 0 
	for (count < len(moves) - 1) {
		if (moves[count] == moves[count + 1]) {
			conflicts = append(conflicts, moves[count])
		}
		count++
	}
	return conflicts
} 

//output[0] is "free to the north", [1] is south, [2], is west, and [3] is east. if all are true, returns all false because that elf won't move
func checkAround(location coords) [4]bool{
var output [4]bool
if (!elfIsAnyHere(coords{location.x, location.y - 1}, coords{location.x - 1, location.y - 1}, coords{location.x + 1, location.y - 1})) {
	output[0] = true
	}
if (!elfIsAnyHere(coords{location.x, location.y + 1}, coords{location.x - 1, location.y + 1}, coords{location.x + 1, location.y + 1}))	{
	output[1] = true
    }
if (!elfIsAnyHere(coords{location.x - 1, location.y}, coords{location.x - 1, location.y - 1}, coords{location.x - 1, location.y + 1})) {
	output[2] = true
}
if (!elfIsAnyHere(coords{location.x + 1, location.y}, coords{location.x + 1, location.y - 1}, coords{location.x + 1, location.y + 1})) {
	output[3] = true
}
if (output[0]&&output[1]&&output[2]&&output[3]) { //if all are true, elf doesn't move
	output = [4]bool{false, false, false, false}
}
return output
}

 


/*Sorts the elves by x-values (primary) and y-values (secondary)
or by y-values (if the isX parameter is false). Will also sort proposals
to check for proposal conflicts. Stability is important in the proposal checks
because it needs to have x-y pairs in roughly double-sorted order. 
*/
func sortElves(isProposal bool) {
	if (isProposal) {
		sort.Slice(elves, func(elf1, elf2 int) bool{
												if (elves[elf1].proposedPos.x == elves[elf2].proposedPos.x) {
													return elves[elf1].proposedPos.y < elves[elf2].proposedPos.y
													} else {
													return elves[elf1].proposedPos.x < elves[elf2].proposedPos.x	
													}
												})
	} else {
		sort.Slice(elves, func(elf1, elf2 int) bool{
												if (elves[elf1].currentPos.x == elves[elf2].currentPos.x) {
													return elves[elf1].currentPos.y < elves[elf2].currentPos.y
													} else {
													return elves[elf1].currentPos.x < elves[elf2].currentPos.x	
												}
											})
	}	
} 

func printAllElves() {
	for _, elf := range elves {
		fmt.Println(elf.currentPos)
	}
}


//returns TRUE if there is an elf in the coordinates specified by 
//the location parameter. can also check for proposal conflicts if the second parameter is TRUE

func elfIsHere(location coords, isProposal bool) bool{
	sortElves(isProposal) //sorts elves by x and y
	if (isProposal) {
		count := sort.Search(len(elves), func(elfIndex int) bool {
																if (elves[elfIndex].proposedPos.x == location.x) {
																	return (elves[elfIndex].proposedPos.y >= location.y)
																} else {
																	return (elves[elfIndex].proposedPos.x > location.x) //normally this would require >=, but equality has already been checked so only > is necessary
																} 
															})
		if (count == len(elves)) {
			return false //there is no elf here, and it's to the right of all other elves
		}
		if (elves[count].proposedPos == location) {
			return true //there is an elf here
		} else {
			return false
		}
	} else {
		count := sort.Search(len(elves), func(elfIndex int) bool {
																if (elves[elfIndex].currentPos.x == location.x) {
																	return (elves[elfIndex].currentPos.y >= location.y)
																} else {
																	return (elves[elfIndex].currentPos.x > location.x)
																} 
															})
		if (count == len(elves)) {
			return false //there is no elf here, and it's to the right of all other elves
		}
		if (elves[count].currentPos == location) {
			return true //there is an elf here
		} else {
			return false
		}
	}
}


//takes multiple locations, returns true if there is an elf in any of them. No option for 
//proposal checking because there proposal conflicts only ever need to check one spot
func elfIsAnyHere(locations ...coords) bool{ 
	for _, spot := range locations {
		if (elfIsHere(spot, false)) {
			return true
		}
	}
	return false
}