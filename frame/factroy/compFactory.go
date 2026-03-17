package factroy

import (
	"fucker/frame/config"
	"fucker/frame/execute"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
)

func Create(name *config.Child) fyne.CanvasObject {
	log.Println("create " + name.Type + "key:" + name.Key)
	if name.Type == "widget" {
		return createWight(name)
	}

	if name.Type == "layout" {
		return createLayout(name)
	}

	panic("not support type")
}

func createLayout(w *config.Child) *fyne.Container {
	switch w.Key {
	case "vbox":
		return createVbox(w)
		break
	case "hbox":
		return createHbox(w)
		break
	case "grid":
		return createGrid(w)
		break
	case "vscroll":
		return createVScroll(w)
		break
	case "hscroll":
		return createhScroll(w)
		break
	}

	return nil
}

func createhScroll(w *config.Child) *fyne.Container {
	return nil
}

func createVScroll(w *config.Child) *fyne.Container {
	return nil
}

func createGrid(w *config.Child) *fyne.Container {
	grid := container.NewAdaptiveGrid(w.Attr.Row)

	return grid
}

func createButton(w *config.Child) *widget.Button {
	button := widget.NewButton(w.Attr.Text, func() {
		if w.Function != nil {

			if w.Function.Type == "js" {
				execute.ExecJs(w.Function.FuncName, w.Function.Param)
			}
		}
	})

	return button
}

func createHbox(w *config.Child) *fyne.Container {
	return container.NewHBox()
}

func createVbox(w *config.Child) *fyne.Container {
	return container.NewVBox()
}

func createWight(w *config.Child) fyne.CanvasObject {
	switch w.Key {
	case "label":
		return createLabel(w)
		break
	case "entry":
		return createInput(w)
		break
	case "button":
		return createButton(w)
		break
	case "sep":
		return createSep(w)
		break
	}

	return nil
}

func createSep(w *config.Child) fyne.CanvasObject {
	separator := widget.NewSeparator()
	return separator
}

func createInput(w *config.Child) *widget.Entry {
	entry := widget.NewEntry()
	entry.SetText(w.Attr.Text)

	return entry
}

func createLabel(w *config.Child) *widget.Label {

	label := widget.NewLabel(w.Attr.Text)

	return label
}
