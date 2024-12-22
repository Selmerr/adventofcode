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
	fmt.Println("part 2")
	part2()

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
			n1, _ := strconv.Atoi(string(matches[i][g1]))
			n2, _ := strconv.Atoi(string(matches[i][g2]))
			sum += n1 * n2
		}
	}
	fmt.Printf("Final sum: %d\n", sum)
}

func part2() {
	re := regexp.MustCompile(`(?:mul\((?P<g1>\d{1,3}),(?P<g2>\d{1,3})\)|(?P<action>do\(\)|don't\(\)))`)

	content, err := os.Open("day3/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer content.Close()

	scanner := bufio.NewScanner(content)
	var mulEnabled bool = true
	sum := 0
	actioncounter := 0

	for scanner.Scan() {
		text := scanner.Text()
		matches := re.FindAllSubmatch([]byte(text), -1)
		g1 := re.SubexpIndex("g1")
		g2 := re.SubexpIndex("g2")
		action := re.SubexpIndex("action")
		for _, match := range matches {
			//fmt.Printf("g1: %q g2: %q\n", matches[i][g1], matches[i][g2])
			//fmt.Printf("action: %q\n", matches[i][action])
			if mulEnabled {
				n1, _ := strconv.Atoi(string(match[g1]))
				n2, _ := strconv.Atoi(string(match[g2]))
				sum += n1 * n2
			}
			if string(match[action]) == "do()" {
				actioncounter++
				//fmt.Println("---------------MUL ENABLED------------------")
				mulEnabled = true
			}
			if string(match[action]) == "don't()" {
				actioncounter++
				//fmt.Println("---------------MUL DISABLED-----------------")
				mulEnabled = false
			}
			//fmt.Printf("sum is: %d\n Number of do's and don'ts: %d\n", sum, actioncounter)
		}
	}
	fmt.Printf("Final sum: %d\n", sum)
}
