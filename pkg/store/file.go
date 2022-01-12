package store

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

type Type string

const (
	FileType Type = "file"
)

func New(store Type, fileName string) Store {
	switch store {
	case FileType:
		return &FileStore{
			FileName: fileName,
			Mock:     &Mock{},
		}
	}
	return nil
}

type FileStore struct {
	FileName string
	Mock     *Mock
}
type Mock struct {
	Data []byte
	Err  error
}

func (fs *FileStore) AddMock(mock *Mock) {
	fs.Mock = mock
}
func (fs *FileStore) ClearMock() {
	fs.Mock = nil
}

func (fs *FileStore) Write(data interface{}) (err error) {
	if fs.Mock != nil {
		if fs.Mock.Err != nil {
			return fs.Mock.Err
		}
		return json.Unmarshal(fs.Mock.Data, data)
	}
	fileData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	f, err := os.OpenFile(fs.FileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	_, err = f.Write(fileData)
	if err != nil {
		return err
	}
	return nil
}

func (fs *FileStore) Read(data interface{}) error {
	if fs.Mock != nil {
		if fs.Mock.Err != nil {
			return fs.Mock.Err
		}
		return json.Unmarshal(fs.Mock.Data, data)
	}
	file, err := ioutil.ReadFile(fs.FileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, data)

	/*
		file, err := os.ReadFile(fs.FileName)
		if err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				return err
			}
			file = []byte("[]")
		}
		err = json.Unmarshal(file, &data)
		return err
	*/
}
