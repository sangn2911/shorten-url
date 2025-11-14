package service

import (
	"context"
	"shorten-url/internal/port"
)

type service struct {
	repository port.Repository
}

func NewService(repository port.Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) Encode(context.Context, string) (string, error) {
	panic("unimplemented")
}

func (s *service) Decode(context.Context, string) (string, error) {
	panic("unimplemented")
}

// Use order number id to get shortenId string.
// From Shorten id string get int number of char (97 - 122) => (1 - 26)
// Example 1 -> a a a a a a -> 97 97 97 97 97 97 -> 1 1 1 1 1 1
// 1 - 26 a - z (97 - 122) diff 96
// 27 - 52 A - Z (65 - 90) diff 38
// 53 - 62 0 - 9 diff 53
