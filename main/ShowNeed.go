package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	_ "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowAnswer(Want Need, app fyne.App) {
	w := app.NewWindow("Your choice ! ")
	labelofprotein := widget.NewLabel("The Protein you need : " + fmt.Sprintf("%.2f", Want.Needmap["Protein"]))
	labelofva := widget.NewLabel("The VitaminA you need : " + fmt.Sprintf("%.2f", Want.Needmap["Va"]))
	labelofvb := widget.NewLabel("The VitaminB you need : " + fmt.Sprintf("%.2f", Want.Needmap["Vb"]))
	labelofvc := widget.NewLabel("The VitaminC you need : " + fmt.Sprintf("%.2f", Want.Needmap["Vc"]))
	labeloffat := widget.NewLabel("The Fat you need : " + fmt.Sprintf("%.2f", Want.Needmap["Fat"]))
	labelofcarbon := widget.NewLabel("The Carbohydrate complex you need : " + fmt.Sprintf("%.2f", Want.Needmap["Carbon"]))
	w.SetContent(container.NewVBox(
		labelofprotein,
		labelofva,
		labelofvb,
		labelofvc,
		labeloffat,
		labelofcarbon,
	))
	w.Show()
}
