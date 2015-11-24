package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func read_string(filename string) []byte {
	f, _ := os.Open(filename)
	defer f.Close()
	r := bufio.NewReader(f)
	x, _, _ := r.ReadLine()
	return x
}

func less(s []byte, i, j int) bool {
	k := 0

	for i+k < len(s) && j+k < len(s) && s[i+k] == s[j+k] {
		k++
	}

	if i+k < len(s) && j+k < len(s) {
		return s[i+k] < s[j+k]
	} else {
		return i > j
	}

}

func main() {

	string_file := os.Args[1]
	suffix_array_file := os.Args[2]

	s := read_string(string_file)

	fmt.Println(string_file)
	fmt.Println(suffix_array_file)

	f, _ := os.Open(suffix_array_file)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	x := scanner.Text()
	prev, _ := strconv.Atoi(x)

	for scanner.Scan() {
		x = scanner.Text()
		curr, _ := strconv.Atoi(x)

		if !less(s, prev, curr) {
			fmt.Println("Bad positions ", prev, curr)
		}

		prev = curr
	}
}
