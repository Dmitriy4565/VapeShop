package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/services/productService"
)

// ProductController - контроллер для обработки запросов к продуктам
type ProductController struct {
	productService *productService.ProductService
	validate       *validator.Validate
}

// NewProductController - конструктор контроллера
func NewProductController(productService *productService.ProductService) *ProductController {
	return &ProductController{
		productService: productService,
		validate:       validator.New(),
	}
}

// GetProductsHandler - обработчик запроса на получение всех продуктов
func (c *ProductController) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := c.productService.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}

// GetProductByIDHandler - обработчик запроса на получение продукта по ID
func (c *ProductController) GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID продукта не указан", http.StatusBadRequest)
		return
	}

	product, err := c.productService.GetProductByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

// CreateProductHandler - обработчик запроса на создание нового продукта
func (c *ProductController) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product productService.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Валидация данных
	err = c.validate.Struct(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newProduct, err := c.productService.CreateProduct(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newProduct)
}

// UpdateProductHandler - обработчик запроса на обновление продукта
func (c *ProductController) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product productService.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Валидация данных
	err = c.validate.Struct(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.productService.UpdateProduct(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteProductHandler - обработчик запроса на удаление продукта
func (c *ProductController) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID продукта не указан", http.StatusBadRequest)
		return
	}

	err := c.productService.DeleteProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
