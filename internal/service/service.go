package service

import (
	"context"
	"fmt"
	"shorten-url/internal/model"
	"shorten-url/internal/port"
	"shorten-url/package/logger"
	"strings"
)

type service struct {
	repository         port.Repository
	publicDomain       string
	minShortenIdLength int
}

func NewService(
	repository port.Repository,
	publicDomain string,
	minShortenIdLength int,
) *service {
	return &service{
		repository:         repository,
		publicDomain:       publicDomain,
		minShortenIdLength: minShortenIdLength,
	}
}

func handleEndOfTransaction(ctx context.Context, repository port.Repository, err error) error {
	if panicErr := recover(); panicErr != nil {
		repository.RollBack(ctx)
		panic(panicErr)
	}
	if err != nil {
		repository.RollBack(ctx)
	}
	return repository.Commit(ctx)
}

func (s *service) Encode(ctx context.Context, url string) (shortenUrl string, err error) {
	logger := logger.WithContext(ctx)
	logger.Info("Service.Encode")
	ctx, err = s.repository.Begin(ctx)
	if err != nil {
		return shortenUrl, err
	}
	defer func() {
		err = handleEndOfTransaction(ctx, s.repository, err)
	}()
	var latestUrlEncode model.UrlEncodeModel
	latestUrlEncode, err = s.repository.GetLatestUrlEncode(ctx)
	if err != nil {
		return shortenUrl, err
	}
	// TODO: check long url exists
	logger.Info("Service.Encode GetLatestUrlEncode Successfully")
	nextShortenNumber := latestUrlEncode.ShortenNumber + 1
	if len(latestUrlEncode.ShortenId) == 0 {
		nextShortenNumber = 0
	}
	newUrlEncode := model.UrlEncodeModel{
		ShortenNumber: nextShortenNumber,
		ShortenId: generateShortenIdByShortenNumber(
			nextShortenNumber,
			s.minShortenIdLength,
		),
		LongUrl: url,
	}
	if err = s.repository.InsertUrlEncode(ctx, newUrlEncode); err != nil {
		return shortenUrl, err
	}
	logger.Info("Service.Encode InsertUrlEncode Successfully")
	return fmt.Sprintf(
		"http://%s/%s",
		s.publicDomain,
		newUrlEncode.ShortenId,
	), nil
}

func (s *service) Decode(ctx context.Context, url string) (string, error) {
	logger := logger.WithContext(ctx)
	logger.Info("Service.Decode")
	panic("unimplemented")
}

func generateShortenIdByShortenNumber(shortenNumber int, length int) string {
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	baseNumber := len(charSet)
	var builder strings.Builder
	for range length {
		builder.WriteByte(charSet[shortenNumber%baseNumber])
		shortenNumber /= baseNumber
	}
	return builder.String()
}
