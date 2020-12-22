package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
)

/*
CenterLayout将容器内的所有控件显示在中心位置，按传入的顺序显示。最后传入的控件显示最上层。
CenterLayout中所有控件将保持它们的最小尺寸（大小能容纳其内容）。
*/
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Center Layout")

	image := canvas.NewImageFromResource(theme.FyneLogo())
	image.FillMode = canvas.ImageFillOriginal
	text := canvas.NewText("Fyne Logo", color.Black)

	container := fyne.NewContainerWithLayout(
		layout.NewCenterLayout(),
		image, text,
		//text, image,
	)
	myWindow.SetContent(container)
	myWindow.ShowAndRun()
}
