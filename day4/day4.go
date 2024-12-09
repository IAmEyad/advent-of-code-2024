package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	file_string := string(file)
	file_lines := strings.Split(file_string, "\n")

	//fmt.Println(file_lines)
	rows := len(file_lines)
	col := len(file_lines[0])
	//fmt.Println(file_lines[0])
	i_try_dfs(file_lines, col, rows)

}

func i_try_dfs(data []string, c int, r int) {
	count := 0
	for i := range 8 {
		for rows := range r {
			for col := range c {
				//fmt.Println(dfs(data, 0, col, rows))
				ret := dfs(data, 0, col, rows, i)

				count += ret
			}
		}
	}
	fmt.Println(count)

}

func dfs(data []string, index int, c int, r int, dir int) int {
	word := "XMAS"
	fmt.Println(dir)

	if index >= len(word) {
		//fmt.Println("Count went up by 1")
		return 1
	}
	if c >= len(data[0])-1 || r >= len(data) || c < 0 || r < 0 || string(data[r][c]) != string(word[index]) {
		return 0
	}

	// up
	if dir == 0 {
		return dfs(data, index+1, c, r-1, dir)
	}
	// down
	if dir == 1 {
		return dfs(data, index+1, c, r+1, dir)
	}
	// left
	if dir == 2 {
		return dfs(data, index+1, c-1, r, dir)
	}
	// right
	if dir == 3 {
		return dfs(data, index+1, c+1, r, dir)
	}
	// right and up
	if dir == 4 {
		return dfs(data, index+1, c+1, r-1, dir)
	}
	// left and up
	if dir == 5 {
		return dfs(data, index+1, c-1, r-1, dir)
	}
	// left and down
	if dir == 6 {
		return dfs(data, index+1, c-1, r+1, dir)
	} else {
		// right and down
		return dfs(data, index+1, c+1, r+1, dir)
	}
}
