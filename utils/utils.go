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

func RemoveIndex[T comparable](slice []T, s int) []T {
	newList := make([]T, len(slice))
	copy(newList, slice)
	// Remove the element at the given index
	return append(newList[:s], newList[s+1:]...)
}
