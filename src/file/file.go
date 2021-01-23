package file

import "net/http"

func RegisterFileHandlers(docRoot string) {
	fs := http.FileServer(http.Dir(docRoot))
	http.Handle("/", fs)
}