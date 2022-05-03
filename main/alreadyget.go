package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

type Already struct {
	Get map[string]float64
}

func Getalready(app fyne.App) *Already {
	get := new(Already)
	win := app.NewWindow("Hello Niki ! ^-^ ")
	sayhi := widget.NewLabel("Say Hi to Niki ! Struggling in figuring out what to eat this night ? Niki will help you .")
	asktoinput := widget.NewLabel("What did you eat this noon ?")
	//Entry for rice
	//
	inputrice := widget.NewEntry()
	inputrice.SetPlaceHolder("The rice you eat :")
	//Entry for pork
	//
	inputpork := widget.NewEntry()
	inputpork.SetPlaceHolder("The pork you eat :")
	//Entry for banana
	//
	inputbanana := widget.NewEntry()
	inputbanana.SetPlaceHolder("The banana you eat :")
	//The window
	//
	win.SetContent(container.NewVBox(
		sayhi,
		asktoinput,
		inputrice,
		inputpork,
		inputbanana,
		widget.NewButton("Done",
			func() {
				get.Get = make(map[string]float64)
				get.Get["Protein"] = CaculateProtein("rice", GetFromEntry(inputrice)) + CaculateProtein("pork", GetFromEntry(inputpork)) + CaculateFat("banana", GetFromEntry(inputbanana))
				get.Get["Va"] = CaculateVa("rice", GetFromEntry(inputrice)) + CaculateVa("pork", GetFromEntry(inputpork)) + CaculateVa("banana", GetFromEntry(inputbanana))
				get.Get["Vb"] = CaculateVb("rice", GetFromEntry(inputrice)) + CaculateVb("pork", GetFromEntry(inputpork)) + CaculateVb("banana", GetFromEntry(inputbanana))
				get.Get["Vc"] = CaculateVc("rice", GetFromEntry(inputrice)) + CaculateVc("pork", GetFromEntry(inputpork)) + CaculateVc("banana", GetFromEntry(inputbanana))
				get.Get["Fat"] = CaculateFat("rice", GetFromEntry(inputrice)) + CaculateFat("pork", GetFromEntry(inputpork)) + CaculateFat("banana", GetFromEntry(inputbanana))
				get.Get["Carbon"] = CaculateCarbon("rice", GetFromEntry(inputrice)) + CaculateCarbon("pork", GetFromEntry(inputpork)) + CaculateCarbon("banana", GetFromEntry(inputbanana))
				Need := GetNeed(get)
				ShowAnswer(*Need, app)
			}),
	))
	win.ShowAndRun()
	return get
}
func GetFromEntry(En *widget.Entry) float64 {
	//Use the method strconv.ParseFloat to transform a string into float (The same to int )
	data, err := strconv.ParseFloat(En.Text, 64)
	//Judge error
	if err != nil {
		fmt.Println("strconv.ParseFloat() error : ", err)
		return 0
	} else {
		return data
	}
}

//For protein
//
func CaculateProtein(Name string, Num float64) float64 {
	var Back float64
	for p := Getfooddata(); p != nil; p = p.Next {
		if p.Name == Name {
			Back = Num * p.Protein
			return Back
		}
	}
	return 8888
}

//For Va
//
func CaculateVa(Name string, Num float64) float64 {
	var Back float64
	for p := Getfooddata(); p != nil; p = p.Next {
		if p.Name == Name {
			Back = Num * p.Vita.Va
			return Back
		}
	}
	return 8888
}

//For Vb
//
func CaculateVb(Name string, Num float64) float64 {
	var Back float64
	for p := Getfooddata(); p != nil; p = p.Next {
		if p.Name == Name {
			Back = Num * p.Vita.Vb
			return Back
		}
	}
	return 8888
}

//For Vc
//
func CaculateVc(Name string, Num float64) float64 {
	var Back float64
	for p := Getfooddata(); p != nil; p = p.Next {
		if p.Name == Name {
			Back = Num * p.Vita.Vc
			return Back
		}
	}
	return 8888
}

//For fat
//
func CaculateFat(Name string, Num float64) float64 {
	var Back float64
	for p := Getfooddata(); p != nil; p = p.Next {
		if p.Name == Name {
			Back = Num * p.Fat
			return Back
		}
	}
	return 8888
}

//For carbon
//
func CaculateCarbon(Name string, Num float64) float64 {
	var Back float64
	for p := Getfooddata(); p != nil; p = p.Next {
		if p.Name == Name {
			Back = Num * p.Carbon
			return Back
		}
	}
	return 8888
}
