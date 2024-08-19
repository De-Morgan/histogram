package pipe

import (
	"maps"
	"sync"
	"testing"
	"time"
)

func TestHistogram(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "repeated_word",
			args: args{
				words: []string{"first", "line", "second", "line"},
			},
			want: map[string]int{
				"first":  1,
				"second": 1,
				"line":   2,
			},
		},
		{
			name: "unique_words",
			args: args{
				words: []string{"first", "line", "second"},
			},
			want: map[string]int{
				"first":  1,
				"second": 1,
				"line":   1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make(chan string, len(tt.args.words))
			var wg sync.WaitGroup
			wg.Add(1)
			go func() {
				defer wg.Done()
				for _, word := range tt.args.words {
					input <- word
				}
				close(input)
			}()
			wg.Wait()

			result := Histogram(input)

			select {
			case got := <-result:
				if !maps.Equal(tt.want, got) {
					t.Errorf("Histogram() = %v, want %v", got, tt.want)
				}
			case <-time.After(1 * time.Second):
				t.Error("time out")
			}

		})
	}
}
