package main

import (
	"errors"
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

	return download(launcherURL)
}
