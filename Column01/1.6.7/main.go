package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// INPUTFILE defines input filename.
const INPUTFILE = "input.txt"

// OUTPUTFILE defines output filename.
const OUTPUTFILE = "output.txt"

// BUFSIZE defines read buffer size.
const BUFSIZE = 1024

// N is max+1 number.
const N = 10000000

func main() {
	log.Println("start.")

	bits := initBits()

	bits, err := read(bits)
	if err != nil {
		log.Fatal(err)
	}

	write(bits)

	log.Println("finish.")
}

func write(bits []byte) error {
	file, err := os.Create(OUTPUTFILE)
	if err != nil {
		return err
	}
	defer file.Close()

	for i, v := range bits {
		if v == 1 {
			file.Write(([]byte)(strconv.Itoa(i) + "\n"))
		}
	}

	return nil
}

func initBits() []byte {
	bits := make([]byte, N)
	return bits
}

func read(bits []byte) ([]byte, error) {
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
		v := string(line)
		i, _ := strconv.Atoi(v)
		if !validate(bits, i) {
			continue
		}
		bits[i] = 1
	}

	return bits, nil
}

func validate(bits []byte, i int) bool {
	if i < 0 {
		errorMsg(fmt.Sprintf("error : %v is minus.", i))
		return false
	}
	if i >= N {
		errorMsg(fmt.Sprintf("error : %v is too large.", i))
		return false
	}
	if bits[i] == 1 {
		errorMsg(fmt.Sprintf("error : %v is deplicate.", i))
		return false
	}
	return true
}

func errorMsg(msg string) {
	log.Println(msg)
}
