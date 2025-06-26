package port

import (
	"context"

	"com.fernando/pkg/app/ecomm/payment/model"
)

type PaymentRepository interface {
	Insert(ctx context.Context, orderID string, payment *model.Payment) error
	UpdateStatus(ctx context.Context, id string, status model.PaymentStatus) error
}

type PaymentSrv interface {
	// de momento general podria ser UpdateStatusToPay
	UpdateStatus(ctx context.Context, id string, status model.PaymentStatus) error
}
