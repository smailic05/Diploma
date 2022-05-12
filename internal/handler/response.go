package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func writeResponse(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"Internal server error"}`))
		return
	}
	w.WriteHeader(code)
	w.Write([]byte(b))
}

func writeResponseFile(w http.ResponseWriter, code int, filename string, v []byte) {
	w.Header().Set("Content-Disposition",
		fmt.Sprintf("attachment; filename=%s", filename))
	w.WriteHeader(code)
	w.Write(v)
}
