package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

// INPUTFILE defines input filename.
const INPUTFILE = "input.txt"

// BUFSIZE defines read buffer size.
const BUFSIZE = 1024

func main() {
	nums, err := read()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("nums = %v\n", nums)
	sort.Sort(sort.IntSlice(nums))
	log.Printf("sorted = %v\n", nums)
	target := 157
	i, v := binarySearch(target, nums)
	log.Printf("i = %v, v = %v\n", i, v)
	target = 9999
	i, v = binarySearch(target, nums)
	log.Printf("i = %v, v = %v\n", i, v)
}

func read() ([]int, error) {
	file, err := os.Open(INPUTFILE)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []int
	reader := bufio.NewReaderSize(file, BUFSIZE)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		i, _ := strconv.Atoi(string(line))
		result = append(result, i)
	}
	return result, nil
}
