package model

type UrlEncodeModel struct {
	ShortenId     string `json:"shortenId"`     // DB column: shorten_id
	ShortenNumber int    `json:"shortenNumber"` // DB column: shorten_number
	LongUrl       string `json:"shortenUrl"`    // DB column: long_url
}
