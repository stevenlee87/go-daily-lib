package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

/*
fyne的使用很简单。每个fyne程序都包括两个部分，一个是应用程序对象myApp，通过app.New()创建。另一个是窗口对象，通过应用程序对象myApp来
创建myApp.NewWindow("Hello")。myApp.NewWindow()方法中传入的字符串就是窗口标题。
*/
func main() {
	myApp := app.New()

	myWin := myApp.NewWindow("Hello")
	myWin.SetContent(widget.NewLabel("Hello Fyne!"))
	myWin.Resize(fyne.NewSize(200, 200))
	myWin.ShowAndRun()
}
