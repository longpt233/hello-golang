package main

import (
	"fmt"

	"testing"
)

type MyInterface interface {
	MyMethod() string
}

type MyStruct struct{}

func (m *MyStruct) MyMethod() string {
	return "Hello, World!"
}

type MyMockStruct struct{}

func (m *MyMockStruct) MyMethod() string {
	return "Mocked Method"
}

func TestMyMethod(t *testing.T) {

	var myInterface MyInterface

	myStruct := &MyStruct{}
	myInterface = myStruct
	result := myInterface.MyMethod()
	fmt.Println(result)

	myMockStruct := &MyMockStruct{}
	myInterface = myMockStruct
	result = myInterface.MyMethod()
	fmt.Println(result)

}

func main() {
	TestMyMethod(&testing.T{})
}
