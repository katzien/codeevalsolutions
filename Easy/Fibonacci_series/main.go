package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/katzien/fibonacci/fibonacci"
)

// InputReader represents the reader of input values
type InputReader interface {
	ReadInput() []int64
}

func main() {

	filename := os.Args[1]

	fr := FileInputReader{filename}

	numbers := ReadInputFromFile(fr)

	for _, n := range numbers {
		answer := fibonacci.Fibonacci(n)
		fmt.Println(answer)
	}
}

// ReadInputFromFile will read the input values using a given file reader
func ReadInputFromFile(fr InputReader) []int64 {
	return fr.ReadInput()
}

// FileInputReader reads input values from a file
type FileInputReader struct {
	filename string
}

// ReadInput is a file-based implementation of the InputReader's interface function
func (fr FileInputReader) ReadInput() []int64 {
	var input []int64

	if fr.filename == "" {
		log.Fatal("No input file specified")
	}

	file, err := os.Open(fr.filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		input = append(input, number)
	}

	return input
}
