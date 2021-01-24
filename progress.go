package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

type progressBar struct {
	*widget.ProgressBar
	Infinite *widget.ProgressBarInfinite
}

func newProgressBar() *progressBar {
	regular := widget.NewProgressBar()
	infinite := widget.NewProgressBarInfinite()
	infinite.Hide()

	return &progressBar{ProgressBar: regular, Infinite: infinite}
}

func (p *progressBar) container() *fyne.Container {
	return container.NewVBox(p.ProgressBar, p.Infinite)
}

func (p *progressBar) add(n float64) {
	p.SetValue(p.Value + n)
}

func (p *progressBar) startInfinite() {
	p.SetValue(p.Min)
	p.Hide()
	p.Infinite.Show()
	p.Infinite.Start()
}

func (p *progressBar) stopInfinite() {
	p.SetValue(p.Max)
	p.Infinite.Stop()
	p.Infinite.Hide()
	p.Show()
}
