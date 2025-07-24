package main

import (
	"os"

	"github.com/e1z0/miqt/qt"
	"github.com/e1z0/miqt/qt/svg"
)

func main() {

	qt.NewQApplication(os.Args)

	w := svg.NewQSvgWidget3("../../../doc/logo.svg")
	w.Show()

	qt.QApplication_Exec()
}
