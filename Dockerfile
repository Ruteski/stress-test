# Usar uma imagem base do Go para compilar o aplicativo
FROM golang:1.23.4 AS builder

# Definir o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copiar o código-fonte para o contêiner
COPY . .

# Compilar o aplicativo
RUN go build -o stresstest .

# Usar uma imagem base do Go
FROM golang:1.23.4

# Definir o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copiar o código-fonte para o contêiner
COPY . .

# Compilar o aplicativo
RUN go build -o stresstest .

# Definir o comando padrão para executar o aplicativo
ENTRYPOINT ["./stresstest", "loadtest"]