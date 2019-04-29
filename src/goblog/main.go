package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

const ROOT = "/Users/alan/goblog"

func main() {
	savePid()
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":8081", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world")
}

//save pid
func savePid() {
	pidFilename := ROOT + "/pid/" + filepath.Base(os.Args[0]) + ".pid"
	pid := os.Getpid()
	ioutil.WriteFile(pidFilename, []byte(strconv.Itoa(pid)), 0755)
}
