package main

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("not find key in the dictionart")

func Search(dic Dictionary, key string) (string, error) {
	value1, ok := dic[key]
	if !ok {
		return "", ErrNotFound
	}
	return value1, nil
}
