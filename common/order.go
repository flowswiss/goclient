package common

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/flowswiss/goclient"
)

var ErrOrderFailed = errors.New("order failed")

var orderIdentifierRegex = regexp.MustCompile(`/orders/(\d+)$`)

const (
	OrderStatusCreated = iota + 1
	OrderStatusProcessing
	OrderStatusSucceeded
	OrderStatusFailed
)

type OrderStatus struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Ordering struct {
	Ref string `json:"ref"`
}

func (o Ordering) ExtractIdentifier() (int, error) {
	data := orderIdentifierRegex.FindStringSubmatch(o.Ref)

	if len(data) < 2 {
		return 0, fmt.Errorf("invalid order identifier")
	}

	id, err := strconv.ParseInt(data[1], 10, 64)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

type Order struct {
	ID        int         `json:"id"`
	Status    OrderStatus `json:"status"`
	Product   Product     `json:"product_instance"`
	CreatedAt Time        `json:"created_at"`
}

type OrderService struct {
	client goclient.Client
}

func NewOrderService(client goclient.Client) OrderService {
	return OrderService{client: client}
}

func (o OrderService) Get(ctx context.Context, id int) (order Order, err error) {
	err = o.client.Get(ctx, getSpecificOrderPath(id), &order)
	return
}

// Deprecated: use WaitUntilProcessed instead
func (o OrderService) WaitForCompletion(ctx context.Context, ordering Ordering) error {
	_, err := o.WaitUntilProcessed(ctx, ordering)
	return err
}

func (o OrderService) WaitUntilProcessed(ctx context.Context, ordering Ordering) (order Order, err error) {
	id, err := ordering.ExtractIdentifier()
	if err != nil {
		return Order{}, fmt.Errorf("extract ordering identifier: %w", err)
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		order, err = o.Get(ctx, id)
		if err != nil {
			return Order{}, fmt.Errorf("fetch order: %w", err)
		}

		if order.Status.ID == OrderStatusSucceeded {
			return order, nil
		}

		if order.Status.ID == OrderStatusFailed {
			return order, ErrOrderFailed
		}

		select {
		case <-ticker.C:

		case <-ctx.Done():
			return Order{}, ctx.Err()
		}
	}
}

const ordersSegment = "/v4/orders"

func getSpecificOrderPath(orderID int) string {
	return goclient.Join(ordersSegment, orderID)
}
