package main

import (
	"fmt"
	gui "mcc/interface"
	"mcc/keys"
	"mcc/structs"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/atotto/clipboard"
)

func main() {
	a := app.New()
	w := a.NewWindow("MCC")
	w.Resize(fyne.Size{Width: 400.0, Height: 200.0})

	icons := gui.Icons{}
	icons.LoadResources()
	w.SetIcon(icons.Logo)

	options := keys.Check(w)
	fmt.Println(options)

	userPath, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	os.Chdir(userPath)

	FE := structs.FileExplorer{
		Label: widget.NewLabelWithStyle(
			userPath,
			fyne.TextAlignLeading,
			fyne.TextStyle{Bold: true},
		),
		Current: userPath,
		History: []string{userPath},
	}

	log := widget.NewMultiLineEntry() //log field
	log.Wrapping = fyne.TextWrapWord
	UI := gui.UI{
		Scale:   0,
		LogData: log,
	}

	openDir, backDir, forwardDir, changeDir, upDir, action := gui.BaseWidgets(&UI, icons, &FE)

	pathField := container.NewBorder( // path label and copy/open buttons
		nil, nil,
		openDir,
		widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
			clipboard.WriteAll(FE.Current)
		}),
		FE.Label,
	)
	pathControl := container.NewBorder( // back/forward/up dir, change dir and action buttons
		nil, nil,
		container.NewHBox(backDir, forwardDir),
		container.NewHBox(upDir, action),
		changeDir,
	)

	w.SetContent( // packing all containers into main window
		container.NewBorder(
			container.NewVBox(
				pathField,
				pathControl,
			),
			nil, nil, nil,
			log,
		),
	)

	UI.LogData.SetText("MCC started successfully!")
	w.ShowAndRun()
}
