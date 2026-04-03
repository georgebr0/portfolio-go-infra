package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func main() {
	// Configuração de Infra: Ler credenciais de ENV (Segurança)
	sa := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	app, err := firebase.NewApp(context.Background(), nil, sa)
	if err != nil {
		log.Fatalf("Erro ao iniciar Firebase: %v", err)
	}

	_, err = app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("Erro ao conectar Firestore: %v", err)
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		// Simula um check de infraestrutura
		fmt.Fprintf(w, "Status: OK - Timestamp: %d", time.Now().Unix())
	})

	log.Println("Servidor rodando na porta 8080")
	http.ListenAndServe(":8080", nil)
}