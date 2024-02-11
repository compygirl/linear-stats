package main

import (
	"fmt"
	"linearstat/functions"
	"os"
	"strings"
)

func main() {
	files := os.Args[1:]
	if functions.IsFileValid(files) {
		text, err := os.ReadFile(files[0])
		if err != nil {
			fmt.Printf("ERROR: the file %v coudn't be read: %v", files[0], err.Error())
		}

		// fmt.Println(text)

		datastr := strings.Fields(string(text))
		if functions.IsDataValid(datastr) {
			data := functions.SliceAtoi(datastr) // convert str to int

			// avg := functions.Average(data)
			// med := functions.Median(data)
			// variance := functions.Variance(data)
			// stdv := functions.StdDev(data)
			// fmt.Printf("Average: %v\n", int(math.Round(avg)))
			// fmt.Printf("Median: %v\n", int(math.Round(med)))
			// fmt.Printf("Variance: %v\n", int(math.Round(variance)))
			// fmt.Printf("Standard Deviation: %v\n", int(math.Round(stdv)))

			coef_k, coef_b := functions.LinearRegLine(data)
			fmt.Print("Linear Regression Line: y = ")
			fmt.Printf("%.6f", coef_k)
			fmt.Print("x + ")
			fmt.Printf("%.6f", coef_b)
			fmt.Println()

			pearson := functions.PearsonCoef(data)
			fmt.Print("Pearson Correlation Coefficient: ")
			fmt.Printf("%.10f", pearson)
			fmt.Println()
		}

	}
}
