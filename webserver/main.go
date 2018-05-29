package main

import (
	"net/http"
	"io"
)

func firstPage(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "<p>你好，我是golang</p>");
}

func main() {
	http.HandleFunc("/", firstPage);
	http.ListenAndServe(":80", nil);

}
