package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
)

/*
格子布局（GridLayout）每一行有固定的列，添加的控件数量超过这个值时，后面的控件将会在新的行显示。
创建方法layout.NewGridLayout(cols)，传入每行的列数。
*/
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Grid Layout")

	img1 := canvas.NewImageFromResource(theme.FyneLogo())
	img2 := canvas.NewImageFromResource(theme.FyneLogo())
	img3 := canvas.NewImageFromResource(theme.FyneLogo())
	myWindow.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayout(2),
		img1, img2, img3))
	myWindow.Resize(fyne.NewSize(300, 300))
	myWindow.ShowAndRun()
}
