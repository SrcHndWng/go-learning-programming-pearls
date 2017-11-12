package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

// INPUTFILE defines input filename.
const INPUTFILE = "input.txt"

// BUFSIZE defines read buffer size.
const BUFSIZE = 1024

func main() {
	flag.Parse()
	pushed := flag.Args()[0]
	log.Printf("pushed dials = %s\n", pushed)

	names, err := read()
	if err != nil {
		log.Fatal(err)
	}

	sort.Sort(names)

	result := searchName(pushed, names, 0, len(names)-1)
	log.Printf("result = %v\n", result)
}

func read() (personNames, error) {
	var result personNames
	file, err := os.Open(INPUTFILE)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReaderSize(file, BUFSIZE)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		n := strings.Split(string(line), " ")
		p := personName{n[0], n[1], nameToNum(n[0], n[1])}
		result = append(result, p)
	}
	return result, nil
}
