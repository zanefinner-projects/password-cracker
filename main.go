package main

import (
	"fmt"
	"strings"
)

func main() {
	answer := "a"
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	tries := 1
	charsSlice := strings.Split(chars, "")
	var pos []string
	fmt.Println("Finding all possible combinations...")
	for _, val1 := range charsSlice {
		tries++
		pos = append(pos, val1)
		for _, val2 := range charsSlice {
			tries++
			pos = append(pos, val1+val2)
			for _, val3 := range charsSlice {
				tries++
				pos = append(pos, val1+val2+val3)
				for _, val4 := range charsSlice {
					tries++
					pos = append(pos, val1+val2+val3+val4)
				}
			}
		}
	}
	x := whereAt(pos, answer)
	fmt.Println("Solved on try", x+1, " out of ", tries)
}
func whereAt(a []string, x string) int {
	for k, n := range a {
		if x == n {
			return k
		}
	}
	return -1
}
