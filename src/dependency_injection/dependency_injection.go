package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	_, _ = fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	//	fmt.Println("aaaaaa")
	error := http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
	fmt.Printf("unexpected error : %q", error)
}
