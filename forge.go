package main

import (
	"os"
	"path"
	"path/filepath"
)

const forgeInstallerPath = "forge-installer.jar"

func downloadForge() (string, error) {
	poly, err := connect()
	if err != nil {
		return "", err
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	forge, err := poly.Open(forgeInstallerPath)
	if err != nil {
		return "", err
	}
	defer forge.Close()

	name := path.Base(forgeInstallerPath)
	destPath := filepath.Join(filepath.Join(home, "Desktop"), name)

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
