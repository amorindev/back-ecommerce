package port

import (
	"context"

	"com.fernando/pkg/app/stripe_customer/model"
)

// no seria mejor stripe provide repo ver
type StripeCustomerRepo interface {
	Insert(ctx context.Context, pProvider *model.StripeCustomer) error
	GetByUserID(ctx context.Context, userID string) (*model.StripeCustomer, error)
}
