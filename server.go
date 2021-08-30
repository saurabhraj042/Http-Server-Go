package main

import (
	"fmt"
	"net/http"
	"strings"
)

// PlayerServer currently returns Hello, world given _any_ request.
//server.go
func PlayerServer(w http.ResponseWriter, r *http.Request) {
    player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, GetPlayerScore(player))
}

func GetPlayerScore(name string) string {
	if name == "Pepper" {
        return "20"
    }

    if name == "Floyd" {
        return "10"
    }

	return ""
}