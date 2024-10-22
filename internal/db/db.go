package db

import (
 "database/sql"
 "fmt"

 _ "github.com/lib/pq" // Драйвер PostgreSQL
)

// DB структура для работы с базой данных
type DB struct {
 *sql.DB
}

// NewDB - конструктор для создания нового соединения с базой данных
func NewDB(dbURL string) (*DB, error) {
 db, err := sql.Open("postgres", dbURL)
 if err != nil {
  return nil, fmt.Errorf("error opening database connection: %w", err)
 }

 // Проверка соединения
 if err := db.Ping(); err != nil {
  return nil, fmt.Errorf("error pinging database: %w", err)
 }

 return &DB{db}, nil
}

// Close - закрывает соединение с базой данных
func (db *DB) Close() error {
 return db.DB.Close()
}
