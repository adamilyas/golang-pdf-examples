package main

import (
	"fmt"
	"os"

	"github.com/ledongthuc/pdf"
	"github.com/phpdave11/gofpdf"
	"github.com/phpdave11/gofpdf/contrib/gofpdi"
)

func main() {

	inputPath, outputDir := "your_pdf.pdf", "output" // TODO: CHANGE ME

	SplitPdf(inputPath, outputDir)

}

// SplitPdf
func SplitPdf(inputPath string, outputDir string) {

	if !checkFileExists(inputPath) {
		panic(fmt.Sprintf("%s does not exist", inputPath))
	}

	if !isPdf(inputPath) {
		panic(fmt.Sprintf("%s is not a pdf", inputPath))
	}

	if !checkFileExists(outputDir) {
		err := os.Mkdir(outputDir, 0775)
		if err != nil {
			panic("Error occurred while creating output directory")
		}
	}

	_, inputReader, err := pdf.Open(inputPath)
	if err != nil {
		panic("Error occurred while opening file, ")
	}
	pageCount := inputReader.NumPage()

	pdfImporter := gofpdi.NewImporter()

	for i := 1; i <= pageCount; i++ {

		pageNo := i

		outputPath := fmt.Sprintf("%v/%v.pdf", outputDir, pageNo)

		if i%20 == 0 {
			//fmt.Println("Extracting to ", outputPath)
			pdfImporter = gofpdi.NewImporter() // faster
		}

		// extract page and output
		newPdf := gofpdf.New("P", "mm", "A4", "")

		// Import example-pdf.pdf with gofpdi free pdf document importer
		templateID := pdfImporter.ImportPage(newPdf, inputPath, pageNo, "/MediaBox")

		newPdf.AddPage()
		newPdf.SetFillColor(255, 255, 255)

		// Draw imported template onto page
		pdfImporter.UseImportedTemplate(newPdf, templateID, 0, 0, 200, 0)

		err := newPdf.OutputFileAndClose(outputPath)
		if err != nil {
			fmt.Println("Error saving ", outputPath)
		}
	}
}

func checkFileExists(inputPath string) bool {
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		// path/to/whatever does not exist
		return false
	}
	return true
}

func isPdf(inputPath string) bool {
	if inputPath[len(inputPath)-4:] != ".pdf" {
		return false
	}
	return true
}
