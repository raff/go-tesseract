package tesseract

/*
#include <leptonica/allheaders.h>
#include <tesseract/capi.h>
#include <stdlib.h>
#cgo LDFLAGS: -llept -ltesseract

typedef int Bool;
typedef struct Pix Pix;
typedef struct Pixa Pixa;
*/
import "C"

import (
	"unsafe"
)

// TessBaseAPI
type BaseAPI struct {
}

// TessBaseAPIAdaptToWordStr
func (handle *BaseAPI) AdaptToWordStr(mode PageSegMode, wordstr string) (ret bool) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_mode := *(*C.TessPageSegMode)(unsafe.Pointer(&mode))
	_wordstr := C.CString(wordstr)
	defer C.free(unsafe.Pointer(_wordstr))
	_ret := C.TessBaseAPIAdaptToWordStr(_handle, _mode, _wordstr)
	ret = bool(_ret == 1)
	return
}

func (handle *BaseAPI) AllWordConfidences() (ret *int32) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_ret := C.TessBaseAPIAllWordConfidences(_handle)
	ret = (*int32)(unsafe.Pointer(_ret))
	return
}

// TessBaseAPIAnalyseLayout
func (handle *BaseAPI) AnalyseLayout() (ret *PageIterator) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_ret := C.TessBaseAPIAnalyseLayout(_handle)
	ret = (*PageIterator)(unsafe.Pointer(_ret))
	return
}

// TessBaseAPIClear
func (handle *BaseAPI) Clear() {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	C.TessBaseAPIClear(_handle)
}

// TessBaseAPIClearAdaptiveClassifier
func (handle *BaseAPI) ClearAdaptiveClassifier() {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	C.TessBaseAPIClearAdaptiveClassifier(_handle)
}

// TessBaseAPIClearPersistentCache
func (handle *BaseAPI) ClearPersistentCache() {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	C.TessBaseAPIClearPersistentCache(_handle)
}

// TessBaseAPIDelete
func (handle *BaseAPI) Delete() {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	C.TessBaseAPIDelete(_handle)
}

// TessBaseAPIDetectOrientationScript
func (handle *BaseAPI) DetectOrientationScript() (orientDeg int32, orientConf float32, scriptName *int8, scriptConf float32, ret bool) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_orientDeg := (*C.int)(unsafe.Pointer(&orientDeg))
	_orientConf := (*C.float)(unsafe.Pointer(&orientConf))
	_scriptName := (**C.char)(unsafe.Pointer(&scriptName))
	_scriptConf := (*C.float)(unsafe.Pointer(&scriptConf))
	_ret := C.TessBaseAPIDetectOrientationScript(_handle, _orientDeg, _orientConf, _scriptName, _scriptConf)
	ret = bool(_ret == 1)
	return
}

// TessBaseAPIEnd
func (handle *BaseAPI) End() {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	C.TessBaseAPIEnd(_handle)
}

// TessBaseAPIGetAltoText
func (handle *BaseAPI) GetAltoText(pageNumber int32) (ret string) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_pageNumber := C.int(pageNumber)
	_ret := C.TessBaseAPIGetAltoText(_handle, _pageNumber)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// TessBaseAPIGetAvailableLanguagesAsVector
func (handle *BaseAPI) GetAvailableLanguagesAsVector() (ret **int8) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_ret := C.TessBaseAPIGetAvailableLanguagesAsVector(_handle)
	ret = (**int8)(unsafe.Pointer(_ret))
	return
}

// TessBaseAPIGetBoolVariable
func (handle *BaseAPI) GetBoolVariable(name string) (value int32, ret bool) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_name := C.CString(name)
	defer C.free(unsafe.Pointer(_name))
	_value := (*C.Bool)(unsafe.Pointer(&value))
	_ret := C.TessBaseAPIGetBoolVariable(_handle, _name, _value)
	ret = bool(_ret == 1)
	return
}

// TessBaseAPIGetBoxText
func (handle *BaseAPI) GetBoxText(pageNumber int32) (ret string) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_pageNumber := C.int(pageNumber)
	_ret := C.TessBaseAPIGetBoxText(_handle, _pageNumber)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// TessBaseAPIGetComponentImages
func (handle *BaseAPI) GetComponentImages(level PageIteratorLevel, textOnly bool) (pixa *[0]byte, blockids *int32, ret *[0]byte) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_level := *(*C.TessPageIteratorLevel)(unsafe.Pointer(&level))
	_textOnly := C.Bool(0)
	if textOnly {
		_textOnly = C.Bool(1)
	}
	_pixa := (**C.Pixa)(unsafe.Pointer(&pixa))
	_blockids := (**C.int)(unsafe.Pointer(&blockids))
	_ret := C.TessBaseAPIGetComponentImages(_handle, _level, _textOnly, _pixa, _blockids)
	ret = (*[0]byte)(unsafe.Pointer(_ret))
	return
}

// TessBaseAPIGetComponentImages1
func (handle *BaseAPI) GetComponentImages1(level PageIteratorLevel, textOnly bool, rawImage bool, rawPadding int32) (pixa *[0]byte, blockids *int32, paraids *int32, ret *[0]byte) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_level := *(*C.TessPageIteratorLevel)(unsafe.Pointer(&level))
	_textOnly := C.Bool(0)
	if textOnly {
		_textOnly = C.Bool(1)
	}
	_rawImage := C.Bool(0)
	if rawImage {
		_rawImage = C.Bool(1)
	}
	_rawPadding := C.int(rawPadding)
	_pixa := (**C.Pixa)(unsafe.Pointer(&pixa))
	_blockids := (**C.int)(unsafe.Pointer(&blockids))
	_paraids := (**C.int)(unsafe.Pointer(&paraids))
	_ret := C.TessBaseAPIGetComponentImages1(_handle, _level, _textOnly, _rawImage, _rawPadding, _pixa, _blockids, _paraids)
	ret = (*[0]byte)(unsafe.Pointer(_ret))
	return
}

// TessBaseAPIGetConnectedComponents
func (handle *BaseAPI) GetConnectedComponents() (cc *[0]byte, ret *[0]byte) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_cc := (**C.Pixa)(unsafe.Pointer(&cc))
	_ret := C.TessBaseAPIGetConnectedComponents(_handle, _cc)
	ret = (*[0]byte)(unsafe.Pointer(_ret))
	return
}

