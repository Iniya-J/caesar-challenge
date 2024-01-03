Here is the unit test code based on the given "Source_code", "Explanation" and "Test_Scenarios":

```go
package crypto

import (
	"crypto/sha1"
	"fmt"
	"strings"
	"testing"

	"github.com/wesleyholiveira/caesar-challenge/request"
	"github.com/wesleyholiveira/caesar-challenge/writer"
)

func TestDecrypt(t *testing.T) {
	tests := []struct {
		name          string
		cryptedText   string
		places        int
		expectedText  string
		expectedHash  string
	}{
		{"Normal text input", "bcdef", 1, "abcde", fmt.Sprintf("%x", sha1.Sum([]byte("abcde")))},
		{"Text with spaces and periods", "b c.d e.f", 1, "a b.c d.e", fmt.Sprintf("%x", sha1.Sum([]byte("a b.c d.e")))},
		{"Text with numbers", "b2c3d4", 1, "a2b3c4", fmt.Sprintf("%x", sha1.Sum([]byte("a2b3c4")))},
		{"Text with uppercase letters", "BCD", 1, "abc", fmt.Sprintf("%x", sha1.Sum([]byte("abc")))},
		{"Empty text", "", 1, "", fmt.Sprintf("%x", sha1.Sum([]byte("")))},
		{"Edge case with shift of 0", "abc", 0, "abc", fmt.Sprintf("%x", sha1.Sum([]byte("abc")))},
		{"Edge case with shift of 27", "abc", 27, "zab", fmt.Sprintf("%x", sha1.Sum([]byte("zab")))},
		{"Negative shift values", "abc", -1, "bcd", fmt.Sprintf("%x", sha1.Sum([]byte("bcd")))},
		{"Non-ASCII characters", "abç", 1, "zaç", fmt.Sprintf("%x", sha1.Sum([]byte("zaç")))},
		{"Special characters", "a@b#c$d", 1, "z@b#c$d", fmt.Sprintf("%x", sha1.Sum([]byte("z@b#c$d")))},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &request.ChallengeResponse{
				CryptedText: tt.cryptedText,
				Places:      tt.places,
			}
			w := &writer.WriterAnswer{
				Response: r,
			}

			Decrypt(w)

			if w.Response.(*request.ChallengeResponse).DecryptedText != tt.expectedText {
				t.Errorf("Expected decrypted text to be %s, but got %s", tt.expectedText, w.Response.(*request.ChallengeResponse).DecryptedText)
			}

			if w.Response.(*request.ChallengeResponse).SummaryCrypto != tt.expectedHash {
				t.Errorf("Expected summary crypto to be %s, but got %s", tt.expectedHash, w.Response.(*request.ChallengeResponse).SummaryCrypto)
			}
		})
	}
}

func TestDecryptWithVeryLongText(t *testing.T) {
	longText := strings.Repeat("b", 1000000)
	decryptedText := strings.Repeat("a", 1000000)

	r := &request.ChallengeResponse{
		CryptedText: longText,
		Places:      1,
	}
	w := &writer.WriterAnswer{
		Response: r,
	}

	Decrypt(w)

	if w.Response.(*request.ChallengeResponse).DecryptedText != decryptedText {
		t.Errorf("Expected decrypted text to be %s, but got %s", decryptedText, w.Response.(*request.ChallengeResponse).DecryptedText)
	}
}

func TestDecryptWithVeryLargeShiftValues(t *testing.T) {
	r := &request.ChallengeResponse{
		CryptedText: "abc",
		Places:      1000000,
	}
	w := &writer.WriterAnswer{
		Response: r,
	}

	Decrypt(w)

	if w.Response.(*request.ChallengeResponse).DecryptedText != "abc" {
		t.Errorf("Expected decrypted text to be abc, but got %s", w.Response.(*request.ChallengeResponse).DecryptedText)
	}
}

func TestDecryptWithNullInput(t *testing.T) {
	r := &request.ChallengeResponse{
		CryptedText: "",
		Places:      1,
	}
	w := &writer.WriterAnswer{
		Response: r,
	}

	Decrypt(w)

	if w.Response.(*request.ChallengeResponse).DecryptedText != "" {
		t.Errorf("Expected decrypted text to be empty, but got %s", w.Response.(*request.ChallengeResponse).DecryptedText)
	}
}
```

This test suite covers all the scenarios provided in the "Test_Scenarios". It tests the `Decrypt` function with normal text input, text with spaces and periods, text with numbers, text with uppercase letters, empty text, edge cases, negative shift values, non-ASCII characters, special characters, very long text, very large shift values, and null input. Each test checks that the decrypted text and the SHA-1 hash of the decrypted text are as expected.