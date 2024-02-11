package functions

import (
	"fmt"
	"strconv"
)

// this function converts the array of string
// to array of integers
// input: array of strings
// output: converted array of integers
func SliceAtoi(inp []string) []int {
	var res []int
	for _, s := range inp {
		num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
			break
		}
		res = append(res, num)
	}
	return res
}
