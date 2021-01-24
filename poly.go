package main

import (
	"path"
	"strings"

	"github.com/han-tyumi/fync"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

const modsPath = "mods"

type poly struct {
	*sftp.Client
}

var p *poly

func connect() (*poly, error) {
	if p != nil {
		return p, nil
	}

	l.Println("connecting to server ...")

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", address, config)
	if err != nil {
		return nil, err
	}

	client, err := sftp.NewClient(conn)
	if err != nil {
		return nil, err
	}

	p = &poly{client}
	return p, nil
}

func (p *poly) Mods() ([]fync.ServerFile, error) {
	progress.SetValue(0)
	l.Println("getting list of server mods ...")

	files, err := p.ReadDir(modsPath)
	if err != nil {
		return nil, err
	}

	progress.Max = float64(len(files)) * 2
	mods := make([]fync.ServerFile, 0)
	for i := range files {
		if !files[i].IsDir() && strings.HasSuffix(files[i].Name(), ".jar") {
			file, err := p.Open(path.Join(modsPath, files[i].Name()))
			if err != nil {
				return nil, err
			}
			mods = append(mods, file)
		}
		progress.add(1)
	}

	return mods, nil
}

func closePoly() {
	if p != nil {
		p.Close()
	}
}
