package pipe

import (
	"fmt"
	"io"
	"text/tabwriter"
)

// ResultFormatter prints out the result to out
func ResultFormatter(out io.Writer, histogram <-chan map[string]int) error {
	const format = "%s\t => \t%d\n"
	w := tabwriter.NewWriter(out, 1, 8, 2, ' ', 0)
	for word, count := range <-histogram {
		fmt.Fprintf(w, format, word, count)
	}
	return w.Flush()
}
