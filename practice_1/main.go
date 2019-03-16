package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// scanner from stdin

var scanner = bufio.NewScanner(os.Stdin)

func scanNextString() string {
	scanner.Scan()
	return scanner.Text()
}

func scanNextInt() int {
	s := scanNextString()
	i, err := strconv.Atoi(s)
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
	a := scanNextInt()
	b := scanNextInt()
	c := scanNextInt()
	s := scanNextString()

	fmt.Printf("%v %v\n", a+b+c, s)
}
