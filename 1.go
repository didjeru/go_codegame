package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	// n: the number of temperatures to analyse
	var n int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n)

	scanner.Scan()
	inputs := strings.Split(scanner.Text(), " ")
	var arr [500]int64

	for i := 0; i < n; i++ {
		// t: a temperature expressed as an integer ranging from -273 to 5526
		t, _ := strconv.ParseInt(inputs[i], 10, 32)
		arr[i] = t
	}

	result := findNearestValue(arr, n)
	//fmt.Fprintln(os.Stderr, "Debug messages...")
	//Write answer to stdout
	fmt.Println(result)
}

func findNearestValue(arr [500]int64, n int) int64 {
	min := arr[0]
	for i := 0; i < n; i++ {
		if math.Abs(float64(arr[i])) < math.Abs(float64(min)) {
			min = arr[i]
		} else {
			if math.Abs(float64(arr[i])) == math.Abs(float64(min)) && arr[i] > 0 {
				min = arr[i]
			}
		}
	}
	return min
}
