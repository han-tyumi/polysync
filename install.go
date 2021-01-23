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

var launcherURL string

func init() {
	switch runtime.GOOS {
	case "windows":
		launcherURL = "https://launcher.mojang.com/download/MinecraftInstaller.msi"
	case "darwin":
		launcherURL = "https://launcher.mojang.com/download/Minecraft.dmg"
	case "linux":
		launcherURL = "https://launcher.mojang.com/download/Minecraft.tar.gz"
	}
}

func downloadLauncher() (string, error) {
	if launcherURL == "" {
		return "", errors.New("platform not supported")
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	l.Printf("downloading %s ...\n", launcherURL)
	res, err := http.Get(launcherURL)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", errors.New(res.Status)
	}

	name := path.Base(launcherURL)
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
