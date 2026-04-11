# Passo 1: Pegar uma cozinha emprestada (Imagem do Go) para compilar o código
FROM golang:1.26-alpine AS builder

# Passo 2: Definir a pasta de trabalho dentro da cozinha
WORKDIR /app

# Passo 3: Copiar os arquivos de dependências e baixar tudo
COPY go.mod go.sum ./
RUN go mod download

# Passo 4: Copiar o resto do código e transformar em um arquivo executável (binário)
COPY . .
RUN go build -o monitor .

# Passo 5: Criar a marmita final, bem leve (apenas com o executável)
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/monitor .

# Passo 6: Dar a ordem para o robô começar a vigiar
CMD ["./monitor"]