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

func TestServerInfo(t *testing.T) {
	router := src.InitRouter()
	t.Run("Create Server Info", func(t *testing.T) {
		params := []byte(`domain=live-03-hcm.fcam.vn&port=1935`)
		req, _ := http.NewRequest("POST", "/server/create", bytes.NewBuffer(params))

		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		// compare response status code
		assert.Equal(t, http.StatusOK, res.Code)

		// create json variable
		var response map[string]string

		// Convert the JSON response to a map
		err := json.Unmarshal([]byte(res.Body.String()), &response)
		// compare an error with nil
		assert.Nil(t, err)
		// check key exist in response reply
		value, exists := response["result"]
		// compare a key is existed
		assert.True(t, exists)
		// compare value
		assert.Equal(t, "Domain is created", value)

	})
}
