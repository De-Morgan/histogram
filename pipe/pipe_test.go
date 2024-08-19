package pipe

import (
	"context"
	"slices"
	"testing"
)

func TestWordProducer(t *testing.T) {
	type args struct {
		ctx   context.Context
		lines []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "all_lowercase",
			args: args{
				ctx: context.Background(),
				lines: []string{
					"first line",
					"second line",
				},
			},
			want: []string{"first", "line", "second", "line"},
		},
		{
			name: "case_mixture",
			args: args{
				ctx: context.Background(),
				lines: []string{
					"First line",
					"second Line",
				},
			},
			want: []string{"first", "line", "second", "line"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := make([]string, 0, len(tt.want))
			producer := WordProducer(tt.args.ctx, tt.args.lines)
			for word := range producer {
				got = append(got, word)
			}
			if !slices.Equal(tt.want, got) {
				t.Errorf("WordProducer() = %v, want %v", got, tt.want)
			}
		})
	}
}
