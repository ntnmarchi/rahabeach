package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// TestReconstructItinerary tests the itinerary reconstruction endpoint
func TestReconstructItinerary(t *testing.T) {
	e := echo.New()
	tests := []struct {
		name     string
		input    [][]string
		want     []string
		wantCode int
	}{
		{
			name:     "basic itinerary",
			input:    [][]string{{"JFK", "LAX"}, {"LAX", "SFO"}},
			want:     []string{"JFK", "LAX", "SFO"},
			wantCode: http.StatusOK,
		},
		{
			name:     "disconnected itinerary",
			input:    [][]string{{"JFK", "LAX"}, {"SFO", "ORD"}},
			want:     nil, // depending on your logic, you might want to handle errors
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "empty input",
			input:    [][]string{},
			want:     nil,
			wantCode: http.StatusBadRequest,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup
			rec := httptest.NewRecorder()
			reqBody, _ := json.Marshal(map[string][][]string{"tickets": tc.input})
			req := httptest.NewRequest(http.MethodPost, "/itinerary", bytes.NewBuffer(reqBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			// Execution
			c := e.NewContext(req, rec)
			if assert.NoError(t, handlePostItinerary(c)) {
				assert.Equal(t, tc.wantCode, rec.Code)

				var got []string
				if err := json.Unmarshal(rec.Body.Bytes(), &got); err == nil {
					assert.Equal(t, tc.want, got)
				}
			}
		})
	}
}
