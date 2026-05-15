package service

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"

	inventoryv1 "go-kozadayev-exercise/task4/api/proto"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Order struct {
	ID        int64
	ProductID string
	Status    string
}

type InventoryService struct {
	inventoryv1.UnimplementedInventoryServiceServer
	mu     sync.RWMutex
	orders map[int64]*Order
	nextID int64
	ch     *amqp.Channel
}

func NewInventoryService() *InventoryService {
	s := &InventoryService{
		orders: make(map[int64]*Order),
		nextID: 1,
	}

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Printf("RabbitMQ connection error: %v. Falling back to internal channel.", err)
		return s
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	s.ch = ch

	q, _ := ch.QueueDeclare("orders", false, false, false, false, nil)

	go func() {
		msgs, _ := ch.Consume(q.Name, "", true, false, false, false, nil)
		for d := range msgs {
			var id int64
			json.Unmarshal(d.Body, &id)
			s.processOrder(id)
		}
	}()

	return s
}

func (s *InventoryService) CreateOrder(ctx context.Context, req *inventoryv1.CreateOrderRequest) (*inventoryv1.CreateOrderResponse, error) {
	s.mu.Lock()
	id := s.nextID
	s.nextID++
	order := &Order{ID: id, ProductID: req.ProductId, Status: "PENDING"}
	s.orders[id] = order
	s.mu.Unlock()

	if s.ch != nil {
		body, _ := json.Marshal(id)
		s.ch.PublishWithContext(ctx, "", "orders", false, false, amqp.Publishing{Body: body})
	}

	return &inventoryv1.CreateOrderResponse{
		Order: &inventoryv1.Order{Id: order.ID, ProductId: order.ProductID, Status: order.Status},
	}, nil
}

func (s *InventoryService) GetOrderStatus(ctx context.Context, req *inventoryv1.GetOrderStatusRequest) (*inventoryv1.GetOrderStatusResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	order, ok := s.orders[req.OrderId]
	if !ok {
		return nil, context.DeadlineExceeded
	}
	return &inventoryv1.GetOrderStatusResponse{Status: order.Status}, nil
}

func (s *InventoryService) processOrder(id int64) {
	time.Sleep(2 * time.Second)
	s.mu.Lock()
	if order, ok := s.orders[id]; ok {
		order.Status = "COMPLETED"
	}
	s.mu.Unlock()
}
