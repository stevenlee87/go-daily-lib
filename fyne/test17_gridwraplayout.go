package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
)

/*
GridWrapLayout是GridLayout的扩展。GridWrapLayout创建时会指定一个初始size，这个size会应用到所有的子控件上，每个子控件都保持这个size。
初始，每行一个控件。如果界面大小变化了，这些子控件会重新排列。例如宽度翻倍了，那么一行就可以排两个控件了。有点像流动布局：
*/
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Grid Wrap Layout")

	img1 := canvas.NewImageFromResource(theme.FyneLogo())
	img2 := canvas.NewImageFromResource(theme.FyneLogo())
	img3 := canvas.NewImageFromResource(theme.FyneLogo())
	myWindow.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewGridWrapLayout(fyne.NewSize(150, 150)),
			img1, img2, img3))
	myWindow.ShowAndRun()
}
