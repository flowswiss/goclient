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

const (
	OrderStatusCreated = iota + 1
	OrderStatusProcessing
	OrderStatusSucceeded
	OrderStatusFailed
)

type OrderStatus struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Ordering struct {
	Ref string `json:"ref"`
}

func (o Ordering) ExtractIdentifier() (int, error) {
	regex := regexp.MustCompile("/orders/(\\d+)$")
	data := regex.FindStringSubmatch(o.Ref)

	id, err := strconv.ParseInt(data[1], 10, 64)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

type Order struct {
	Id     int         `json:"id"`
	Status OrderStatus `json:"status"`
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

func (o OrderService) WaitForCompletion(ctx context.Context, ordering Ordering) error {
	id, err := ordering.ExtractIdentifier()
	if err != nil {
		return fmt.Errorf("extract ordering identifier: %w", err)
	}

	for {
		order, err := o.Get(ctx, id)
		if err != nil {
			return err
		}

		if order.Status.Id == OrderStatusSucceeded {
			return nil
		}

		if order.Status.Id == OrderStatusFailed {
			return ErrOrderFailed
		}

		<-time.After(time.Second)
	}
}

const ordersSegment = "/v4/orders"

func getSpecificOrderPath(orderId int) string {
	return goclient.Join(ordersSegment, orderId)
}
