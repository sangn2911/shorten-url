package service

import (
	"context"
	"errors"
	"fmt"
	"math"
	"regexp"
	"shorten-url/internal/adapter/repository"
	"shorten-url/package/database"
	"shorten-url/package/database/mysql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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
			name:          "Test shorten id when shorten number is 0 and length is 6",
			shortenNumber: 0,
			length:        6,
			want:          "aaaaaa",
		},
		{
			name:          "Test shorten id when shorten number is 3843 and length is 2",
			shortenNumber: 3843,
			length:        2,
			want:          "99",
		},
		{
			name:          "Test avoiding collision when shorten number is 3844 and length is 2",
			shortenNumber: 3844,
			length:        2,
			want:          "aab",
		},
		{
			name:          "Test avoiding collision when shorten number is 62^6 and length is 6",
			shortenNumber: int(math.Pow(float64(62), 6)),
			length:        6,
			want:          "aaaaaab",
		},
		{
			name:          "Test generate shorten id GeAi9K",
			shortenNumber: 33884145296,
			length:        6,
			want:          "GeAi9K",
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

type sqlMockHandler func(mock sqlmock.Sqlmock)

func Test_handler_Encode(t *testing.T) {
	const (
		publicDomain       = "public_domain.com"
		orginalUrl         = "https://orginal_url.com/info"
		orginalUrl2        = "https://orginal_url.com/v2/info"
		minShortenIdLength = 2
		table_url_encode   = "url_encode"
	)
	txErr := errors.New("fail to commit transaction")
	columns := []string{
		"shorten_id",
		"shorten_number",
		"original_url",
	}
	tests := []struct {
		name           string
		mockHandler    sqlMockHandler
		wantShortenUrl string
		wantErr        error
	}{
		{
			name:           "Test encoding new url when database is empty",
			wantShortenUrl: fmt.Sprintf("http://%s/aa", publicDomain),
			wantErr:        nil,
			mockHandler: func(mock sqlmock.Sqlmock) {

				mock.
					ExpectQuery(regexp.QuoteMeta(fmt.Sprintf(
						"SELECT %s FROM %s WHERE original_url = ? LIMIT 1",
						database.JoinColumns(columns),
						table_url_encode,
					))).
					WithArgs(orginalUrl).
					WillReturnRows(&sqlmock.Rows{})
				mock.ExpectBegin()
				mock.
					ExpectQuery(regexp.QuoteMeta(fmt.Sprintf(
						"SELECT %s FROM %s ORDER BY shorten_number DESC LIMIT 1",
						database.JoinColumns(columns),
						table_url_encode,
					))).
					WithoutArgs().
					WillReturnRows(&sqlmock.Rows{})
				mock.
					ExpectExec(regexp.QuoteMeta(fmt.Sprintf(
						"INSERT INTO %s (%s) VALUES (%s)",
						table_url_encode,
						database.JoinColumns(columns),
						database.GetValueHolder(len(columns)),
					))).
					WithArgs("aa", 0, orginalUrl).
					WillReturnResult(sqlmock.NewResult(0, 0))
				mock.ExpectCommit()
			},
		},
		{
			name:           "Test encoding new url when database reach collision",
			wantShortenUrl: fmt.Sprintf("http://%s/aab", publicDomain),
			wantErr:        nil,
			mockHandler: func(mock sqlmock.Sqlmock) {

				collisionNumber := int(math.Pow(62, 2))
				table_url_encode := "url_encode"
				mock.
					ExpectQuery(regexp.QuoteMeta(fmt.Sprintf(
						"SELECT %s FROM %s WHERE original_url = ? LIMIT 1",
						database.JoinColumns(columns),
						table_url_encode,
					))).
					WithArgs(orginalUrl).
					WillReturnRows(&sqlmock.Rows{})
				mock.ExpectBegin()
				mock.
					ExpectQuery(regexp.QuoteMeta(fmt.Sprintf(
						"SELECT %s FROM %s ORDER BY shorten_number DESC LIMIT 1",
						database.JoinColumns(columns),
						table_url_encode,
					))).
					WithoutArgs().
					WillReturnRows(sqlmock.NewRows(columns).AddRow("99", collisionNumber-1, orginalUrl2))
				mock.
					ExpectExec(regexp.QuoteMeta(fmt.Sprintf(
						"INSERT INTO %s (%s) VALUES (%s)",
						table_url_encode,
						database.JoinColumns(columns),
						database.GetValueHolder(len(columns)),
					))).
					WithArgs("aab", collisionNumber, orginalUrl).
					WillReturnResult(sqlmock.NewResult(0, 0))
				mock.ExpectCommit()
			},
		},
		{
			name:           "Test encoding url encoded",
			wantShortenUrl: fmt.Sprintf("http://%s/99", publicDomain),
			wantErr:        nil,
			mockHandler: func(mock sqlmock.Sqlmock) {

				collisionNumber := int(math.Pow(62, 2))
				table_url_encode := "url_encode"
				mock.
					ExpectQuery(regexp.QuoteMeta(fmt.Sprintf(
						"SELECT %s FROM %s WHERE original_url = ? LIMIT 1",
						database.JoinColumns(columns),
						table_url_encode,
					))).
					WithArgs(orginalUrl).
					WillReturnRows(sqlmock.NewRows(columns).AddRow("99", collisionNumber-1, orginalUrl))
			},
		},
		{
			name:           "Test commit error",
			wantShortenUrl: fmt.Sprintf("http://%s/aa", publicDomain),
			wantErr:        txErr,
			mockHandler: func(mock sqlmock.Sqlmock) {

				table_url_encode := "url_encode"
				mock.
					ExpectQuery(regexp.QuoteMeta(fmt.Sprintf(
						"SELECT %s FROM %s WHERE original_url = ? LIMIT 1",
						database.JoinColumns(columns),
						table_url_encode,
					))).
					WithArgs(orginalUrl).
					WillReturnRows(&sqlmock.Rows{})
				mock.ExpectBegin()
				mock.
					ExpectQuery(regexp.QuoteMeta(fmt.Sprintf(
						"SELECT %s FROM %s ORDER BY shorten_number DESC LIMIT 1",
						database.JoinColumns(columns),
						table_url_encode,
					))).
					WithoutArgs().
					WillReturnRows(&sqlmock.Rows{})
				mock.
					ExpectExec(regexp.QuoteMeta(fmt.Sprintf(
						"INSERT INTO %s (%s) VALUES (%s)",
						table_url_encode,
						database.JoinColumns(columns),
						database.GetValueHolder(len(columns)),
					))).
					WithArgs("aa", 0, orginalUrl).
					WillReturnResult(sqlmock.NewResult(0, 0))
				mock.ExpectCommit().WillReturnError(txErr)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatal("error when create mock sql", err)
			}
			defer db.Close()
			tt.mockHandler(mock)
			repository := repository.NewRepository(&mysql.Client{DB: db})
			service := NewService(repository, publicDomain, minShortenIdLength)
			shortenUrl, err := service.Encode(context.Background(), orginalUrl)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Encode() return err %v, want %v", err, tt.wantErr)
			}
			if shortenUrl != tt.wantShortenUrl {
				t.Errorf("Encode() return shortenUrl %s, want %s", shortenUrl, tt.wantShortenUrl)
			}
		})
	}
}

