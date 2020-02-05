package dictionary

import "errors"

type Dictionary map[string]string

var errNotFound = errors.New("Word not found")

func (d Dictionary) Search(key string) (string, error) {
	definition := d[key]

	if definition == "" {
		return "", errNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(key, value string) {
	d[key] = value
}
