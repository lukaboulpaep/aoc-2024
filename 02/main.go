package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ParseCSV(path string) ([][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var report [][]int
	for scanner.Scan() {
		line := scanner.Text()
		lineArray := strings.Split(line, " ")

		var numericArray []int

		for i := 0; i < len(lineArray); i++ {
			numericValue, err := strconv.Atoi(lineArray[i])
			if err != nil {
				return nil, err
			}
			numericArray = append(numericArray, numericValue)
		}

		report = append(report, numericArray)
	}

	return report, nil
}

func CalculateAmountOfSafeReports(report [][]int) int {
	amountOfSafeReports := 0

	for i := 0; i < len(report); i++ {
		line := report[i]
		isValidLine := false
		var previousNumber int

		for j := 0; j < len(line)-1; j++ {
			if j != 0 {
				previousNumber = line[j-1]
			}

			if line[j] < line[j+1] {
				if j > 0 && previousNumber > line[j] {
					isValidLine = false
					break
				}

				diff := line[j+1] - line[j]

				if diff > 3 || diff == 0 {
					isValidLine = false
					break
				}
			} else {
				if j > 0 && previousNumber < line[j] {
					isValidLine = false
					break
				}

				diff := line[j] - line[j+1]

				if diff > 3 || diff == 0 {
					isValidLine = false
					break
				}
			}

			isValidLine = true
		}

		if isValidLine {
			amountOfSafeReports++
		}
	}

	return amountOfSafeReports
}
