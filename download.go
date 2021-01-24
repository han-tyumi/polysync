package main

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

var downloadPath string

func init() {
	if home, err := os.UserHomeDir(); err == nil {
		downloadPath = filepath.Join(home, "Desktop")
	} else if wd, err := os.Getwd(); err == nil {
		downloadPath = wd
	}
}

func download(url string) (string, error) {
	l.Printf("downloading %s ...\n", url)
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", errors.New(res.Status)
	}

	name := path.Base(res.Request.URL.Path)
	dest := filepath.Join(downloadPath, name)

	l.Printf("copying to %s ...\n", dest)
	file, err := os.Create(dest)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if _, err := io.Copy(file, res.Body); err != nil {
		return "", err
	}

	return dest, nil
}
