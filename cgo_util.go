package clipboard

import "C"
import (
	sc "golang.org/x/text/encoding/simplifiedchinese"
	"unsafe"
)

// Go字符串转成指针，该指针可被外部调用
func GoStr2CPtr(s string) uintptr {
	return uintptr(unsafe.Pointer(Go2CStr(s)))
}

// Go字符串转C字符串
func Go2CStr(str string) *C.char {
	gbstr, _ := sc.GB18030.NewEncoder().String(str)
	return C.CString(gbstr)
}

// C字符串转Go字符串
func C2GoStr(str *C.char) string {
	utf8str, _ := sc.GB18030.NewDecoder().String(C.GoString(str))
	return utf8str
}
