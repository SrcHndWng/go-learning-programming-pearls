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

// OUTPUTFILE defines output filename.
const OUTPUTFILE = "output.txt"

type zipNum struct {
	Key   int
	Value string
}

type zipNums []zipNum

func (z zipNums) Len() int {
	return len(z)
}

func (z zipNums) Swap(i, j int) {
	z[i], z[j] = z[j], z[i]
}

func (z zipNums) Less(i, j int) bool {
	return z[i].Key < z[j].Key
}

func main() {
	log.Println("start.")
	z, err := read()
	if err != nil {
		log.Fatal(err)
	}
	sort.Sort(z)
	err = write(z)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("finish.")
}

func write(z zipNums) error {
	file, err := os.Create(OUTPUTFILE)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, d := range z {
		file.Write(([]byte)(d.Value + "\n"))
	}
	return nil
}

func read() (zipNums, error) {
	var result zipNums
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
			log.Fatal(err)
		}
		v := string(line)
		k, _ := strconv.Atoi(v)
		z := zipNum{k, v}
		result = append(result, z)
	}
	return result, nil
}
