package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alochym01/ftp-api/src"
	"github.com/stretchr/testify/assert"
)

type Response struct {
	Message string `form:"message" binding:"required"`
	Result  bool   `form:"result" binding:"required"`
	Ftp     string `form:"ftp" binding:"required"`
	Live    string `form:"live" binding:"required"`
	Port    int    `form:"port" binding:"required"`
}

type DeleteResponse struct {
	Message string `form:"message" binding:"required"`
	Result  bool   `form:"result" binding:"required"`
}

func TestAccount(t *testing.T) {
	router := src.InitRouter()

	t.Run("Create Account", func(t *testing.T) {
		// prepare a request
		params := []byte(`username=alochym1&password=password`)
		req, _ := http.NewRequest("POST", "/account/create", bytes.NewBuffer(params))

		// prepare request headers
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		// create recording response
		res := httptest.NewRecorder()

		router.ServeHTTP(res, req)

		// compare response status code
		assert.Equal(t, http.StatusOK, res.Code)

		// create json variable
		var response Response

		// Convert the JSON response to a map
		err := json.Unmarshal([]byte(res.Body.String()), &response)
		// compare an error with nil
		assert.Nil(t, err)
		// check key exist in response reply
		value := response.Message
		// compare a key is existed
		// assert.True(t, exists)
		// compare value
		assert.Equal(t, "Account is created", value)
	})

	t.Run("Check Account", func(t *testing.T) {
		// prepare a request
		params := []byte(`username=alochym1&password=password`)
		req, _ := http.NewRequest("POST", "/account/check", bytes.NewBuffer(params))

		// prepare request headers
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		// create recording response
		res := httptest.NewRecorder()

		router.ServeHTTP(res, req)

		// compare response status code
		assert.Equal(t, http.StatusOK, res.Code)

		// compare value
		assert.Equal(t, "OK", res.Body.String())
	})

	t.Run("Delete Account", func(t *testing.T) {
		// prepare a request
		params := []byte(`username=alochym1&password=password`)
		req, _ := http.NewRequest("POST", "/account/delete", bytes.NewBuffer(params))

		// prepare request headers
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		// create recording response
		res := httptest.NewRecorder()

		router.ServeHTTP(res, req)

		// compare response status code
		assert.Equal(t, http.StatusOK, res.Code)

		// create json variable
		var response DeleteResponse

		// // Convert the JSON response to a map
		err := json.Unmarshal([]byte(res.Body.String()), &response)
		// // compare an error with nil
		assert.Nil(t, err)
		// // check key exist in response reply
		value := response.Message
		// // compare a key is existed
		// assert.True(t, exists)
		// // compare value
		assert.Equal(t, "Account is deleted", value)

	})

}
