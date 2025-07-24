package main

import (
	"os"

	"github.com/e1z0/miqt/qt"
	"github.com/e1z0/miqt/qt/pdf"
)

func main() {

	qt.NewQApplication(os.Args)

	d := pdf.NewQPdfDocument()
	d.Load("example.pdf")

	p := pdf.NewQPdfView2()
	p.SetWindowTitle("Qt 5 PDF Example")
	p.SetMinimumSize2(650, 600)
	p.SetPageMode(pdf.QPdfView__MultiPage)
	p.SetZoomMode(pdf.QPdfView__FitInView)
	p.SetDocument(d)

	p.Show()

	qt.QApplication_Exec()
}
