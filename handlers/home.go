package handlers

import (
	"net/http"

	"github.com/DimRev/Go-Htmx-Tailwind-Starter/views/pages"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, pages.Home())
}
