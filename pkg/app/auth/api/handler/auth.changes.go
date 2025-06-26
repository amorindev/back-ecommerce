package handler

import (
	"net/http"
	"time"
)

// see handler
func (h Handler) AuthChanges(w http.ResponseWriter, r *http.Request) {
	// admite solo cierta candidad de connecciones cliente con sse?
	// ver si solo se admite text
	w.Header().Set("Content-Type", "json/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	w.Header().Set("Access-Control-Allow-Origin", "*")

	memT := time.NewTicker(time.Second)
	defer memT.Stop()

	cpuT := time.NewTicker(time.Second)
	defer cpuT.Stop()

	clientGone := r.Context().Done()

	for {
		select {
		case <-clientGone:
			//fmt.Printf("client has disconnected")
		case <-memT.C:
		case <-cpuT.C:
		}
	}
}
