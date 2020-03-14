package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alochym01/ftp-api/src"
	"github.com/stretchr/testify/assert"
)

func TestHomeController(t *testing.T) {
	t.Run("Home function w GET Method", func(t *testing.T) {
		r := src.InitRouter()
		// Create new request to "/", with options
		// Method
		// URL
		// request body
		req, _ := http.NewRequest("GET", "/", nil)

		// create recording response
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		// create response variable
		var response map[string]string
		// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to
		// Convert the JSON response to a map
		err := json.Unmarshal([]byte(res.Body.String()), &response)
		// compare an error with nil
		assert.Nil(t, err)

		// compare response status code
		assert.Equal(t, http.StatusOK, res.Code)

		// check key exist in response reply
		value, exists := response["mesg"]
		// compare a key is existed
		assert.True(t, exists)
		// compare value
		assert.Equal(t, "hello world", value)
	})
}
