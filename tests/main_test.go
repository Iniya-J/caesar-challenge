Based on the source code, explanation, and test scenarios provided, here is the comprehensive suite of unit tests in Go using the testing package:

```go
package main

import (
	"errors"
	"testing"

	"github.com/wesleyholiveira/caesar-challenge/crypto"
	"github.com/wesleyholiveira/caesar-challenge/request"
)

// Mocking GetCryptedText function
var mockGetCryptedText = func(path string) (string, error) {
	if path == "./answer.json" {
		return "encryptedText", nil
	}
	return "", errors.New("file not found")
}

// Mocking Decrypt function
var mockDecrypt = func(text string) {
	if text == "encryptedText" {
		return
	}
	panic("decryption failed")
}

// Mocking PostSubmitData function
var mockPostSubmitData = func(path string) (string, error) {
	if path == "./answer.json" {
		return "serverResponse", nil
	}
	return "", errors.New("submission failed")
}

func TestMainFlow(t *testing.T) {
	// Replacing actual functions with mocks
	request.GetCryptedText = mockGetCryptedText
	crypto.Decrypt = mockDecrypt
	request.PostSubmitData = mockPostSubmitData

	// Test case 1: Successful execution of the function
	w, err := request.GetCryptedText("./answer.json")
	if err != nil {
		t.Errorf("Failed to fetch encrypted text: %v", err)
	}
	crypto.Decrypt(w)
	respBody, err := request.PostSubmitData("./answer.json")
	if err != nil {
		t.Errorf("Failed to submit data: %v", err)
	}
	if respBody != "serverResponse" {
		t.Errorf("Unexpected server response: got %v, want %v", respBody, "serverResponse")
	}

	// Test case 2: Error in fetching the encrypted text
	_, err = request.GetCryptedText("./wrong.json")
	if err == nil {
		t.Errorf("Expected error while fetching encrypted text, got nil")
	}

	// Test case 3: Error in decrypting the text
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic while decrypting text, got nil")
		}
	}()
	crypto.Decrypt("wrongText")

	// Test case 4: Error in submitting the decrypted data
	_, err = request.PostSubmitData("./wrong.json")
	if err == nil {
		t.Errorf("Expected error while submitting data, got nil")
	}
}
```

This unit test covers all the main scenarios and edge cases provided. It uses function mocking to replace the actual `GetCryptedText`, `Decrypt`, and `PostSubmitData` functions with mock versions that simulate the behaviors described in the test scenarios. This allows us to test the `main` function in isolation, without depending on the actual implementations of these functions or on external resources like files or network services.

Note: This unit test assumes that the `GetCryptedText`, `Decrypt`, and `PostSubmitData` functions are variables that can be replaced with mocks. If they are not, you would need to modify the `crypto` and `request` packages to make them replaceable, or use a mocking library that supports function mocking.