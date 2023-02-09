package fileserver

import (
	"net/http"
)

func NewFileServer() http.Handler {
	return http.FileServer(http.Dir("./static/"))
}
