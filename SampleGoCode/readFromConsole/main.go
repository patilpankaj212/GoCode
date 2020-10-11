package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
read two lines from user,
1st line contains an integer
second line contains a string

print integer * 2
print string
*/

func main() {

	//define reader to read from standard input
	reader := bufio.NewReader(os.Stdin)

	//read two lines from user
	integer, _ := reader.ReadString('\n')
	text, _ := reader.ReadString('\n')

	// replace newline character from the inputs
	integer = strings.Replace(integer, "\n", "", -1)
	text = strings.Replace(text, "\n", "", -1)

	//convert string to integer
	i, _ := strconv.Atoi(integer)

	fmt.Println(i * 2)
	fmt.Println(text)

	// Using Scanf function from fmt package
	var number int
	var userString string

	fmt.Scanf("%d", &number)
	fmt.Scanf("%s", &userString)

	fmt.Println(number * 2)
	fmt.Println(userString)
}
