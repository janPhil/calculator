package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/janPhil/calculator/calculator"
	"google.golang.org/grpc"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Missing statement for calculation: <number><operation><number>")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		fmt.Println("Please dont use any spaces: <number><operation><number>")
		os.Exit(1)
	}

	input := os.Args[1]
	fmt.Println(input)
	inputTerm := parseInput(input)

	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "no connection to server. %v\n", err)
	}

	client := calculator.NewCalculatorServiceClient(conn)

	res, err := client.Calculate(context.Background(), inputTerm)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	fmt.Println(res.Result)

}

func parseInput(input string) *calculator.Term {

	var operator string
	var indexOperator int

	var parts [3]string

	for pos, char := range input {
		if char == '+' || char == '-' || char == '*' || char == '/' {
			if pos == 0 {
				continue
			}
			operator = string(char)
			indexOperator = pos
			break
		}
	}
	parts[0] = input[:indexOperator]
	left, errLeft := strconv.ParseFloat(parts[0], 32)
	if errLeft != nil {
		log.Fatal("cannot convert l: ", parts[0])
	}
	parts[2] = input[indexOperator+1:]
	parts[1] = operator
	right, errRight := strconv.ParseFloat(parts[2], 32)
	if errRight != nil {
		log.Fatal("cannot convert r: ", parts[2])
	}
	return &calculator.Term{Left: float32(left), Right: float32(right), Operator: operator}
}
