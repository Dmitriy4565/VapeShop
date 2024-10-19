package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/services/manufacturerService"
)

// ManufacturerController - контроллер для обработки запросов к производителям
type ManufacturerController struct {
	manufacturerService *manufacturerService.ManufacturerService
	validate            *validator.Validate
}

// NewManufacturerController - конструктор контроллера
func NewManufacturerController(manufacturerService *manufacturerService.ManufacturerService) *ManufacturerController {
	return &ManufacturerController{
		manufacturerService: manufacturerService,
		validate:            validator.New(),
	}
}

// GetManufacturersHandler - обработчик запроса на получение всех производителей
func (c *ManufacturerController) GetManufacturersHandler(w http.ResponseWriter, r *http.Request) {
	manufacturers, err := c.manufacturerService.GetAllManufacturers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(manufacturers)
}

// GetManufacturerByIDHandler - обработчик запроса на получение производителя по ID
func (c *ManufacturerController) GetManufacturerByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID производителя не указан", http.StatusBadRequest)
		return
	}

	manufacturer, err := c.manufacturerService.GetManufacturerByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(manufacturer)
}

// CreateManufacturerHandler - обработчик запроса на создание нового производителя
func (c *ManufacturerController) CreateManufacturerHandler(w http.ResponseWriter, r *http.Request) {
	var manufacturer manufacturerService.Manufacturer
	err := json.NewDecoder(r.Body).Decode(&manufacturer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Валидация данных
	err = c.validate.Struct(manufacturer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newManufacturer, err := c.manufacturerService.CreateManufacturer(manufacturer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newManufacturer)
}

// UpdateManufacturerHandler - обработчик запроса на обновление производителя
func (c *ManufacturerController) UpdateManufacturerHandler(w http.ResponseWriter, r *http.Request) {
	var manufacturer manufacturerService.Manufacturer
	err := json.NewDecoder(r.Body).Decode(&manufacturer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Валидация данных
	err = c.validate.Struct(manufacturer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.manufacturerService.UpdateManufacturer(manufacturer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteManufacturerHandler - обработчик запроса на удаление производителя
func (c *ManufacturerController) DeleteManufacturerHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID производителя не указан", http.StatusBadRequest)
		return
	}

	err := c.manufacturerService.DeleteManufacturer(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
