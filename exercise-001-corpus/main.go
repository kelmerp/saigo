package main

import wc "github.com/saigo/exercise-001-corpus/wordcount"

func main() {
	// get data from the file
	fileData := wc.ReadFile(wc.GetFilename())

	// convert file data in bytes into strings
	words := string(fileData)

	// remove non-letter characters at the end of a word
	words = wc.RemoveNonLetters(words)

	// Count takes a string and outputs the occurence of words in that string
	wc.Count(words)
}
