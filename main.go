package main

import (
	"database/sql"
	"fmt"
	"time" // Nova peça: serve para o robô esperar um pouco entre as rondas
	_ "github.com/lib/pq"
)

func checkDatabase(serviceName string, port int) {
	// 1. Definimos variáveis vazias
	var user, password, dbname string

	// 2. Escolhemos os dados certos baseados na porta
	if port == 5433 {
		user = "postgres"
		password = "postgres"
		dbname = "microservicesdb"
	} else {
		// DADOS DA KOTLIN (Baseado no seu application.yml)
		user = "demo"
		password = "demo123"
		dbname = "demo_db"
	}

	// 3. Montamos a string de conexão com os dados corretos
	connStr := fmt.Sprintf("host=localhost port=%d user=%s password=%s dbname=%s sslmode=disable", 
		port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("❌ [%s] Erro de config: %v\n", serviceName, err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Printf("❌ [%s] OFFLINE (Porta %d) - Erro: %v\n", serviceName, port, err)
	} else {
		fmt.Printf("✅ [%s] ONLINE (Porta %d)\n", serviceName, port)
	}
}

// Função auxiliar para saber qual banco conectar baseado na porta
func getDBName(port int) string {
	if port == 5433 {
		return "microservicesdb" // Banco da Java
	}
	return "demo_db" // Banco da Kotlin
}

func main() {
	fmt.Println("🚀 Iniciando Monitor de Infraestrutura (Go)...")
	fmt.Println("----------------------------------------------")

	// Criamos um loop infinito (o vigia nunca dorme!)
	for {
		// Checa o banco da Java
		checkDatabase("API JAVA (Audit)", 5433)

		// Checa o banco da Kotlin
		checkDatabase("API KOTLIN (User)", 5434)

		fmt.Println("----------------------------------------------")
		
		// Espera 10 segundos antes da próxima ronda
		time.Sleep(10 * time.Second)
	}
}