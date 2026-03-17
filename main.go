package main

import (
	"fucker/frame/comp"
	"fucker/frame/config"
)

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

	path := config.GetDefaultConfigPath()
	comp.StartFromConfigFile(path)
}
