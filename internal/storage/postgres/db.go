package postgres

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// NewDB создаёт подключение к PostgreSQL через переменные окружения
func NewDB() (*sql.DB, error) {
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    if dbHost == "" {
        dbHost = "localhost"
    }
    if dbPort == "" {
        dbPort = "5432"
    }

    connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
        dbUser, dbPassword, dbHost, dbPort, dbName)

    return sql.Open("postgres", connStr)
}

// InitDB создаёт таблицы, если их нет
func InitDB(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS movies (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		slug TEXT UNIQUE NOT NULL,
		image TEXT,
		description TEXT,
		year INT,
		director TEXT
	);

	CREATE TABLE IF NOT EXISTS actors (
		id SERIAL PRIMARY KEY,
		name TEXT UNIQUE NOT NULL
	);

	CREATE TABLE IF NOT EXISTS movie_actors (
		movie_id INT REFERENCES movies(id) ON DELETE CASCADE,
		actor_id INT REFERENCES actors(id) ON DELETE CASCADE,
		PRIMARY KEY (movie_id, actor_id)
	);

	CREATE TABLE IF NOT EXISTS genres (
		id SERIAL PRIMARY KEY,
		name TEXT UNIQUE NOT NULL
	);

	CREATE TABLE IF NOT EXISTS movie_genres (
		movie_id INT REFERENCES movies(id) ON DELETE CASCADE,
		genre_id INT REFERENCES genres(id) ON DELETE CASCADE,
		PRIMARY KEY (movie_id, genre_id) 
	);

	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	);
	`

	_, err := db.Exec(query)
	return err
}