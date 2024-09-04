package main

import "testing"

var tests = []struct {
	name string
	dividend float32
	divisor float32
	expectedResult float32
	isErr bool
} {
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		result, err := divide(tt.dividend, tt.divisor)
		if tt.isErr && err ==nil {
			t.Errorf("Expected when dividing %f by %f", tt.dividend, tt.divisor)
		} else if !tt.isErr && err != nil {
			t.Errorf("Expected no error when dividing %f by %f but got %s", tt.dividend, tt.divisor, err)
		}

		if result != tt.expectedResult {
			t.Errorf("Expected %f but got %f", tt.expectedResult, result)
		}
	}
}

// func TestDivide(t *testing.T) {
// 	_, err := divide(10.0, 1.0)

// 	if err != nil {
// 		t.Error("Division should work for values 10.0 and 1.0")
// 	}
// }

// func TestErrorDivide(t *testing.T) {
// 	_, err := divide(10.0, 0)

// 	if err == nil {
// 		t.Error("Division should not work when dividing by 0")
// 	}
// }