package string_sum

import (
	"errors"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	var op1, op2 string
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return output, errorEmptyInput
	}
	nums := strings.Split(input, "+")
	if len(nums) != 2 {
		nums := strings.Split(input, "-")
		if len(nums) < 2 || len(nums) > 3 {
			return output, errorNotTwoOperands
		}
		if len(nums) == 2 {
			op1 = strings.TrimSpace(nums[0])
			op2 = "-" + strings.TrimSpace(nums[1])
		} else if len(nums) == 3 {
			op3 := strings.TrimSpace(nums[0])
			if len(op3) == 0 {
				op1 = "-" + strings.TrimSpace(nums[1])
				op2 = "-" + strings.TrimSpace(nums[2])
			}
		} else {
			return output, errorNotTwoOperands
		}
	} else {
		op1 = strings.TrimSpace(nums[0])
		op2 = strings.TrimSpace(nums[1])
	}
	num1, err := strconv.ParseInt(op1, 10, 32)
	if err != nil {
		return output, err
	}
	num2, err := strconv.ParseInt(op2, 10, 32)
	if err != nil {
		return output, err
	}
	sum := num1 + num2
	sumOfNums := strconv.Itoa(int(sum))
	return sumOfNums, nil
}
