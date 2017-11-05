package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
)

// OUTPUTFILE defines output filename.
const OUTPUTFILE = "output.txt"

// N is max+1 number.
const N = 10000000

// COUNT is create number count.
const COUNT = 1000000

func main() {
	log.Println("start.")

	file, err := os.Create(OUTPUTFILE)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for i := 0; i < COUNT; i++ {
		r := rand.Intn(N - 1)
		file.Write(([]byte)(strconv.Itoa(r) + "\n"))
	}

	log.Println("stop.")
}
