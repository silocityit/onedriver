// +build linux,cgo

package graph

/*
#cgo linux pkg-config: webkit2gtk-4.0
#include "stdlib.h"
#include "oauth2_gtk.h"
*/
import "C"

import (
	"regexp"
	"unsafe"

	log "github.com/sirupsen/logrus"
)

// Fetch the auth code required as the first part of oauth2 authentication. Uses
// webkit2gtk to create a popup browser.
func getAuthCode() string {
	cAuthURL := C.CString(getAuthURL())
	cResponse := C.webkit_auth_window(cAuthURL)
	response := C.GoString(cResponse)
	C.free(unsafe.Pointer(cAuthURL))
	C.free(unsafe.Pointer(cResponse))

	rexp := regexp.MustCompile("code=([a-zA-Z0-9-_])+")
	code := rexp.FindString(response)
	if len(code) == 0 {
		log.Fatal("No validation code returned, or code was invalid. " +
			"Please restart the application and try again.")
	}
	return code[5:]
}
