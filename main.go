package main

import "fucker/frame/comp"

func main() {
	//a := app.New()
	//w := a.NewWindow("Hello")
	//
	//hello := widget.NewLabel("Hello Fyne!")
	//w.SetContent(container.NewVBox(
	//	hello,
	//	widget.NewButton("Hi!", func() {
	//		hello.SetText("Welcome :)")
	//	}),
	//))
	//
	//w.ShowAndRun()
	p := "/Users/james/Desktop/config.json"
	comp.StartFromConfigFile(p)
}
