package main

import (
 "context"
 "fmt"
 "log"
 "net/http"
 "os"
 "os/signal"
 "syscall"
 "time"

 "github.com/VapeShop/cmd/server"
 "github.com/VapeShop/internal/db"
)

func main() {
 // Получение строки подключения к базе данных из окружения
 dbURL := os.Getenv("DATABASE_URL")
 if dbURL == "" {
  log.Fatal("DATABASE_URL is not set")
 }

 // Создание соединения с базой данных
 db, err := db.NewDB(dbURL)
 if err != nil {
  log.Fatal("Error connecting to database:", err)
 }
 defer db.Close()

 // Запуск сервера
 srv := server.NewServer(db) // Передача соединения с базой данных серверу
 go func() {
  // Получение порта из окружения (по умолчанию 8080)
  port := os.Getenv("PORT")
  if port == "" {
   port = "8080"
  }
  if err := srv.Run(":" + port); err != nil && err != http.ErrServerClosed {
   log.Fatalf("error starting server: %s\n", err)
  }
 }()

 // Обработка сигналов для грациозного завершения
 quit := make(chan os.Signal, 1)
 signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
 <-quit
 log.Println("Shutting down server...")

 // Ожидание завершения сервера
 ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
 defer cancel()
 if err := srv.Shutdown(ctx); err != nil {
  log.Fatalf("error shutting down server: %s\n", err)
 }
 log.Println("Server exited gracefully")
}
