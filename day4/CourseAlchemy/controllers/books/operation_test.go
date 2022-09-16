package books

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestSubtestOperation(t *testing.T) {
	var createdId = 0

	t.Run("createBooks", func(t *testing.T) {
		body := map[string]interface{}{
			"name":      "erwin",
			"author":    "erwin",
			"publisher": "1234",
			"year":      1995,
		}
		jsonbody, _ := json.Marshal(body)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(jsonbody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, CreateBooks(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			result := make(map[string]interface{})
			json.Unmarshal([]byte(rec.Body.String()), &result)
			assert.Equal(t, "success", result["message"])
			fmt.Println(result["data"])
			createdId = int(result["data"].(map[string]interface{})["ID"].(float64))
		}
	})

	t.Run("getBook", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		id := createdId
		c.SetParamValues(strconv.Itoa(id))

		if assert.NoError(t, GetBooks(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			result := make(map[string]interface{})
			json.Unmarshal([]byte(rec.Body.String()), &result)
			assert.Equal(t, "success", result["message"])
			fmt.Println(result["data"])
			assert.Equal(t, id, int(result["data"].(map[string]interface{})["ID"].(float64)))
		}
	})

	t.Run("showBooks", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, ShowBooks(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			result := make(map[string]interface{})
			json.Unmarshal([]byte(rec.Body.String()), &result)
			fmt.Println(result["data"])
			assert.Equal(t, "success", result["message"])
			// assert.Equal(t, "array", reflect.TypeOf(result["data"].([]interface{})).Kind())
		}
	})

	t.Run("editBooks", func(t *testing.T) {
		body := map[string]interface{}{
			"name":      "tested",
			"author":    "tested",
			"publisher": "1234",
			"year":      1995,
		}
		jsonbody, _ := json.Marshal(body)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(jsonbody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		id := createdId
		c.SetParamValues(strconv.Itoa(id))

		if assert.NoError(t, UpdateBooks(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			result := make(map[string]interface{})
			json.Unmarshal([]byte(rec.Body.String()), &result)
			assert.Equal(t, "success", result["message"])
			fmt.Println(result["data"])
			assert.Equal(t, id, int(result["data"].(map[string]interface{})["ID"].(float64)))
		}
	})

	t.Run("deleteBooks", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		id := createdId
		c.SetParamValues(strconv.Itoa(id))

		if assert.NoError(t, DeleteBooks(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			result := make(map[string]interface{})
			json.Unmarshal([]byte(rec.Body.String()), &result)
			assert.Equal(t, "success", result["message"])
			fmt.Println(result["data"])
			assert.Equal(t, id, int(result["data"].(map[string]interface{})["ID"].(float64)))
		}
	})
}
