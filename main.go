package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Resposta struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

func projetoKorpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	resposta := Resposta{
		Nome:    "Projeto Korp",
		Horario: time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(resposta); err != nil {
		log.Printf("erro ao gerar resposta JSON: %v", err)
	}
}

func main() {
	http.HandleFunc("/projeto-korp", projetoKorpHandler)

	log.Println("http-server-projeto-korp iniciado na porta 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
