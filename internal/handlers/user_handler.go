package handlers

import (
	"fmt"
	"net/http"
)


func UserHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "This is user handler")
}


func Register(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Registration Handler")
}