func Test_handler_Decode(t *testing.T) {
	columns := []string{
		"shorten_id",
		"shorten_number",
		"original_url",
	}
	const (
		publicDomain       = "public_domain.com"
		orginalUrl         = "https://orginal_url.com/info"
		table_url_encode   = "url_encode"
		minShortenIdLength = 2
	)
	tests := []struct {
		name        string
		shortenUrl  string
		mockHandler sqlMockHandler
		wantUrl     string
		wantErr     error
	}{
		{
			name:       "Decode registered shorten url",
			shortenUrl: fmt.Sprintf("http://%s/GeAi9K", publicDomain),
			wantUrl:    orginalUrl,
			wantErr:    nil,
			mockHandler: func(mock sqlmock.Sqlmock) {
				mock.
					ExpectQuery(regexp.QuoteMeta(fmt.Sprintf(
						"SELECT %s FROM %s WHERE shorten_id = ? LIMIT 1",
						database.JoinColumns(columns),
						table_url_encode,
					))).
					WithArgs("GeAi9K").
					WillReturnRows(sqlmock.NewRows(columns).AddRow("GeAi9K", 33884145296, orginalUrl))
			},
		},
		{
			name:       "Decode unregistered shorten url",
			shortenUrl: fmt.Sprintf("http://%s/GeAi9K", publicDomain),
			wantUrl:    "",
			wantErr:    ErrShortenUrlNotFound,
			mockHandler: func(mock sqlmock.Sqlmock) {
				mock.
					ExpectQuery(regexp.QuoteMeta(fmt.Sprintf(
						"SELECT %s FROM %s WHERE shorten_id = ? LIMIT 1",
						database.JoinColumns(columns),
						table_url_encode,
					))).
					WithArgs("GeAi9K").
					WillReturnRows(&sqlmock.Rows{})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatal("error when create mock sql", err)
			}
			defer db.Close()
			tt.mockHandler(mock)
			repository := repository.NewRepository(&mysql.Client{DB: db})
			service := NewService(repository, publicDomain, minShortenIdLength)
			url, err := service.Decode(context.Background(), tt.shortenUrl)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Decode() return err %v, want %v", err, tt.wantErr)
			}
			if url != tt.wantUrl {
				t.Errorf("Decode() return shortenUrl %s, want %s", url, tt.wantUrl)
			}
		})
	}
}
