package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var count int
var mu sync.Mutex

func main() {

	// 最长匹配原则
	// 所以访问counter 不会匹配到handler
	http.HandleFunc("/", handler)
	http.HandleFunc("/counter", counter)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()

	// %q 该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
	fmt.Fprintf(w, "URL.Path=%q", r.URL.Path)

}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "count=%d", count)
	mu.Unlock()

}
