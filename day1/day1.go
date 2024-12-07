package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var inputLines []string

	for scanner.Scan() {
		inputLines = append(inputLines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var list1 []int
	var list2 []int

	for _, line := range inputLines {
		split_line := strings.Split(line, "   ")
		split_int1, err := strconv.Atoi(split_line[0])

		if err != nil {
			log.Fatal(err)
		}

		split_int2, err := strconv.Atoi(split_line[1])

		if err != nil {
			log.Fatal(err)
		}

		list1 = append(list1, split_int1)
		list2 = append(list2, split_int2)
	}

	slices.Sort(list1)
	slices.Sort(list2)
	distance := 0

	for a, _ := range list1 {

		fmt.Println(a)
		dif := math.Abs(float64(list1[a] - list2[a]))
		distance += int(dif)
	}

	fmt.Println(distance)

}
