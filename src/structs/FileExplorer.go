package structs

import (
	"mcc/utils"
	"os"

	"fyne.io/fyne/v2/widget"
)

type FileExplorer struct {
	Label      *widget.Label
	Current    string
	History    []string
	HistoryPos int
}

func (fe *FileExplorer) Cd(path string) {
	if fe.Current == path || path == "." {
		return
	}

	_ = os.Chdir(path)
	if len(fe.History) != 0 && utils.Pwd() == fe.History[len(fe.History)-1] {
		_ = os.Chdir(fe.History[len(fe.History)-1])
		return
	}
	fe.Current = utils.Pwd()

	if fe.HistoryPos < len(fe.History)-1 {
		fe.History = fe.History[:fe.HistoryPos+1]
	}

	fe.History = append(fe.History, fe.Current)
	if len(fe.History) > 255 {
		fe.History = fe.History[1:]
		if fe.HistoryPos > 0 {
			fe.HistoryPos--
		}
	}

	fe.Label.SetText(fe.Current)
	fe.HistoryPos = len(fe.History) - 1
}

func (fe *FileExplorer) Back() {
	if fe.HistoryPos > 0 {
		fe.HistoryPos--
		_ = os.Chdir(fe.History[fe.HistoryPos])
		fe.Current = fe.History[fe.HistoryPos]
		fe.Label.SetText(fe.Current)
	}
}

func (fe *FileExplorer) Forward() {
	if fe.HistoryPos < len(fe.History)-1 {
		fe.HistoryPos++
		_ = os.Chdir(fe.History[fe.HistoryPos])
		fe.Current = fe.History[fe.HistoryPos]
		fe.Label.SetText(fe.Current)
	}
}
