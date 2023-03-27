package pdfgen

import (
	"bytes"
	"html/template"
	"io"
	"os"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type DemoData struct {
	Name string
	Age  int
}

func PanicError(err error) (exists bool) {
	if err != nil {
		exists = true
		panic(err.Error())
	}
	return
}

func GeneratePdf() {
	tmpl, err := template.ParseFiles("template/invoice.html")
	PanicError(err)
	var buf bytes.Buffer

	err = tmpl.Execute(&buf, DemoData{"Beyond", 12})
	PanicError(err)

	pdfG, err := wkhtmltopdf.NewPDFGenerator()
	PanicError(err)

	page := wkhtmltopdf.NewPageReader(bytes.NewReader(buf.Bytes()))
	page.EnableLocalFileAccess.Set(true)
	pdfG.AddPage(page)
	err = pdfG.Create()
	PanicError(err)

	file, err := os.Create("assets/test.pdf")
	PanicError(err)
	defer file.Close()

	_, err = io.Copy(file, bytes.NewReader(pdfG.Bytes()))
	PanicError(err)
}
