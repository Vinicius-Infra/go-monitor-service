package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/lib/pq"
<<<<<<< Updated upstream
	"net/http"
    "github.com/prometheus/client_golang/prometheus/promhttp"
=======
	"github.com/prometheus/client_golang/prometheus/promhttp"
>>>>>>> Stashed changes
)

func checkDatabase(serviceName string, port int) {
	var user, password, dbname string

	if port == 5433 {
		user = "postgres"
		password = "postgres"
		dbname = "microservicesdb"
	} else {
		user = "demo"
		password = "demo123"
		dbname = "demo_db"
	}

	// Como o Go rodará DENTRO do Docker agora, usamos host.docker.internal 
	// para apontar para os bancos mapeados no seu localhost do Windows
	connStr := fmt.Sprintf("host=host.docker.internal port=%d user=%s password=%s dbname=%s sslmode=disable", 
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

func main() {
	fmt.Println("🚀 Iniciando Monitor de Infraestrutura (Go)...")
	fmt.Println("----------------------------------------------")

	// 1. Cria o servidor HTTP para o Prometheus em background (Goroutine)
	// Isso permite que o endpoint /metrics fique ouvindo na porta 8081
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		fmt.Println("📊 Endpoint de métricas ativo em http://localhost:8081/metrics")
		if err := http.ListenAndServe(":8081", nil); err != nil {
			fmt.Printf("❌ Erro ao iniciar servidor de métricas: %v\n", err)
		}
	}()

	// 2. Loop principal do Monitor (Vigia)
	for {
		checkDatabase("API JAVA (Audit)", 5433)
		checkDatabase("API KOTLIN (User)", 5434)
		fmt.Println("----------------------------------------------")
		time.Sleep(10 * time.Second)

		// Rota padrão do Prometheus
    http.Handle("/metrics", promhttp.Handler())
    
    // Inicia o servidor na porta 8081
    http.ListenAndServe(":8081", nil)
	}
}