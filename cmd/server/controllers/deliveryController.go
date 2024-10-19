package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/services/deliveryService"
)

// DeliveryController - контроллер для обработки запросов к доставке
type DeliveryController struct {
	deliveryService *deliveryService.DeliveryService
	validate        *validator.Validate
}

// NewDeliveryController - конструктор контроллера
func NewDeliveryController(deliveryService *deliveryService.DeliveryService) *DeliveryController {
	return &DeliveryController{
		deliveryService: deliveryService,
		validate:        validator.New(),
	}
}

// GetDeliveriesHandler - обработчик запроса на получение всех доставок
func (c *DeliveryController) GetDeliveriesHandler(w http.ResponseWriter, r *http.Request) {
	deliveries, err := c.deliveryService.GetAllDeliveries()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(deliveries)
}

// GetDeliveryByIDHandler - обработчик запроса на получение доставки по ID
func (c *DeliveryController) GetDeliveryByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID доставки не указан", http.StatusBadRequest)
		return
	}

	delivery, err := c.deliveryService.GetDeliveryByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(delivery)
}

// CreateDeliveryHandler - обработчик запроса на создание новой доставки
func (c *DeliveryController) CreateDeliveryHandler(w http.ResponseWriter, r *http.Request) {
	var delivery deliveryService.Delivery
	err := json.NewDecoder(r.Body).Decode(&delivery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Валидация данных
	err = c.validate.Struct(delivery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newDelivery, err := c.deliveryService.CreateDelivery(delivery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newDelivery)
}

// UpdateDeliveryHandler - обработчик запроса на обновление доставки
func (c *DeliveryController) UpdateDeliveryHandler(w http.ResponseWriter, r *http.Request) {
	var delivery deliveryService.Delivery
	err := json.NewDecoder(r.Body).Decode(&delivery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Валидация данных
	err = c.validate.Struct(delivery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.deliveryService.UpdateDelivery(delivery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteDeliveryHandler - обработчик запроса на удаление доставки
func (c *DeliveryController) DeleteDeliveryHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID доставки не указан", http.StatusBadRequest)
		return
	}

	err := c.deliveryService.DeleteDelivery(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
