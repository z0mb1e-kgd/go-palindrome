package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mx sync.Mutex
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

	intInput, err := strconv.Atoi(input)
	if err != nil {
		log.Fatalln("Conversion to integer failed.")
	}

	palindrome := 0

	for next := intInput + 1; palindrome == 0; next++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if isPalindrome(i) {
				mx.Lock()
				defer mx.Unlock()
				palindrome = i
			}
		}(next)
	}

	wg.Wait()

	if palindrome <= intInput {
		log.Fatalln("Failed to count difference.")
	}
	fmt.Println(palindrome - intInput)

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
	if v, ok := value.(int); ok {
		str = strconv.Itoa(v)
	} else if v, ok := value.(string); ok {
		str = v
	} else {
		log.Fatalln("Palindrome check failed.")
	}
	return str == reverseString(str)
}
