package main

import (
	"sync"

	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

type console struct {
	Scroll  *container.Scroll
	Content *widget.Label
	mu      sync.Mutex
}

func newConsole() *console {
	content := widget.NewLabel("")
	content.TextStyle.Monospace = true

	scroll := container.NewScroll(content)

	l := &console{Scroll: scroll, Content: content}

	return l
}

func (l *console) Write(p []byte) (n int, err error) {
	l.mu.Lock()
	l.Content.SetText(l.Content.Text + string(p))
	l.Scroll.ScrollToBottom()
	l.mu.Unlock()
	return len(p), nil
}
