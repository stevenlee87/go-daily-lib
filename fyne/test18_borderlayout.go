package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

/*
边框布局（BorderLayout）比较常用于构建用户界面，上面例子中的Toolbar一般都和BorderLayout搭配使用。
创建方法layout.NewBorderLayout(top, bottom, left, right)，分别传入顶部、底部、左侧、右侧的控件对象。
添加到容器中的控件如果是这些边界对象，则显示在对应位置，其他都显示在中心：
*/
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Border Layout")

	left := canvas.NewText("left", color.Black)
	right := canvas.NewText("right", color.Black)
	top := canvas.NewText("top", color.Black)
	bottom := canvas.NewText("bottom", color.Black)
	content := widget.NewLabel(`Lorem ipsum dolor, 
  sit amet consectetur adipisicing elit.
  Quidem consectetur ipsam nesciunt,
  quasi sint expedita minus aut,
  porro iusto magnam ducimus voluptates cum vitae.
  Vero adipisci earum iure consequatur quidem.`)

	container := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(top, bottom, left, right),
		top, bottom, left, right, content,
	)
	myWindow.SetContent(container)
	myWindow.ShowAndRun()
}
