package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	age0_2    = "Still in Mama's arms"
	age3_4    = "Preschool Maniac"
	age5_11   = "Elementary school"
	age12_14  = "Middle school"
	age15_18  = "High school"
	age19_22  = "College"
	age23_65  = "Working for the man"
	age66_100 = "The Golden Years"
	alien     = "This program is for humans"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		a := scanner.Text()
		age, err := strconv.Atoi(a)
		if err != nil {
			log.Fatalf("Couldn't parse the age %s as an int", a)
		}

		switch {
		case age < 0 || age > 100:
			fmt.Println(alien)
			break

		case age >= 66 && age <= 100:
			fmt.Println(age66_100)
			break

		case age >= 23 && age <= 65:
			fmt.Println(age23_65)
			break

		case age >= 5 && age <= 11:
			fmt.Println(age5_11)
			break

		case age >= 19 && age <= 22:
			fmt.Println(age19_22)
			break

		case age >= 15 && age <= 18:
			fmt.Println(age15_18)
			break

		case age >= 12 && age <= 14:
			fmt.Println(age12_14)
			break

		case age >= 0 && age <= 2:
			fmt.Println(age0_2)
			break

		case age >= 3 && age <= 4:
			fmt.Println(age3_4)
			break
		}
	}
}
