package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
)

func main() {
	var input string

	n, err := fmt.Scanf("%s", &input)
	if n != 1 {
		log.Fatalln("Only one argument is accepted.")
	}
	if err != nil {
		log.Fatalln(err)
	}
	if !checkInput(input) {
		log.Fatalln("Input is not correct.")
	}
	intInput, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		log.Fatalln("Conversion to integer failed.")
	}

	palindrome := make(chan uint64, 1)
	palindrome <- 0

	for next := intInput + 1; <-palindrome == 0 && next <= math.MaxUint64; next++ {
		go func(i uint64) {
			if isPalindrome(i) {
				palindrome <- i
			}
		}(next)
	}

	if <-palindrome <= intInput {
		log.Fatalln("Failed to count difference.")
	}
	fmt.Println(<-palindrome - intInput)

}

func checkInput(i string) bool {
	re := regexp.MustCompile(`^[0-9]{2,}$`)
	return re.MatchString(i)
}

func reverseString(str string) string {
	var revStr string
	for i := len(str) - 1; i >= 0; i-- {
		revStr += string(str[i])
	}
	return revStr
}

func isPalindrome(value interface{}) bool {
	var str string
	if v, ok := value.(uint64); ok {
		str = strconv.FormatUint(v, 10)
	} else if v, ok := value.(string); ok {
		str = v
	} else {
		log.Fatalln("Palindrome check failed.")
	}
	return str == reverseString(str)
}
