package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// download input from https://adventofcode.com/2024/day/1/input

	file, err := os.Open("./input.txt")

	if err != nil {
		fmt.Printf("Not able to read the file with err: %s", err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	arr1 := make([]int, 0)

	mp2 := make(map[int]int)

	for {

		str, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" { // Check for end of file
				if len(str) > 0 {
					fmt.Println(str) // Print the last line if it doesn't end with '\n'
				}
				break
			}
			fmt.Println("Error reading file:", err)
			break
		}

		splitStrs := strings.Split(str, "   ")

		num1, err := strconv.Atoi(splitStrs[0])
		arr1 = append(arr1, num1)

		num2, err := strconv.Atoi(strings.Split(splitStrs[1], "\n")[0])
		mp2[num2]++
	}

	ans := 0

	for _, val := range arr1 {
		ans += val * mp2[val]
	}
	fmt.Println(ans)

}
