package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMetadataHandler(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		remoteAddr     string
		userAgent      string
		xForwardedFor  string
		wantIP         string
		wantUserAgent  string
		wantForwarded  string
		wantStatusCode int
	}{
		{
			name:           "returns metadata from headers and remote addr",
			remoteAddr:     "203.0.113.10:45678",
			userAgent:      "curl/8.7.1",
			xForwardedFor:  "198.51.100.9",
			wantIP:         "203.0.113.10",
			wantUserAgent:  "curl/8.7.1",
			wantForwarded:  "198.51.100.9",
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "falls back to raw remote addr when split fails",
			remoteAddr:     "malformed-address",
			userAgent:      "test-agent",
			xForwardedFor:  "10.0.0.1",
			wantIP:         "malformed-address",
			wantUserAgent:  "test-agent",
			wantForwarded:  "10.0.0.1",
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "missing headers return empty strings",
			remoteAddr:     "192.0.2.1:9999",
			wantIP:         "192.0.2.1",
			wantUserAgent:  "",
			wantForwarded:  "",
			wantStatusCode: http.StatusOK,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.RemoteAddr = tc.remoteAddr
			if tc.userAgent != "" {
				req.Header.Set("User-Agent", tc.userAgent)
			}
			if tc.xForwardedFor != "" {
				req.Header.Set("X-Forwarded-For", tc.xForwardedFor)
			}

			recorder := httptest.NewRecorder()
			metadataHandler(recorder, req)

			if recorder.Code != tc.wantStatusCode {
				t.Fatalf("status code mismatch: got %d want %d", recorder.Code, tc.wantStatusCode)
			}

			contentType := recorder.Header().Get("Content-Type")
			if !strings.Contains(contentType, "application/json") {
				t.Fatalf("content-type mismatch: got %q", contentType)
			}

			var got response
			if err := json.Unmarshal(recorder.Body.Bytes(), &got); err != nil {
				t.Fatalf("failed to decode response body: %v", err)
			}

			if got.IP != tc.wantIP {
				t.Fatalf("ip mismatch: got %q want %q", got.IP, tc.wantIP)
			}
			if got.UserAgent != tc.wantUserAgent {
				t.Fatalf("user-agent mismatch: got %q want %q", got.UserAgent, tc.wantUserAgent)
			}
			if got.XForwardedFor != tc.wantForwarded {
				t.Fatalf("x-forwarded-for mismatch: got %q want %q", got.XForwardedFor, tc.wantForwarded)
			}
		})
	}
}

