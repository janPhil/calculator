/*
Implementation of the calculator server.
Calculator_server runs by default on port 8888.
It provides a method Calculate to calculate two
numbers. Supported operations are Addition, Subtraction,
Multiplication and Division.
*/
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/janPhil/calculator/calculator"
	"google.golang.org/grpc"
)

// The Server is communicating over gRPC Calls and provides a method to calculate two numbers.
func main() {

	srv := grpc.NewServer()

	var calc calculatorServiceServer

	calculator.RegisterCalculatorServiceServer(srv, calc)
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("Could not listen to port 8888", err)
	}
	fmt.Println("Server started")
	log.Fatal(srv.Serve(l))

}

type calculatorServiceServer struct{}

// Calculate takes in an argument of typ *calculator.Term.
// It computes the result of the operation and sends back a *calculator.Result.
// Should the operation be unsupported an error is thrown, also if there is an
// attempt to do a division through zero
func (c calculatorServiceServer) Calculate(ctx context.Context, term *calculator.Term) (*calculator.Result, error) {

	left := term.GetLeft()
	right := term.GetRight()
	operator := term.GetOperator()

	fmt.Printf("Received request for %v %v %v\n", left, operator, right)

	var response calculator.Result

	switch operator {
	case "+":
		response.Result = left + right
	case "-":
		response.Result = left - right
	case "*":
		response.Result = left * right
	case "/":
		if right == 0 {
			return nil, errors.New("division through zero not allowed")
		}
		response.Result = left / right
	default:
		return nil, errors.New("Unsupported operation")
	}

	return &response, nil

}
