package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {}

func ParseCSV(path string) ([]int, []int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var leftColumn []int
	var rightColumn []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Split(line, "   ")

		if len(columns) != 2 {
			return nil, nil, fmt.Errorf("a row with more than 2 values was not expected")
		}

		leftValue, err := strconv.Atoi(columns[0])
		if err != nil {
			return nil, nil, err
		}
		leftColumn = append(leftColumn, leftValue)

		rightValue, err := strconv.Atoi(columns[1])
		if err != nil {
			return nil, nil, err
		}
		rightColumn = append(rightColumn, rightValue)
	}

	return leftColumn, rightColumn, nil
}

func CalculatePairwiseDistance(leftColumn, rightColumn *[]int) (int, error) {
	left := *leftColumn
	right := *rightColumn

	sort.Ints(left)
	sort.Ints(right)

	if len(left) != len(right) {
		return 0, errors.New("slices must be of the same length")
	}

	distance := 0
	for i := 0; i < len(left); i++ {
		distance += int(math.Abs(float64(left[i] - right[i])))
	}

	return distance, nil
}
