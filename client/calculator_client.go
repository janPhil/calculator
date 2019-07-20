/*
Implementation of the calculator client.
Calculator_client listens by default on port 8888.
WARNING the gRPC dial is Insecure (dont use in production)
*/
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

	// The command-line Arguments are checked. If it is less then 2 that means that the user didnt provide a term for calculation.
	// If it is more then two it might be that the user tried to use whitespaces which is not supported.
	if len(os.Args) < 2 {
		log.Fatal("Missing statement for calculation: <number><operation><number>")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		log.Fatal("Please dont use any spaces: <number><operation><number>")
		os.Exit(1)
	}

	input := os.Args[1]
	inputTerm := parseInput(input)

	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		log.Fatal("no connection to server.", err)
	}

	client := calculator.NewCalculatorServiceClient(conn)

	res, err := client.Calculate(context.Background(), inputTerm)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	fmt.Println(res.Result)

}

// parseInput takes the provided input and tries to parse it into a term which can be used for a calculation.
// The operator and both sides are put into a *calculator.Term.
// If a conversion is not possible a message is logged, otherwise the term is returned.
func parseInput(input string) *calculator.Term {

	var operator string
	var indexOperator int
	var parts [3]string

	supportedOperations := []rune{'+', '-', '*', '/'}

	for pos, char := range input {
		if isSupportedOperation(char, &supportedOperations) {
			if pos == 0 {
				continue
			}
			operator = string(char)
			indexOperator = pos
			break
		}
	}
	if operator == "" {
		log.Fatal("Operation not supported")
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

// isSupportedOperations receives a rune and the reference to a slice of runes representing the supported operations
// if the rune is found in the supported operations true is returned, otherwise false
func isSupportedOperation(x rune, supportedOperations *[]rune) bool {
	for _, n := range *supportedOperations {
		if x == n {
			return true
		}
	}
	return false
}
