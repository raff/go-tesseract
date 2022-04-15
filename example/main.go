package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/raff/go-tesseract"
)

func main() {
	imagebytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	tess := tesseract.BaseAPICreate()
	defer func() {
		tess.End()  // relase Tesseract allocated data
	}()

	if ret := tess.Init3("", "eng"); ret != 0 {
		log.Fatal("Tesseract init", ret)
	}

	// imagebytes should be a byte buffer
	// containing an image in a format supported by leptonica (i.e. png)
	pbytes := tess.SetImageBytes(imagebytes)

	if text := tess.GetUTF8Text(); text != "" {
		fmt.Println(text)
	}

	tesseract.FreeImageBytes(pbytes)    // release image data (pix converted from imagebytes)
        tesseract.Clear()                   // release data allocated for OCR
}
