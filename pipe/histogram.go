package pipe

// histogram counts the occurrence of each word in words.
func histogram(words <-chan string, out chan<- map[string]int) {
	result := make(map[string]int, len(words))
	for word := range words {
		result[word]++
	}
	out <- result
}

// Histogram runs histogram in a goroutine.
func Histogram(words <-chan string) <-chan map[string]int {
	out := make(chan map[string]int)
	go func() {
		defer close(out)
		histogram(words, out)
	}()
	return out
}
