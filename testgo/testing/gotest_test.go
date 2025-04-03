package main

import (
	"fmt"
	"reflect"
	"testing"
)

// the test are cached
// go test ./ -v TestCalculateValues -count -1

func TestEqualPlayer(t *testing.T) {
	expected := Player{
		name: "Mark",
		hp:   100,
	}
	have := Player{
		name: "Mark",
		hp:   100,
	}

	if !reflect.DeepEqual(expected, have) {
		t.Errorf("expected %+v have %+v ", expected, have)
	}
}

func TestCalculateValues(t *testing.T) {
	var (
		expected = 10
		a        = 5
		b        = 4
	)
	have := calculateValues(a, b)
	if have != expected {
		t.Errorf("expected %d but have %d", expected, have)
	}
}

func TestCalculateValuesSpecial(t *testing.T) {
	fmt.Println("special test case")
}
