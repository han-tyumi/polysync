package main

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

var javaURL string

func init() {
	switch runtime.GOOS {
	case "windows":
		javaURL = "https://javadl.oracle.com/webapps/download/AutoDL?BundleId=244065_89d678f2be164786b292527658ca1605"
	case "darwin":
		javaURL = "https://javadl.oracle.com/webapps/download/AutoDL?BundleId=244059_89d678f2be164786b292527658ca1605"
	case "linux":
		switch runtime.GOARCH {
		case "amd64":
			javaURL = "https://javadl.oracle.com/webapps/download/AutoDL?BundleId=244058_89d678f2be164786b292527658ca1605"
		case "386":
			javaURL = "https://javadl.oracle.com/webapps/download/AutoDL?BundleId=244056_89d678f2be164786b292527658ca1605"
		}
	}
}

func downloadJava() (string, error) {
	if javaURL == "" {
		return "", errors.New("platform not supported")
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	l.Printf("downloading %s ...\n", javaURL)
	res, err := http.Get(javaURL)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", errors.New(res.Status)
	}

	name := path.Base(res.Request.URL.Path)
	dest := filepath.Join(filepath.Join(home, "Desktop"), name)

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
