package service

import (
	"context"
	"errors"
	"fmt"
	"math"
	netUrl "net/url"
	"shorten-url/internal/model"
	"shorten-url/internal/port"
	"shorten-url/package/logger"
	"strings"
)

var (
	ErrShortenUrlNotFound = errors.New("shorten url is not found")
	ErrUrlEmpty           = errors.New("url is empty")
	ErrUrlInvalid         = errors.New("url is invalid")
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
	if err != nil {
		return repository.RollBack(ctx)
	}
	return repository.Commit(ctx)
}

func (s *service) Encode(ctx context.Context, url string) (shortenUrl string, err error) {
	logger := logger.WithContext(ctx)
	logger.Info("Service.Encode")
	if len(url) == 0 {
		return shortenUrl, ErrUrlEmpty
	}
	if !isValidURLWithHost(url, "") {
		return shortenUrl, ErrUrlInvalid
	}
	urlEncode, err := s.repository.GetUrlEncodeByOriginalUrl(ctx, url)
	if err != nil {
		return shortenUrl, err
	}
	logger.Info("Service.Encode GetUrlEncodeByOriginalUrl Successfully")
	if len(urlEncode.ShortenId) > 0 {
		return fmt.Sprintf(
			"http://%s/%s",
			s.publicDomain,
			urlEncode.ShortenId,
		), nil
	}
	ctx, err = s.repository.Begin(ctx)
	if err != nil {
		return shortenUrl, err
	}
	defer func() {
		if errTx := handleEndOfTransaction(ctx, s.repository, err); errTx != nil {
			err = errTx
		}
	}()
	var latestUrlEncode model.UrlEncodeModel
	latestUrlEncode, err = s.repository.GetLatestUrlEncode(ctx)
	if err != nil {
		return shortenUrl, err
	}
	logger.Info("Service.Encode GetLatestUrlEncode Successfully")
	var nextShortenNumber int
	currentShortenIdLength := s.minShortenIdLength
	if len(latestUrlEncode.ShortenId) > 0 {
		nextShortenNumber = latestUrlEncode.ShortenNumber + 1
		currentShortenIdLength = len(latestUrlEncode.ShortenId)
	}
	newUrlEncode := model.UrlEncodeModel{
		ShortenNumber: nextShortenNumber,
		ShortenId: generateShortenIdByShortenNumber(
			nextShortenNumber,
			currentShortenIdLength,
		),
		OriginalUrl: url,
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

func (s *service) Decode(ctx context.Context, url string) (originalUrl string, err error) {
	logger := logger.WithContext(ctx)
	logger.Info("Service.Decode")
	if len(url) == 0 {
		return originalUrl, ErrUrlEmpty
	}
	if !isValidURLWithHost(url, s.publicDomain) {
		return originalUrl, ErrUrlInvalid
	}
	parsedUrl, err := netUrl.Parse(url)
	if err != nil {
		return originalUrl, err
	}
	urlEncode, err := s.repository.GetUrlEncodeByShortenId(ctx, parsedUrl.Path[1:])
	if err != nil {
		return originalUrl, err
	}
	if len(urlEncode.OriginalUrl) == 0 {
		return originalUrl, ErrShortenUrlNotFound
	}
	return urlEncode.OriginalUrl, nil
}

func generateShortenIdByShortenNumber(shortenNumber int, length int) string {
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	baseNumber := len(charSet)
	if shortenNumber == int(math.Pow(float64(baseNumber), float64(length))) {
		length = length + 1
	}
	var builder strings.Builder
	for range length {
		builder.WriteByte(charSet[shortenNumber%baseNumber])
		shortenNumber /= baseNumber
	}
	return builder.String()
}

func isValidURLWithHost(s string, host string) bool {
	u, err := netUrl.ParseRequestURI(s)
	if err != nil {
		return false
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}
	if host != "" && u.Host != host {
		return false
	}
	return true
}
