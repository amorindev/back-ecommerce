package port

import (
	"context"

	"com.fernando/pkg/app/ecomm/delivery-orders/model"
)

type DeliveryOrderRepository interface {
	Insert(ctx context.Context, deliveryO *model.DeliveryOrder) error
}