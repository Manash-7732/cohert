package main

import (
	"fmt"
	"net/http"
)

func manash(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w ,"Hello World");
}

func main() {
      http.HandleFunc("/",manash);
	  fmt.Println("Starting server on :8080");
	  http.ListenAndServe("",nil);
}
