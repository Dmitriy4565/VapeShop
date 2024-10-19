package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/services/categoryService"
)

// CategoryController - контроллер для обработки запросов к категориям
type CategoryController struct {
	categoryService *categoryService.CategoryService
	validate        *validator.Validate
}

// NewCategoryController - конструктор контроллера
func NewCategoryController(categoryService *categoryService.CategoryService) *CategoryController {
	return &CategoryController{
		categoryService: categoryService,
		validate:        validator.New(),
	}
}

// GetCategoriesHandler - обработчик запроса на получение всех категорий
func (c *CategoryController) GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := c.categoryService.GetAllCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(categories)
}

// CreateCategoryHandler - обработчик запроса на создание новой категории
func (c *CategoryController) CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var category categoryService.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Валидация данных
	err = c.validate.Struct(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newCategory, err := c.categoryService.CreateCategory(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newCategory)
}

// UpdateCategoryHandler - обработчик запроса на обновление категории
func (c *CategoryController) UpdateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var category categoryService.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Валидация данных
	err = c.validate.Struct(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.categoryService.UpdateCategory(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteCategoryHandler - обработчик запроса на удаление категории
func (c *CategoryController) DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID категории не указан", http.StatusBadRequest)
		return
	}

	err := c.categoryService.DeleteCategory(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