// TessBaseAPIGetDatapath
func (handle *BaseAPI) GetDatapath() (ret string) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_ret := C.TessBaseAPIGetDatapath(_handle)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// TessBaseAPIGetDoubleVariable
func (handle *BaseAPI) GetDoubleVariable(name string) (value float64, ret bool) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_name := C.CString(name)
	defer C.free(unsafe.Pointer(_name))
	_value := (*C.double)(unsafe.Pointer(&value))
	_ret := C.TessBaseAPIGetDoubleVariable(_handle, _name, _value)
	ret = bool(_ret == 1)
	return
}

// TessBaseAPIGetHOCRText
func (handle *BaseAPI) GetHOCRText(pageNumber int32) (ret string) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_pageNumber := C.int(pageNumber)
	_ret := C.TessBaseAPIGetHOCRText(_handle, _pageNumber)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// TessBaseAPIGetInitLanguagesAsString
func (handle *BaseAPI) GetInitLanguagesAsString() (ret string) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_ret := C.TessBaseAPIGetInitLanguagesAsString(_handle)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// TessBaseAPIGetInputImage
func (handle *BaseAPI) GetInputImage() (ret *[0]byte) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_ret := C.TessBaseAPIGetInputImage(_handle)
	ret = (*[0]byte)(unsafe.Pointer(_ret))
	return
}

// TessBaseAPIGetInputName
func (handle *BaseAPI) GetInputName() (ret string) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_ret := C.TessBaseAPIGetInputName(_handle)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// TessBaseAPIGetIntVariable
func (handle *BaseAPI) GetIntVariable(name string) (value int32, ret bool) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_name := C.CString(name)
	defer C.free(unsafe.Pointer(_name))
	_value := (*C.int)(unsafe.Pointer(&value))
	_ret := C.TessBaseAPIGetIntVariable(_handle, _name, _value)
	ret = bool(_ret == 1)
	return
}

// TessBaseAPIGetIterator
func (handle *BaseAPI) GetIterator() (ret *ResultIterator) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_ret := C.TessBaseAPIGetIterator(_handle)
	ret = (*ResultIterator)(unsafe.Pointer(_ret))
	return
}

// TessBaseAPIGetLSTMBoxText
func (handle *BaseAPI) GetLSTMBoxText(pageNumber int32) (ret string) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_pageNumber := C.int(pageNumber)
	_ret := C.TessBaseAPIGetLSTMBoxText(_handle, _pageNumber)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// TessBaseAPIGetLoadedLanguagesAsVector
func (handle *BaseAPI) GetLoadedLanguagesAsVector() (ret **int8) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_ret := C.TessBaseAPIGetLoadedLanguagesAsVector(_handle)
	ret = (**int8)(unsafe.Pointer(_ret))
	return
}

// TessBaseAPIGetMutableIterator
func (handle *BaseAPI) GetMutableIterator() (ret *MutableIterator) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_ret := C.TessBaseAPIGetMutableIterator(_handle)
	ret = (*MutableIterator)(unsafe.Pointer(_ret))
	return
}

// TessBaseAPIGetOpenCLDevice
func (handle *BaseAPI) GetOpenCLDevice() (device uintptr, ret int) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_device := (*unsafe.Pointer)(unsafe.Pointer(&device))
	_ret := C.TessBaseAPIGetOpenCLDevice(_handle, _device)
	ret = int(_ret)
	return
}

// TessBaseAPIGetPageSegMode
func (handle *BaseAPI) GetPageSegMode() (ret PageSegMode) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_ret := C.TessBaseAPIGetPageSegMode(_handle)
	ret = *(*PageSegMode)(unsafe.Pointer(&_ret))
	return
}

// TessBaseAPIGetRegions
func (handle *BaseAPI) GetRegions() (pixa *[0]byte, ret *[0]byte) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_pixa := (**C.Pixa)(unsafe.Pointer(&pixa))
	_ret := C.TessBaseAPIGetRegions(_handle, _pixa)
	ret = (*[0]byte)(unsafe.Pointer(_ret))
	return
}

// TessBaseAPIGetSourceYResolution
func (handle *BaseAPI) GetSourceYResolution() (ret int32) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_ret := C.TessBaseAPIGetSourceYResolution(_handle)
	ret = int32(_ret)
	return
}

// TessBaseAPIGetStringVariable
func (handle *BaseAPI) GetStringVariable(name string) (ret string) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_name := C.CString(name)
	defer C.free(unsafe.Pointer(_name))
	_ret := C.TessBaseAPIGetStringVariable(_handle, _name)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// TessBaseAPIGetStrips
func (handle *BaseAPI) GetStrips() (pixa *[0]byte, blockids *int32, ret *[0]byte) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_pixa := (**C.Pixa)(unsafe.Pointer(&pixa))
	_blockids := (**C.int)(unsafe.Pointer(&blockids))
	_ret := C.TessBaseAPIGetStrips(_handle, _pixa, _blockids)
	ret = (*[0]byte)(unsafe.Pointer(_ret))
	return
}

// TessBaseAPIGetTextDirection
func (handle *BaseAPI) GetTextDirection() (outOffset int32, outSlope float32, ret bool) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_outOffset := (*C.int)(unsafe.Pointer(&outOffset))
	_outSlope := (*C.float)(unsafe.Pointer(&outSlope))
	_ret := C.TessBaseAPIGetTextDirection(_handle, _outOffset, _outSlope)
	ret = bool(_ret == 1)
	return
}

// TessBaseAPIGetTextlines
func (handle *BaseAPI) GetTextlines() (pixa *[0]byte, blockids *int32, ret *[0]byte) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_pixa := (**C.Pixa)(unsafe.Pointer(&pixa))
	_blockids := (**C.int)(unsafe.Pointer(&blockids))
	_ret := C.TessBaseAPIGetTextlines(_handle, _pixa, _blockids)
	ret = (*[0]byte)(unsafe.Pointer(_ret))
	return
}

// TessBaseAPIGetTextlines1
func (handle *BaseAPI) GetTextlines1(rawImage bool, rawPadding int32) (pixa *[0]byte, blockids *int32, paraids *int32, ret *[0]byte) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_rawImage := C.Bool(0)
	if rawImage {
		_rawImage = C.Bool(1)
	}
	_rawPadding := C.int(rawPadding)
	_pixa := (**C.Pixa)(unsafe.Pointer(&pixa))
	_blockids := (**C.int)(unsafe.Pointer(&blockids))
	_paraids := (**C.int)(unsafe.Pointer(&paraids))
	_ret := C.TessBaseAPIGetTextlines1(_handle, _rawImage, _rawPadding, _pixa, _blockids, _paraids)
	ret = (*[0]byte)(unsafe.Pointer(_ret))
	return
}

