package day1

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Run() {

	content, err := os.Open("day1/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer content.Close()

	scanner := bufio.NewScanner(content)
	scanner.Split(bufio.ScanLines)

	slice1 := []int{}
	slice2 := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		list := strings.Fields(line)
		i, _ := strconv.Atoi(list[0])
		j, _ := strconv.Atoi(list[1])
		slice1 = append(slice1, i)
		slice2 = append(slice2, j)
	}

	sort.Slice(slice1, func(i, j int) bool {
		return slice1[i] < slice1[j]
	})

	sort.Slice(slice2, func(i, j int) bool {
		return slice2[i] < slice2[j]
	})

	var sum float64 = 0

	var simscore int = 0

	for i := 0; i < len(slice1); i++ {
		sim := 0
		sum += math.Abs(float64(slice1[i] - slice2[i]))
		for j := 0; j < len(slice2); j++ {
			if slice1[i] == slice2[j] {
				sim += 1
			}
		}
		simscore += sim * slice1[i]
	}

	fmt.Printf("%.0f \n", sum)
	fmt.Printf("simscore: %d \n", simscore)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
