package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	arr2 := make([]int, 0)

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
		arr2 = append(arr2, num2)
	}

	sort.Ints(arr1)
	sort.Ints(arr2)

	ans := 0

	for idx := range arr1 {
		if arr1[idx] > arr2[idx] {
			ans += arr1[idx] - arr2[idx]
		} else {
			ans += arr2[idx] - arr1[idx]
		}
	}
	fmt.Println(ans)

}
