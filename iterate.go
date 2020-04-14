package main

import (
	"errors"
	"fmt"
	"math"
)

type Construct struct {
	x int
	y float64
}

func main() {
	// define a slice of an array
	arr := []int{3, 35}

	// define a dictionary like object
	dict := make(map[string]int)

	dict["ff"] = 23
	arr = append(arr, 22)
	fmt.Println(arr)
	fmt.Println(dict)

	// use a for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// iterate through an array or a dict
	for index, value := range dict {
		fmt.Println("index", index, "value", value)
	}

	result := sum(2, 3)
	fmt.Println("result:", result)

	sqrtresult, err := sqrt(-4)
	if err != nil {
		fmt.Println("An error occured:", err)
	} else {
		fmt.Println("Sqrt of a number is:", sqrtresult)
	}

	c := Construct{x: 5, y: 4.6}
	fmt.Println(c)

}

// function with single return types
func sum(x int, y int) int {
	return x + y
}

// function with multiple return types
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("undefined")
	}
	return math.Sqrt(x), nil
}
