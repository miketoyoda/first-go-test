package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CalculateSumOfSquares(testValuesString *[]string, params ...int) (int, error) {
	index := 0
	if len(params) == 1 {
		index = params[0]
	}

	head := (*testValuesString)[index]
	value, err := strconv.Atoi(head)

	if err != nil {
		return value, err
	}

	if value > 100 || value < -100 {
		return value, fmt.Errorf("Test input cannot exceed the range [-100, 100]. Value: %d", value)
	}

	squares := 0
	if value > 0 {
		squares = value * value
	}

	if (len(*testValuesString) - 1) == index {
		return squares, nil
	}

	index++
	accumulated, err := CalculateSumOfSquares(testValuesString, index)

	if err != nil {
		return accumulated, err
	}

	return accumulated + squares, nil
}

func FetchTestCases(remainingTestCases int, resultsList *[]int) error {
	if 0 == remainingTestCases {
		return nil
	}

	var numElements uint
	fmt.Scanln(&numElements)
	if numElements == 0 || numElements > 100 {
		return fmt.Errorf("Number of input values is out of range [0, 100]: X=%d", numElements)
	}

	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	values := strings.Fields(line)

	if numElements != uint(len(values)) {
		return fmt.Errorf("Number of input values is not equal to X=%d, values length=%d", numElements, len(values))
	}

	sumOfSquares, err := CalculateSumOfSquares(&values)

	if err != nil {
		return err
	}

	(*resultsList)[len((*resultsList))-int(remainingTestCases)] = sumOfSquares

	remainingTestCases--
	return FetchTestCases(remainingTestCases, resultsList)
}

func PrintResults(results *[]int, params ...int) {
	index := 0
	if len(params) == 1 {
		index = params[0]
	}

	if len(*results) == index {
		return
	}

	fmt.Println((*results)[index])

	index++
	PrintResults(results, index)
}

func main() {
	var numTestCases uint
	fmt.Scanln(&numTestCases)

	if numTestCases == 0 || numTestCases > 100 {
		fmt.Println("Number of test cases is out of range [0, 100]: X=", numTestCases)
		os.Exit(1)
	}

	results := make([]int, numTestCases)
	err := FetchTestCases(int(numTestCases), &results)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	PrintResults(&results)
}
