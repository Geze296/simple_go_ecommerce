package handlers

import (
	"fmt"
	"net/http"
)


func OrderHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "This is an Order Handler")
}