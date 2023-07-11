package tesseract

import (
	"fmt"
	"io/ioutil"
)

var (
	TextPSM = PSM_SINGLE_COLUMN
	TextOEM = OEM_LSTM_ONLY
)

// TextFromFile extract text from an image file
func TextFromFile(imagefile string) (string, error) {
	imagebytes, err := ioutil.ReadFile(imagefile)
	if err != nil {
		return "", err
	}

	return TextFromBytes(imagebytes)
}

// TextFromBytes extract text from an image (image file read into a byte buffer)
func TextFromBytes(imagebytes []byte) (string, error) {
	tess := BaseAPICreate()
	defer func() {
		tess.End() // relase Tesseract allocated data
	}()

	if ret := tess.Init2("", "eng", TextOEM); ret != 0 {
		return "", fmt.Errorf("tesseract init: %v", ret)
	}

	tess.SetPageSegMode(TextPSM)

	// imagebytes should be a byte buffer
	// containing an image in a format supported by leptonica (i.e. png)
	pbytes := tess.SetImageBytes(imagebytes)

	text := tess.GetUTF8Text()

	FreeImageBytes(pbytes) // release image data (pix converted from imagebytes)
	tess.Clear()           // release data allocated for OCR
	return text, nil
}
