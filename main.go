package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println()
	start := time.Now()
	answer := "hey!"
	chars := "`~!@#$%^&*()_+-={[}]\\|;:/?.>,<qazwsxedcrfvtgbyhnujmikolpQAZWSXEDCRFVTGBYHNUJMIKOLP"
	tries := 0
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
	elapsed := time.Since(start)
	comPerSec := float64(tries+1) / elapsed.Seconds()
	if x != -1 {
		fmt.Println("Solved on try", x+1, " out of ", tries)
		fmt.Println("Stats")
		fmt.Println("  | Exec Time     : ", elapsed)
		fmt.Println("  | Tries per Sec :  ~", strconv.FormatFloat(comPerSec, 'f', 0, 32))
		fmt.Println()
	} else {
		fmt.Print("not found...")
	}
}
func whereAt(a []string, x string) int {
	for k, n := range a {
		if x == n {
			return k
		}
	}
	return -1
}
