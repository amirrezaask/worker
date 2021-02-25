package worker

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type httpServer struct {
	mux  *http.ServeMux
	port int
	w    Worker
}

func NewHttpServer(w Worker, port int) *httpServer {
	mux := http.NewServeMux()
	h := &httpServer{}
	mux.HandleFunc("/recent", h.RecentJobs)
	mux.HandleFunc("/failed", h.Failed)
	mux.HandleFunc("/stats", h.Stats)
	h.mux = mux
	h.port = port
	h.w = w
	return h
}

func (h *httpServer) Start() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", h.port), h.mux)
}

func (h *httpServer) RecentJobs(w http.ResponseWriter, r *http.Request) {
	n, err := strconv.Atoi(r.URL.Query().Get("number"))
	if err != nil {
		n = 40
	}
	jobs := h.w.RecentJobs(n)
	bs, err := json.Marshal(jobs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	w.Write(bs)
}

func (h *httpServer) Failed(w http.ResponseWriter, r *http.Request) {
	jobs := h.w.Failed()
	bs, err := json.Marshal(jobs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	w.Write(bs)
}

func (h *httpServer) Stats(w http.ResponseWriter, r *http.Request) {}
