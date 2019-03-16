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
	x := scanInt() * scanInt()

	if x%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
