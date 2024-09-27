package main

import "fmt"

func main() {
	var minimumWage int = 10
	var workingHour int = 30

	var income int = minimumWage * workingHour

	fmt.Println("최저 시급:", minimumWage)
	fmt.Println("근무 시간:", workingHour)
	fmt.Println("급여:", income)
}
