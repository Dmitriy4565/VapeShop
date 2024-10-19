package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/services/storeService"
)

// StoreController - контроллер для обработки запросов к магазинам
type StoreController struct {
	storeService *storeService.StoreService
	validate     *validator.Validate
}

// NewStoreController - конструктор контроллера
func NewStoreController(storeService *storeService.StoreService) *StoreController {
	return &StoreController{
		storeService: storeService,
		validate:     validator.New(),
	}
}

// GetStoresHandler - обработчик запроса на получение всех магазинов
func (c *StoreController) GetStoresHandler(w http.ResponseWriter, r *http.Request) {
	stores, err := c.storeService.GetAllStores()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(stores)
}

// GetStoreByIDHandler - обработчик запроса на получение магазина по ID
func (c *StoreController) GetStoreByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID магазина не указан", http.StatusBadRequest)
		return
	}

	store, err := c.storeService.GetStoreByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(store)
}

// CreateStoreHandler - обработчик запроса на создание нового магазина
func (c *StoreController) CreateStoreHandler(w http.ResponseWriter, r *http.Request) {
	var store storeService.Store
	err := json.NewDecoder(r.Body).Decode(&store)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Валидация данных
	err = c.validate.Struct(store)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newStore, err := c.storeService.CreateStore(store)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newStore)
}

// UpdateStoreHandler - обработчик запроса на обновление магазина
func (c *StoreController) UpdateStoreHandler(w http.ResponseWriter, r *http.Request) {
	var store storeService.Store
	err := json.NewDecoder(r.Body).Decode(&store)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Валидация данных
	err = c.validate.Struct(store)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.storeService.UpdateStore(store)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteStoreHandler - обработчик запроса на удаление магазина
func (c *StoreController) DeleteStoreHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID магазина не указан", http.StatusBadRequest)
		return
	}

	err := c.storeService.DeleteStore(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
