package main

import "fmt"


var items [8][40]int //kinda want to use slices but it gets janky
var inspections [8]int

func main() {
	/* Input was hard-coded but has been redacted as AoC creator Eric Wastl 
	dislikes people publishing their inputs. It was in the form
	items[0][0] = 22
	items[0][1] = 42
	....
	items[7][7] = 421*/
	
	round := 1
	count := 0
	var throw int
	const primes = 1000000
	for (round <= 20) {	
	//there's probably a way to make this a function to handle the different operations
	//and tests and monkeys but i'd rather re-write
	inspections[0] += endOfArray(items[0])
	throw = endOfArray(items[0])
	count = 0
	for (count < throw) {
		if (items[0][count] % 5 == 0) {
			//monkey 0 multiplies by three then divides by three so nothing happens
			items[2][endOfArray(items[2])] = items[0][count]
		} else {
			items[7][endOfArray(items[7])] = items[0][count]
		}
		items[0][count] = 0
		count++
	}
	inspections[1] += endOfArray(items[1])
	throw = endOfArray(items[1])
	count = 0
	for (count < throw) {
		items[1][count] += 7
		items[1][count] /= 3
		if (items[1][count] % 2 == 0) {
			items[3][endOfArray(items[3])] = items[1][count]
		} else {
			items[6][endOfArray(items[6])] = items[1][count]
		}
		items[1][count] = 0
		count++
	}
	
	inspections[2] += endOfArray(items[2])
	throw = endOfArray(items[2])
	count = 0
	for (count < throw) {
		items[2][count]	+= 5
		items[2][count] /= 3
		if (items[2][count] % 13 == 0) {
			items[5][endOfArray(items[5])] = items[2][count]
		} else {
			items[4][endOfArray(items[4])] = items[2][count]
		}
		items[2][count] = 0
		count++
	}

	inspections[3] += endOfArray(items[3])
	throw = endOfArray(items[3])
	count = 0
	for (count < throw) {
		items[3][count]	+= 8
		items[3][count] /= 3
		if (items[3][count] % 19 == 0) {
			items[6][endOfArray(items[6])] = items[3][count]
		} else {
			items[0][endOfArray(items[0])] = items[3][count]
		}
		items[3][count] = 0
		count++
	}
	
	inspections[4] += endOfArray(items[4])
	throw = endOfArray(items[4])
	count = 0
	for (count < throw) {
		items[4][count]	+= 4
		items[4][count] /= 3
		if (items[4][count] % 11 == 0) {
			items[3][endOfArray(items[3])] = items[4][count]
		} else {
			items[1][endOfArray(items[1])] = items[4][count]
		}
		items[4][count] = 0
		count++
	}
	
	inspections[5] += endOfArray(items[5])
	throw = endOfArray(items[5])
	count = 0
	for (count < throw) {
		items[5][count]	*= 2
		items[5][count] /= 3
		if (items[5][count] % 3 == 0) {
			items[4][endOfArray(items[4])] = items[5][count]
		} else {
			items[1][endOfArray(items[1])] = items[5][count]
		}
		items[5][count] = 0
		count++
	}
	
	inspections[6] += endOfArray(items[6])
	throw = endOfArray(items[6])
	count = 0
	for (count < throw) {
		items[6][count]	+= 6
		items[6][count] /= 3
		if (items[6][count] % 7 == 0) {
			items[7][endOfArray(items[7])] = items[6][count]
		} else {
			items[0][endOfArray(items[0])] = items[6][count]
		}
		items[6][count] = 0
		count++
	}

	inspections[7] += endOfArray(items[7])
	throw = endOfArray(items[7])
	count = 0
	for (count < throw) {
		items[7][count]	*= items[7][count]
		items[7][count] /= 3
		if (items[7][count] % 17 == 0) {
			items[2][endOfArray(items[2])] = items[7][count]
		} else {
			items[5][endOfArray(items[5])] = items[7][count]
		}
		items[7][count] = 0
		count++
	}
	round++	
	}
	fmt.Println(inspections) //only 8 of them so find the top 2 by "inspection"
}

//returns the position of the first empty (0) element in the array. I really hope
//no worry levels go to 0
func endOfArray(arr [40]int) int{
	count := 0 
	for (arr[count] != 0) {
		count++
	}
	return count
	
}