package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("TabContainer")

	nameLabel := widget.NewLabel("Name: steven")
	sexLabel := widget.NewLabel("Sex: male")
	ageLabel := widget.NewLabel("Age: 18")
	addressLabel := widget.NewLabel("City: beijing")
	addressLabel.Hide()
	profile := container.NewVBox(nameLabel, sexLabel, ageLabel, addressLabel)

	musicRadio := widget.NewRadioGroup([]string{"on", "off"}, func(string) {})
	showAddressCheck := widget.NewCheck("show address?", func(value bool) {
		if !value {
			addressLabel.Hide()
		} else {
			addressLabel.Show()
		}
	})
	memberTypeSelect := widget.NewSelect([]string{"junior", "senior", "admin"}, func(string) {})

	setting := widget.NewForm(
		&widget.FormItem{Text: "music", Widget: musicRadio},
		&widget.FormItem{Text: "check", Widget: showAddressCheck},
		&widget.FormItem{Text: "member type", Widget: memberTypeSelect},
	)

	tabs := widget.NewTabContainer(
		widget.NewTabItem("Profile", profile),
		widget.NewTabItem("Setting", setting),
	)

	myWindow.SetContent(tabs)
	myWindow.Resize(fyne.NewSize(200, 200))
	myWindow.ShowAndRun()
}
