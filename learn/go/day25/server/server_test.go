package server

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEchoHandler(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		body       string
		wantStatus int
		wantJSON   string
	}{
		{"ok", http.MethodPost, `{"message":"hi"}`, http.StatusOK, `{"message":"hi"}` + "\n"},
		{"bad method", http.MethodGet, "", http.StatusMethodNotAllowed, "method not allowed\n"},
		{"bad json", http.MethodPost, `{`, http.StatusBadRequest, "invalid json\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/echo", bytes.NewBufferString(tt.body))
			rec := httptest.NewRecorder()

			EchoHandler(rec, req)

			if rec.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
			if rec.Body.String() != tt.wantJSON {
				t.Errorf("body = %q, want %q", rec.Body.String(), tt.wantJSON)
			}
		})
	}
}
