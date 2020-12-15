package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

/*
widget.NewHBox has deprecated!!! Use container.NewVBox() or container.NewHBox().
*/
func main() {
	myApp := app.New()
	myWin := myApp.NewWindow("Box")

	content := container.NewVBox(
		widget.NewLabel("The top row of VBox"),
		container.NewHBox(
			widget.NewLabel("Label 1"),
			widget.NewLabel("Label 2"),
		),
	)
	//content.Append(widget.NewButton("Append", func() {
	//	content.Append(widget.NewLabel("Appended"))
	//}))
	//content.Append(widget.NewButton("Prepend", func() {
	//	content.Prepend(widget.NewLabel("Prepended"))
	//}))

	content.Add(widget.NewButton("Add", func() {
		content.Add(widget.NewLabel("Added"))
	}))

	myWin.SetContent(content)
	myWin.Resize(fyne.NewSize(150, 150))
	myWin.ShowAndRun()
}
