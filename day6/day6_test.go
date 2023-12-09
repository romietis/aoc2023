package main

import (
	"reflect"
	"testing"
)

func TestQuadraticRaceMinMax(t *testing.T) {
	var tests = []struct {
		name           string
		time           int
		distanceToBeat int
		expectedMin    float64
		expectedMax    float64
	}{
		{"7 and 9", 7, 9, 1.6972243622680054, 5.302775637731995},
		{"15 and 40", 15, 40, 3.4688711258507254, 11.531128874149275},
		{"30 and 200", 30, 200, 10, 20},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			gotMin, gotMax := getQuadraticRaceValues(testCase.time, testCase.distanceToBeat)
			if gotMin != testCase.expectedMin && gotMax != testCase.expectedMax {
				t.Errorf("Expected: %v and %v. Got: %v and %v \n", testCase.expectedMin, testCase.expectedMax, gotMin, gotMax)
			}
		})
	}
}

func TestNaturalMinMax(t *testing.T) {
	var tests = []struct {
		name        string
		inputMinVal float64
		inputMaxVal float64
		expectedMin int
		expectedMax int
	}{
		{"1.6972243622680054 and 5.302775637731995", 1.6972243622680054, 5.302775637731995, 2, 5},
		{"3.4688711258507254 and 11.531128874149275", 3.4688711258507254, 11.531128874149275, 4, 11},
		{"10, 20", 10, 20, 11, 19},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			gotMin, gotMax := getNaturalMinMax(testCase.inputMinVal, testCase.inputMaxVal)
			if gotMin != testCase.expectedMin && gotMax != testCase.expectedMax {
				t.Errorf("Expected %v and %v. Got: %v and %v \n", testCase.expectedMin, testCase.expectedMax, gotMax, gotMax)
			}
		})
	}
}

func TestWaysOfWinning(t *testing.T) {
	var tests = []struct {
		name     string
		inputMin int
		inputMax int
		expected []int
	}{
		{"2 and 5", 2, 5, []int{2, 3, 4, 5}},
		{"4 and 11", 4, 11, []int{4, 5, 6, 7, 8, 9, 10, 11}},
		{"11 19", 11, 19, []int{11, 12, 13, 14, 15, 16, 17, 18, 19}},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			gotWayOfWinning := waysOfWinning(testCase.inputMin, testCase.inputMax)
			if !reflect.DeepEqual(gotWayOfWinning, testCase.expected) {
				t.Errorf("Expected: %v. Got: %v", testCase.expected, gotWayOfWinning)
			}
		})
	}
}
