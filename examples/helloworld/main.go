package main

import (
	"fmt"
	"os"

	"github.com/e1z0/miqt/qt"
)

func main() {

	qt.NewQApplication(os.Args)

	btn := qt.NewQPushButton3("Hello world!")
	btn.SetFixedWidth(320)

	var counter int = 0

	btn.OnPressed(func() {
		counter++
		btn.SetText(fmt.Sprintf("You have clicked the button %d time(s)", counter))
	})

	btn.Show()

	qt.QApplication_Exec()

	fmt.Println("OK!")
}
