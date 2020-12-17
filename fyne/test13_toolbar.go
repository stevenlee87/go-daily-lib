package main

import (
	"fmt"
	"log"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

/*
工具栏（Toolbar）是很多 GUI 应用程序必备的部分。工具栏将常用命令用图标的方式很形象地展示出来，方便使用。
创建方法widget.NewToolbar()，传入多个widget.ToolbarItem作为参数。
最常使用的ToolbarItem有命令（Action）、分隔符（Separator）和空白（Spacer），分别通过
widget.NewToolbarItemAction(resource, callback)
widget.NewToolbarSeparator()
widget.NewToolbarSpacer()创建。命令需要指定回调，点击时触发。
*/
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Toolbar")

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			fmt.Println("New document")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {
			fmt.Println("Cut")
		}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {
			fmt.Println("Copy")
		}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {
			fmt.Println("Paste")
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Display help")
		}),
	)

	content := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(toolbar, nil, nil, nil),
		toolbar, widget.NewLabel(`spring goeth all in white,
    crowned with milk-white may;
    in fleecy flocks of light,
    o'er heaven the white clouds stray;
    white butterflies in the air;
    white daisies prank the ground;
    the cherry and hoary pear,
    scatter their snow around.`),
	)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
