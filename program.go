package main

import (
	"fmt"
)

func main() {

	callFunction()

	helloWorld()

	bitwiseShiftLeft()

	// return of function with multiple return values
	fmt.Println(helloMultipleVals())

	iterateForLoop()

	iterateUntilTrue()

	switchEx()

	deferPrint()

	pointersExample1()

	pointersStructExample()
	pointersStructLiteralExample()

	primes := []int { 3, 5, 7, 11, 13, 17,  19, 23, 29, 31}
	sliceArray(primes)
	dynamicSlice(primes)


}