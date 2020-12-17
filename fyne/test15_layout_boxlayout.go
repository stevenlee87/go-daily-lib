package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
)

/*
盒状布局（BoxLayout）是最常使用的一个布局。它将控件都排在一行或一列。在fyne中，我们可以通过layout.NewHBoxLayout()创建
一个水平盒装布局，通过layout.NewVBoxLayout()创建一个垂直盒装布局。水平布局中的控件都排列在一行中，每个控件的宽度等于其
内容的最小宽度（MinSize().Width），它们都拥有相同的高度，即所有控件的最大高度（MinSize().Height）。

垂直布局中的控件都排列在一列中，每个控件的高度等于其内容的最小高度，它们都拥有相同的宽度，即所有控件的最大宽度。

一般地，在BoxLayout中使用layout.NewSpacer()辅助布局，它会占满剩余的空间。对于水平盒状布局来说，
第一个控件前添加一个layout.NewSpacer()，所有控件右对齐。最后一个控件后添加一个layout.NewSpacer()，
所有控件左对齐。前后都有，那么控件中间对齐。如果在中间有添加一个layout.NewSpacer()，那么其它控件两边对齐。
*/

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
