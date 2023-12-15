package main

import (
	"log"

	"github.com/wesleyholiveira/caesar-challenge/crypto"
	"github.com/wesleyholiveira/caesar-challenge/request"
)

func main() {
	w, err := request.GetCryptedText("./answer.json")

	if err != nil {
		log.Panicln(err)
	}

	crypto.Decrypt(w)

	respBody, err := request.PostSubmitData("./answer.json")
	if err != nil {
		log.Panicln(err)
	}

	log.Println(string(respBody))
}
