package service

import "com.fernando/pkg/app/ecomm/orders/port"

var _ port.OrderService = &Service{}

type Service struct {
	OrderRepo port.OrderRepository // se usara para los get
	// ?? Ad Payment service de la tabla payment si es asi es una transaccion
	OrderTx port.OrderTransaction
}

func NewService(orderRepo port.OrderRepository, orderTx port.OrderTransaction) *Service {
	return &Service{
		OrderRepo: orderRepo,
		OrderTx: orderTx,
	}
}
