package mongo

import (
	"com.fernando/pkg/app/ecomm/delivery-orders/port"
	orderPort "com.fernando/pkg/app/ecomm/orders/port"
	paymentPort "com.fernando/pkg/app/ecomm/payment/port"
	pickupP "com.fernando/pkg/app/ecomm/pickup-orders/port"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ orderPort.OrderTransaction = &Transaction{}

// * demomento las collecciones de muchos a muchos ambos tienen la colleccion intermedia
// * el otro tema es que payment se creara desde order service
// * o debo inyectarle un PaymentSrv y no un repo de momento sensicillo
type Transaction struct {
	Client       *mongo.Client
	OrderRepo    orderPort.OrderRepository
	PaymentRepo  paymentPort.PaymentRepository
	PickupRepo   pickupP.PickupOrderRepository
	DeliveryRepo port.DeliveryOrderRepository
}

func NewTransaction(client *mongo.Client, orderRepo orderPort.OrderRepository, paymentRepo paymentPort.PaymentRepository, pickupRepo pickupP.PickupOrderRepository, deliveryRepo port.DeliveryOrderRepository) *Transaction {
	return &Transaction{
		Client:       client,
		OrderRepo:    orderRepo,
		PaymentRepo:  paymentRepo,
		PickupRepo:   pickupRepo,
		DeliveryRepo: deliveryRepo,
	}
}
