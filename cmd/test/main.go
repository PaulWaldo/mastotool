package main

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	tags := []string{"jhdjhg", "hgfehjqgheq", "jhdfkjhjkdhgjhjhg", "jdhfjkhgjkhg"}
	var sb strings.Builder
	sb.WriteString("These tags will be removed from your following list and not seen in your feed:")
	for _, t := range tags {
		sb.WriteString("\n* " + t)
	}
	rt := widget.NewRichTextFromMarkdown(sb.String())
	rt.Wrapping = fyne.TextWrapWord
	scroll := container.NewVScroll(rt)
	scroll.SetMinSize(fyne.Size{Width: 50, Height: 150})

	w.Resize(fyne.Size{Width: 250, Height: 200})
	d := dialog.NewCustomConfirm("Confirm", "Confirm", "Cancel", scroll, func(b bool) {}, w)
	d.Show()
	w.ShowAndRun()
}
