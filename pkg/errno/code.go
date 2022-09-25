/**
@author: yeebing
@date: 2022/9/24
**/

package errno

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	unknownCoder Coder = Coder{1, http.StatusInternalServerError, "An internal server error occurred"}
)

type Coder struct {
	// C refers to the integer code of the ErrCode.
	C int

	// HTTP status that should be used for the associated error code.
	HTTP int

	Msg string
}

// New create an instance of Coder
func New(code int, httpStatus int, msg string) Coder {
	return Coder{
		C:    code,
		HTTP: httpStatus,
		Msg:  msg,
	}
}

// Code returns the integer code of the coder.
func (coder Coder) Code() int {
	return coder.C

}

// String implements stringer. String returns the external error message,
// if any.
func (coder Coder) String() string {
	return coder.Msg
}

// HTTPStatus returns the associated HTTP status code, if any. Otherwise,
// returns 200.
func (coder Coder) HTTPStatus() int {
	if coder.HTTP == 0 {
		return 500
	}

	return coder.HTTP
}

// codes contains a map of error codes to metadata.
var codes = map[int]Coder{}
var codeMux = &sync.Mutex{}

// Support httpStatus
// StatusOK                           = 200
// StatusBadRequest                   = 400
// StatusUnauthorized                 = 401
// StatusForbidden                    = 403
// StatusNotFound                     = 404
// StatusInternalServerError          = 500

// RegisterWithArgs register code, httpStatus and msg into map.
func RegisterWithArgs(code int, httpStatus int, msg string) {
	Register(New(code, httpStatus, msg))
}

// Register a user define error code.
// It will override to exist code.
func Register(coder Coder) {
	if coder.Code() == 0 {
		panic("code `0` is reserved as unknownCode error code")
	}

	codeMux.Lock()
	defer codeMux.Unlock()

	codes[coder.Code()] = coder
}

// MustRegister register a user define error code.
// It will panic when the same Code already exist.
func MustRegister(coder Coder) {
	if coder.Code() == 0 {
		panic("code '0' is reserved as ErrUnknown error code")
	}

	codeMux.Lock()
	defer codeMux.Unlock()

	if _, ok := codes[coder.Code()]; ok {
		panic(fmt.Sprintf("code: %d already exist", coder.Code()))
	}

	codes[coder.Code()] = coder
}

// ParseCoder parse any error into *withCode.
// nil error will return nil direct.
// None withStack error will be parsed as ErrUnknown.
func ParseCoder(err error) Coder {
	if err == nil {
		return Coder{}
	}

	if v, ok := err.(*withCode); ok {
		if coder, ok := codes[v.code]; ok {
			return coder
		}
	}

	return unknownCoder
}

func init() {
	codes[unknownCoder.Code()] = unknownCoder
}
