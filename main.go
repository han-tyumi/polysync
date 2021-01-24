package main

import (
	"log"
	"os/exec"
	"runtime"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
	"github.com/han-tyumi/fync"
)

var l *log.Logger
var progress *progressBar

func main() {
	a := app.New()
	w := a.NewWindow("PolySync")

	console := newConsole()
	l = log.New(console, "", log.Ltime)

	modsDir, err := fync.ModsDir()
	if err != nil {
		l.Printf("error: %s\n", err)
	}
	l.Printf("mods directory: %s\n", modsDir)

	backupDir, err := fync.BackupDir()
	if err != nil {
		l.Printf("error: %s\n", err)
	}
	l.Printf("backup directory: %s\n", backupDir)

	keepExistingCheck := widget.NewCheck("Keep Existing", nil)
	forceCheck := widget.NewCheck("Force", nil)

	progress = newProgressBar()

	w.SetContent(
		container.NewBorder(
			nil,
			container.NewVBox(
				progress.container(),
				container.NewHBox(keepExistingCheck, forceCheck),
				container.NewHBox(
					widget.NewButton("Sync", func() {
						l.Println("starting sync ...")
						n, err := syncMods(keepExistingCheck.Checked, forceCheck.Checked)
						if err != nil {
							l.Printf("error syncing: %s\n", err)
						} else if n == 0 {
							l.Println("already up to date")
						} else {
							l.Println("finished sync")
						}
					}),

					widget.NewButton("Download Forge Installer", func() {
						progress.startInfinite()
						l.Println("starting forge installer download ...")
						if _, err := downloadForge(); err != nil {
							l.Printf("error downloading: %s\n", err)
						} else {
							l.Println("finished forge installer download")
						}
						progress.stopInfinite()
					}),

					widget.NewButton("Download Minecraft Launcher", func() {
						progress.startInfinite()
						l.Println("starting launcher download ...")
						if _, err := downloadLauncher(); err != nil {
							l.Printf("error downloading: %s\n", err)
						} else {
							l.Println("finished launcher download")
						}
						progress.stopInfinite()
					}),

					widget.NewButton("Mods Folder", func() {
						var cmd string

						switch runtime.GOOS {
						case "windows":
							cmd = "explorer"
						case "darwin":
							cmd = "open"
						case "linux":
							cmd = "xdg-open"
						default:
							return
						}

						err := exec.Command(cmd, modsDir).Run()
						if _, ok := err.(*exec.ExitError); !ok {
							l.Printf("error opening: %s\n", err)
						}
					}),
				),
			),
			nil,
			nil,
			console.Scroll,
		),
	)

	w.Resize(fyne.NewSize(900, 400))

	w.ShowAndRun()

	closePoly()
}
