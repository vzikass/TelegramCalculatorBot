package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func calculator(expr string) string {
	if expr == "" {
		return ""
	}

	factorialResult, err := processFactorial(expr)
	if err == nil {
		return factorialResult
	}

	re := regexp.MustCompile(`^\s*(-?\d+)\s*([-+*/])\s*(-?\d+)\s*$`)
	matches := re.FindStringSubmatch(expr)
	if len(matches) != 4 {
		return "Invalid expression format"
	}
	num1, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return "Invalid operand: " + matches[1]
	}
	num2, err := strconv.ParseFloat(matches[3], 64)
	if err != nil {
		return "invalid operand: " + matches[3]
	}
	operator := matches[2]
	idx := 0
	strnum1 := fmt.Sprintf("%f", num1)
	strnum2 := fmt.Sprintf("%f", num2)
	if strnum1[1] == byte(idx) {
		return "Number starts from zero"
	}
	if strnum2[3] == byte(idx) {
		return "Number starts from zero"
	}
	switch {
	case num1 == 0:
		return "division by zero!"
	case num2 == 0:
		return "division by zero!"
	}
	var result float64
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		return "Unsupported operator: " + operator
	}
	return strconv.FormatFloat(result, 'f', -1, 64)
}

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func processFactorial(input string) (string, error) {
	re := regexp.MustCompile(`(\d+)!`)
	matches := re.FindStringSubmatch(input)

	if len(matches) > 0 {
		numberStr := matches[1]
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			return "", fmt.Errorf("error converting string to number: %v", err)
		}
		result := factorial(number)
		return fmt.Sprintf("%d", result), nil
	}
	return "", fmt.Errorf("no factorial expression found")
}
