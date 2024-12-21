package day2

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Run() {
	fmt.Println("Advent of Code 2024, Day 1")
	fmt.Print("Part 1\n\n")
	part1()
	fmt.Print("Part 2\n\n")
	part2()

}

func part1() {

	content, err := os.Open("day2/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer content.Close()

	var count int = 0
	scanner := bufio.NewScanner(content)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		strlist := strings.Fields(line)
		intlist, err := utils.StringsToIntegers(strlist)
		if err != nil {
			fmt.Print(err)
		}
		if intlist[1] > intlist[0] {
			if isSafe, _ := isIncreasing(intlist); isSafe {
				count++
			}
		} else if intlist[1] < intlist[0] {
			if isSafe, _ := isDecreasing(intlist); isSafe {
				count++
			}
		}
	}
	fmt.Printf("total number of safe rows: %d \n", count)
}

func part2() {
	content, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer content.Close()

	var safeSlices [][]int
	scanner := bufio.NewScanner(content)
	scanner.Split(bufio.ScanLines)
	var scanline int = 0
	for scanner.Scan() {
		line := scanner.Text()
		strlist := strings.Fields(line)
		intlist, err := utils.StringsToIntegers(strlist)
		if err != nil {
			fmt.Println("Error converting line:", err)
			continue
		}
		scanline++

		fmt.Printf("Processing line %d: %v\n", scanline, intlist)

		// Track if the line has been added
		added := false

		// Check if already safe
		if ascSafeCheck(intlist) || descSafeCheck(intlist) {
			fmt.Printf("Line %d is already safe: %v\n", scanline, intlist)
			safeSlices = append(safeSlices, intlist)
			added = true
		}

		// Check for removal-based safety
		if !added && removeAndCheckAsc(intlist) {
			fmt.Printf("Line %d is safe after removal for ascending: %v\n", scanline, intlist)
			safeSlices = append(safeSlices, intlist)
			added = true
		}
		if !added && removeAndCheckDesc(intlist) {
			fmt.Printf("Line %d is safe after removal for descending: %v\n", scanline, intlist)
			safeSlices = append(safeSlices, intlist)
		}
	}
	fmt.Printf("Total number of safe lines: %d\n", len(safeSlices))
}

func isIncreasing(list []int) (bool, int) {
	for i := 1; i < len(list); i++ {
		if 0 >= list[i]-list[i-1] || list[i]-list[i-1] > 3 {
			return false, i
		}
	}
	return true, 0
}

func isDecreasing(list []int) (bool, int) {
	for i := 1; i < len(list); i++ {
		if 0 <= list[i]-list[i-1] || list[i]-list[i-1] < -3 {
			return false, i
		}
	}
	return true, 0
}

func removeError(list []int, index int) []int {
	// Create a new slice to avoid modifying the original list
	newList := make([]int, len(list))
	copy(newList, list)

	// Remove the element at the given index
	return append(newList[:index], newList[index+1:]...)
}

func ascSafeCheck(list []int) bool {
	for i := 0; i < len(list)-1; i++ {
		//if the difference between two indexes are beyond the safe rules
		if list[i+1]-list[i] < 1 || list[i+1]-list[i] > 3 {
			return false
		}
	}
	return true
}

func descSafeCheck(list []int) bool {
	for i := 0; i < len(list)-1; i++ {
		if list[i+1]-list[i] >= 0 || list[i+1]-list[i] < -3 {
			return false
		}
	}
	return true
}

func removeAndCheckAsc(list []int) bool {
	for i := 0; i < len(list); i++ {
		newList := removeError(list, i)
		safe := ascSafeCheck(newList)
		fmt.Printf("NewList: %v safe: %v\n", newList, safe)
		if safe {
			return true
		}
	}
	return false
}

func removeAndCheckDesc(list []int) bool {
	for i := 0; i < len(list); i++ {
		newList := removeError(list, i)
		if safe := descSafeCheck(newList); safe {
			return true
		}
	}
	return false
}
