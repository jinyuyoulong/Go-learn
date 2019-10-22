package main

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound  = errors.New("not find key in the dictionart")
	ErrNotExists = errors.New("cannot add word because it already exists")
)

func (d Dictionary) Search(key string) (string, error) {
	value1, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return value1, nil
}

func (d Dictionary) Add(key, v string) error {
	d[key] = v
	return nil
}
