package main

import "fmt"

// iVisitor interface
type iVisitor interface {
	visitForPdf(*pdfFile)
	visitForPpt(*pptFile)
	visitForWord(*wordFile)
}

// extractor concrete visitor
type extractor struct {
}

func (e *extractor) visitForPdf(p *pdfFile) {
	fmt.Println("Extract PDF.")
}

func (e *extractor) visitForPpt(p *pptFile) {
	fmt.Println("Extract PPT.")
}

func (e *extractor) visitForWord(w *wordFile) {
	fmt.Println("Extract WORD.")
}

// compressor concrete visitor
type compressor struct {
}

func (c *compressor) visitForPdf(p *pdfFile) {
	fmt.Println("Compress PDF.")
}

func (c *compressor) visitForPpt(p *pptFile) {
	fmt.Println("Compress PPT.")
}

func (c *compressor) visitForWord(w *wordFile) {
	fmt.Println("Compress WORD.")
}

// iResourceFile interface
type iResourceFile interface {
	getType() string
	accept(iVisitor)
}

// pptFile concrete resourceFile
type pptFile struct {
	filePath string
}

func (p *pptFile) getType() string {
	return "ppt"
}

func (p *pptFile) accept(v iVisitor) {
	v.visitForPpt(p)
}

// pdfFile concrete resourceFile
type pdfFile struct {
	filePath string
}

func (p *pdfFile) getType() string {
	return "pdf"
}

func (p *pdfFile) accept(v iVisitor) {
	v.visitForPdf(p)
}

// wordFile concrete resourceFile
type wordFile struct {
	filePath string
}

func (w *wordFile) getType() string {
	return "ppt"
}

func (w *wordFile) accept(v iVisitor) {
	v.visitForWord(w)
}

// listAllResourceFiles
func listAllResourceFiles() []iResourceFile {
	var resourceFiles []iResourceFile
	resourceFiles = append(resourceFiles, &pdfFile{"a.pdf"})
	resourceFiles = append(resourceFiles, &wordFile{"b.word"})
	resourceFiles = append(resourceFiles, &pptFile{"c.ppt"})
	return resourceFiles
}

// main
func main() {
	extractor := &extractor{}
	resourceFiles := listAllResourceFiles()

	for _, resouceFile := range resourceFiles {
		resouceFile.accept(extractor)
	}

	compressor := &compressor{}
	for _, resouceFile := range resourceFiles {
		resouceFile.accept(compressor)
	}

}
