package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	calculator "github.com/janPhil/calculator/calculator"
	"google.golang.org/grpc"
)

func main() {

	srv := grpc.NewServer()

	var calc calculatorServiceServer

	calculator.RegisterCalculatorServiceServer(srv, calc)
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("Could not listen to port 8888")
	}
	log.Fatal(srv.Serve(l))

}

type calculatorServiceServer struct{}

// Calculate function takes an argument of typ *calculator.Term.
// It computes the result of the operation and sends back a *calculator.Result
func (c calculatorServiceServer) Calculate(ctx context.Context, term *calculator.Term) (*calculator.Result, error) {

	left := term.GetLeft()
	right := term.GetRight()
	operator := term.GetOperator()

	if right == 0 && operator == "/" {
		err := errors.New("division through zero not allowed")
		return nil, err
	}

	fmt.Printf("Received request for %v %v %v\n", left, operator, right)

	response := calculator.Result{}

	switch operator {
	case "+":
		response.Result = left + right
		return &response, nil
	case "-":
		response.Result = left - right
		return &response, nil
	case "*":
		response.Result = left * right
		return &response, nil
	case "/":
		response.Result = left / right
		return &response, nil
	default:
		err := errors.New("Unsupported operation")
		return nil, err
	}

}
