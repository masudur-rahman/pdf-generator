package main

import (
	"fmt"
	pdf "github.com/adrg/go-wkhtmltopdf"
	"github.com/codeskyblue/go-sh"
)

func main() {
	htmlPath := "/tmp/pdf/simple.html"
	title := fmt.Sprintf("Invoice for the month of May")
	pdfPath := "/tmp/pdf/simple.pdf"

	err := sh.Command("wkhtmltopdf", "--title", title, "--page-size", string(pdf.Legal), "--orientation", string(pdf.Portrait), htmlPath, pdfPath).Run()
	if err != nil {
		panic(err)
	}
}
