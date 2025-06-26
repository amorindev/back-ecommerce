package port

import (
	"context"

	"com.fernando/pkg/app/ecomm/orders/model"
)

type OrderRepository interface {
	// es engorroso que se llame create y para buscar el archivo se combinan entre servicio handler
	Insert(ctx context.Context, order *model.Order) error
	InsertOrderProduct(ctx context.Context, orderID string, items []*model.OrderItem) error
	// es obio que es con el userID o no se, recuerda en el repos filtrar
	GetAll(ctx context.Context, userID string) ([]*model.Order, error)
}

type OrderService interface {
	// enserio chat-gpt save-encerio
	//CreateOrder(ctx context.Context, userID string, items []OrderItem) (*Order, error)
	Create(ctx context.Context, order *model.Order) error
	// páginacion , userID desde el token
	GetAll(ctx context.Context, userID string) ([]*model.Order, error)
}

type OrderTransaction interface {
	// * de momento vamos dos collecciones o tablas order yorder_product
	// * ver si se va agregar el Payment como agregado
	// * ver que coincida todo
	Create(ctx context.Context, order *model.Order) error
}
