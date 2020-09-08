package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Three constants: string, int, bool
const (
	Author = "kev"
	bigNum = int(math.MaxUint32)
	isTutorial = true
)

func helloWorld() {
	// Date Time now function
	fmt.Println("Hello ", Author, ". The time is ", time.Now(), ". This is a tutorial? ", isTutorial, " level ", bigNum)
}

func callFunction() {
	var firstNum int = rand.Intn(50)
	var secondNum int = rand.Intn(72)


	// Simple function call
	fmt.Println("The sum of ", firstNum, " and ", secondNum, " is ", add(firstNum, secondNum))
}

// equivalent of x * 2^y (n times 2, y times) where 1 is x and 5 is y. >> is otherwise x * 2^-y (n divided by 2, y times)
func bitwiseShiftLeft() {

	shiftedVal := 1 << 5
	fmt.Println("shiftedVal", shiftedVal)
}

// Return function with multiple (2) values
func helloMultipleVals() (string, string) {
	return "Hello", "Sir"
}

// For loop
func iterateForLoop() {
	// for loop
	for i := 0; i < 10; i++ {
		fmt.Print("Counter ", i, ", ")
	}

	fmt.Print("\n")
}

// While loop
func iterateUntilTrue() {

	sumIt := true
	sum := 1
	for sumIt == true {
		sum += sum + 1

		fmt.Print("sum = ", sum, ". ")
		if sum > 50 {
			sumIt = false
		}
	}

	fmt.Println("Sum value = ", sum)
}

// switch case
func switchEx() {

	todayDay := time.Now().Weekday()
	fmt.Println("Current day = ", todayDay)

	switch time.Tuesday {
	case todayDay + 0:
		fmt.Println("Today is Tuesday")
	case todayDay + 1:
		fmt.Println("Tuesday is tomorrow")
	case todayDay + 2:
		fmt.Println("Tuesday is in two days")
	default:
		fmt.Println("Tuesday is further than 2 days")
	}
}

func deferPrint() {
	defer fmt.Println("(Delayed) ....not") // multiple calls of defer within a block are stacked and executed in LIFO order

	fmt.Print("This suit is black........")
}

func pointersExample1() {

	i, j := 22, 2020

	p := &i

	fmt.Println("Value of p", p)
	fmt.Println("Value of i through pointer p", *p)
	*p = 25
	fmt.Println("Value of i", i)

	p = &j
	*p = *p / 200
	fmt.Println("Value of j divided by 200", j)
}

func pointersStructExample() {

	bmw := Car{"bmw", "m5", "e39", 30000}

	fmt.Println("Car price", bmw.price)

	p := &bmw
	p.price = 32000

	fmt.Println("Value of p", p)
	fmt.Println("Car price change via pointer p", bmw.price)
}

func pointersStructLiteralExample() {

	v1 := Car{"bmw", "m5", "e39", 30000}
	v2 := Car{brand: "bmw", model: "m5", trim: "e39"}

	fmt.Println("Cars (Literal): ", v1, v2)
}

func sliceArray(numbers []int) {

	fmt.Println("Slice - Input array: ", numbers)

	if len(numbers) > 6 {

		fmt.Println("Sliced array: ", numbers[1:6])
	}

	fmt.Println("Make'd a slice array of 5 length: ", make([]int, 5))
}

func dynamicSlice(numbers []int) {

	fmt.Println("Dynamic Slice - Input array: ", numbers)

	if len(numbers) > 6 {

		printArrayDetails("Sliced to zero length:", numbers[:0])
		printArrayDetails("Sliced to zero length:", numbers[:4])
		printArrayDetails("Sliced to zero length:", numbers[2:])
	}
}

