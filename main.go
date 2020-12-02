package main

import "fmt"

func main() {
	day1Data := getData()
	fmt.Println(
		onePart1(day1Data),
		onePart2(day1Data),
		twoPart1(),
		twoPart2(),
	)
}
