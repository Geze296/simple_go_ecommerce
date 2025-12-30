package handlers

import (
	"fmt"
	"net/http"
)

func ProductHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "This is a product handler")
}