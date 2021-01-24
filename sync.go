package main

import (
	"os"

	"github.com/han-tyumi/fync"
)

func syncMods(keepExisting, force bool) (int, error) {
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
			if curr != 0 {
				progress.add(1)
			}
		},
	}

	return fync.Sync(poly, options)
}
