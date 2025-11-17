package dto

import "errors"

var ErrUrlEmpty = errors.New("url is empty")

type ConvertUrlRequest struct {
	Url string `json:"url"`
}

type DecodeUrlResponse struct {
	OriginalUrl string `json:"originalUrl"`
}

type EncodeUrlResponse struct {
	ShortenUrl string `json:"shortenUrl"`
}
