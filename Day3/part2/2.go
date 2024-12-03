package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		fmt.Printf("unable to read the input file %s", err)
	}

	reader := bufio.NewReader(file)
	ans := 0
	history := 1
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Printf("String in last line %s", str)
			}
			break
		}

		startpts := getMUL(str, &history)

		for _, val := range startpts {
			// start from the val MUL then next ( be found at val+3
			if val+3 < len(str) && str[val+3] == '(' {
				// from now try finding ')'
				for i := val + 3; i < len(str); i++ {
					if str[i] == ')' {
						values := validate(val+3, i, str)
						ans += values[0] * values[1]
						break
					}
				}
			}
		}

	}
	fmt.Println(ans)
}

func getMUL(str string, history *int) []int {

	starts := make([]int, 0)
	for idx := 0; idx < len(str); idx++ {
		if idx+2 < len(str) && str[idx] == 'm' && str[idx+1] == 'u' && str[idx+2] == 'l' && *history == 1 {
			starts = append(starts, idx)
		}
		if idx+3 < len(str) && str[idx] == 'd' && str[idx+1] == 'o' && str[idx+2] == '(' && str[idx+3] == ')' {
			*history = 1
		}
		if idx+6 < len(str) && str[idx] == 'd' && str[idx+1] == 'o' && str[idx+2] == 'n' && str[idx+3] == '\'' && str[idx+4] == 't' && str[idx+5] == '(' && str[idx+6] == ')' {
			*history = 0
		}
	}
	return starts
}

//117549707

func validate(start int, end int, str string) []int {
	// comma seperated split
	x := str[start : end+1]
	substrs := strings.Split(x, ",")
	if len(substrs) != 2 {
		return []int{0, 0}
	} else {
		a := checkNums(substrs[0][1:])
		b := checkNums(substrs[1][:len(substrs[1])-1])

		return []int{a, b}
	}
}

func checkNums(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error converting character:", err)
		return 0
	}

	return num
}
