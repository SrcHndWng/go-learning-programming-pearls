package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

// INPUTFILE defines input filename.
const INPUTFILE = "input.txt"

// BUFSIZE defines read buffer size.
const BUFSIZE = 1024

func main() {
	data, err := read()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("input = %v\n", data)

	result := reshuffle(data)
	log.Printf("result = %v\n", result)
}

func reshuffle(data [][]string) [][]string {
	rows := len(data)
	cols := len(data[0])

	var result [][]string
	for c := 0; c < cols; c++ {
		var line []string
		for r := 0; r < rows; r++ {
			line = append(line, data[r][c])
		}
		result = append(result, line)
	}
	return result
}

func read() ([][]string, error) {
	file, err := os.Open(INPUTFILE)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result [][]string
	reader := bufio.NewReaderSize(file, BUFSIZE)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		ary := strings.Split(string(line), ",")
		result = append(result, ary)
	}
	return result, nil
}
