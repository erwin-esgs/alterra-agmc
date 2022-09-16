package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestSubtestOperation(t *testing.T) {
	t.Run("login", func(t *testing.T) {
		e := echo.New()
		body := map[string]interface{}{
			"email":    "erwin",
			"password": "1234",
		}
		jsonbody, _ := json.Marshal(body)
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(jsonbody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, Login(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			result := make(map[string]interface{})
			json.Unmarshal([]byte(rec.Body.String()), &result)
			assert.Equal(t, "success", result["message"])
		}
	})
}
