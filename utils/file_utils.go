package utils

import (
	"io/ioutil"
	"strings"
)

// ListMIFiles lists all `.mip` files in the current directory
func ListMIFiles() ([]string, error) {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		return nil, err
	}

	var miFiles []string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".mip") {
			miFiles = append(miFiles, file.Name())
		}
	}
	return miFiles, nil
}
