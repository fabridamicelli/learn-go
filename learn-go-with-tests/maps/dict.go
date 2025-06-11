package main

import "errors"

var WordNotFoundError = errors.New("word not found")
var WordAlreadyExistsError = errors.New("word already exists")

type Dict map[string]string

func (d Dict) Search(key string) (string, error) {
	got, ok := d[key]
	if !ok {
		return "", WordNotFoundError
	}

	return got, nil

}

func (d Dict) Add(key, val string) error {

	_, wordPresent := d[key]
	if wordPresent {
		return WordAlreadyExistsError
	}

	d[key] = val
	return nil
}

func (d Dict) Update(key, val string) {
	d[key] = val
}

func (d Dict) Delete(key string) {
	delete(d, key)
}
