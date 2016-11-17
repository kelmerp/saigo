package wordcount

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
)

type wordCount struct {
	word  string
	count int
}

func (wc wordCount) String() string {
	return fmt.Sprintf("%d %s", wc.count, wc.word)
}

type byCount []wordCount

func (bc byCount) Len() int      { return len(bc) }
func (bc byCount) Swap(i, j int) { bc[i], bc[j] = bc[j], bc[i] }
func (bc byCount) Less(i, j int) bool {
	if bc[i].count < bc[j].count {
		return true
	}
	if bc[i].count > bc[j].count {
		return false
	}
	return bc[i].word < bc[j].word
}
func (bc byCount) print() {
	for _, item := range bc {
		fmt.Println(item)
	}
}

// GetFilename gets the file name from the command line input
func GetFilename() string {
	return os.Args[1]
}

// ReadFile reads in the file
func ReadFile(filename string) []byte {
	fileData, readErr := ioutil.ReadFile(filename)
	if readErr != nil {
		panic(readErr)
	}
	return fileData
}

// RemoveNonLetters replaces non-letter characters at the end of a word with a space
func RemoveNonLetters(words string) string {
	charactersToRemove := regexp.MustCompile(`([^a-zA-Z]\W)`)
	return charactersToRemove.ReplaceAllString(words, " ")
}

func wordsToMap(words string) map[string]int {
	// splits the string into a slice of words
	wordSlice := strings.Fields(words)

	// transform slice into a map for easier word counting
	var wordMap = make(map[string]int)

	// counts word occurence
	for _, word := range wordSlice {
		wordMap[strings.ToLower(word)]++
	}

	return wordMap
}

// Count takes a string and outputs the occurence of words in that string
func Count(words string) {
	wordMap := wordsToMap(words)

	// create a slice of wordCount in order to sort
	wordCounter := []wordCount{}
	for w, c := range wordMap {
		wordCounter = append(wordCounter, wordCount{word: w, count: c})
	}

	// create a byCount from wordCounter in order to sort
	bycount := byCount(wordCounter)

	// reverse sort based upon count
	sort.Sort(sort.Reverse(bycount))

	// print out the results
	bycount.print()
}
