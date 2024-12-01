package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, _ := ReadFile("input")

	l1 := make([]int, 0)
	l2 := make([]int, 0)

	for _, v := range input {
		l1 = append(l1, v[0])
		l2 = append(l2, v[1])
	}

	sort.Ints(l1)
	sort.Ints(l2)

	diff := 0

	for i := range l1 {
		diff += int(math.Abs(float64(l1[i] - l2[i])))
	}

	fmt.Println("Part one:", diff)

	similar := 0

	for _, v := range l1 {
		similar += v * Count(l2, v)
	}

	fmt.Println("Part two:", similar)
}

func Count(l2 []int, v int) int {
	count := 0
	for _, a := range l2 {
		if a == v {
			count++
		}
	}
	return count
}

func ReadFile(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		// Split line into string values
		strValues := strings.Fields(line)

		// Convert string values to integers
		rowInts := make([]int, len(strValues))
		for i, strVal := range strValues {
			num, err := strconv.Atoi(strVal)
			if err != nil {
				return nil, fmt.Errorf("error converting %s: %v", strVal, err)
			}
			rowInts[i] = num
		}

		result = append(result, rowInts)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
