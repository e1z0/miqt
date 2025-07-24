package main

import (
	"os"

	"github.com/e1z0/miqt/qt"
	"github.com/e1z0/miqt/qt-extras/scintillaedit"
)

// n.b. May need LD_LIBRARY_PATH= env var set to find the necessary .so file

func main() {

	qt.NewQApplication(os.Args)

	area := scintillaedit.NewScintillaEdit2()
	area.SetFixedSize2(640, 480)
	area.Show()

	qt.QApplication_Exec()
}
