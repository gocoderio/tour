package main

import "testing"

func TestAbs(t *testing.T) {
	v := Vertex{3, 4}
	expected := 5.0
	actual := Abs(v)

	if actual != expected {
		t.Errorf("Abs(%v) = %f; expected %f", v, actual, expected)
	}
}
