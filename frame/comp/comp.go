package comp

import (
	"fucker/frame/config"
	"fucker/frame/factroy"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var Gcontext *FContext

func init() {
	Gcontext = &FContext{
		C: make(map[string]*GParam),
	}
}
func StartFromConfigFile(file string) {
	config := config.ParseConfig(file)

	log.Println(config)

	app := GetApp(config)

	comp := GetComp(config)

	app.SetContent(comp)
	app.ShowAndRun()
}

func GetComp(c *config.App) fyne.CanvasObject {
	child := c.Child

	comp := DoGetComp2(child[0])
	return comp
}

func DoGetComp2(node *config.Child) fyne.CanvasObject {
	root := factroy.Create(node)

	container, err := root.(*fyne.Container)

	if !err {
		w, e := root.(fyne.CanvasObject)

		if !e {
			panic(e)
		}

		return w
	}
	for _, subNode := range node.Child {
		childTree := DoGetComp2(subNode)
		container.Add(childTree)
	}

	return root
}

func GetApp(c *config.App) fyne.Window {
	a := app.NewWithID(c.AppId)
	window := a.NewWindow(c.Title)

	window.Resize(fyne.NewSize(c.Weight, c.Height))

	return window
}
