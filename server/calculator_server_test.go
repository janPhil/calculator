package main

import (
	"context"
	"testing"

	"github.com/janPhil/calculator/calculator"
)

func TestAddition(t *testing.T) {

	s := calculatorServiceServer{}

	testTerm := calculator.Term{Left: 1, Right: 2, Operator: "+"}
	want := calculator.Result{Result: 3}

	resp, err := s.Calculate(context.Background(), &testTerm)
	if err != nil {
		t.Errorf("TestAddition got unexpected error")
	}
	if resp.Result != want.Result {
		t.Errorf("TestAddition = %v, wanted %v", resp.Result, want.Result)
	}
}

func TestSubtraction(t *testing.T) {

	s := calculatorServiceServer{}

	testTerm := calculator.Term{Left: 2, Right: 1, Operator: "-"}
	want := calculator.Result{Result: 1}

	resp, err := s.Calculate(context.Background(), &testTerm)
	if err != nil {
		t.Errorf("TestSubtraction got unexpected error")
	}
	if resp.Result != want.Result {
		t.Errorf("TestSubtraction = %v, wanted %v", resp.Result, want.Result)
	}
}

func TestMultiplication(t *testing.T) {

	s := calculatorServiceServer{}

	testTerm := calculator.Term{Left: 3, Right: 8, Operator: "*"}
	want := calculator.Result{Result: 24}

	resp, err := s.Calculate(context.Background(), &testTerm)
	if err != nil {
		t.Errorf("TestMultiplication got unexpected error")
	}
	if resp.Result != want.Result {
		t.Errorf("TestMultiplication = %v, wanted %v", resp.Result, want.Result)
	}
}

func TestDevision(t *testing.T) {

	s := calculatorServiceServer{}

	testTerm := calculator.Term{Left: 60, Right: 15, Operator: "/"}
	want := calculator.Result{Result: 4}

	resp, err := s.Calculate(context.Background(), &testTerm)
	if err != nil {
		t.Errorf("TestDevision got unexpected error")
	}
	if resp.Result != want.Result {
		t.Errorf("TestDevision = %v, wanted %v", resp.Result, want.Result)
	}
}
