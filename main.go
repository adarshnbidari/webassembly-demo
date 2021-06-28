package main

import (
	"net/http"
)

func main() {

	s := http.FileServer(http.Dir("examples"))

	http.ListenAndServe(":80", s)

}
