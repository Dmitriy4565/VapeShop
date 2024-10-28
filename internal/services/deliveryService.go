package services //блять я рот ебал этих бд, не забудь потом пройтись по всему коду и заменить

import (
	"context"
	"errors"
	"time"

	"database/sql"
)

type Delivery struct {
	ID         string    `json:"id"`
	CustomerID string    `json:"customerId"`
	StoreID    string    `json:"storeId"`
	Address    string    `json:"address"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type DeliveryService interface {
	GetAllDeliveries() ([]Delivery, error)
	GetDeliveryByID(id string) (*Delivery, error)
	CreateDelivery(delivery Delivery) (*Delivery, error)
	UpdateDelivery(delivery Delivery) error
	DeleteDelivery(id string) error
}

type DeliveryServiceImpl struct {
	db *sql.DB // Ссылка на объект базы данных
}

func NewDeliveryService(db *sql.DB) *DeliveryServiceImpl {
	return &DeliveryServiceImpl{
		db: db,
	}
}

func (s *DeliveryServiceImpl) GetAllDeliveries() ([]Delivery, error) {
	rows, err := s.db.QueryContext(context.Background(), "SELECT * FROM deliveries")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var deliveries []Delivery
	for rows.Next() {
		var delivery Delivery
		if err := rows.Scan(&delivery.ID, &delivery.CustomerID, &delivery.StoreID, &delivery.Address, &delivery.Status, &delivery.CreatedAt, &delivery.UpdatedAt); err != nil {
			return nil, err
		}
		deliveries = append(deliveries, delivery)
	}

	return deliveries, nil
}

func (s *DeliveryServiceImpl) GetDeliveryByID(id string) (*Delivery, error) {
	var delivery Delivery
	err := s.db.QueryRowContext(context.Background(), "SELECT * FROM deliveries WHERE id = $1", id).Scan(&delivery.ID, &delivery.CustomerID, &delivery.StoreID, &delivery.Address, &delivery.Status, &delivery.CreatedAt, &delivery.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("доставка не найдена")
		}
		return nil, err
	}
	return &delivery, nil
}

func (s *DeliveryServiceImpl) CreateDelivery(delivery Delivery) (*Delivery, error) {
	ctx := context.Background()
	result, err := s.db.ExecContext(ctx, "INSERT INTO deliveries (customerId, storeId, address, status) VALUES ($1, $2, $3, $4)", delivery.CustomerID, delivery.StoreID, delivery.Address, delivery.Status)
	if err != nil {
		return nil, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	delivery.ID = lastInsertID
	return &delivery, nil
}

func (s *DeliveryServiceImpl) UpdateDelivery(delivery Delivery) error {
	ctx := context.Background()
	_, err := s.db.ExecContext(ctx, "UPDATE deliveries SET customerId = $1, storeId = $2, address = $3, status = $4 WHERE id = $5", delivery.CustomerID, delivery.StoreID, delivery.Address, delivery.Status, delivery.ID)
	return err
}

func (s *DeliveryServiceImpl) DeleteDelivery(id string) error {
	ctx := context.Background()
	_, err := s.db.ExecContext(ctx, "DELETE FROM deliveries WHERE id = $1", id)
	return err
}
