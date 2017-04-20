package main

import (
	"log"
	"os"

	"github.com/saigo/exercise-001-corpus/corpus"
)

func main() {
	// make sure a filename was passed in
	if len(os.Args) < 2 {
		log.Fatal("Wrong number of arguments, please call word_count with a text file.")

	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	corpus.Analyze(file)
}
