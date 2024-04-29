package main

import (
	"testing"
)



func TestAdd(t *testing.T) {
	testCases := []struct {
		a, b, expected int
	}{
		{5, 7, 12},
		{10, -3, 7},
		{0, 0, 0},
		{-5, -5, -10},
	}

	for _, tc := range testCases {
		result := Add(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Add(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expected)
		}
	}
}