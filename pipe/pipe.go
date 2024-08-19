package pipe

import (
	"context"
	"strings"
)

// wordProducer extracts each word in lines,
// convert it into lowercase and sent it to out.
func wordProducer(ctx context.Context, lines []string, out chan<- string) {
	for _, line := range lines {
		words := strings.Fields(line)
		for _, word := range words {
			select {
			case <-ctx.Done():
				return
			case out <- strings.ToLower(word):
			}
		}
	}
}

// WordProducer runs wordProducer in a goroutine.
func WordProducer(ctx context.Context, lines []string) <-chan string {
	out := make(chan string, len(lines))
	go func() {
		defer close(out)
		wordProducer(ctx, lines, out)
	}()
	return out
}
