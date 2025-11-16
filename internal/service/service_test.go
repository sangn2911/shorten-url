package service

import (
	"math"
	"testing"
)

func Test_generateShortenIdByShortenNumber(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		shortenNumber int
		length        int
		want          string
	}{
		{
			name:          "TEST_SHORTEN_ID_WHEN_SHORT_NUMBER_0_AND_CURR_LENGTH_6",
			shortenNumber: 0,
			length:        6,
			want:          "aaaaaa",
		},
		{
			name:          "TEST_SHORTEN_ID_WHEN_SHORT_NUMBER_3843_AND_CURR_LENGTH_2",
			shortenNumber: 3843,
			length:        2,
			want:          "99",
		},
		{
			name:          "TEST_COLLISION_WHEN_SHORT_NUMBER_3844_AND_CURR_LENGTH_2",
			shortenNumber: 3844,
			length:        2,
			want:          "aab",
		},
		{
			name:          "TEST_COLLISION_WHEN_SHORT_NUMBER_62^6_AND_CURR_LENGTH_6",
			shortenNumber: int(math.Pow(float64(62), 6)),
			length:        6,
			want:          "aaaaaab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateShortenIdByShortenNumber(tt.shortenNumber, tt.length)
			if got != tt.want {
				t.Errorf("generateShortenIdByShortenNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
