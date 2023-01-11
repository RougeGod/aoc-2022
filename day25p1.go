package main

import "fmt"
import "strings"
import "io/ioutil"
import "os"


func main() {
filecontents, err := ioutil.ReadFile("./" + os.Args[1])
if (err != nil) {
	fmt.Println("oh no it broke")
	return
}
file := string(filecontents)
process(file)
}

func process(file string) {
	lines := strings.Split(file, "\n")
	count := 0
	totalScore := 0
	for (count < len(lines)) {
		totalScore += toInt(lines[count])
		count++
	}
	fmt.Println(toSnafu(totalScore))
}

func toInt(input string) int{
	output := 0
	chopped := strings.Split(rev(input), "")
	count := 0
	power := 1
	for (count < len(chopped)) {
		switch (chopped[count]) {
			case "=": output -= (2*power)
					  break
			case "-": output -= power
					  break
			case "1": output += power
					  break
			case "2": output += (2*power)
					  break
			
		}
		count++
		power *= 5
	}
	return output
}
func toSnafu(input int) string{
	digit := -1
	reversedOutput := make([]string, 30)
	for (input != 0) {
		digit++
		switch (input % 5) {
			case 0: reversedOutput[digit] = "0"
					break
			case 1: reversedOutput[digit] = "1"
					input--
					break
			case 2: reversedOutput[digit] = "2"
					input -= 2
					break
			case 3: reversedOutput[digit] = "="
					input += 2
					break
			case 4: reversedOutput[digit] = "-"
					input++
					break
		   default: fmt.Println("something is terribly wrong")
		}
		input /= 5 //integer division is fine because number is made divisible by 5 through the switch
	}
	return rev(strings.Join(reversedOutput[0:digit+1], "")) //end index exclusive
}



func rev(input string) string{
	s := strings.Split(input, "")
	count := 0
	var temp string
	for (count < (len(s)/2)) {
		temp = s[count]
		s[count] = s[len(s) - count - 1]
		s[len(s) - count - 1] = temp
		count++
	}
	return strings.Join(s, "")
	}