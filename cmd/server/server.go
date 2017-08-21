package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func imgHandler(w http.ResponseWriter, r *http.Request) {
	// Yeah, we could use gorilla/schema for the query handling (https://godoc.org/github.com/gorilla/schema) but this
	// is so simple, we might as well do it here so it is immediately understandable.

	// we know "w" should ALWAYS be specified
	width, err := strconv.Atoi(r.FormValue("w"))
	if err != nil {
		http.Error(w, "Invalid 'w' (width).", 400)
		return
	}

	height := 0
	if r.FormValue("h") == "" {
		height = width
	} else {
		height, err = strconv.Atoi(r.FormValue("h"))
		if err != nil {
			http.Error(w, "Invalid 'h' (height).", 400)
			return
		}
	}

	saturation := "colour"
	if r.FormValue("s") == "" {
		saturation = "colour"
	} else if r.FormValue("s") == "c" {
		saturation = "colour"
	} else if r.FormValue("s") == "g" {
		saturation = "grey"
	} else {
		http.Error(w, "Invalid 's' (saturation), must be 'c' or 'g' (optional: default 'c').", 400)
		return
	}

	fmt.Fprintf(w, "<h2>Image %q</h2>\n", html.EscapeString(r.URL.Path))
	fmt.Fprintf(w, "<ul>\n")
	fmt.Fprintf(w, "  <li>w=%v</li>\n", width)
	fmt.Fprintf(w, "  <li>h=%v</li>\n", height)
	fmt.Fprintf(w, "  <li>s=%v</li>\n", saturation)
	fmt.Fprintf(w, "</ul>\n")
	fmt.Fprintf(w, "<p>Ends</p>\n")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home %q.", html.EscapeString(r.URL.Path))
}

func main() {
	// setup
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// the handler - there are only two routes
	mux := http.NewServeMux()

	mux.HandleFunc("/img", imgHandler)
	mux.HandleFunc("/", homeHandler)

	// the server
	s := &http.Server{
		Addr:           ":" + port,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
