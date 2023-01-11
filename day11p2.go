package main

import "fmt"


var items [8][40]int64 //kinda want to use slices but it gets janky
var inspections [8]int

func main() {
	/* Input was hard-coded but has been redacted as AoC creator Eric Wastl 
	dislikes people publishing their inputs. It was in the form
	items[0][0] = 24
	items[0][1] = 42
	....
	items[7][7] = 421*/
	
	round := 0
	count := 0
	var throw int
	const primes = 9699690
	for (round < 10000) {	
	//there's probably a way to make this a function to handle the different operations
	//and tests and monkeys but i'd rather re-write
	inspections[0] += endOfArray(items[0])
	throw = endOfArray(items[0])
	count = 0
	for (count < throw) {
		items[0][count] *= 3
		if (items[0][count] % 5 == 0) {
			items[2][endOfArray(items[2])] = (items[0][count] % primes)
		} else {
			items[7][endOfArray(items[7])] = (items[0][count] % primes)
		}
		items[0][count] = 0
		count++
	}
	inspections[1] += endOfArray(items[1])
	throw = endOfArray(items[1])
	count = 0
	for (count < throw) {
		items[1][count] += 7
		if (items[1][count] % 2 == 0) {
			items[3][endOfArray(items[3])] = (items[1][count] % primes)
		} else {
			items[6][endOfArray(items[6])] = (items[1][count] % primes)
		}
		items[1][count] = 0
		count++
	}
	
	inspections[2] += endOfArray(items[2])
	throw = endOfArray(items[2])
	count = 0
	for (count < throw) {
		items[2][count]	+= 5
		if (items[2][count] % 13 == 0) {
			items[5][endOfArray(items[5])] = (items[2][count] % primes)
		} else {
			items[4][endOfArray(items[4])] = (items[2][count] % primes)
		}
		items[2][count] = 0
		count++
	}

	inspections[3] += endOfArray(items[3])
	throw = endOfArray(items[3])
	count = 0
	for (count < throw) {
		items[3][count]	+= 8
		if (items[3][count] % 19 == 0) {
			items[6][endOfArray(items[6])] = (items[3][count] % primes)
		} else {
			items[0][endOfArray(items[0])] = (items[3][count] % primes)
		}
		items[3][count] = 0
		count++
	}
	
	inspections[4] += endOfArray(items[4])
	throw = endOfArray(items[4])
	count = 0
	for (count < throw) {
		items[4][count]	+= 4
		if (items[4][count] % 11 == 0) {
			items[3][endOfArray(items[3])] = (items[4][count] % primes)
		} else {
			items[1][endOfArray(items[1])] = (items[4][count] % primes)
		}
		items[4][count] = 0
		count++
	}
	
	inspections[5] += endOfArray(items[5])
	throw = endOfArray(items[5])
	count = 0
	for (count < throw) {
		items[5][count] *= 2
		if (items[5][count] % 3 == 0) {
			items[4][endOfArray(items[4])] = (items[5][count] % primes)
		} else {
			items[1][endOfArray(items[1])] = (items[5][count] % primes)
		}
		items[5][count] = 0
		count++
	}
	
	inspections[6] += endOfArray(items[6])
	throw = endOfArray(items[6])
	count = 0
	for (count < throw) {			
		items[6][count]	+= 6
		if (items[6][count] % 7 == 0) {
			items[7][endOfArray(items[7])] = (items[6][count] % primes)
		} else {
			items[0][endOfArray(items[0])] = (items[6][count] % primes)
		}
		items[6][count] = 0
		count++
	}

	inspections[7] += endOfArray(items[7])
	throw = endOfArray(items[7])
	count = 0
	for (count < throw) {
		items[7][count]	*= items[7][count]
		if (items[7][count] % 17 == 0) {
			items[2][endOfArray(items[2])] = (items[7][count] % primes)
		} else {
			items[5][endOfArray(items[5])] = (items[7][count] % primes)
		}
		items[7][count] = 0
		count++
	}
	round++	
	}

	//janky code to find the top 2 and multiply them, with variable reuse
	count = 0
	throw = 0
	round = 0
	for (count < 8) {
		if (inspections[count] > throw) {
			throw = inspections[count]
		}
		count++
	}
	count = 0
	for (count < 8) {
		if (inspections[count] > round)&&(inspections[count] != throw) {
			round = inspections[count]
		}
		count++
	}
	fmt.Println(throw * round) 
}

//returns the position of the first empty (0) element in the array. I really hope
//no worry levels go to 0
func endOfArray(arr [40]int64) int{
	count := 0 
	for (arr[count] != 0) {
		count++
	}
	return count
	
}