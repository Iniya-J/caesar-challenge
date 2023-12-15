package crypto

import (
	"crypto/sha1"
	"fmt"
	"strings"

	"github.com/wesleyholiveira/caesar-challenge/request"
	"github.com/wesleyholiveira/caesar-challenge/writer"
)

func Decrypt(w *writer.WriterAnswer) {
	r := w.Response.(*request.ChallengeResponse)
	h := sha1.New()
	places := r.Places
	decryptedBytes := make([]byte, 0, len(r.CryptedText))
	r.CryptedText = strings.ToLower(r.CryptedText)

	for _, char := range r.CryptedText {
		ascii := int(char)
		if (char != ' ' && char != '.') && (ascii < 48 || ascii > 57) {
			ascii -= places
		}
		decryptedBytes = append(decryptedBytes, byte(ascii))
	}

	h.Write(decryptedBytes)
	r.DecryptedText = string(decryptedBytes)
	r.SummaryCrypto = fmt.Sprintf("%x", h.Sum(nil))

	writer.WriteAnswer(w)
}
