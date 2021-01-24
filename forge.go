package main

import (
	"os"
	"path/filepath"
)

const forgeInstaller = "forge-installer.jar"

func downloadForge() (string, error) {
	poly, err := connect()
	if err != nil {
		return "", err
	}

	forge, err := poly.Open(forgeInstaller)
	if err != nil {
		return "", err
	}
	defer forge.Close()

	destPath := filepath.Join(downloadPath, forgeInstaller)

	l.Printf("copying to %s ...\n", destPath)
	dest, err := os.Create(destPath)
	if err != nil {
		return "", err
	}
	defer dest.Close()

	if _, err := forge.WriteTo(dest); err != nil {
		return "", err
	}

	return destPath, nil
}
