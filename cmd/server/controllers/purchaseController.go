package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/services/purchaseService"
)

// PurchaseController - контроллер для обработки запросов к покупкам
type PurchaseController struct {
	purchaseService *purchaseService.PurchaseService
	validate        *validator.Validate
}

// NewPurchaseController - конструктор контроллера
func NewPurchaseController(purchaseService *purchaseService.PurchaseService) *PurchaseController {
	return &PurchaseController{
		purchaseService: purchaseService,
		validate:        validator.New(),
	}
}

// GetPurchasesHandler - обработчик запроса на получение всех покупок
func (c *PurchaseController) GetPurchasesHandler(w http.ResponseWriter, r *http.Request) {
	purchases, err := c.purchaseService.GetAllPurchases()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(purchases)
}

// GetPurchaseByIDHandler - обработчик запроса на получение покупки по ID
func (c *PurchaseController) GetPurchaseByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID покупки не указан", http.StatusBadRequest)
		return
	}

	purchase, err := c.purchaseService.GetPurchaseByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(purchase)
}

// CreatePurchaseHandler - обработчик запроса на создание новой покупки
func (c *PurchaseController) CreatePurchaseHandler(w http.ResponseWriter, r *http.Request) {
	var purchase purchaseService.Purchase
	err := json.NewDecoder(r.Body).Decode(&purchase)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Валидация данных
	err = c.validate.Struct(purchase)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newPurchase, err := c.purchaseService.CreatePurchase(purchase)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newPurchase)
}

// UpdatePurchaseHandler - обработчик запроса на обновление покупки
func (c *PurchaseController) UpdatePurchaseHandler(w http.ResponseWriter, r *http.Request) {
	var purchase purchaseService.Purchase
	err := json.NewDecoder(r.Body).Decode(&purchase)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Валидация данных
	err = c.validate.Struct(purchase)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.purchaseService.UpdatePurchase(purchase)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeletePurchaseHandler - обработчик запроса на удаление покупки
func (c *PurchaseController) DeletePurchaseHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID покупки не указан", http.StatusBadRequest)
		return
	}

	err := c.purchaseService.DeletePurchase(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
