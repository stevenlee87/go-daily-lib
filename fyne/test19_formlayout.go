package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Border Layout")

	nameLabel := canvas.NewText("Name", color.Black)
	nameValue := canvas.NewText("steven", color.Black)
	ageLabel := canvas.NewText("Age", color.Black)
	ageValue := canvas.NewText("18", color.Black)

	container := fyne.NewContainerWithLayout(
		layout.NewFormLayout(),
		nameLabel, nameValue, ageLabel, ageValue,
	)
	myWindow.SetContent(container)
	myWindow.Resize(fyne.NewSize(200, 200))
	myWindow.ShowAndRun()
}
