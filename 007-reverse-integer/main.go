package main

import (
	"fmt"
	"math"
)

func main() {
	res := reverse(1463847412)
	fmt.Println(res)
}

func reverse(num int) (revNum int) {
	for num != 0 {
		pop := num % 10
		num /= 10

        if revNum > math.MaxInt32/10 || (revNum == math.MaxInt32/10 && pop > 7) { return 0 }
		if revNum < math.MinInt32/10 || (revNum == math.MinInt32/10 && pop < -8) { return 0 }

		revNum = revNum*10 + pop
	}
	return
}
