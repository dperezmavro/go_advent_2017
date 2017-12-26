package main

import "fmt"

func calculateSideSize(number int) int {
	var previousLayer = 1

	for {
		if number > previousLayer*previousLayer {
			previousLayer += 2
		} else {
			return previousLayer
		}
	}
}

func calculateXYOnSquare(number int, gridSize int) (int, int) {
	var x = gridSize
	var y = gridSize - 1
	var firstNumber = ((gridSize - 2) * (gridSize - 2)) + 1
	var currentNumber int
	for currentNumber = firstNumber; currentNumber < number; currentNumber++ {
		x++
		y++
	}
	return x, y
}

func main() {
	fmt.Println(calculateSideSize(1))
	fmt.Println(calculateSideSize(8))
	fmt.Println(calculateSideSize(25))
	fmt.Println(calculateSideSize(46))
	fmt.Println(calculateSideSize(289326))

}
