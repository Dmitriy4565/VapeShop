package server

import (
  "context"
  "fmt"
  "log"
  "net/http"
  "time"

  "github.com/gin-gonic/gin" // Используем Gin для HTTP-обработки
  "github.com/your-project/internal/db"
  "github.com/your-project/internal/services"
)

// Server структура, описывающая сервер
type Server struct {
  router     *gin.Engine
  categoryService services.CategoryService
}

// NewServer - конструктор сервера
func NewServer(db *db.DB) *Server {
  // Инициализация Gin
  router := gin.Default()

  // Инициализация сервиса (CategoryService)
  categoryService := services.NewCategoryService(db)

  // Создание контроллера 
  categoryController := NewCategoryController(categoryService)

  // Регистрация маршрутов (пример)
  router.GET("/categories", categoryController.GetCategoriesHandler)
  // ... другие маршруты ...

  return &Server{
    router:     router,
    categoryService: categoryService,
  }
}

// Run - запуск сервера
func (s *Server) Run(addr string) error {
  // Запуск сервера
  return http.ListenAndServe(addr, s.router)
}

// NewCategoryController - конструктор контроллера
func NewCategoryController(categoryService services.CategoryService) *CategoryController {
  return &CategoryController{
    categoryService: categoryService,
  }
}

// GetCategoriesHandler - обработчик запроса на получение всех категорий (пример)
func (c *CategoryController) GetCategoriesHandler(ctx *gin.Context) {
  categories, err := c.categoryService.GetAllCategories(ctx)
  if err != nil {
    ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  ctx.JSON(http.StatusOK, categories)
}

// ... другие контроллеры ...
