package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"sort"

	"./utils"
)

// INPUTFILE defines input filename.
const INPUTFILE = "input.txt"

// BUFSIZE defines read buffer size.
const BUFSIZE = 1024

type word struct {
	Value  string
	Sorted string
}

type words []word

func (w words) Len() int {
	return len(w)
}

func (w words) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func (w words) Less(i, j int) bool {
	return w[i].Sorted < w[j].Sorted
}

func main() {
	log.Println("start.")

	words, err := read()
	if err != nil {
		log.Fatal(err)
	}

	sort.Sort(words)

	for i, w := range words {
		if i == (len(words) - 1) {
			continue
		}

		if (w.Sorted == words[i+1].Sorted) && (w.Value != words[i+1].Value) {
			log.Printf("anagram : %v, %v\n", w, words[i+1])
		}
	}

	log.Println("end.")
}

func read() (words, error) {
	file, err := os.Open(INPUTFILE)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result words
	reader := bufio.NewReaderSize(file, BUFSIZE)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		v := string(line)
		w := word{v, utils.SortStr(v)}
		result = append(result, w)
	}

	return result, nil
}
