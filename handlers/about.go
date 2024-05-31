package handlers

import (
	"net/http"

	"github.com/DimRev/Go-Htmx-Tailwind-Starter/views/pages"
)

func HandleAbout(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, pages.About())
}
