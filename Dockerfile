# Build stage
FROM golang:1.24-alpine AS builder

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar go.mod y go.sum para descargar las dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -o dominote-api main.go

# Final stage (imagen más ligera)
FROM alpine:latest

WORKDIR /app

# Copiar el ejecutable desde la etapa de construcción
COPY --from=builder /app/dominote-api .

# Exponer el puerto por defecto de Gin
EXPOSE 8080

# Comando por defecto para ejecutar la aplicación
CMD ["./dominote-api"]
