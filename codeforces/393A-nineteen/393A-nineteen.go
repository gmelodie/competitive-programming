package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var n, i, e, t int
	var min int
	var div []int

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	for _, ch := range line {
		switch ch {
		case 'n':
			n++
		case 'i':
			i++
		case 'e':
			e++
		case 't':
			t++
		}
	}

	div = append(div, n/3)
	div = append(div, i)
	div = append(div, e/3)
	div = append(div, t)

	fmt.Println(div)

	min = math.MaxInt32
	for _, v := range div {
		if v < min {
			min = v
		}
	}

	fmt.Println(min)
}
