package main

import (
	gui "mcc/interface"
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

	icons := gui.Icons{}
	icons.LoadResources()
	w.SetIcon(icons.Logo)

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

	log := widget.NewMultiLineEntry()
	log.Wrapping = fyne.TextWrapWord
	UI := gui.UI{
		Scale:   0,
		LogData: log,
	}

	openDir, backDir, forwardDir, changeDir, upDir, action := gui.BaseWidgets(&UI, icons, &FE)

	pathField := container.NewBorder(
		nil, nil,
		openDir,
		widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
			clipboard.WriteAll(FE.Current)
		}),
		FE.Label,
	)
	pathControl := container.NewBorder(
		nil, nil,
		container.NewHBox(backDir, forwardDir),
		container.NewHBox(upDir, action),
		changeDir,
	)

	w.SetContent(
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
	w.Resize(fyne.Size{Width: 400.0, Height: 200.0})
	w.ShowAndRun()
}
