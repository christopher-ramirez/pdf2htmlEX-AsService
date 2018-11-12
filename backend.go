package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// TemporalFileName makes and return a random filename in the OS temp directory
func TemporalFileName(extension string) string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)

	return filepath.Join(
		os.TempDir(),
		fmt.Sprintf("%s.%s", hex.EncodeToString(randBytes), extension))
}

// ConvertPDFToHTML converts a PDF file into HTML returning the final HTML
func ConvertPDFToHTML(pdfFile string) string {
	output := TemporalFileName("html")
	fmt.Printf("Trying to convert PDF %s to HTML %s\n", pdfFile, output)
	cmd := exec.Command("pdf2htmlEX", "--dest-dir", "", "--zoom", "0", "--process-outline", "0", pdfFile, output)
	co, err := cmd.CombinedOutput()

	if err != nil {
		log.Printf("Command finished with error: %v. %s", err, co)
	}

	html, err := ioutil.ReadFile(output)
	if err != nil {
		log.Printf("Could not read output HTML file %s. Error: %v", output, err)
	}

	err = os.Remove(output)
	if err != nil {
		log.Printf("Could not delete file %s", output)
	}

	return string(html)
}
