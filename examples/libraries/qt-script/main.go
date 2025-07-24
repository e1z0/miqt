package main

import (
	"fmt"
	"os"

	"github.com/e1z0/miqt/qt"
	"github.com/e1z0/miqt/qt/script"
)

func main() {

	qt.NewQApplication(os.Args)

	inputProgram := "1 + 2"

	eng := script.NewQScriptEngine()
	result := eng.Evaluate(inputProgram)

	fmt.Printf("%s = %1.f\n", inputProgram, result.ToNumber())

	// qt.QApplication_Exec()
}
