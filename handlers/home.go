package handlers

import (
	"net/http"

	"github.com/DimRev/Go-Htmx-Tailwind-Starter/app/page"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, page.Home())
}
