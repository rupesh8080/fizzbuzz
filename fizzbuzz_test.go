package fizzbuzz_test

import (
	"fmt"
	"testing"

	"github.com/rupesh8080/fizzbuzz"
)

func TestFizz(t *testing.T) {
	t.Log("starting TestFizz")
	_, ok := fizzbuzz.Fizzbuzz(1)
	if ok {
		t.Logf("The value %v shouldn't have returned true", 1)
		t.Fail()
	}
	result, ok := fizzbuzz.Fizzbuzz(3)
	if !ok {
		t.Logf("The value %v shouldn't have returned false", 3)
		t.FailNow()
	}
	if result != "fizz" {
		t.Log("Result should have been fizz")
		t.Fail()
	}
}

func ExampleFizzbuzz() {
	result, _ := fizzbuzz.Fizzbuzz(3)
	fmt.Println(result)
	//Output: fizz
}
