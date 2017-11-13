package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"strings"
)

// INPUTFILE defines input filename.
const INPUTFILE = "input.txt"

// BUFSIZE defines read buffer size.
const BUFSIZE = 1024

type rule struct {
	key        string
	separation string
}

func main() {
	flag.Parse()
	word := flag.Args()[0]
	log.Printf("word = %s\n", word)

	rules, err := read()
	if err != nil {
		log.Fatal(err)
	}

	result := separate(word, rules)
	log.Printf("result = %s\n", result)
}

func separate(word string, rules []rule) string {
	for _, r := range rules {
		if strings.Index(word, r.key) > 0 {
			return strings.Replace(word, r.key, r.separation, -1)
		}
	}
	return word
}

func read() ([]rule, error) {
	file, err := os.Open(INPUTFILE)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []rule
	reader := bufio.NewReaderSize(file, BUFSIZE)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		separation := string(line)
		r := rule{strings.Replace(separation, "-", "", -1), separation}
		result = append(result, r)
	}
	return result, nil
}
