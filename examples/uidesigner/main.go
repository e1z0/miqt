package main

import (
	"os"

	"github.com/e1z0/miqt/qt"
)

func main() {
	qt.NewQApplication(os.Args)

	ui := NewMainWindowUi()
	ui.MainWindow.Show()

	qt.QApplication_Exec()
}
