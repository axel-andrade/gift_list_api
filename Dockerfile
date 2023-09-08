# Use uma imagem mínima do Go como imagem base, como a imagem Alpine
FROM golang:1.16-alpine

# Defina variáveis de ambiente para otimizar a compilação
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Crie um diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie apenas o arquivo go.mod e go.sum para permitir o cache das dependências
COPY go.mod go.sum ./

# Baixe e instale as dependências
RUN go mod download

# Copie o restante dos arquivos de origem para o diretório de trabalho
COPY . .

# Compile a aplicação Go
RUN go build -o app

# Remova qualquer dependência desnecessária que possa ter sido instalada durante a compilação
RUN apk --no-cache add ca-certificates

# Exponha a porta em que a aplicação irá rodar
# EXPOSE 8080

# Comando para executar a aplicação
CMD ["./app"]
