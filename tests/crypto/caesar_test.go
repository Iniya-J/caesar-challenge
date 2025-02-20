{insert imports as needed}
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
	testCases := []struct {
		name           string
		cryptedText    string
		places         int
		expectedText   string
		expectedSHA1   string
	}{
		{"Normal Case", "bcd", 1, "abc", "a9993e364706816aba3e25717850c26c9cd0d89d"},
		{"Edge Case - Empty String", "", 1, "", "5ba93c9db0cff93f52b521d7420e43f6eda2784f"},
		{"Edge Case - Spaces and Punctuation", "bcd efg.", 1, "abc def.", "03cfd743661f07975fa2f1220c5194cbaff48451"},
		{"Edge Case - Numbers", "bcd123", 1, "abc123", "6367c48dd193d56ea7b0baad25b19455e529f5ee"},
		{"Edge Case - Large Shift", "abc", 25, "bcd", "3e23e8160039594a33894f6564e1b1348bbd7a0088d42c4acb73eeaed59c009d"},
		{"Edge Case - Case Sensitivity", "BCD", 1, "abc", "a9993e364706816aba3e25717850c26c9cd0d89d"},
		{"Edge Case - Non-Alphabetic Characters", "bcd!@#", 1, "abc!@#", "3c01bdbb26f358bab27f267924aa2c9a03fcfdb8"},
		{"Edge Case - Negative Shift", "abc", -1, "bcd", "3e23e8160039594a33894f6564e1b1348bbd7a0088d42c4acb73eeaed59c009d"},
		{"Edge Case - Zero Shift", "abc", 0, "abc", "a9993e364706816aba3e25717850c26c9cd0d89d"},
		{"Edge Case - Non-ASCII Characters", "abc£€¥", 1, "abc£€¥", "9e0e6f57b0752f714f52464a72b3e8a50b93b3ba"},
		{"Edge Case - Very Long Text", strings.Repeat("a", 1000000), 1, strings.Repeat("z", 1000000), "5c8b2a63c65d20f0c16c1b682c1415a91e47b0df"},
		{"Edge Case - Special Characters", "abc\n", 1, "abc\n", "3c363836cf4e16666669a25da280a1865c2d2874"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := &writer.WriterAnswer{
				Response: &request.ChallengeResponse{
					CryptedText: tc.cryptedText,
					Places:      tc.places,
				},
			}
			Decrypt(w)
			r := w.Response.(*request.ChallengeResponse)
			if r.DecryptedText != tc.expectedText {
				t.Errorf("expected decrypted text to be %s, but got %s", tc.expectedText, r.DecryptedText)
			}
			if r.SummaryCrypto != tc.expectedSHA1 {
				t.Errorf("expected SHA1 hash to be %s, but got %s", tc.expectedSHA1, r.SummaryCrypto)
			}
		})
	}
}
```
{insert unit test code here}