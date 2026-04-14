# 🚀 GoLang Microservice - [Go-Monitor-Service]

Este microsserviço é responsável por monitorar a saúde e as métricas de performance do ecossistema de microsserviços, integrando-se via Apache Kafka para coletar eventos de auditoria em tempo real.

## 🛠️ Tecnologias
- **GoLang** (1.21+)
- **Gin Gonic** (ou o framework que você estiver usando)
- **Kafka Go** (segmentio/kafka-go)

## 🏗️ Arquitetura
Este serviço consome/produz eventos para o tópico `audit-log-topic` no Apache Kafka, permitindo a comunicação assíncrona com os serviços Kotlin e Java.

## 🚀 Como rodar
1. Garanta que o Kafka esteja de pé no Docker.
2. Configure as variáveis de ambiente:
   ```bash
   export KAFKA_BROKER="localhost:9092"

3. Execute: 
go run main.go




