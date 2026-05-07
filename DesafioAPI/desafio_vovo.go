package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Cliente representa a estrutura de dados da vovó
type Cliente struct {
	ID    string `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// Simulando um banco de dados em memória
var clientes []Cliente

func main() {
	// Dados iniciais para teste
	clientes = append(clientes, Cliente{ID: "1", Nome: "Bento", Email: "bento@email.com"})

	router := mux.NewRouter()

	// Definição das Rotas (Endpoints)
	router.HandleFunc("/clientes", getClientes).Methods("GET")
	router.HandleFunc("/clientes/{id}", getCliente).Methods("GET")
	router.HandleFunc("/clientes", createCliente).Methods("POST")
	router.HandleFunc("/clientes/{id}", updateCliente).Methods("PUT")
	router.HandleFunc("/clientes/{id}", deleteCliente).Methods("DELETE")

	fmt.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// --- Funções Handlers ---

func getClientes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clientes)
}

func getCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range clientes {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Cliente{})
}

func createCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cliente Cliente
	_ = json.NewDecoder(r.Body).Decode(&cliente)
	clientes = append(clientes, cliente)
	json.NewEncoder(w).Encode(cliente)
}

func updateCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range clientes {
		if item.ID == params["id"] {
			clientes = append(clientes[:index], clientes[index+1:]...)
			var cliente Cliente
			_ = json.NewDecoder(r.Body).Decode(&cliente)
			cliente.ID = params["id"]
			clientes = append(clientes, cliente)
			json.NewEncoder(w).Encode(cliente)
			return
		}
	}
}

func deleteCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range clientes {
		if item.ID == params["id"] {
			clientes = append(clientes[:index], clientes[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(clientes)
}
