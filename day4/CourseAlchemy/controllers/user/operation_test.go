package user

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

	t.Run("createUser", func(t *testing.T) {
		body := map[string]interface{}{
			"name":     "erwin",
			"email":    "asd@asd.asd",
			"password": "1234",
		}
		jsonbody, _ := json.Marshal(body)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(jsonbody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, CreateUser(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			result := make(map[string]interface{})
			json.Unmarshal([]byte(rec.Body.String()), &result)
			assert.Equal(t, "success", result["message"])
			fmt.Println(result["data"])
			createdId = int(result["data"].(map[string]interface{})["ID"].(float64))
		}
	})

	t.Run("getUser", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		id := createdId
		c.SetParamValues(strconv.Itoa(id))

		if assert.NoError(t, GetUser(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			result := make(map[string]interface{})
			json.Unmarshal([]byte(rec.Body.String()), &result)
			assert.Equal(t, "success", result["message"])
			fmt.Println(result["data"])
			assert.Equal(t, id, int(result["data"].(map[string]interface{})["ID"].(float64)))
		}
	})

	t.Run("showUser", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, ShowUser(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			result := make(map[string]interface{})
			json.Unmarshal([]byte(rec.Body.String()), &result)
			fmt.Println(result["data"])
			assert.Equal(t, "success", result["message"])
			// assert.Equal(t, "array", reflect.TypeOf(result["data"].([]interface{})).Kind())
		}
	})

	t.Run("editUser", func(t *testing.T) {
		body := map[string]interface{}{
			"name":     "tested",
			"email":    "a@a.a",
			"password": "1",
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

		if assert.NoError(t, UpdateUser(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			result := make(map[string]interface{})
			json.Unmarshal([]byte(rec.Body.String()), &result)
			assert.Equal(t, "success", result["message"])
			fmt.Println(result["data"])
			assert.Equal(t, id, int(result["data"].(map[string]interface{})["ID"].(float64)))
		}
	})

	t.Run("deleteUser", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		id := createdId
		c.SetParamValues(strconv.Itoa(id))

		if assert.NoError(t, DeleteUser(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			result := make(map[string]interface{})
			json.Unmarshal([]byte(rec.Body.String()), &result)
			assert.Equal(t, "success", result["message"])
			fmt.Println(result["data"])
			assert.Equal(t, id, int(result["data"].(map[string]interface{})["ID"].(float64)))
		}
	})
}
