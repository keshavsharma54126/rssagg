package main

import (
	"net/http"
)

func handlerReadniness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
