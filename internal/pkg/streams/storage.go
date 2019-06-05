package streams

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

type FileReader struct {
	// assuming to use simple data scan instead of map with id
	handler []*StoredRecord
}

func NewFileReader() (*FileReader, error) {
	// read a file
	f, err := os.OpenFile("./data.json", os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	rec := FileReader{
		handler: loadData(f),
	}
	return &rec, nil
}

func (s *FileReader) ReadByID(idStr string) (*StoredRecord, error) {
	if s.handler == nil {
		// just no data there, no need to return error
		return nil, nil
	}
	for _, s := range s.handler {
		if strings.EqualFold(s.ID, idStr) {
			return s, nil
		}
	}
	return nil, nil
}

func loadData(hd *os.File) []*StoredRecord {
	b := make([]*StoredRecord, 0)
	data, err := ioutil.ReadAll(hd)
	if err != nil {
		return b
	}
	if err := json.Unmarshal(data, &b); err != nil {
		return b
	}
	return b
}
