package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Resposta struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

var (
	requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total de requisicoes HTTP recebidas pela aplicacao.",
		},
		[]string{"method", "path", "status"},
	)

	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duracao das requisicoes HTTP em segundos.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)
)

func init() {
	prometheus.MustRegister(requestsTotal)
	prometheus.MustRegister(requestDuration)
}

func projetoKorpHandler(w http.ResponseWriter, r *http.Request) {
	inicio := time.Now()

	if r.Method != http.MethodGet {
		requestsTotal.WithLabelValues(
			r.Method,
			r.URL.Path,
			"405",
		).Inc()

		http.Error(
			w,
			"Metodo nao permitido",
			http.StatusMethodNotAllowed,
		)
		return
	}

	resposta := Resposta{
		Nome:    "Projeto Korp",
		Horario: time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(resposta); err != nil {
		log.Printf("erro ao gerar resposta JSON: %v", err)
		return
	}

	requestsTotal.WithLabelValues(
		r.Method,
		r.URL.Path,
		"200",
	).Inc()

	requestDuration.WithLabelValues(
		r.Method,
		r.URL.Path,
	).Observe(time.Since(inicio).Seconds())
}

func main() {
	http.HandleFunc("/projeto-korp", projetoKorpHandler)

	http.Handle("/metrics", promhttp.Handler())

	log.Println("http-server-projeto-korp iniciado na porta 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
