package database

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

var DB *sql.DB

func InitDB() error {
	host := getEnvWithDefault("DB_HOST", "localhost")
	port := getEnvWithDefault("DB_PORT", "5432")
	user := getEnvWithDefault("DB_USER", "dominote")
	password := getEnvWithDefault("DB_PASSWORD", "dominote_pass")
	dbname := getEnvWithDefault("DB_NAME", "dominote_db")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbname)

	var err error
	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Printf("Error conectando a la base de datos: %v", err)
		return err
	}

	if err := DB.Ping(); err != nil {
		log.Printf("Error haciendo ping a la base de datos: %v", err)
		return err
	}

	log.Println("Conexión a la base de datos establecida exitosamente")

	// Configurar goose para usar el sistema de archivos embebido
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("error configurando el dialecto de goose: %w", err)
	}

	log.Println("Ejecutando migraciones de base de datos...")
	if err := goose.Up(DB, "migrations"); err != nil {
		return fmt.Errorf("error ejecutando migraciones: %w", err)
	}
	log.Println("Migraciones ejecutadas correctamente")

	return nil
}

// getEnvWithDefault obtiene una variable de entorno o retorna un valor por defecto si está vacía.
func getEnvWithDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
