package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func Run() {

	fmt.Println("Day 3")
	fmt.Println("part 1")
	part1()

}

func part1() {
	re := regexp.MustCompile(`mul\((?P<g1>\d{1,3})\,(?P<g2>\d{1,3})\)`)

	content, err := os.Open("day3/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer content.Close()

	scanner := bufio.NewScanner(content)
	sum := 0

	for scanner.Scan() {
		text := scanner.Text()
		matches := re.FindAllSubmatch([]byte(text), -1)
		g1 := re.SubexpIndex("g1")
		g2 := re.SubexpIndex("g2")
		for i := range matches {
			fmt.Printf("g1: %q g2: %q\n", matches[i][g1], matches[i][g2])
			n1, _ := strconv.Atoi(string(matches[i][g1]))
			n2, _ := strconv.Atoi(string(matches[i][g2]))
			fmt.Printf("n1: %d n2: %d\n", n1, n2)
			sum += n1 * n2
		}
	}
	fmt.Printf("Final sum: %d", sum)
}
