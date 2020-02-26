package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func Search(basePath string, extension string, only_basename bool, include_extension bool) ([]string, error) {
	var files []string

	fullPath, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	root := fullPath + basePath

	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(strings.TrimSpace(path)) == extension {
			if only_basename {
				path = strings.Replace(path, root, "", 1)
			}

			if !include_extension {
				path = strings.Replace(path, extension, "", 1)
			}

			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func ReadFile(basePath string) ([]byte, error) {
	fullPath, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	root := fullPath + basePath

	return ioutil.ReadFile(root)
}

func WriteFile(basePath string, content []byte) error {
	fullPath, err := os.Getwd()

	if err != nil {
		return err
	}

	root := fullPath + basePath

	return ioutil.WriteFile(root, content, 0600)
}

func GetTmpPath() string {
	path := "/tmp/"

	if os.Getenv("APP_ENV") == "testing" {
		path += "/_tests/"
	}

	return path
}
