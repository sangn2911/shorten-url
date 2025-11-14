package model

const SHORTEN_URL_MIN_LENGTH = 6

type ConvertUrlRequest struct {
	Url string `json:"url"`
}

type DecodeUrlResponse struct {
	OriginalUrl string `json:"originalUrl"`
}

type EncodeUrlResponse struct {
	ShortenUrl string `json:"shortenUrl"`
}
