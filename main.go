package main

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
