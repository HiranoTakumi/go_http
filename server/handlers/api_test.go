package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func Test_ListApi(t *testing.T) {
	t.Run("sample_test", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/list", nil)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req, rec)

		assert.NoError(t, ListApi(c))
		assert.Equal(t, http.StatusOK, rec.Code)

		var res Persons
		assert.NoError(t, json.NewDecoder(rec.Body).Decode(&res))
		assert.Len(t, res.Persons, 0)
		expect := Persons{}
		expect.Persons = []Person{}
		assert.Equal(t, expect, res)
		// assert.Equal(t, http.StatusBadRequest, res.Code)
		// assert.Equal(t, `{"admin":true,"name":"John Doe","sub":"1234567890"}`+"\n", res.Body.String())
	})

}
