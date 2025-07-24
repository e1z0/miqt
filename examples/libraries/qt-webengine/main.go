package main

import (
	"os"

	"github.com/e1z0/miqt/qt"
	"github.com/e1z0/miqt/qt/webengine"
)

func main() {

	qt.NewQApplication(os.Args)

	w := webengine.NewQWebEngineView2()
	w.Load(qt.NewQUrl3("https://www.github.com/mappu/miqt"))
	w.Show()

	qt.QApplication_Exec()
}
