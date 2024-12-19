package utils

import (
	"fmt"
	"strconv"
)

func StringsToIntegers(list []string) ([]int, error) {
	ints := make([]int, len(list))
	for i := range list {
		num, err := strconv.Atoi(list[i])
		if err != nil {
			return nil, fmt.Errorf("failed to convert %q to int: %w", list[i], err)
		}
		ints[i] = num
	}
	return ints, nil
}
