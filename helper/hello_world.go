package helper

import "errors"

func HelloWorld(name string) (string, error) {
	if len(name) == 0 {
		return "", errors.New("Name is empty")
	}
	return "Hello world, " + name, nil
}
