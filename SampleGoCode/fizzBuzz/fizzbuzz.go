package main

import (
	"fmt"
	"strconv"
)

/*
if a number is divisible by 3 then print 'fizz'
if a number is divisible by 5 then print 'buzz'
if a number is divisible by both 3 and 5 then print 'fizzbuzz'
for every other number just print the number
*/
func main() {

	fmt.Print("enter a number: ")
	var number int
	fmt.Scanf("%d", &number)

	value := fizzbuzz(number)
	fmt.Println(value)
}

func fizzbuzz(n int) string {

	if n%5 == 0 && n%3 == 0 {
		return "fizzbuzz"
	} else if n%5 == 0 {
		return "buzz"
	} else if n%3 == 0 {
		return "fizz"
	} else {
		return strconv.FormatInt(int64(n), 10)
	}
}
