package repository

import (
	"io/ioutil"
	"os"
)

type Repository struct {
}

func New() *Repository {
	return &Repository{}
}

func (r *Repository) Load(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil
	}
	return data
}
