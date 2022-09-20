package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/Bisns/fyne-iconpark/iconpark/Base"
)

var icons = []struct {
	name     string
	resource fyne.Resource
}{
	{"ConfiIcon", Base.ConfiIcon()},
	{"LikeIcon", Base.LikeIcon()},
	{"HomeIcon", Base.HomeIcon()},
}

func main() {
	a := app.New()
	w := a.NewWindow("Fyne iconpark")
	w.Resize(fyne.NewSize(800, 400))
	img := widget.NewIcon(nil)
	list := container.NewVBox()
	for _, icon := range icons {
		i := icon
		button := widget.NewButton(i.name, func() {
			img.SetResource(i.resource)
		})
		list.Add(button)
	}
	w.SetContent(
		container.New(
			layout.NewAdaptiveGridLayout(2),
			container.NewScroll(list),
			container.NewScroll(img),
		),
	)
	w.CenterOnScreen()
	w.ShowAndRun()
}