// TessBaseAPIGetThresholdedImage
func (handle *BaseAPI) GetThresholdedImage() (ret *[0]byte) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_ret := C.TessBaseAPIGetThresholdedImage(_handle)
	ret = (*[0]byte)(unsafe.Pointer(_ret))
	return
}

// TessBaseAPIGetThresholdedImageScaleFactor
func (handle *BaseAPI) GetThresholdedImageScaleFactor() (ret int32) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_ret := C.TessBaseAPIGetThresholdedImageScaleFactor(_handle)
	ret = int32(_ret)
	return
}

// TessBaseAPIGetTsvText
func (handle *BaseAPI) GetTsvText(pageNumber int32) (ret string) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_pageNumber := C.int(pageNumber)
	_ret := C.TessBaseAPIGetTsvText(_handle, _pageNumber)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// TessBaseAPIGetUNLVText
func (handle *BaseAPI) GetUNLVText() (ret string) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_ret := C.TessBaseAPIGetUNLVText(_handle)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// TessBaseAPIGetUTF8Text
func (handle *BaseAPI) GetUTF8Text() (ret string) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_ret := C.TessBaseAPIGetUTF8Text(_handle)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// TessBaseAPIGetUnichar
func (handle *BaseAPI) GetUnichar(unicharId int32) (ret string) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_unicharId := C.int(unicharId)
	_ret := C.TessBaseAPIGetUnichar(_handle, _unicharId)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// TessBaseAPIGetWordStrBoxText
func (handle *BaseAPI) GetWordStrBoxText(pageNumber int32) (ret string) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_pageNumber := C.int(pageNumber)
	_ret := C.TessBaseAPIGetWordStrBoxText(_handle, _pageNumber)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// TessBaseAPIGetWords
func (handle *BaseAPI) GetWords() (pixa *[0]byte, ret *[0]byte) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_pixa := (**C.Pixa)(unsafe.Pointer(&pixa))
	_ret := C.TessBaseAPIGetWords(_handle, _pixa)
	ret = (*[0]byte)(unsafe.Pointer(_ret))
	return
}

// TessBaseAPIInit1
func (handle *BaseAPI) Init1(datapath string, language string, oem OcrEngineMode, configsSize int32) (configs *int8, ret int32) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_datapath := C.CString(datapath)
	defer C.free(unsafe.Pointer(_datapath))
	_language := C.CString(language)
	defer C.free(unsafe.Pointer(_language))
	_oem := *(*C.TessOcrEngineMode)(unsafe.Pointer(&oem))
	_configs := (**C.char)(unsafe.Pointer(&configs))
	_configsSize := C.int(configsSize)
	_ret := C.TessBaseAPIInit1(_handle, _datapath, _language, _oem, _configs, _configsSize)
	ret = int32(_ret)
	return
}

// TessBaseAPIInit2
func (handle *BaseAPI) Init2(datapath string, language string, oem OcrEngineMode) (ret int32) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_datapath := C.CString(datapath)
	defer C.free(unsafe.Pointer(_datapath))
	_language := C.CString(language)
	defer C.free(unsafe.Pointer(_language))
	_oem := *(*C.TessOcrEngineMode)(unsafe.Pointer(&oem))
	_ret := C.TessBaseAPIInit2(_handle, _datapath, _language, _oem)
	ret = int32(_ret)
	return
}

// TessBaseAPIInit3
func (handle *BaseAPI) Init3(datapath string, language string) (ret int32) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_datapath := C.CString(datapath)
	defer C.free(unsafe.Pointer(_datapath))
	_language := C.CString(language)
	defer C.free(unsafe.Pointer(_language))
	_ret := C.TessBaseAPIInit3(_handle, _datapath, _language)
	ret = int32(_ret)
	return
}

// TessBaseAPIInit4
func (handle *BaseAPI) Init4(datapath string, language string, mode OcrEngineMode, configsSize int32, varsVecSize int, setOnlyNonDebugParams bool) (configs *int8, varsVec *int8, varsValues *int8, ret int32) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_datapath := C.CString(datapath)
	defer C.free(unsafe.Pointer(_datapath))
	_language := C.CString(language)
	defer C.free(unsafe.Pointer(_language))
	_mode := *(*C.TessOcrEngineMode)(unsafe.Pointer(&mode))
	_configs := (**C.char)(unsafe.Pointer(&configs))
	_configsSize := C.int(configsSize)
	_varsVec := (**C.char)(unsafe.Pointer(&varsVec))
	_varsValues := (**C.char)(unsafe.Pointer(&varsValues))
	_varsVecSize := C.size_t(varsVecSize)
	_setOnlyNonDebugParams := C.Bool(0)
	if setOnlyNonDebugParams {
		_setOnlyNonDebugParams = C.Bool(1)
	}
	_ret := C.TessBaseAPIInit4(_handle, _datapath, _language, _mode, _configs, _configsSize, _varsVec, _varsValues, _varsVecSize, _setOnlyNonDebugParams)
	ret = int32(_ret)
	return
}

// TessBaseAPIInitForAnalysePage
func (handle *BaseAPI) InitForAnalysePage() {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	C.TessBaseAPIInitForAnalysePage(_handle)
}

// TessBaseAPIIsValidWord
func (handle *BaseAPI) IsValidWord(word string) (ret int32) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_word := C.CString(word)
	defer C.free(unsafe.Pointer(_word))
	_ret := C.TessBaseAPIIsValidWord(_handle, _word)
	ret = int32(_ret)
	return
}

// TessBaseAPIMeanTextConf
func (handle *BaseAPI) MeanTextConf() (ret int32) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_ret := C.TessBaseAPIMeanTextConf(_handle)
	ret = int32(_ret)
	return
}

// TessBaseAPINumDawgs
func (handle *BaseAPI) NumDawgs() (ret int32) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_ret := C.TessBaseAPINumDawgs(_handle)
	ret = int32(_ret)
	return
}

// TessBaseAPIOem
func (handle *BaseAPI) Oem() (ret OcrEngineMode) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_ret := C.TessBaseAPIOem(_handle)
	ret = *(*OcrEngineMode)(unsafe.Pointer(&_ret))
	return
}

// TessBaseAPIPrintVariables
func (handle *BaseAPI) PrintVariables(fp *[152]byte) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_fp := (*C.FILE)(unsafe.Pointer(fp))
	C.TessBaseAPIPrintVariables(_handle, _fp)
}

// TessBaseAPIPrintVariablesToFile
func (handle *BaseAPI) PrintVariablesToFile(filename string) (ret bool) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(_filename))
	_ret := C.TessBaseAPIPrintVariablesToFile(_handle, _filename)
	ret = bool(_ret == 1)
	return
}

