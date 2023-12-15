package writer

import (
	"encoding/json"
	"io/ioutil"
)

type WriterAnswer struct {
	File     string
	Response interface{}
	Data     []byte
}

func New() *WriterAnswer {
	return &WriterAnswer{}
}

func WriteAnswer(w *WriterAnswer) error {
	strStruct, err := json.Marshal(w.Response)

	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(w.File, strStruct, 0755); err != nil {
		return err
	}

	return nil
}
