package main

import (
	"os"

	"github.com/e1z0/miqt/qt"
	"github.com/e1z0/miqt/qt/printsupport"
)

func main() {

	qt.NewQApplication(os.Args)

	btn := qt.NewQPushButton3("QPrintSupport sample")
	btn.SetFixedWidth(320)

	btn.OnPressed(func() {

		dlg := printsupport.NewQPrintDialog3()
		dlg.OnFinished(func(int) {
			dlg.DeleteLater()
		})
		dlg.Show()
	})

	btn.Show()

	qt.QApplication_Exec()
}
