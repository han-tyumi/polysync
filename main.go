package main

import (
	"log"
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
	"github.com/han-tyumi/fync"
)

var l *log.Logger
var progress *widget.ProgressBar

func runSync(keepExisting, force bool) (int, error) {
	poly, err := connect()
	if err != nil {
		return 0, err
	}

	options := &fync.SyncOptions{
		KeepExisting: keepExisting,
		Force:        force,
		OnWrite: func(from os.FileInfo, _ string) {
			l.Printf("copying %s ...\n", from.Name())
		},
		OnBackup: func(name, _, _ string) {
			l.Printf("backing up %s ... \n", name)
		},
		OnProgress: func(_ string, curr, total int) {
			progress.Max = float64(total)
			progress.SetValue(float64(curr))
		},
	}

	return fync.Sync(poly, options)
}

func main() {
	a := app.New()
	w := a.NewWindow("Polygondwanaland Sync")

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

	progress = widget.NewProgressBar()

	w.SetContent(
		container.NewBorder(
			nil,
			container.NewVBox(
				progress,
				container.NewHBox(keepExistingCheck, forceCheck),
				widget.NewButton("Sync", func() {
					progress.SetValue(0)
					l.Println("starting sync ...")
					n, err := runSync(keepExistingCheck.Checked, forceCheck.Checked)
					if err != nil {
						l.Printf("error syncing: %s", err)
					} else if n == 0 {
						l.Println("already up to date")
					} else {
						l.Println("finished sync")
					}
				}),
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
