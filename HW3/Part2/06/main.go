// https://projecteuler.net
// Number letter counts
// Problem 17
// If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
// If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?

// NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20 letters.
// The use of "and" when writing out numbers is in compliance with British usage.

package main

import (
	"fmt"

	"github.com/divan/num2words"
)

func main() {
	letterCount := 0
	for i := 1; i <= 1000; i++ {
		str := num2words.Convert(i)
		fmt.Println(str)
		for _, r := range str {
			c := string(r)
			if c != " " || c != "-" {
				letterCount++
			}
		}
	}
	fmt.Println(letterCount)
}