// TessBaseAPIProcessPage
func (handle *BaseAPI) ProcessPage(pix *[0]byte, pageIndex int32, filename string, retryConfig string, timeoutMillisec int32, renderer *ResultRenderer) (ret bool) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_pix := (*C.Pix)(unsafe.Pointer(pix))
	_pageIndex := C.int(pageIndex)
	_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(_filename))
	_retryConfig := C.CString(retryConfig)
	defer C.free(unsafe.Pointer(_retryConfig))
	_timeoutMillisec := C.int(timeoutMillisec)
	_renderer := (*C.TessResultRenderer)(unsafe.Pointer(renderer))
	_ret := C.TessBaseAPIProcessPage(_handle, _pix, _pageIndex, _filename, _retryConfig, _timeoutMillisec, _renderer)
	ret = bool(_ret == 1)
	return
}

// TessBaseAPIProcessPages
func (handle *BaseAPI) ProcessPages(filename string, retryConfig string, timeoutMillisec int32, renderer *ResultRenderer) (ret bool) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(_filename))
	_retryConfig := C.CString(retryConfig)
	defer C.free(unsafe.Pointer(_retryConfig))
	_timeoutMillisec := C.int(timeoutMillisec)
	_renderer := (*C.TessResultRenderer)(unsafe.Pointer(renderer))
	_ret := C.TessBaseAPIProcessPages(_handle, _filename, _retryConfig, _timeoutMillisec, _renderer)
	ret = bool(_ret == 1)
	return
}

// TessBaseAPIReadConfigFile
func (handle *BaseAPI) ReadConfigFile(filename string) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(_filename))
	C.TessBaseAPIReadConfigFile(_handle, _filename)
}

// TessBaseAPIReadDebugConfigFile
func (handle *BaseAPI) ReadDebugConfigFile(filename string) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(_filename))
	C.TessBaseAPIReadDebugConfigFile(_handle, _filename)
}

// TessBaseAPIRecognize
func (handle *BaseAPI) Recognize(monitor *[0]byte) (ret int32) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_monitor := (*C.ETEXT_DESC)(unsafe.Pointer(monitor))
	_ret := C.TessBaseAPIRecognize(_handle, _monitor)
	ret = int32(_ret)
	return
}

// TessBaseAPIRect
func (handle *BaseAPI) Rect(imagedata *byte, bytesPerPixel int32, bytesPerLine int32, left int32, top int32, width int32, height int32) (ret string) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_imagedata := (*C.uchar)(unsafe.Pointer(imagedata))
	_bytesPerPixel := C.int(bytesPerPixel)
	_bytesPerLine := C.int(bytesPerLine)
	_left := C.int(left)
	_top := C.int(top)
	_width := C.int(width)
	_height := C.int(height)
	_ret := C.TessBaseAPIRect(_handle, _imagedata, _bytesPerPixel, _bytesPerLine, _left, _top, _width, _height)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// TessBaseAPISetDebugVariable
func (handle *BaseAPI) SetDebugVariable(name string, value string) (ret bool) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_name := C.CString(name)
	defer C.free(unsafe.Pointer(_name))
	_value := C.CString(value)
	defer C.free(unsafe.Pointer(_value))
	_ret := C.TessBaseAPISetDebugVariable(_handle, _name, _value)
	ret = bool(_ret == 1)
	return
}

// TessBaseAPISetImage
func (handle *BaseAPI) SetImage(imagedata *byte, width int32, height int32, bytesPerPixel int32, bytesPerLine int32) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_imagedata := (*C.uchar)(unsafe.Pointer(imagedata))
	_width := C.int(width)
	_height := C.int(height)
	_bytesPerPixel := C.int(bytesPerPixel)
	_bytesPerLine := C.int(bytesPerLine)
	C.TessBaseAPISetImage(_handle, _imagedata, _width, _height, _bytesPerPixel, _bytesPerLine)
}

// TessBaseAPISetImage2
func (handle *BaseAPI) SetImage2(pix *[0]byte) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_pix := (*C.Pix)(unsafe.Pointer(pix))
	C.TessBaseAPISetImage2(_handle, _pix)
}

// TessBaseAPISetInputImage
func (handle *BaseAPI) SetInputImage(pix *[0]byte) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_pix := (*C.Pix)(unsafe.Pointer(pix))
	C.TessBaseAPISetInputImage(_handle, _pix)
}

// TessBaseAPISetInputName
func (handle *BaseAPI) SetInputName(name string) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_name := C.CString(name)
	defer C.free(unsafe.Pointer(_name))
	C.TessBaseAPISetInputName(_handle, _name)
}

// TessBaseAPISetMinOrientationMargin
func (handle *BaseAPI) SetMinOrientationMargin(margin float64) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_margin := C.double(margin)
	C.TessBaseAPISetMinOrientationMargin(_handle, _margin)
}

// TessBaseAPISetOutputName
func (handle *BaseAPI) SetOutputName(name string) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_name := C.CString(name)
	defer C.free(unsafe.Pointer(_name))
	C.TessBaseAPISetOutputName(_handle, _name)
}

// TessBaseAPISetPageSegMode
func (handle *BaseAPI) SetPageSegMode(mode PageSegMode) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_mode := *(*C.TessPageSegMode)(unsafe.Pointer(&mode))
	C.TessBaseAPISetPageSegMode(_handle, _mode)
}

// TessBaseAPISetRectangle
func (handle *BaseAPI) SetRectangle(left int32, top int32, width int32, height int32) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_left := C.int(left)
	_top := C.int(top)
	_width := C.int(width)
	_height := C.int(height)
	C.TessBaseAPISetRectangle(_handle, _left, _top, _width, _height)
}

// TessBaseAPISetSourceResolution
func (handle *BaseAPI) SetSourceResolution(ppi int32) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_ppi := C.int(ppi)
	C.TessBaseAPISetSourceResolution(_handle, _ppi)
}

// TessBaseAPISetVariable
func (handle *BaseAPI) SetVariable(name string, value string) (ret bool) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_name := C.CString(name)
	defer C.free(unsafe.Pointer(_name))
	_value := C.CString(value)
	defer C.free(unsafe.Pointer(_value))
	_ret := C.TessBaseAPISetVariable(_handle, _name, _value)
	ret = bool(_ret == 1)
	return
}

