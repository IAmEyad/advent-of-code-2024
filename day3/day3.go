package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	// I said I wasn't going to use regex but here I am
	r, _ := regexp.Compile(`(mul)+\(+[[:digit:]]+,[[:digit:]]+\)`)
	n, _ := regexp.Compile(`[[:digit:]]+,[[:digit:]]+`)

	//fmt.Println(r.FindAllString(string(file), -1))
	filtered_file := r.FindAllString(string(file), -1)

	result := 0
	// maybe?
	for _, i := range filtered_file {

		clean_i := n.FindAllString(string(i), -1)
		clean_i = strings.Split(clean_i[0], ",")

		in_1, _ := strconv.Atoi(clean_i[0])
		in_2, _ := strconv.Atoi(clean_i[1])
		result += (in_1 * in_2)
	}

	fmt.Println(result)

}
