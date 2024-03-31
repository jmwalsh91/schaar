package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func main() {
	inputFile := flag.String("input", "", "Path to the input PDF file")
	outputDir := flag.String("output", "output", "Output directory for split PDF files")
	splitPages := flag.Int("pages", 1, "Number of pages per split PDF")

	flag.Parse()

	if *inputFile == "" {
		log.Fatal("Please provide the path to the input PDF file using the --input flag")
	}

	outputFilePattern := filepath.Join(*outputDir, fmt.Sprintf("%%d_%d.pdf", *splitPages))

	conf := model.NewDefaultConfiguration()
	err := api.SplitFile(*inputFile, outputFilePattern, *splitPages, conf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PDF splitting completed successfully!")
}
