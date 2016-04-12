package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"fibonacci"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, e := strconv.Atoi(scanner.Text())
		if e != nil {
			fmt.Println(e)
		}

		answer := fibonacci.Fibonacci(number)
		fmt.Println(answer)
	}
}
