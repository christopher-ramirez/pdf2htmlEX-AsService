package main

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

/**
 * Declare the Web service handler. This is a frontend service
 * that receives PDF files and convert them to HTML using PDF2HTMLEx
 * returning the generated HTML source.
 */
func main() {
	println("Starting server at :8080...")
	http.HandleFunc("/", transformer)
	http.HandleFunc("/healthcheck", healthcheck)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func transformer(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(405)
		fmt.Fprintf(w, "This method is not allowed")
		return
	}

	r.ParseMultipartForm(32 << 20)
	uploaded, _, err := r.FormFile("pdf")
	if err != nil && err == http.ErrMissingFile {
		w.WriteHeader(400)
		fmt.Fprintf(w, "%v", err)
		return
	}
	defer uploaded.Close()

	// Make a copy of the received pdf to a temporary file
	tempFile, err := uploadedContentsToTempFile(uploaded)
	if err != nil {
		log.Printf("Could not store uploaded contents. %v", err)
		fmt.Fprintf(w, "Could not store uploaded contents. %v", err)
		return
	}

	// ConvertPDFToHTML uses PDF2HTMLEx to convert the uploaded file
	HTMLOutput := ConvertPDFToHTML(tempFile)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, HTMLOutput)

	// Free resources
	err = os.Remove(tempFile)
	if err != nil {
		log.Printf("Could not delete file %s", tempFile)
	}
}

func uploadedContentsToTempFile(uploadedFile multipart.File) (string, error) {
	tempFile := TemporalFileName("pdf")
	finalFile, err := os.Create(tempFile)

	if err != nil {
		return "", err
	}
	io.Copy(finalFile, uploadedFile)
	finalFile.Close()
	return tempFile, nil
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "healthy")
}
