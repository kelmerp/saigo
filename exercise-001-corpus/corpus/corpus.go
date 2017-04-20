package corpus

import (
	"bufio"
	"fmt"
	"os"
)

// WordCounter ...
type WordCounter struct {
	words   map[string]int
	counter int
}

// NewWordCounter ...
func NewWordCounter() *WordCounter {
	var w WordCounter
	w.words = make(map[string]int)
	return &w
}

func (w WordCounter) Len() int {
	return len(s)
}

func (w WordCounter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (w WordCounter) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// Analyze ...
func Analyze(file *os.File) {
	wc := NewWordCounter()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		wc.counter++
		wc.words[scanner.Text()]++
	}
	fmt.Println("There are", wc.counter, "words.")
	fmt.Println(wc.words)
}
