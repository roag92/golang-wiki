package test

import (
	"os"
	"testing"
)

type TestState struct {
	Prev string
	Next string
	Test *testing.T
}

func NewTestState(t *testing.T) *TestState {
	os.Setenv("APP_ENV", "testing")

	fullPath, err := os.Getwd()

	if err != nil {
		t.Fatal(err)
	}

	currentState := TestState{Prev: fullPath, Next: "./../..", Test: t}

	currentState.chdir()

	prev := currentState.Prev
	next := currentState.Next

	currentState.Next = prev
	currentState.Prev = next

	return &currentState
}

func (currentState *TestState) ResetState() {
	currentState.chdir()
}

func (currentState *TestState) chdir() {
	err := os.Chdir(currentState.Next)

	if err != nil {
		currentState.Test.Fatal(err)
	}
}
