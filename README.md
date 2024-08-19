# Word Histogram Using Go's Channel and Pipeline Pattern

This repository demonstrates how to compute a word histogram from a list of lines using Go's channel and pipes pattern. The channel and pipes pattern allows us to process data concurrently, making our program more efficient, especially with large datasets.

## Overview

The program reads a list of lines, splits them into words, counts the occurrences of each word, and then prints a histogram showing the frequency of each word. The entire process is implemented using Go's concurrency primitives: goroutines and channels.

## How It Works

1. **WordProducer**: A goroutine reads the input lines, splits them into words, and sends each word through a channel.
2. **Histogram**: Another goroutine reads words from the channel, counts the occurrence of each word and send them result through another channel.
3. **ResultFormatter**: Runs on the main goroutine and outputs the result to io.Writer.


## Implementation

Here's a simple implementation of the channel and pipes pattern to compute the word histogram.

### main.go

```
import (
	"context"
	"os"

	"github.com/demorgan/histogram/pipe"
)

func main() {
	ctx := context.TODO()

	sentence := []string{
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit",
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque ac dolor posuere",
		"Nunc molestie tempus ante, vitae viverra felis fermentum ut",
	}
	producer := pipe.WordProducer(ctx, sentence)
	histogram := pipe.Histogram(producer)
	pipe.ResultFormatter(os.Stdout, histogram)

}

```
