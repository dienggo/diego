package render

import (
	"net/http"
)

// Json writes the response headers and calls JSON to render data.
func Json(w http.ResponseWriter, code int, data any) {
	General(w, code, JSON{data})
}

// General writes the response headers and calls render.Render to render data.
func General(w http.ResponseWriter, code int, r Render) {
	w.WriteHeader(code)
	r.WriteContentType(w)
	if err := r.Render(w); err != nil {
		panic(err)
	}
}
