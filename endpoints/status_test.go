package endpoints

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestStatusCheck(t *testing.T) {
	// Write test cases for endpoint
	tests := map[string]bool{
		"Success": false,
	}

	for name, wantErr := range tests {
		// Setup a test echo server
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		t.Run(name, func(t *testing.T) {
			err := StatusCheck(ctx)
			// assert the results will not be error
			assert.Equal(t, wantErr, err != nil)
		})
	}
}
