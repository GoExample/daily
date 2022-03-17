package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", sayHello)
	http.HandleFunc("/calc", bitOperation)
	err := http.ListenAndServe("0.0.0.0:9000", nil)
	if err != nil {
		fmt.Printf("Err msg is %s\n", err)
		return
	}

}

func bitOperation(writer http.ResponseWriter, request *http.Request) {
	i := 1<<16 - 1
	fprintf, err := fmt.Fprintf(writer, "2的16次方 %d", i)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(request.URL)
	fmt.Printf("fprintf %d\n\n", fprintf)
}

func sayHello(writer http.ResponseWriter, request *http.Request) {
	fprintf, err := fmt.Fprintf(writer, "<h1>hello golang!</h1><button id=\"bt1\">picture</button>")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(request.URL)
	fmt.Printf("fprintf %d\n\n", fprintf)
}
