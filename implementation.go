package lab2

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/apaxa-go/eval"
)

func splice(array []string, startIndex int, count int) []string {
	copy(array[startIndex:], array[startIndex + count:])
	array = array[:len(array) - count]
	return array
}

func CalculatePolishNotation(input string) (string, error) {
	inputNormalized := strings.Split(input, " ")

	for {

		for i := len(inputNormalized) - 1; i >= 0; i-- {

			if (inputNormalized[i] == "+") || (inputNormalized[i] == "-") || (inputNormalized[i] == "*") || (inputNormalized[i] == "/") {

				stringExpression := "float64(" + inputNormalized[i + 1] + inputNormalized[i] + inputNormalized[i + 2] + ")"
				expression, _ := eval.ParseString(stringExpression, "")
				result, err := expression.EvalToInterface(nil)
				if err != nil {
					return "Consider that expression must contain only numbers and operands separated by empty string", fmt.Errorf("Invalid expression")
				}

				finalResult := result.(float64)
				inputNormalized[i] = fmt.Sprint(finalResult)
				inputNormalized = splice(inputNormalized, i + 1, 2)
				break

			} else if inputNormalized[i] == "^" {

				val1, err := strconv.ParseFloat(inputNormalized[i + 1], 64)
				val2, err := strconv.ParseFloat(inputNormalized[i + 2], 64)
				if err != nil {
					return "Consider that expression must contain only numbers and operands separated by empty string", fmt.Errorf("Invalid expression")
				}

				res := math.Pow(val1, val2)
				inputNormalized[i] = fmt.Sprint(res)
				inputNormalized = splice(inputNormalized, i + 1, 2)
				break

			} else {

				_, err := strconv.ParseFloat(inputNormalized[i], 64)
				if err != nil {
					return inputNormalized[i] + " is not a number or operand\n",fmt.Errorf("Expression contains invalid characters")
				}
			}

		}

		if len(inputNormalized) == 1 {
			result := inputNormalized[0]
			return result, nil
		}

	}
}
