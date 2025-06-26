package port

import (
	"context"

	"com.fernando/pkg/app/ecomm/pickup-orders/model"
)

type PickupOrderRepository interface {
	Insert(ctx context.Context, pickupO *model.PickupOrder) error
}
