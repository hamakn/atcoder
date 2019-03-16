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
	scanner.Split(bufio.ScanRunes)
}

// write code below

func main() {
	count := 0
	for i := 0; i < 3; i++ {
		if scanInt() == 1 {
			count++
		}
	}
	fmt.Println(count)
}
