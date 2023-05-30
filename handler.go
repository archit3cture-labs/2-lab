package lab2

import (
	"bytes"
	"io"
	"strings"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	bufRead := make([]byte, 128)
	_, err := ch.Input.ReadAll(bufRead)
	if err != nil {
		return err
	}
	bufRead = bytes.Trim(bufRead, "\x00")

	expression := string(bufRead)
	trimmed := strings.Trim(expression, " \n")
	res, err := CalculatePolishNotation(trimmed)
	if err != nil {
		return err
	}
	ch.Output.Write([]byte(res))
	return nil
}
