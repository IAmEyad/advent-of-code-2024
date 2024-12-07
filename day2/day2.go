package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	var inputLines []string
	file_string := string(file)

	inputLines = strings.Split(file_string, "\n")

	safe := 0

	for _, i := range inputLines {
		fmt.Println(i)
		left := 0
		right := 1

		increasing := 0
		decreasing := 0
		split_line := strings.Fields(i)
		for right < len(split_line) {
			// determine if the difference is between 1-3
			l, _ := strconv.Atoi(string(split_line[left]))
			r, _ := strconv.Atoi(string(split_line[right]))
			diff := math.Abs(float64(l - r))

			// if it is we need to know if we are decerasing or increasing
			if diff <= 3 && diff > 0 {
				if l < r {
					increasing += 1

				} else if l > r {
					decreasing += 1

				}
			} else {
				increasing = decreasing
				break
			}
			// incremenet l and right
			left += 1
			right += 1
		}
		// check if we where only increasing or only decreasing, otherwise we aren't safe
		if increasing > 0 && decreasing == 0 {
			safe += 1
		} else if decreasing > 0 && increasing == 0 {
			safe += 1
		}

	}
	fmt.Println(safe)

}
