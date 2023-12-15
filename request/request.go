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

// ChallengeResponse struct deals with the http response
type ChallengeResponse struct {
	Places        int    `json:"numero_casas"`
	Token         string `json:"token"`
	CryptedText   string `json:"cifrado"`
	DecryptedText string `json:"decifrado"`
	SummaryCrypto string `json:"resumo_criptografico"`
}

func getRequest(url string) ([]byte, error) {
	log.Printf("Making request to %s", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func parseResponse(body []byte) (*ChallengeResponse, error) {
	response := &ChallengeResponse{}
	err := json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetCryptedText sends request to codenation and return a struct with the json parsed
func GetCryptedText(file string, getRequest func(string) ([]byte, error), parseResponse func([]byte) (*ChallengeResponse, error)) (*writer.WriterAnswer, error) {
	w := writer.New()
	url := fmt.Sprintf("%s?token=%s", config.GenerateUrl, config.TokenCodeNation)

	body, err := getRequest(url)
	if err != nil {
		return nil, err
	}

	response, err := parseResponse(body)
	if err != nil {
		return nil, err
	}

	w.File = file
	w.Response = response
	w.Data = body
	writer.WriteAnswer(w)

	return w, nil
}

func postRequest(url string, body *bytes.Buffer) ([]byte, error) {
	log.Printf("Making request to %s", url)

	resp, err := http.Post(url, "multipart/form-data", body)
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

// PostSubmitData sends a POST request to submit the data
func PostSubmitData(file string, postRequest func(string, *bytes.Buffer) ([]byte, error)) ([]byte, error) {
	url := fmt.Sprintf("%s?token=%s", config.SubmitUrl, config.TokenCodeNation)

	r, err := reader.ReadAnswer(file)
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

	respBody, err := postRequest(url, body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}
