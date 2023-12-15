package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/wesleyholiveira/caesar-challenge/config"
	"github.com/wesleyholiveira/caesar-challenge/reader"
	"github.com/wesleyholiveira/caesar-challenge/writer"
)

type Requester interface {
	Get(url string) (*http.Response, error)
	Post(url, contentType string, body *bytes.Buffer) (*http.Response, error)
}

type FileHandler interface {
	Open(string) (*os.File, error)
	Stat(name string) (os.FileInfo, error)
}

// ChallengeResponse struct deals with the http response
type ChallengeResponse struct {
	Places        int    `json:"numero_casas"`
	Token         string `json:"token"`
	CryptedText   string `json:"cifrado"`
	DecryptedText string `json:"decifrado"`
	SummaryCrypto string `json:"resumo_criptografico"`
}

// GetCryptedText sends request to codenation and return a struct with the json parsed
func GetCryptedText(requester Requester, fileHandler FileHandler, file string) (*writer.WriterAnswer, error) {
	w := writer.New()
	url := fmt.Sprintf("%s?token=%s", config.GenerateUrl, config.TokenCodeNation)
	log.Printf("Making request to %s", url)

	response := &ChallengeResponse{}
	resp, err := requester.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	w.File = file
	w.Response = response
	w.Data = body
	writer.WriteAnswer(w, fileHandler)

	return w, nil
}

func PostSubmitData(requester Requester, fileHandler FileHandler, file string) ([]byte, error) {
	url := fmt.Sprintf("%s?token=%s", config.SubmitUrl, config.TokenCodeNation)
	log.Printf("Making request to %s", url)

	r, err := reader.ReadAnswer(fileHandler, file)
	if err != nil {
		return nil, err
	}

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("answer", r.Info.Name())
	if err != nil {
		return nil, err
	}

	if _, err := part.Write(r.Data); err != nil {
		return nil, err
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	fmt.Println(body.String())
	resp, err := requester.Post(url, writer.FormDataContentType(), body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}
