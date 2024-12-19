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
		ascDesc := checkAscendingOrDescending(intlist)
		if ascDesc {
			fmt.Printf("\n(1) Ascending List: %v\n", intlist)
			isSafe, index := isIncreasing(intlist)
			if !isSafe {
				errList := removeError(intlist, index)
				fmt.Printf("Ascending List with error: %v\n", intlist)
				fmt.Printf("Ascending List without error: %v\n", errList)
				if errorSafe, _ := isIncreasing(errList); errorSafe {
					fmt.Printf("(2) Ascending List: %v is safe\n", errList)
					count++
				}
			} else {
				fmt.Printf("(3) Ascending List: %v is safe\n", intlist)
				count++
			}
		}
		if !ascDesc {
			fmt.Printf("\n(1) Descending List: %v\n", intlist)
			isSafe, index := isDecreasing(intlist)
			if !isSafe {
				errList := removeError(intlist, index)
				fmt.Printf("Decreasing list with error: %v\n", intlist)
				fmt.Printf("Decreasing list without error: %v\n", errList)
				if errorSafe, _ := isDecreasing(errList); errorSafe {
					fmt.Printf("(2) Descending List: %v is safe\n", errList)
					count++
				}
			} else if isSafe {
				fmt.Printf("(3) Descending List: %v is safe\n", intlist)
				count++
			}
		}
	}
	fmt.Printf("total number of safe rows: %d \n", count)
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
	fmt.Printf("List: %v Remove index: %d\n", list, index)
	return append(list[:index], list[index+1:]...)
}

// Returns true for ascending, and false for not descending (which has to be descending)
func checkAscendingOrDescending(list []int) bool {
	sum := 0
	sum += list[1] - list[0]
	sum += list[2] - list[1]

	if sum > 0 {
		return true
	} else {
		return false
	}
}
