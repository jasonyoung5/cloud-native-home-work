package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func MyHandleRequest(w http.ResponseWriter, req *http.Request){
	var version string
	version = os.Getenv("JAVA_HOME")
	fmt.Println(version)

	params := req.Header.Get("User-Agent")
	remoteIp := req.Header.Get("RemoteAddr")
	fmt.Println("log remote ip:",remoteIp)
	w.Header().Set("VERSION",version)
	w.Header().Set("MYRETURN",params)
	io.WriteString(w,"ok")
}

func main() {
	http.HandleFunc("/healthz", MyHandleRequest)
	http.ListenAndServe(":80", nil)
}
