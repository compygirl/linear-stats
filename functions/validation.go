package functions

import (
	"fmt"
	"regexp"
	"strings"
)

func IsFileValid(args []string) bool {
	if len(args) != 1 {
		fmt.Println("ERROR: the number of arguments is incorrect: SHOULD BE 1")
		return false
	} else {
		if !strings.HasSuffix(args[0], ".txt") {
			fmt.Println("ERROR: the file name has invalid extension. Should be .txt")
			return false
		}
		return true
	}
}

func IsDataValid(data []string) bool {
	if len(data) == 0 {
		fmt.Println("ERROR: the file is empty!")
		return false
	} else {
		pattern := `[^0-9+-]+`
		regExp := regexp.MustCompile(pattern)
		for _, str := range data {
			if regExp.MatchString(str) {
				fmt.Printf("ERROR: the data is not numeric")
				return false
			}
		}
	}
	return true
}
