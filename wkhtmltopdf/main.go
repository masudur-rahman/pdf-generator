package main

import (
	pdf "github.com/adrg/go-wkhtmltopdf"
	"github.com/masudur-rahman/pdf-generator/templates"
	"log"
	"os"
	"path/filepath"
)

func main() {
	for idx := 0; idx < 1; idx++ {
		if err := generatePDF(); err != nil {
			log.Fatal(err)
		}
		log.Printf("PDF #%d generation successfull\n\n", idx+1)
	}
}

func generatePDF() error {
	// Initialize library.
	if err := pdf.Init(); err != nil {
		return err
	}
	defer pdf.Destroy()

	// Create object from file.
	object, err := pdf.NewObject(filepath.Join(templates.Directory, "simple.html"))
	if err != nil {
		return err
	}
	defer object.Destroy()

	object.Header.ContentCenter = "[title]"
	object.Header.DisplaySeparator = true

	// Create object from URL.
	object2, err := pdf.NewObject("https://google.com")
	if err != nil {
		return err
	}
	defer object2.Destroy()

	object.Footer.ContentLeft = "[date]"
	object.Footer.ContentCenter = "Sample footer information"
	object.Footer.ContentRight = "[page]"
	object.Footer.DisplaySeparator = true

	// Create converter.
	converter, err := pdf.NewConverter()
	if err != nil {
		return err
	}
	defer converter.Destroy()

	// Add created objects to the converter.
	converter.Add(object)
	converter.Add(object2)

	// Set converter options.
	converter.Title = "Invoice"
	converter.PaperSize = pdf.Legal
	converter.Orientation = pdf.Landscape

	//converter.MarginTop = "0cm"
	//converter.MarginBottom = "0cm"
	converter.MarginLeft = "0mm"
	converter.MarginRight = "0mm"

	// Convert objects and save the output PDF document.
	outFile, err := os.Create(filepath.Join(templates.Directory, "/simple.pdf"))
	if err != nil {
		return err
	}
	defer outFile.Close()

	if err := converter.Run(outFile); err != nil {
		return err
	}

	return nil
}
