package main

import (
	"os"

	qt "github.com/e1z0/miqt/qt6"
	"github.com/e1z0/miqt/qt6/webengine"
)

func main() {

	qt.NewQApplication(os.Args)

	w := webengine.NewQWebEngineView2()
	w.Load(qt.NewQUrl3("https://www.github.com/mappu/miqt"))
	w.Show()

	qt.QApplication_Exec()
}
