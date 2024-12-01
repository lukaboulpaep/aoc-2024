package main

import (
	"testing"
)

func TestCalculatePairwiseDistance(t *testing.T) {
	tests := []struct {
		name           string
		inputPath      string
		expectedOutput int
	}{
		{
			name:           "example input data should result to 11",
			inputPath:      "data/example_input.csv",
			expectedOutput: 11,
		},
		{
			name:           "aoc input data should result to 1506483",
			inputPath:      "data/input.csv",
			expectedOutput: 1506483,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			leftColumn, rightColumn, err := ParseCSV(tt.inputPath)
			if err != nil {
				t.Errorf("Error occurred when parsing the CSV file %v", err)
			}

			distance, err := CalculatePairwiseDistance(&leftColumn, &rightColumn)
			if err != nil {
				t.Errorf("Error occurred when calculating the distance %v", err)
			}

			if distance != tt.expectedOutput {
				t.Errorf("Expected output to be %d, instead got %d", tt.expectedOutput, distance)
			}
		})
	}
}
