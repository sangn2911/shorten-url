package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"shorten-url/internal/adapter/handler/dto"
	"shorten-url/internal/service"
	serviceMock "shorten-url/internal/service/mock"
	"shorten-url/package/httputils"
	"testing"

	"github.com/stretchr/testify/mock"
)

type mockServiceHandler func(mockService *serviceMock.MockService)

func Test_handler_Encode(t *testing.T) {
	const (
		publicDomain = "public_domain.com"
		orginalUrl   = "https://orginal_url.com/info"
	)
	tests := []struct {
		name           string
		url            string
		mockHandler    mockServiceHandler
		wantCode       int
		wantShortenUrl string
		wantMessage    string
	}{
		{
			name:           "Test when service encodes successfully",
			url:            orginalUrl,
			wantCode:       200,
			wantShortenUrl: fmt.Sprintf("http://%s/GeAi9K", publicDomain),
			mockHandler: func(mockService *serviceMock.MockService) {
				mockService.
					On("Encode", mock.Anything, orginalUrl).
					Return(fmt.Sprintf("http://%s/GeAi9K", publicDomain), nil)
			},
		},
		{
			name:        "Test when service fail to encodes",
			url:         orginalUrl,
			wantCode:    500,
			wantMessage: "something is wrong",
			mockHandler: func(mockService *serviceMock.MockService) {
				mockService.
					On("Encode", mock.Anything, orginalUrl).
					Return("", errors.New("something is wrong"))
			},
		},
		{
			name:        "Test when service receives an empty url",
			url:         "",
			wantCode:    400,
			wantMessage: service.ErrUrlEmpty.Error(),
			mockHandler: func(mockService *serviceMock.MockService) {
				mockService.
					On("Encode", mock.Anything, "").
					Return("", service.ErrUrlEmpty)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(serviceMock.MockService)
			tt.mockHandler(mockService)
			handler := NewHandler(mockService)
			body, err := json.Marshal(dto.ConvertUrlRequest{Url: tt.url})
			if err != nil {
				t.Fatal("error when marshal request body", err.Error())
			}
			req := httptest.NewRequest(http.MethodPost, "/encode", bytes.NewReader(body))
			rec := httptest.NewRecorder()
			handler.Encode(rec, req)
			if rec.Code != tt.wantCode {
				t.Errorf("Encode() return code %d, want %d", rec.Code, tt.wantCode)
			}
			if tt.wantCode == http.StatusOK {
				var response dto.EncodeUrlResponse
				if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
					t.Fatal("error when unmarshal response body", err.Error())
				}
				if response.ShortenUrl != tt.wantShortenUrl {
					t.Errorf("Encode() return shortenUrl %s, want %s", response.ShortenUrl, tt.wantShortenUrl)
				}
			} else {
				var response httputils.Response
				if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
					t.Fatal("error when unmarshal response body", err.Error())
				}
				if response.Message != tt.wantMessage {
					t.Errorf("Encode() return shortenUrl %s, want %s", response.Message, tt.wantMessage)
				}
			}
		})
	}
}

func Test_handler_Decode(t *testing.T) {
	const (
		publicDomain = "public_domain.com"
		shortenUrl   = "http://public_domain.com/GeAi9K"
		orginalUrl   = "https://orginal_url.com/info"
	)
	tests := []struct {
		name            string
		url             string
		mockHandler     mockServiceHandler
		wantCode        int
		wantOriginalUrl string
		wantMessage     string
	}{
		{
			name:            "Test when service decodes successfully",
			url:             shortenUrl,
			wantCode:        200,
			wantOriginalUrl: orginalUrl,
			mockHandler: func(mockService *serviceMock.MockService) {
				mockService.
					On("Decode", mock.Anything, shortenUrl).
					Return(orginalUrl, nil)
			},
		},
		{
			name:        "Test when service fail to decodes",
			url:         shortenUrl,
			wantCode:    500,
			wantMessage: "something is wrong",
			mockHandler: func(mockService *serviceMock.MockService) {
				mockService.
					On("Decode", mock.Anything, shortenUrl).
					Return("", errors.New("something is wrong"))
			},
		},
		{
			name:        "Test when service receives an empty url",
			url:         "",
			wantCode:    400,
			wantMessage: service.ErrUrlEmpty.Error(),
			mockHandler: func(mockService *serviceMock.MockService) {
				mockService.
					On("Decode", mock.Anything, "").
					Return("", service.ErrUrlEmpty)
			},
		},
		{
			name:        "Test when service receives an unregistered url",
			url:         shortenUrl,
			wantCode:    404,
			wantMessage: service.ErrShortenUrlNotFound.Error(),
			mockHandler: func(mockService *serviceMock.MockService) {
				mockService.
					On("Decode", mock.Anything, shortenUrl).
					Return("", service.ErrShortenUrlNotFound)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(serviceMock.MockService)
			tt.mockHandler(mockService)
			handler := NewHandler(mockService)
			body, err := json.Marshal(dto.ConvertUrlRequest{Url: tt.url})
			if err != nil {
				t.Fatal("error when marshal request body", err.Error())
			}
			req := httptest.NewRequest(http.MethodPost, "/decode", bytes.NewReader(body))
			rec := httptest.NewRecorder()
			handler.Decode(rec, req)
			if rec.Code != tt.wantCode {
				t.Errorf("Decode() return code %d, want %d", rec.Code, tt.wantCode)
			}
			if tt.wantCode == http.StatusOK {
				var response dto.DecodeUrlResponse
				if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
					t.Fatal("error when unmarshal response body", err.Error())
				}
				if response.OriginalUrl != tt.wantOriginalUrl {
					t.Errorf("Decode() return shortenUrl %s, want %s", response.OriginalUrl, tt.wantOriginalUrl)
				}
			} else {
				var response httputils.Response
				if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
					t.Fatal("error when unmarshal response body", err.Error())
				}
				if response.Message != tt.wantMessage {
					t.Errorf("Decode() return shortenUrl %s, want %s", response.Message, tt.wantMessage)
				}
			}
		})
	}
}
