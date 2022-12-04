package part1

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func getRange(rng []string) ([]int, error) {
	rangeStart, err := strconv.Atoi(rng[0])
	if err != nil {
		return nil, fmt.Errorf("failed to parse range start: %w", err)
	}

	rangeEnd, err := strconv.Atoi(rng[1])
	if err != nil {
		return nil, fmt.Errorf("failed to parse range end: %w", err)
	}

	return []int{rangeStart, rangeEnd}, nil
}

func Solution(scanner *bufio.Scanner) (int, error) {
	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		ranges := strings.Split(line, ",")

		firstRange, err := getRange(strings.Split(ranges[0], "-"))
		if err != nil {
			return 0, fmt.Errorf("failed to parse first range: %w", err)
		}

		secondRange, err := getRange(strings.Split(ranges[1], "-"))
		if err != nil {
			return 0, fmt.Errorf("failed to parse second range: %w", err)
		}

		if secondRange[0]-firstRange[0] >= 0 && secondRange[1]-firstRange[1] <= 0 {
			total++

			continue
		}

		if secondRange[0]-firstRange[0] <= 0 && secondRange[1]-firstRange[1] >= 0 {
			total++
		}

	}

	return total, nil
}
