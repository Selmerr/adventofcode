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
	fmt.Println("dag 2")

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
		if intlist[1] > intlist[0] && isIncreasing(intlist) {
			count++
		} else if intlist[1] < intlist[0] && isDecreasing(intlist) {
			count++
		}
	}
	fmt.Printf("total number of safe rows: %d \n", count)
}

func isIncreasing(list []int) bool {
	for i := 1; i < len(list); i++ {
		if 0 >= list[i]-list[i-1] || list[i]-list[i-1] > 3 {
			return false
		}
	}
	return true
}

func isDecreasing(list []int) bool {
	for i := 1; i < len(list); i++ {
		if 0 <= list[i]-list[i-1] || list[i]-list[i-1] < -3 {
			return false
		}
	}
	return true
}
