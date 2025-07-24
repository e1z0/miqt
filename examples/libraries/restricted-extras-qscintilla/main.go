package main

import (
	"os"

	"github.com/e1z0/miqt/qt"
	"github.com/e1z0/miqt/qt-restricted-extras/qscintilla"
)

func main() {

	qt.NewQApplication(os.Args)

	area := qscintilla.NewQsciScintilla2()
	area.SetFixedSize2(640, 480)
	area.Show()

	qt.QApplication_Exec()
}
