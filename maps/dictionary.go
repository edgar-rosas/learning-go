package maps

import (
	"errors"
)

const (
	ErrNotFound        = DictionaryErr("could not find the key you were looking for")
	ErrKeyExists       = DictionaryErr("key already exists")
	ErrKeyDoesNotExist = DictionaryErr("key not found in dictionary")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

// Search returns the value of the provided key. Returns error if not found
func (d Dictionary) Search(key string) (string, error) {
	definition, found := d[key]
	if !found {
		return "", ErrNotFound
	}

	return definition, nil
}

// AddKey adds a new key with its definition to the dictionary
func (d Dictionary) AddKey(key, definition string) error {
	_, err := d.Search(key)

	switch {
	case errors.Is(err, ErrNotFound):
		d[key] = definition
	case err == nil:
		return ErrKeyExists
	default:
		return err
	}

	return nil
}

// Update updates the value of a key if found.
func (d Dictionary) Update(key, definition string) error {
	_, err := d.Search(key)

	switch {
	case errors.Is(err, ErrNotFound):
		return ErrKeyDoesNotExist
	case err == nil:
		d[key] = definition
	default:
		return err
	}

	return nil
}

// Delete removes a key from the dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
