package main

import (
	"net/http"
	"io"
	"os/exec"
	"log"
)

func reLaunch()  {
	cmd := exec.Command("sh", "./deploy.sh");
	err := cmd.Start()
	if err != nil {
		log.Fatal(err);
	}
	err = cmd.Wait();
}

func firstPage(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "<p>你好，我是自动部署的golang</p>");
	reLaunch()
}

func main() {
	http.HandleFunc("/", firstPage);
	http.ListenAndServe(":8080", nil);

}
