package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func strToBin(s string) (binString string) {
	for _, c := range s {
		binString = fmt.Sprintf("%s%07b", binString, c)
	}
	return
}

func stringToChack(s string) {
	for i := 0; i < len(s); i++ {
		var min [500]string
		c := s[i]
		if c == '1' {
			if i > 0 {
				if s[i-1] == '1' {
					min[i] = "0"
				} else {
					min[i] = " 0 0"
				}
			} else {
				min[i] = "0 0"
			}
		} else {
			if i > 0 {
				if s[i-1] == '0' {
					min[i] = "0"
				} else {
					min[i] = " 00 0"
				}
			} else {
				min[i] = "00 0"
			}
		}
		fmt.Print(min[i])
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	scanner.Scan()
	message := scanner.Text()
	binMessage := strToBin(message)

	stringToChack(binMessage)
}
