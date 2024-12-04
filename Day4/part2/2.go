package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		fmt.Printf("unable to read the input file %s", err)
	}

	reader := bufio.NewReader(file)

	// create a 2d slice...

	grid := make([]string, 0)

	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Printf("String in last line %s", str)
			}
			break
		}
		grid = append(grid, str)
	}

	// process grid.

	ans := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			x := CheckCentric(i, j, grid)
			ans += x

		}
	}

	fmt.Println(ans)
}

func Validate(s1 string, s2 string, s3 string) bool {
	if s1 == "M" && s2 == "A" && s3 == "S" {
		return true
	}

	if s3 == "M" && s2 == "A" && s1 == "S" {
		return true
	}

	return false
}

func CheckCentric(i int, j int, grid []string) int {
	n := len(grid)
	m := len(grid[0])

	if i-1 >= 0 && i+1 < n && j-1 >= 0 && j+1 < m {
		if Validate(string(grid[i-1][j-1]), string(grid[i][j]), string(grid[i+1][j+1])) && Validate(string(grid[i-1][j+1]), string(grid[i][j]), string(grid[i+1][j-1])) {
			return 1
		}
	}

	return 0
}
