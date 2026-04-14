# 🚀 GoLang Microservice - [Nome do Serviço]

Este microsserviço faz parte de um ecossistema distribuído, focado em alta performance e baixa latência.

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

---

### 2. Atualização: README da API Kotlin (Producer)
Você deve remover as referências ao RabbitMQ e adicionar a nova responsabilidade de produção de eventos Kafka.

```markdown
## 📡 Integração com Kafka (Novo)
A API agora atua como um **Producer**. 
Toda vez que um usuário é criado/atualizado, um evento é disparado para o tópico `audit-log-topic`.

**Configuração principal:**
- Broker: `localhost:9092`
- Tópico: `audit-log-topic`



