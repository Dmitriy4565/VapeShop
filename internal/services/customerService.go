package services

import (
	"context"
	"errors"
	"time"

	// Импортируйте пакеты для работы с базой данных
	"database/sql"
)

// Customer - структура, представляющая данные о клиенте
type Customer struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// CustomerService - интерфейс сервиса клиентов
type CustomerService interface {
	// Методы для работы с клиентами:
	GetAllCustomers() ([]Customer, error)
	GetCustomerByID(id string) (*Customer, error)
	CreateCustomer(customer Customer) (*Customer, error)
	UpdateCustomer(customer Customer) error
	DeleteCustomer(id string) error
}

// CustomerServiceImpl - реализация сервиса клиентов
type CustomerServiceImpl struct {
	db *sql.DB // Ссылка на объект базы данных
}

// NewCustomerService - конструктор сервиса клиентов
func NewCustomerService(db *sql.DB) *CustomerServiceImpl {
	return &CustomerServiceImpl{
		db: db,
	}
}

// GetAllCustomers - получение всех клиентов
func (s *CustomerServiceImpl) GetAllCustomers() ([]Customer, error) {
	rows, err := s.db.QueryContext(context.Background(), "SELECT * FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []Customer
	for rows.Next() {
		var customer Customer
		if err := rows.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Phone, &customer.Address, &customer.CreatedAt, &customer.UpdatedAt); err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

// GetCustomerByID - получение клиента по ID
func (s *CustomerServiceImpl) GetCustomerByID(id string) (*Customer, error) {
	var customer Customer
	err := s.db.QueryRowContext(context.Background(), "SELECT * FROM customers WHERE id = $1", id).Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Phone, &customer.Address, &customer.CreatedAt, &customer.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("клиент не найден")
		}
		return nil, err
	}
	return &customer, nil
}

// CreateCustomer - создание нового клиента
func (s *CustomerServiceImpl) CreateCustomer(customer Customer) (*Customer, error) {
	ctx := context.Background()
	result, err := s.db.ExecContext(ctx, "INSERT INTO customers (name, email, phone, address) VALUES ($1, $2, $3, $4)", customer.Name, customer.Email, customer.Phone, customer.Address)
	if err != nil {
		return nil, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	customer.ID = lastInsertID
	return &customer, nil
}

// UpdateCustomer - обновление клиента
func (s *CustomerServiceImpl) UpdateCustomer(customer Customer) error {
	ctx := context.Background()
	_, err := s.db.ExecContext(ctx, "UPDATE customers SET name = $1, email = $2, phone = $3, address = $4 WHERE id = $5", customer.Name, customer.Email, customer.Phone, customer.Address, customer.ID)
	return err
}

// DeleteCustomer - удаление клиента
func (s *CustomerServiceImpl) DeleteCustomer(id string) error {
	ctx := context.Background()
	_, err := s.db.ExecContext(ctx, "DELETE FROM customers WHERE id = $1", id)
	return err
}
