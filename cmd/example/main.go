package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	lab2 "github.com/archit3cture-labs/2-lab"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "File with expression to compute")
	outputFile      = flag.String("o", "", "File to store computed expression")
)

func main() {
	flag.Parse()

	if *inputExpression == "" && *inputFile == "" {
		log.Fatal("no expression provided. use -e \"{expression}\" or -f {file with expression}")
	}

	if *inputExpression != "" && *inputFile != "" {
		log.Fatal("flags -e and -f can't both be used")
	}

	var reader io.Reader

	if *inputExpression != "" {
		reader = strings.NewReader(*inputExpression)
	} else {
		file, err := os.Open(*inputFile)
		if err != nil {
			log.Fatal("no such file")
		}
		reader = file
		defer file.Close()
	}

	var writer io.Writer

	if *outputFile != "" {
		file, err := os.Create(*outputFile)
		if err != nil {
			log.Fatal("something went wrong while creating file")
		}

		writer = file
		defer file.Close()
	} else {
		writer = &Writer{}
	}

	handler := &lab2.ComputeHandler{
		Input:  reader,
		Output: writer,
	}

	err := handler.Compute()

	if err != nil {
		fmt.Errorf("invalid input")
	}

}

type Writer struct{}

func (w *Writer) Write(data []byte) (n int, err error) {
	fmt.Println(string(data))
	return len(data), nil
}
