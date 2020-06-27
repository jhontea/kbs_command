package receiver

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	testCases := []struct {
		name           string
		inputA         int
		inputB         int
		expectedOutput int
	}{
		{"success sum 1 + 2 = 3", 1, 2, 3},
		{"success sum 1000 + 2500 = 3500", 1000, 2500, 3500},
		{"success sum -1000 + 2500 = 1500", -1000, 2500, 1500},
	}

	device1 := NewDevice1Receiver()

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			device1.SetA(testCase.inputA)
			device1.SetB(testCase.inputB)

			actualOutput := device1.Sum()
			if actualOutput != testCase.expectedOutput {
				t.Errorf("got %v should be %v", actualOutput, testCase.expectedOutput)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	testCases := []struct {
		name           string
		inputA         int
		inputB         int
		expectedOutput int
	}{
		{"success sum 1 * 2 = 2", 1, 2, 2},
		{"success sum 1000 * 2500 = 2.500.000", 1000, 2500, 2500000},
		{"success sum -1000 * 2500 = -2.500.000", -1000, 2500, -2500000},
	}

	device1 := NewDevice1Receiver()

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			device1.SetA(testCase.inputA)
			device1.SetB(testCase.inputB)

			actualOutput := device1.Multiply()
			if actualOutput != testCase.expectedOutput {
				t.Errorf("got %v should be %v", actualOutput, testCase.expectedOutput)
			}
		})
	}
}

func TestPrime(t *testing.T) {
	testCases := []struct {
		name           string
		inputN         int
		expectedOutput []int
	}{
		{"fail n prime < 1", 0, []int{}},
		{"success first 4 prime", 4, []int{2, 3, 5, 7}},
		{"success first 10 prime", 10, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}},
	}

	device1 := NewDevice1Receiver()

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			device1.SetN(testCase.inputN)

			actualOutput := device1.Prime()
			if len(actualOutput) != len(testCase.expectedOutput) {
				t.Errorf("got %v should be %v", len(actualOutput), len(testCase.expectedOutput))
			}

			if len(actualOutput) > 0 && !reflect.DeepEqual(actualOutput, testCase.expectedOutput) {
				t.Errorf("got %v should be %v", actualOutput, testCase.expectedOutput)
			}
		})
	}
}

func TestFibonacci(t *testing.T) {
	testCases := []struct {
		name           string
		inputN         int
		expectedOutput []int
	}{
		{"fail n fibonacci < 1", 0, []int{}},
		{"success first 5 fibonacci", 5, []int{0, 1, 1, 2, 3}},
		{"success first 11 fibonacci", 11, []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}},
	}

	device1 := NewDevice1Receiver()

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			device1.SetN(testCase.inputN)

			actualOutput := device1.Fibonacci()
			if len(actualOutput) != len(testCase.expectedOutput) {
				t.Errorf("got %v should be %v", len(actualOutput), len(testCase.expectedOutput))
			}

			if len(actualOutput) > 0 && !reflect.DeepEqual(actualOutput, testCase.expectedOutput) {
				t.Errorf("got %v should be %v", actualOutput, testCase.expectedOutput)
			}
		})
	}
}
