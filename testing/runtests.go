package test

import (
	"testing"
)

func TestFunction(t *testing.T) {
	t.Errorf("")
}

// func TestFooer2(t *testing.T) {
// 	input := 3
// 	result := test.Fooer(3)
// 	t.Logf("The input was %d", input)
// 	if result != "Foo" {
// 		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
// 	}
// 	t.Fatalf("Stop the test now, we have seen enough")
// 	t.Error("This won't be executed")
// }
