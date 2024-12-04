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
			ans += CheckN(i, j, grid)
			ans += CheckNW(i, j, grid)
			ans += CheckW(i, j, grid)
			ans += CheckSW(i, j, grid)
			ans += CheckS(i, j, grid)
			ans += CheckSE(i, j, grid)
			ans += CheckE(i, j, grid)
			ans += CheckNE(i, j, grid)

		}
	}

	fmt.Println(ans)
}

func Validate(s1 string, s2 string, s3 string, s4 string) int {
	if s1 == "X" && s2 == "M" && s3 == "A" && s4 == "S" {
		return 1
	}
	return 0
}

func CheckN(i int, j int, grid []string) int {
	n := len(grid)
	m := len(grid[0])

	if i < n && i-3 >= 0 && j < m && j >= 0 {
		return Validate(string(grid[i][j]), string(grid[i-1][j]), string(grid[i-2][j]), string(grid[i-3][j]))
	}
	return 0
}

func CheckNW(i int, j int, grid []string) int {
	n := len(grid)
	m := len(grid[0])

	if i < n && i-3 >= 0 && j < m && j-3 >= 0 {
		return Validate(string(grid[i][j]), string(grid[i-1][j-1]), string(grid[i-2][j-2]), string(grid[i-3][j-3]))
	}
	return 0
}

func CheckW(i int, j int, grid []string) int {
	n := len(grid)
	m := len(grid[0])

	if i < n && i >= 0 && j < m && j-3 >= 0 {
		return Validate(string(grid[i][j]), string(grid[i][j-1]), string(grid[i][j-2]), string(grid[i][j-3]))
	}
	return 0
}

func CheckSW(i int, j int, grid []string) int {
	n := len(grid)
	m := len(grid[0])

	if i+3 < n && i >= 0 && j < m && j-3 >= 0 {
		return Validate(string(grid[i][j]), string(grid[i+1][j-1]), string(grid[i+2][j-2]), string(grid[i+3][j-3]))
	}
	return 0
}

func CheckS(i int, j int, grid []string) int {
	n := len(grid)
	m := len(grid[0])

	if i+3 < n && i >= 0 && j < m && j >= 0 {
		return Validate(string(grid[i][j]), string(grid[i+1][j]), string(grid[i+2][j]), string(grid[i+3][j]))
	}
	return 0
}

func CheckSE(i int, j int, grid []string) int {
	n := len(grid)
	m := len(grid[0])

	if i+3 < n && i >= 0 && j+3 < m && j >= 0 {
		return Validate(string(grid[i][j]), string(grid[i+1][j+1]), string(grid[i+2][j+2]), string(grid[i+3][j+3]))
	}
	return 0
}

func CheckE(i int, j int, grid []string) int {
	n := len(grid)
	m := len(grid[0])

	if i < n && i >= 0 && j+3 < m && j >= 0 {
		return Validate(string(grid[i][j]), string(grid[i][j+1]), string(grid[i][j+2]), string(grid[i][j+3]))
	}
	return 0
}

func CheckNE(i int, j int, grid []string) int {
	n := len(grid)
	m := len(grid[0])

	if i < n && i-3 >= 0 && j+3 < m && j >= 0 {
		return Validate(string(grid[i][j]), string(grid[i-1][j+1]), string(grid[i-2][j+2]), string(grid[i-3][j+3]))
	}
	return 0
}
