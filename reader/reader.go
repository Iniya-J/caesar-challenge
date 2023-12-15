package reader

import (
	"os"
)

type ReaderAnswer struct {
	Info os.FileInfo
	Data []byte
}

func ReadAnswer(f string) (*ReaderAnswer, error) {
	reader := &ReaderAnswer{}

	file, err := os.Open(f)
	defer file.Close()

	if err != nil {
		return nil, err
	}

	stat, _ := file.Stat()
	data := make([]byte, stat.Size())

	if _, err := file.Read(data); err != nil {
		return nil, err
	}

	reader.Data = data
	reader.Info = stat
	return reader, nil
}
