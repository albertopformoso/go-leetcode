package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	var rows []strings.Builder
	
	for i := 0; i < int(math.Min(float64(numRows), float64(len(s)))); i++ {
		rows = append(rows, strings.Builder{})
	}

	var currRow int
	var goingDown bool

	for _, c := range s {
		rows[currRow].WriteRune(c)
		if currRow == 0 || currRow == numRows - 1 {
			goingDown = !goingDown
		}
		if goingDown {
			currRow++
		} else {
			currRow--
		}
	}

	var ret strings.Builder
	for _, row := range rows {
		ret.WriteString(row.String())
	}

	return ret.String()
}
