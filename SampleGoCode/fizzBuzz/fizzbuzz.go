package main

import "fmt"

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

	fizzbuzz(number)
}

func fizzbuzz(n int) {

	if n%5 == 0 && n%3 == 0 {
		fmt.Println("fizzbuzz")
	} else if n%5 == 0 {
		fmt.Println("buzz")
	} else if n%3 == 0 {
		fmt.Println("fizz")
	} else {
		fmt.Println(n)
	}
}
