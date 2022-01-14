package tesseract

// PageSegMode represents tesseract::PageSegMode.
const (
	// PSM_OSD_ONLY - Orientation and script detection (OSD) only.
	PSM_OSD_ONLY PageSegMode = iota
	// PSM_AUTO_OSD - Automatic page segmentation with OSD.
	PSM_AUTO_OSD
	// PSM_AUTO_ONLY - Automatic page segmentation, but no OSD, or OCR.
	PSM_AUTO_ONLY
	// PSM_AUTO - (DEFAULT) Fully automatic page segmentation, but no OSD.
	PSM_AUTO
	// PSM_SINGLE_COLUMN - Assume a single column of text of variable sizes.
	PSM_SINGLE_COLUMN
	// PSM_SINGLE_BLOCK_VERT_TEXT - Assume a single uniform block of vertically aligned text.
	PSM_SINGLE_BLOCK_VERT_TEXT
	// PSM_SINGLE_BLOCK - Assume a single uniform block of text.
	PSM_SINGLE_BLOCK
	// PSM_SINGLE_LINE - Treat the image as a single text line.
	PSM_SINGLE_LINE
	// PSM_SINGLE_WORD - Treat the image as a single word.
	PSM_SINGLE_WORD
	// PSM_CIRCLE_WORD - Treat the image as a single word in a circle.
	PSM_CIRCLE_WORD
	// PSM_SINGLE_CHAR - Treat the image as a single character.
	PSM_SINGLE_CHAR
	// PSM_SPARSE_TEXT - Find as much text as possible in no particular order.
	PSM_SPARSE_TEXT
	// PSM_SPARSE_TEXT_OSD - Sparse text with orientation and script det.
	PSM_SPARSE_TEXT_OSD
	// PSM_RAW_LINE - Treat the image as a single text line, bypassing hacks that are Tesseract-specific.
	PSM_RAW_LINE

	// PSM_COUNT - Just a number of enum entries. This is NOT a member of PSM ;)
	PSM_COUNT
)

// OcrEngineMode represents tesseract::OcrEngineMode.
const (
	// OEM_TESSERACT_ONLY - Run Tesseract only - fastest; deprecated
	OEM_TESSERACT_ONLY OcrEngineMode = iota
	// OEM_LSTM_ONLY - Run just the LSTM line recognizer.
	OEM_LSTM_ONLY
	// OEM_TESSERACT_LSTM_COMBINED - Run the LSTM recognizer, but allow fallback
	// to Tesseract when things get difficult.
	// deprecated
	OEM_TESSERACT_LSTM_COMBINED
	// OEM_DEFAULT - Specify this mode when calling init_*(),
	// to indicate that any of the above modes
	// should be automatically inferred from the
	// variables in the language-specific config,
	// command-line configs, or if not specified
	// in any of the above should be set to the
	// default OEM_TESSERACT_ONLY.
	OEM_DEFAULT
	// OEM_COUNT - Number of OEMs. This is NOT a member of OEM
	OEM_COUNT
)

// PageIteratorLevel maps directly to tesseracts enum tesseract::PageIteratorLevel
const (
	// RIL_BLOCK - Block of text/image/separator line.
	RIL_BLOCK PageIteratorLevel = iota
	// RIL_PARA - Paragraph within a block.
	RIL_PARA
	// RIL_TEXTLINE - Line within a paragraph.
	RIL_TEXTLINE
	// RIL_WORD - Word within a textline.
	RIL_WORD
	// RIL_SYMBOL - Symbol/character within a word.
	RIL_SYMBOL
)

