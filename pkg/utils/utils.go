package utils

import "net/http"
import "io"

func Error(res http.ResponseWriter, err error) {
	res.WriteHeader(http.StatusInternalServerError)
	io.WriteString(res, err.Error())
}
