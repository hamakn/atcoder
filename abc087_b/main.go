package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// scanner from stdin

var scanner = bufio.NewScanner(os.Stdin)

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func scanInt() int {
	i, err := strconv.Atoi(scanString())
	if err != nil {
		panic(err)
	}
	return i
}

func init() {
	scanner.Split(bufio.ScanWords)
}

// write code below

func main() {
	num500, num100, num50 := scanInt(), scanInt(), scanInt()
	x := scanInt()

	c := 0
LABEL1:
	for c500 := 0; c500 <= num500; c500++ {
		// fmt.Printf("c500: %v\n", c500)
		if c500*500 > x {
			// fmt.Printf("500 continue: %v\n", c500)
			break
		}
	LABEL2:
		for c100 := 0; c100 <= num100; c100++ {
			// fmt.Printf("  c100: %v\n", c100)
			if c500*500+c100*100 > x {
				// fmt.Printf("  100 continue: %v\n", c100)
				continue LABEL1
			}
			for c50 := 0; c50 <= num50; c50++ {
				// fmt.Printf("    c50: %v\n", c50)
				if c500*500+c100*100+c50*50 == x {
					// fmt.Printf("    50 ok: %v\n", c50)
					c++
					continue LABEL2
				}
			}
		}
	}
	fmt.Println(c)
}
