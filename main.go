package main

import (
	"database/sql" // Ferramenta para falar com bancos SQL
	"fmt"          // Ferramenta para escrever mensagens
	"log"          // Ferramenta para registrar erros
	_ "github.com/lib/pq" // O "driver" do Postgres (o _ significa que usamos ele quietinho no fundo)
)

func main() {
	// 1. A nossa "frase mágica" de conexão (Igual a que usamos no application.properties)
	// host, port, user, password, dbname e sslmode
	connStr := "host=localhost port=5433 user=postgres password=postgres dbname=microservicesdb sslmode=disable"

	// 2. O Go tenta abrir a porta do banco
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao configurar o banco: ", err)
	}
	defer db.Close() // Isso garante que a porta feche quando terminarmos (boa educação!)

	// 3. O PING: É aqui que o Vigia pergunta: "Tem alguém aí?"
	err = db.Ping()
	if err != nil {
		fmt.Println("❌ A API Java está fora do ar ou o banco está desligado!")
		log.Fatal(err)
	}

	fmt.Println("✅ Sucesso! O monitor Go conseguiu falar com o banco da API Java na porta 5433.")
}