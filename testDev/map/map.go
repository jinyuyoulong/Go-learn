package main

type Dictionary map[string]string

const (
	ErrNotFound  = DictionaryErr("not find key in the dictionart")
	ErrNotExists = DictionaryErr("cannot add word because it already exists")
)

func (d Dictionary) Search(key string) (string, error) {
	value1, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return value1, nil
}

type DictionaryErr string

func (d DictionaryErr) Error() string {
	return string(d)
}
func (d Dictionary) Add(key, v string) error {
	_, err := d.Search(key)
	switch err {
	case nil:
		return ErrNotExists
	case ErrNotFound:
		return ErrNotFound
	default:
		return err
	}
	d[key] = v
	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case nil:
		return ErrNotExists
	case ErrNotFound:
		return ErrNotFound
	default:
		return err
	}
	d[key] = value
	return nil
}
