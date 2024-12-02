package main

import (
	"testing"
)

func TestCalculateAmountOfSafeReports(t *testing.T) {
	tests := []struct {
		name           string
		inputPath      string
		expectedOutput int
	}{
		{
			name:           "example input data should result to 2",
			inputPath:      "data/example_input.csv",
			expectedOutput: 2,
		},
		{
			name:           "aoc input data should result to 486",
			inputPath:      "data/input.csv",
			expectedOutput: 486,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			report, err := ParseCSV(tt.inputPath)
			if err != nil {
				t.Errorf("Error occurred when parsing the CSV file %v", err)
			}

			result := CalculateAmountOfSafeReports(report)

			if result != tt.expectedOutput {
				t.Errorf("Expected output to be %d, instead got %d", tt.expectedOutput, result)
			}
		})
	}
}
