package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const letterBytes = "0123456789abcdefghigklmnopqrstuvwxyzABCDEFGHIGKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func randString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

func main() {
	host := os.Args[1]
	port := os.Args[2]
	url := fmt.Sprintf("/%s", randString(10))
	fmt.Printf("host is %s\n", host)
	fmt.Printf("port is %s\n", host)
	fmt.Printf("url is %s\n", url)

	http.HandleFunc(url, serverHandler)
	http.HandleFunc("/", hello)
	host = fmt.Sprintf("%s:%s", host, port)
	err := http.ListenAndServe(host, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func hello(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello this fucking world"))
}

func serverHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Welcome to URL.Path = %q\n", r.URL.Path)
}
