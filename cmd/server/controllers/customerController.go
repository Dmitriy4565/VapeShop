package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/services/customerService"
)

// CustomerController - контроллер для обработки запросов к клиентам
type CustomerController struct {
	customerService *customerService.CustomerService
	validate        *validator.Validate
}

// NewCustomerController - конструктор контроллера
func NewCustomerController(customerService *customerService.CustomerService) *CustomerController {
	return &CustomerController{
		customerService: customerService,
		validate:        validator.New(),
	}
}

// GetCustomersHandler - обработчик запроса на получение всех клиентов
func (c *CustomerController) GetCustomersHandler(w http.ResponseWriter, r *http.Request) {
	customers, err := c.customerService.GetAllCustomers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(customers)
}

// GetCustomerByIDHandler - обработчик запроса на получение клиента по ID
func (c *CustomerController) GetCustomerByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID клиента не указан", http.StatusBadRequest)
		return
	}

	customer, err := c.customerService.GetCustomerByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(customer)
}

// CreateCustomerHandler - обработчик запроса на создание нового клиента
func (c *CustomerController) CreateCustomerHandler(w http.ResponseWriter, r *http.Request) {
	var customer customerService.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Валидация данных
	err = c.validate.Struct(customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newCustomer, err := c.customerService.CreateCustomer(customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newCustomer)
}

// UpdateCustomerHandler - обработчик запроса на обновление клиента
func (c *CustomerController) UpdateCustomerHandler(w http.ResponseWriter, r *http.Request) {
	var customer customerService.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Валидация данных
	err = c.validate.Struct(customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.customerService.UpdateCustomer(customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteCustomerHandler - обработчик запроса на удаление клиента
func (c *CustomerController) DeleteCustomerHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID клиента не указан", http.StatusBadRequest)
		return
	}

	err := c.customerService.DeleteCustomer(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
