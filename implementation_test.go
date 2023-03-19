package lab2

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestcalculatePolishNotation(c *C) {
	result, err := CalculatePolishNotation("+ 2 2")
	c.Assert(result, Equals, "4")

	result, err = CalculatePolishNotation("+ 5 * ^ 4 2 3")
	c.Assert(result, Equals, "53")

	result, err = CalculatePolishNotation("+ 5 * - 4 2 3")
	c.Assert(result, Equals, "10")

	result, err = CalculatePolishNotation(" +++ 12 +++")
	c.Assert(err, ErrorMatches, "Expression contains invalid characters")
}

func ExampleCalculatePolishNotation() {
	res, err := CalculatePolishNotation("+ 2 2")
	if err != nil {
		fmt.Println(err)
	} 
	fmt.Println(res)

	// Output:
	// 4
}
