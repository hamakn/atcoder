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
	n := scanInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = scanInt()
	}

	count := 0
	for {
		for i, x := range arr {
			if x&1 == 1 {
				fmt.Println(count)
				return
			}
			arr[i] = x >> 1
		}
		count++
	}
}