// TessBaseGetBlockTextOrientations
/*
func (handle *BaseAPI) BaseGetBlockTextOrientations() (blockOrientation *int32, verticalWriting uintptr) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_blockOrientation := (**C.int)(unsafe.Pointer(&blockOrientation))
	_verticalWriting := (*unsafe.Pointer)(unsafe.Pointer(&verticalWriting))
	C.TessBaseGetBlockTextOrientations(_handle, _blockOrientation, _verticalWriting)
	return
}
*/

// TessCancelFunc
type CancelFunc *[0]byte

// TessChoiceIterator
type ChoiceIterator struct {
}

// TessChoiceIteratorConfidence
func (handle *ChoiceIterator) Confidence() (ret float32) {
	_handle := (*C.TessChoiceIterator)(unsafe.Pointer(handle))
	_ret := C.TessChoiceIteratorConfidence(_handle)
	ret = float32(_ret)
	return
}

// TessChoiceIteratorDelete
func (handle *ChoiceIterator) Delete() {
	_handle := (*C.TessChoiceIterator)(unsafe.Pointer(handle))
	C.TessChoiceIteratorDelete(_handle)
}

// TessChoiceIteratorGetUTF8Text
func (handle *ChoiceIterator) GetUTF8Text() (ret string) {
	_handle := (*C.TessChoiceIterator)(unsafe.Pointer(handle))
	_ret := C.TessChoiceIteratorGetUTF8Text(_handle)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// TessChoiceIteratorNext
func (handle *ChoiceIterator) Next() (ret bool) {
	_handle := (*C.TessChoiceIterator)(unsafe.Pointer(handle))
	_ret := C.TessChoiceIteratorNext(_handle)
	ret = bool(_ret == 1)
	return
}

// TessMutableIterator
type MutableIterator struct {
}

// TessOcrEngineMode
type OcrEngineMode int32

// TessOrientation
type Orientation int32

// TessPageIterator
type PageIterator struct {
}

// TessPageIteratorBaseline
func (handle *PageIterator) Baseline(level PageIteratorLevel) (x1 int32, y1 int32, x2 int32, y2 int32, ret bool) {
	_handle := (*C.TessPageIterator)(unsafe.Pointer(handle))
	_level := *(*C.TessPageIteratorLevel)(unsafe.Pointer(&level))
	_x1 := (*C.int)(unsafe.Pointer(&x1))
	_y1 := (*C.int)(unsafe.Pointer(&y1))
	_x2 := (*C.int)(unsafe.Pointer(&x2))
	_y2 := (*C.int)(unsafe.Pointer(&y2))
	_ret := C.TessPageIteratorBaseline(_handle, _level, _x1, _y1, _x2, _y2)
	ret = bool(_ret == 1)
	return
}

// TessPageIteratorBegin
func (handle *PageIterator) Begin() {
	_handle := (*C.TessPageIterator)(unsafe.Pointer(handle))
	C.TessPageIteratorBegin(_handle)
}

// TessPageIteratorBlockType
func (handle *PageIterator) BlockType() (ret PolyBlockType) {
	_handle := (*C.TessPageIterator)(unsafe.Pointer(handle))
	_ret := C.TessPageIteratorBlockType(_handle)
	ret = *(*PolyBlockType)(unsafe.Pointer(&_ret))
	return
}

// TessPageIteratorBoundingBox
func (handle *PageIterator) BoundingBox(level PageIteratorLevel) (left int32, top int32, right int32, bottom int32, ret bool) {
	_handle := (*C.TessPageIterator)(unsafe.Pointer(handle))
	_level := *(*C.TessPageIteratorLevel)(unsafe.Pointer(&level))
	_left := (*C.int)(unsafe.Pointer(&left))
	_top := (*C.int)(unsafe.Pointer(&top))
	_right := (*C.int)(unsafe.Pointer(&right))
	_bottom := (*C.int)(unsafe.Pointer(&bottom))
	_ret := C.TessPageIteratorBoundingBox(_handle, _level, _left, _top, _right, _bottom)
	ret = bool(_ret == 1)
	return
}

// TessPageIteratorCopy
func (handle *PageIterator) Copy() (ret *PageIterator) {
	_handle := (*C.TessPageIterator)(unsafe.Pointer(handle))
	_ret := C.TessPageIteratorCopy(_handle)
	ret = (*PageIterator)(unsafe.Pointer(_ret))
	return
}

// TessPageIteratorDelete
func (handle *PageIterator) Delete() {
	_handle := (*C.TessPageIterator)(unsafe.Pointer(handle))
	C.TessPageIteratorDelete(_handle)
}

// TessPageIteratorGetBinaryImage
func (handle *PageIterator) GetBinaryImage(level PageIteratorLevel) (ret *[0]byte) {
	_handle := (*C.TessPageIterator)(unsafe.Pointer(handle))
	_level := *(*C.TessPageIteratorLevel)(unsafe.Pointer(&level))
	_ret := C.TessPageIteratorGetBinaryImage(_handle, _level)
	ret = (*[0]byte)(unsafe.Pointer(_ret))
	return
}

// TessPageIteratorGetImage
func (handle *PageIterator) GetImage(level PageIteratorLevel, padding int32, originalImage *[0]byte) (left int32, top int32, ret *[0]byte) {
	_handle := (*C.TessPageIterator)(unsafe.Pointer(handle))
	_level := *(*C.TessPageIteratorLevel)(unsafe.Pointer(&level))
	_padding := C.int(padding)
	_originalImage := (*C.Pix)(unsafe.Pointer(originalImage))
	_left := (*C.int)(unsafe.Pointer(&left))
	_top := (*C.int)(unsafe.Pointer(&top))
	_ret := C.TessPageIteratorGetImage(_handle, _level, _padding, _originalImage, _left, _top)
	ret = (*[0]byte)(unsafe.Pointer(_ret))
	return
}

// TessPageIteratorIsAtBeginningOf
func (handle *PageIterator) IsAtBeginningOf(level PageIteratorLevel) (ret bool) {
	_handle := (*C.TessPageIterator)(unsafe.Pointer(handle))
	_level := *(*C.TessPageIteratorLevel)(unsafe.Pointer(&level))
	_ret := C.TessPageIteratorIsAtBeginningOf(_handle, _level)
	ret = bool(_ret == 1)
	return
}

// TessPageIteratorIsAtFinalElement
func (handle *PageIterator) IsAtFinalElement(level PageIteratorLevel, element PageIteratorLevel) (ret bool) {
	_handle := (*C.TessPageIterator)(unsafe.Pointer(handle))
	_level := *(*C.TessPageIteratorLevel)(unsafe.Pointer(&level))
	_element := *(*C.TessPageIteratorLevel)(unsafe.Pointer(&element))
	_ret := C.TessPageIteratorIsAtFinalElement(_handle, _level, _element)
	ret = bool(_ret == 1)
	return
}

// TessPageIteratorNext
func (handle *PageIterator) Next(level PageIteratorLevel) (ret bool) {
	_handle := (*C.TessPageIterator)(unsafe.Pointer(handle))
	_level := *(*C.TessPageIteratorLevel)(unsafe.Pointer(&level))
	_ret := C.TessPageIteratorNext(_handle, _level)
	ret = bool(_ret == 1)
	return
}

// TessPageIteratorOrientation
func (handle *PageIterator) Orientation(orientation *Orientation, writingDirection *WritingDirection, textlineOrder *TextlineOrder) (deskewAngle float32) {
	_handle := (*C.TessPageIterator)(unsafe.Pointer(handle))
	_orientation := (*C.TessOrientation)(unsafe.Pointer(orientation))
	_writingDirection := (*C.TessWritingDirection)(unsafe.Pointer(writingDirection))
	_textlineOrder := (*C.TessTextlineOrder)(unsafe.Pointer(textlineOrder))
	_deskewAngle := (*C.float)(unsafe.Pointer(&deskewAngle))
	C.TessPageIteratorOrientation(_handle, _orientation, _writingDirection, _textlineOrder, _deskewAngle)
	return
}

// TessPageIteratorParagraphInfo
func (handle *PageIterator) ParagraphInfo(justification *ParagraphJustification) (isListItem int32, isCrown int32, firstLineIndent int32) {
	_handle := (*C.TessPageIterator)(unsafe.Pointer(handle))
	_justification := (*C.TessParagraphJustification)(unsafe.Pointer(justification))
	_isListItem := (*C.Bool)(unsafe.Pointer(&isListItem))
	_isCrown := (*C.Bool)(unsafe.Pointer(&isCrown))
	_firstLineIndent := (*C.int)(unsafe.Pointer(&firstLineIndent))
	C.TessPageIteratorParagraphInfo(_handle, _justification, _isListItem, _isCrown, _firstLineIndent)
	return
}

// TessPageIteratorLevel
type PageIteratorLevel int32

// TessPageSegMode
type PageSegMode int32

// TessParagraphJustification
type ParagraphJustification int32

// TessPolyBlockType
type PolyBlockType int32

// TessProgressFunc
type ProgressFunc *[0]byte

// TessResultIterator
type ResultIterator struct {
}

// TessResultIteratorConfidence
func (handle *ResultIterator) Confidence(level PageIteratorLevel) (ret float32) {
	_handle := (*C.TessResultIterator)(unsafe.Pointer(handle))
	_level := *(*C.TessPageIteratorLevel)(unsafe.Pointer(&level))
	_ret := C.TessResultIteratorConfidence(_handle, _level)
	ret = float32(_ret)
	return
}

// TessResultIteratorCopy
func (handle *ResultIterator) Copy() (ret *ResultIterator) {
	_handle := (*C.TessResultIterator)(unsafe.Pointer(handle))
	_ret := C.TessResultIteratorCopy(_handle)
	ret = (*ResultIterator)(unsafe.Pointer(_ret))
	return
}

// TessResultIteratorDelete
func (handle *ResultIterator) Delete() {
	_handle := (*C.TessResultIterator)(unsafe.Pointer(handle))
	C.TessResultIteratorDelete(_handle)
}

// TessResultIteratorGetChoiceIterator
func (handle *ResultIterator) GetChoiceIterator() (ret *ChoiceIterator) {
	_handle := (*C.TessResultIterator)(unsafe.Pointer(handle))
	_ret := C.TessResultIteratorGetChoiceIterator(_handle)
	ret = (*ChoiceIterator)(unsafe.Pointer(_ret))
	return
}

// TessResultIteratorGetPageIterator
func (handle *ResultIterator) GetPageIterator() (ret *PageIterator) {
	_handle := (*C.TessResultIterator)(unsafe.Pointer(handle))
	_ret := C.TessResultIteratorGetPageIterator(_handle)
	ret = (*PageIterator)(unsafe.Pointer(_ret))
	return
}

// TessResultIteratorGetPageIteratorConst
func (handle *ResultIterator) GetPageIteratorConst() (ret *PageIterator) {
	_handle := (*C.TessResultIterator)(unsafe.Pointer(handle))
	_ret := C.TessResultIteratorGetPageIteratorConst(_handle)
	ret = (*PageIterator)(unsafe.Pointer(_ret))
	return
}

// TessResultIteratorGetUTF8Text
func (handle *ResultIterator) GetUTF8Text(level PageIteratorLevel) (ret string) {
	_handle := (*C.TessResultIterator)(unsafe.Pointer(handle))
	_level := *(*C.TessPageIteratorLevel)(unsafe.Pointer(&level))
	_ret := C.TessResultIteratorGetUTF8Text(_handle, _level)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// TessResultIteratorNext
func (handle *ResultIterator) Next(level PageIteratorLevel) (ret bool) {
	_handle := (*C.TessResultIterator)(unsafe.Pointer(handle))
	_level := *(*C.TessPageIteratorLevel)(unsafe.Pointer(&level))
	_ret := C.TessResultIteratorNext(_handle, _level)
	ret = bool(_ret == 1)
	return
}

// TessResultIteratorSymbolIsDropcap
func (handle *ResultIterator) SymbolIsDropcap() (ret bool) {
	_handle := (*C.TessResultIterator)(unsafe.Pointer(handle))
	_ret := C.TessResultIteratorSymbolIsDropcap(_handle)
	ret = bool(_ret == 1)
	return
}

// TessResultIteratorSymbolIsSubscript
func (handle *ResultIterator) SymbolIsSubscript() (ret bool) {
	_handle := (*C.TessResultIterator)(unsafe.Pointer(handle))
	_ret := C.TessResultIteratorSymbolIsSubscript(_handle)
	ret = bool(_ret == 1)
	return
}

// TessResultIteratorSymbolIsSuperscript
func (handle *ResultIterator) SymbolIsSuperscript() (ret bool) {
	_handle := (*C.TessResultIterator)(unsafe.Pointer(handle))
	_ret := C.TessResultIteratorSymbolIsSuperscript(_handle)
	ret = bool(_ret == 1)
	return
}

// TessResultIteratorWordFontAttributes
func (handle *ResultIterator) WordFontAttributes() (isBold int32, isItalic int32, isUnderlined int32, isMonospace int32, isSerif int32, isSmallcaps int32, pointsize int32, fontId int32, ret string) {
	_handle := (*C.TessResultIterator)(unsafe.Pointer(handle))
	_isBold := (*C.Bool)(unsafe.Pointer(&isBold))
	_isItalic := (*C.Bool)(unsafe.Pointer(&isItalic))
	_isUnderlined := (*C.Bool)(unsafe.Pointer(&isUnderlined))
	_isMonospace := (*C.Bool)(unsafe.Pointer(&isMonospace))
	_isSerif := (*C.Bool)(unsafe.Pointer(&isSerif))
	_isSmallcaps := (*C.Bool)(unsafe.Pointer(&isSmallcaps))
	_pointsize := (*C.int)(unsafe.Pointer(&pointsize))
	_fontId := (*C.int)(unsafe.Pointer(&fontId))
	_ret := C.TessResultIteratorWordFontAttributes(_handle, _isBold, _isItalic, _isUnderlined, _isMonospace, _isSerif, _isSmallcaps, _pointsize, _fontId)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// TessResultIteratorWordIsFromDictionary
func (handle *ResultIterator) WordIsFromDictionary() (ret bool) {
	_handle := (*C.TessResultIterator)(unsafe.Pointer(handle))
	_ret := C.TessResultIteratorWordIsFromDictionary(_handle)
	ret = bool(_ret == 1)
	return
}

// TessResultIteratorWordIsNumeric
func (handle *ResultIterator) WordIsNumeric() (ret bool) {
	_handle := (*C.TessResultIterator)(unsafe.Pointer(handle))
	_ret := C.TessResultIteratorWordIsNumeric(_handle)
	ret = bool(_ret == 1)
	return
}

// TessResultIteratorWordRecognitionLanguage
func (handle *ResultIterator) WordRecognitionLanguage() (ret string) {
	_handle := (*C.TessResultIterator)(unsafe.Pointer(handle))
	_ret := C.TessResultIteratorWordRecognitionLanguage(_handle)
	ret = C.GoString(_ret)
	return
}

// TessResultRenderer
type ResultRenderer struct {
}

// TessDeleteResultRenderer
func (renderer *ResultRenderer) Delete() {
	_renderer := (*C.TessResultRenderer)(unsafe.Pointer(renderer))
	C.TessDeleteResultRenderer(_renderer)
}

// TessResultRendererAddImage
func (renderer *ResultRenderer) AddImage(api *BaseAPI) (ret bool) {
	_renderer := (*C.TessResultRenderer)(unsafe.Pointer(renderer))
	_api := (*C.TessBaseAPI)(unsafe.Pointer(api))
	_ret := C.TessResultRendererAddImage(_renderer, _api)
	ret = bool(_ret == 1)
	return
}

// TessResultRendererBeginDocument
func (renderer *ResultRenderer) BeginDocument(title string) (ret bool) {
	_renderer := (*C.TessResultRenderer)(unsafe.Pointer(renderer))
	_title := C.CString(title)
	defer C.free(unsafe.Pointer(_title))
	_ret := C.TessResultRendererBeginDocument(_renderer, _title)
	ret = bool(_ret == 1)
	return
}

// TessResultRendererEndDocument
func (renderer *ResultRenderer) EndDocument() (ret bool) {
	_renderer := (*C.TessResultRenderer)(unsafe.Pointer(renderer))
	_ret := C.TessResultRendererEndDocument(_renderer)
	ret = bool(_ret == 1)
	return
}

// TessResultRendererExtention
func (renderer *ResultRenderer) Extention() (ret string) {
	_renderer := (*C.TessResultRenderer)(unsafe.Pointer(renderer))
	_ret := C.TessResultRendererExtention(_renderer)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// TessResultRendererImageNum
func (renderer *ResultRenderer) ImageNum() (ret int32) {
	_renderer := (*C.TessResultRenderer)(unsafe.Pointer(renderer))
	_ret := C.TessResultRendererImageNum(_renderer)
	ret = int32(_ret)
	return
}

// TessResultRendererInsert
func (renderer *ResultRenderer) Insert(next *ResultRenderer) {
	_renderer := (*C.TessResultRenderer)(unsafe.Pointer(renderer))
	_next := (*C.TessResultRenderer)(unsafe.Pointer(next))
	C.TessResultRendererInsert(_renderer, _next)
}

// TessResultRendererNext
func (renderer *ResultRenderer) Next() (ret *ResultRenderer) {
	_renderer := (*C.TessResultRenderer)(unsafe.Pointer(renderer))
	_ret := C.TessResultRendererNext(_renderer)
	ret = (*ResultRenderer)(unsafe.Pointer(_ret))
	return
}

// TessResultRendererTitle
func (renderer *ResultRenderer) Title() (ret string) {
	_renderer := (*C.TessResultRenderer)(unsafe.Pointer(renderer))
	_ret := C.TessResultRendererTitle(_renderer)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// TessTextlineOrder
type TextlineOrder int32

// TessWritingDirection
type WritingDirection int32

// TessAltoRendererCreate
func AltoRendererCreate(outputbase string) (ret *ResultRenderer) {
	_outputbase := C.CString(outputbase)
	defer C.free(unsafe.Pointer(_outputbase))
	_ret := C.TessAltoRendererCreate(_outputbase)
	ret = (*ResultRenderer)(unsafe.Pointer(_ret))
	return
}

// TessBaseAPICreate
func BaseAPICreate() (ret *BaseAPI) {
	_ret := C.TessBaseAPICreate()
	ret = (*BaseAPI)(unsafe.Pointer(_ret))
	return
}

// TessBoxTextRendererCreate
func BoxTextRendererCreate(outputbase string) (ret *ResultRenderer) {
	_outputbase := C.CString(outputbase)
	defer C.free(unsafe.Pointer(_outputbase))
	_ret := C.TessBoxTextRendererCreate(_outputbase)
	ret = (*ResultRenderer)(unsafe.Pointer(_ret))
	return
}

// TessDeleteIntArray
func DeleteIntArray(arr []int32) {
	_arr := (*C.int)(nil)
	if len(arr) > 0 {
		_arr = (*C.int)(unsafe.Pointer(&arr[0]))
	}
	C.TessDeleteIntArray(_arr)
}

// TessDeleteText
func DeleteText(text string) {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	C.TessDeleteText(_text)
}

// TessDeleteTextArray
func DeleteTextArray() (arr *int8) {
	_arr := (**C.char)(unsafe.Pointer(&arr))
	C.TessDeleteTextArray(_arr)
	return
}

// TessHOcrRendererCreate
func HOcrRendererCreate(outputbase string) (ret *ResultRenderer) {
	_outputbase := C.CString(outputbase)
	defer C.free(unsafe.Pointer(_outputbase))
	_ret := C.TessHOcrRendererCreate(_outputbase)
	ret = (*ResultRenderer)(unsafe.Pointer(_ret))
	return
}

// TessHOcrRendererCreate2
func HOcrRendererCreate2(outputbase string, fontInfo bool) (ret *ResultRenderer) {
	_outputbase := C.CString(outputbase)
	defer C.free(unsafe.Pointer(_outputbase))
	_fontInfo := C.Bool(0)
	if fontInfo {
		_fontInfo = C.Bool(1)
	}
	_ret := C.TessHOcrRendererCreate2(_outputbase, _fontInfo)
	ret = (*ResultRenderer)(unsafe.Pointer(_ret))
	return
}

// TessLSTMBoxRendererCreate
func LSTMBoxRendererCreate(outputbase string) (ret *ResultRenderer) {
	_outputbase := C.CString(outputbase)
	defer C.free(unsafe.Pointer(_outputbase))
	_ret := C.TessLSTMBoxRendererCreate(_outputbase)
	ret = (*ResultRenderer)(unsafe.Pointer(_ret))
	return
}

// TessMonitorCreate
func MonitorCreate() (ret *[0]byte) {
	_ret := C.TessMonitorCreate()
	ret = (*[0]byte)(unsafe.Pointer(_ret))
	return
}

// TessMonitorDelete
func MonitorDelete(monitor *[0]byte) {
	_monitor := (*C.ETEXT_DESC)(unsafe.Pointer(monitor))
	C.TessMonitorDelete(_monitor)
}

// TessMonitorGetCancelThis
func MonitorGetCancelThis(monitor *[0]byte) (ret uintptr) {
	_monitor := (*C.ETEXT_DESC)(unsafe.Pointer(monitor))
	_ret := C.TessMonitorGetCancelThis(_monitor)
	ret = (uintptr)(unsafe.Pointer(_ret))
	return
}

// TessMonitorGetProgress
func MonitorGetProgress(monitor *[0]byte) (ret int32) {
	_monitor := (*C.ETEXT_DESC)(unsafe.Pointer(monitor))
	_ret := C.TessMonitorGetProgress(_monitor)
	ret = int32(_ret)
	return
}

// TessMonitorSetCancelFunc
func MonitorSetCancelFunc(monitor *[0]byte, cancelFunc CancelFunc) {
	_monitor := (*C.ETEXT_DESC)(unsafe.Pointer(monitor))
	_cancelFunc := (C.TessCancelFunc)(unsafe.Pointer(cancelFunc))
	C.TessMonitorSetCancelFunc(_monitor, _cancelFunc)
}

// TessMonitorSetCancelThis
func MonitorSetCancelThis(monitor *[0]byte, cancelThis uintptr) {
	_monitor := (*C.ETEXT_DESC)(unsafe.Pointer(monitor))
	_cancelThis := unsafe.Pointer(cancelThis)
	C.TessMonitorSetCancelThis(_monitor, _cancelThis)
}

// TessMonitorSetDeadlineMSecs
func MonitorSetDeadlineMSecs(monitor *[0]byte, deadline int32) {
	_monitor := (*C.ETEXT_DESC)(unsafe.Pointer(monitor))
	_deadline := C.int(deadline)
	C.TessMonitorSetDeadlineMSecs(_monitor, _deadline)
}

// TessMonitorSetProgressFunc
func MonitorSetProgressFunc(monitor *[0]byte, progressFunc ProgressFunc) {
	_monitor := (*C.ETEXT_DESC)(unsafe.Pointer(monitor))
	_progressFunc := (C.TessProgressFunc)(unsafe.Pointer(progressFunc))
	C.TessMonitorSetProgressFunc(_monitor, _progressFunc)
}

// TessPDFRendererCreate
func PDFRendererCreate(outputbase string, datadir string, textonly bool) (ret *ResultRenderer) {
	_outputbase := C.CString(outputbase)
	defer C.free(unsafe.Pointer(_outputbase))
	_datadir := C.CString(datadir)
	defer C.free(unsafe.Pointer(_datadir))
	_textonly := C.Bool(0)
	if textonly {
		_textonly = C.Bool(1)
	}
	_ret := C.TessPDFRendererCreate(_outputbase, _datadir, _textonly)
	ret = (*ResultRenderer)(unsafe.Pointer(_ret))
	return
}

// TessTextRendererCreate
func TextRendererCreate(outputbase string) (ret *ResultRenderer) {
	_outputbase := C.CString(outputbase)
	defer C.free(unsafe.Pointer(_outputbase))
	_ret := C.TessTextRendererCreate(_outputbase)
	ret = (*ResultRenderer)(unsafe.Pointer(_ret))
	return
}

// TessTsvRendererCreate
func TsvRendererCreate(outputbase string) (ret *ResultRenderer) {
	_outputbase := C.CString(outputbase)
	defer C.free(unsafe.Pointer(_outputbase))
	_ret := C.TessTsvRendererCreate(_outputbase)
	ret = (*ResultRenderer)(unsafe.Pointer(_ret))
	return
}

// TessUnlvRendererCreate
func UnlvRendererCreate(outputbase string) (ret *ResultRenderer) {
	_outputbase := C.CString(outputbase)
	defer C.free(unsafe.Pointer(_outputbase))
	_ret := C.TessUnlvRendererCreate(_outputbase)
	ret = (*ResultRenderer)(unsafe.Pointer(_ret))
	return
}

// TessVersion
func Version() (ret string) {
	_ret := C.TessVersion()
	ret = C.GoString(_ret)
	return
}

// TessWordStrBoxRendererCreate
func WordStrBoxRendererCreate(outputbase string) (ret *ResultRenderer) {
	_outputbase := C.CString(outputbase)
	defer C.free(unsafe.Pointer(_outputbase))
	_ret := C.TessWordStrBoxRendererCreate(_outputbase)
	ret = (*ResultRenderer)(unsafe.Pointer(_ret))
	return
}

// TessBaseAPISetImageBytes
func (handle *BaseAPI) SetImageBytes(bytes []byte) (ret *C.Pix) {
	_handle := (*C.TessBaseAPI)(unsafe.Pointer(handle))
	_pix := C.pixReadMem((*C.uchar)(unsafe.Pointer(&bytes[0])), C.ulong(len(bytes)))
	C.TessBaseAPISetImage2(_handle, _pix)
        return _pix
}

// pixFreeData
func FreeImageBytes(pix *C.Pix) {
        C.pixFreeData(pix)
}
