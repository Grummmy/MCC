package gui

import (
	"mcc/structs"
	"mcc/utils"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type UI struct {
	Scale            float32
	LogData          *widget.Entry
	ActionMenuOpened bool
	ChangeDirOpened  bool
}

func (ui *UI) Log(text string) {
	ui.LogData.Append("\n" + text)
}

func (ui *UI) ActionMenu(i Icons) {
	if ui.ActionMenuOpened {
		return
	}
	ui.ActionMenuOpened = true

	w := fyne.CurrentApp().NewWindow("actions")

	w.SetContent(
		container.NewVScroll(
			container.NewVBox(
				widget.NewButtonWithIcon("mods", i.Java, func() {}),
			),
		),
	)

	w.SetCloseIntercept(func() {
		ui.ActionMenuOpened = false
		w.Close()
	})
	w.Resize(fyne.Size{Width: 200.0, Height: 75.0})
	w.Show()
}

type Icons struct {
	Logo *fyne.StaticResource
	Java *fyne.StaticResource
}

func (i *Icons) LoadResources() {
	bytes, err := os.ReadFile("resources/logo.svg")
	if err != nil {
		panic(err)
	}
	i.Logo = fyne.NewStaticResource("logo", bytes)

	bytes, err = os.ReadFile("resources/java5.png")
	if err != nil {
		panic(err)
	}
	i.Java = fyne.NewStaticResource("java", bytes)
}

func BaseWidgets(ui *UI, i Icons, fe *structs.FileExplorer) (openDir, backDir, forwardDir, changeDir, upDir, action *widget.Button) {
	openDir = widget.NewButtonWithIcon("", theme.FolderOpenIcon(), func() { utils.OpenDir(fe.Current) })
	backDir = widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() { fe.Back() })
	forwardDir = widget.NewButtonWithIcon("", theme.NavigateNextIcon(), func() { fe.Forward() })

	changeDir = widget.NewButtonWithIcon("change dir", theme.FolderIcon(), func() {
		if ui.ChangeDirOpened {
			return
		}
		ui.ChangeDirOpened = true

		w := fyne.CurrentApp().NewWindow("Choose folder")
		dialog.NewFolderOpen(func(lu fyne.ListableURI, err error) {
			if lu == nil || err != nil {
				return
			}
			fe.Cd(lu.Path())
		}, w).Show()

		w.SetCloseIntercept(func() {
			ui.ChangeDirOpened = false
			w.Close()
		})

		w.Resize(fyne.Size{Width: 500, Height: 300})
		w.Show()
	})

	upDir = widget.NewButtonWithIcon("", theme.MoveUpIcon(), func() { fe.Cd("../") })
	action = widget.NewButtonWithIcon("action", theme.GridIcon(), func() { ui.ActionMenu(i) })

	return
}
