package mymath

import (
	"testing"
)

func Div_Test(t *testing.T) {
	_, err := Div(1, 0)
	if err != StatusDivideZero {
		t.Error("Divide zero")
	}
	v, err := Div(10, 5)
	if v != 2 {
		t.Fatal("10/5 = 2 but ", v)
	}
}
