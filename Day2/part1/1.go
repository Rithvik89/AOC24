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

	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Printf("String in last line %s", str)
			}
			break
		}
		levels_ := strings.Split(str, " ")
		// iterate over the levels_ and place the int values in levels.

		level_int_ := make([]int, 0)

		for index := range levels_ {
			var val int
			if index+1 == len(levels_) {
				val, err = strconv.Atoi(strings.Split(levels_[index], "\n")[0])
			} else {
				val, err = strconv.Atoi(levels_[index])
			}

			if err != nil {
				fmt.Printf("Unable to convert into int")
			}
			level_int_ = append(level_int_, val)
		}

		ans += calc(level_int_)

	}

	fmt.Println(ans)
}

func calc(level_int_ []int) int {
	monotonic_asc := 1
	monotonic_desc := 1
	adj_cond := 1

	for idx, val := range level_int_ {
		if idx > 0 && level_int_[idx-1] > val {
			x := level_int_[idx-1] - val
			if x < 1 || x > 3 {
				adj_cond = 0
			}
			monotonic_asc = 0
		}
		if idx > 0 && level_int_[idx-1] < val {
			x := val - level_int_[idx-1]
			if x < 1 || x > 3 {
				adj_cond = 0
			}
			monotonic_desc = 0
		}
		if idx > 0 && level_int_[idx-1] == val {
			monotonic_asc = 0
			monotonic_desc = 0
		}

	}

	if monotonic_asc == 1 || monotonic_desc == 1 {
		if adj_cond == 1 {
			return 1
		}
	}
	return 0
}
