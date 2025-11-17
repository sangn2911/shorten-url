package model

type UrlEncodeModel struct {
	ShortenId     string `json:"shortenId"`     // DB column: shorten_id
	ShortenNumber int    `json:"shortenNumber"` // DB column: shorten_number
	OriginalUrl   string `json:"orginalUrl"`    // DB column: original_url
}
