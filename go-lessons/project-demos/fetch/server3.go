package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jinyuyoulong/Go-learn/mytest/gostart"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		strcycles := r.FormValue("cycles")
		cycles, err := strconv.Atoi(strcycles)
		if err != nil {
			fmt.Println("%s", err)
		}
		gostart.Lissajous(w, cycles)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
// 	for k, v := range r.Header {
// 		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
// 	}
// 	fmt.Fprintf(w, "Host = %q\n", r.Host)
// 	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
// 	if err := r.ParseForm(); err != nil {
// 		log.Print(err)
// 	}
// 	for k, v := range r.Form {
// 		fmt.Fprint(w, "Form[%q] = %q\n", k, v)
// 	}
// }
