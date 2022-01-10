package main

import (
	"fmt"
	"strconv"
)

func main() {

	var y int

	fmt.Println("Enter a number")
	if _, err := fmt.Scanln(&y); err != nil {
		fmt.Println("Error:", err)
	}
	var sum int
	str := strconv.Itoa(y)
	if len(str)%2 == 0 {

		for i, v := range str {
			if (i+1)%2 == 0 {
				sum += int(v - 48)
			} else {
				sum += doubleDigitValue(int(v - 48))
			}
		}
	} else {
		for i, v := range str {
			if (i+1)%2 != 0 {
				sum += int(v - 48)
			} else {
				sum += doubleDigitValue(int(v - 48))
			}
		}
	}
	if sum%10 == 0 {
		fmt.Printf("Checksum %d is divisible by 10. Valid\n", sum)
	} else {
		fmt.Printf("Checksum %d is not divisible by 10. Invalid\n", sum)
	}
}

func doubleDigitValue(digit int) (sum int) {
	doubleDigit := digit * 2

	if doubleDigit >= 10 {
		sum = 1 + doubleDigit%10
	} else {
		sum = doubleDigit
	}
	return sum
}
