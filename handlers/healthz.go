package handlers

import (
	"net/http"
)

func HandleHealthz(w http.ResponseWriter, r *http.Request) error {
	_, err := w.Write([]byte("OK"))
	return err
}
