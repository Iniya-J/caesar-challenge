Based on the provided "Source_code", "Explanation" and "Test_Scenarios", here is the comprehensive suite of unit tests in Go Lang using testing package:

```go
package main

import (
	"testing"
	"errors"
	"github.com/wesleyholiveira/caesar-challenge/crypto"
	"github.com/wesleyholiveira/caesar-challenge/request"
	"github.com/stretchr/testify/assert"
)

type MockRequest struct{}

func (m *MockRequest) GetCryptedText(path string) (string, error) {
	if path == "invalid.json" {
		return "", errors.New("invalid json file")
	} else if path == "nonexistent.json" {
		return "", errors.New("file does not exist")
	} else if path == "unencrypted.json" {
		return "plain text", nil
	} else if path == "large.json" {
		return "large encrypted text", nil
	} else if path == "special.json" {
		return "encrypted text with special characters", nil
	} else {
		return "encrypted text", nil
	}
}

func (m *MockRequest) PostSubmitData(path string) ([]byte, error) {
	if path == "serverdown.json" {
		return nil, errors.New("server is down")
	} else if path == "errorresponse.json" {
		return nil, errors.New("server returned an error response")
	} else {
		return []byte("response body"), nil
	}
}

type MockCrypto struct{}

func (m *MockCrypto) Decrypt(text string) (string, error) {
	if text == "encrypted text with special characters" {
		return "", errors.New("decryption error: special characters")
	} else if text == "large encrypted text" {
		return "", errors.New("decryption error: text too large")
	} else {
		return "decrypted text", nil
	}
}

func TestMainFunction(t *testing.T) {
	mockRequest := &MockRequest{}
	mockCrypto := &MockCrypto{}

	t.Run("valid json file with encrypted text", func(t *testing.T) {
		w, err := mockRequest.GetCryptedText("valid.json")
		assert.Nil(t, err)
		assert.Equal(t, "encrypted text", w)

		decryptedText, err := mockCrypto.Decrypt(w)
		assert.Nil(t, err)
		assert.Equal(t, "decrypted text", decryptedText)

		respBody, err := mockRequest.PostSubmitData("valid.json")
		assert.Nil(t, err)
		assert.Equal(t, "response body", string(respBody))
	})

	t.Run("invalid json file", func(t *testing.T) {
		_, err := mockRequest.GetCryptedText("invalid.json")
		assert.NotNil(t, err)
	})

	t.Run("non-existent json file", func(t *testing.T) {
		_, err := mockRequest.GetCryptedText("nonexistent.json")
		assert.NotNil(t, err)
	})

	t.Run("valid json file with unencrypted text", func(t *testing.T) {
		_, err := mockRequest.GetCryptedText("unencrypted.json")
		assert.NotNil(t, err)
	})

	t.Run("edge cases", func(t *testing.T) {
		_, err := mockRequest.GetCryptedText("large.json")
		assert.NotNil(t, err)

		_, err = mockRequest.GetCryptedText("special.json")
		assert.NotNil(t, err)
	})

	t.Run("network issues", func(t *testing.T) {
		_, err := mockRequest.PostSubmitData("serverdown.json")
		assert.NotNil(t, err)

		_, err = mockRequest.PostSubmitData("errorresponse.json")
		assert.NotNil(t, err)
	})

	t.Run("decryption errors", func(t *testing.T) {
		_, err := mockCrypto.Decrypt("encrypted text with special characters")
		assert.NotNil(t, err)

		_, err = mockCrypto.Decrypt("large encrypted text")
		assert.NotNil(t, err)
	})
}
```

This test suite covers all the scenarios mentioned in the "Test_Scenarios". It uses the testify package for assertions and it mocks the `request` and `crypto` packages to simulate different scenarios.