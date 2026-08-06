package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"

	"github.com/google/uuid"
)

func main() {
	msg := flag.String("msg", "Hi", "인사 메시지")
	flag.Parse()

	var (
		mu    sync.Mutex
		count int
	)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		count++
		c := count
		mu.Unlock()
		reqID := uuid.New().String()[:8]
		fmt.Fprintf(w, "%s (visit #%d, req %s)\n", *msg, c, reqID)
	})
	http.ListenAndServe(":5000", nil)
}
