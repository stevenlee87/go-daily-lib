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
	myWindow := myApp.NewWindow("Box Layout")

	hcontainer1 := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		canvas.NewText("left1", color.Black),
		canvas.NewText("right1", color.Black))

	// 左对齐
	hcontainer2 := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		canvas.NewText("left2", color.Black),
		canvas.NewText("right2", color.Black),
		layout.NewSpacer())

	// 右对齐
	hcontainer3 := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		layout.NewSpacer(),
		canvas.NewText("left3", color.Black),
		canvas.NewText("right3", color.Black))

	// 中间对齐
	hcontainer4 := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		layout.NewSpacer(),
		canvas.NewText("left4", color.Black),
		canvas.NewText("right4", color.Black),
		layout.NewSpacer())

	// 两边对齐
	hcontainer5 := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		canvas.NewText("left5", color.Black),
		layout.NewSpacer(),
		canvas.NewText("right5", color.Black))

	myWindow.SetContent(fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		hcontainer1, hcontainer2, hcontainer3, hcontainer4, hcontainer5))
	myWindow.Resize(fyne.NewSize(200, 200))
	myWindow.ShowAndRun()
}
