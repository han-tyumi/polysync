package main

import (
	"errors"
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

	return download(javaURL)
}
