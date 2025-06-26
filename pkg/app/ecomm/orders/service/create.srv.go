package service

import (
	"context"
	"time"

	"com.fernando/pkg/app/ecomm/orders/model"
	paymentModel "com.fernando/pkg/app/ecomm/payment/model"
)

// ! Descontar el stock del product aqui o desde el weebhook
// ! como calcular el costo de envio desde aqui o desde mobile backend o mobile
// ! el men de grpc y otros videos tiee un reo de tiempo real en el backend golang
// ! cuando hacer el seguimiento desde el backend
func (s *Service) Create(ctx context.Context, order *model.Order) error {
	// ? verificar si el usuario existe?
	// ? ver si el stock hay stock del producto
	// ? en el getprducts filtrar por productos activos
	// ? una transaccion o trigger para descontar el stock

	//* calcualr denuevo asi le notificamos al cliente
	// de momento no por que tenemos discount en el frontend por delvery
	/* var total float64
	for _, item := range order.Items {
		// ? verificar los nulos
		total += *item.Price * float64(*item.Quantity)
	} */

	// * verificar si desde el frontend se caluclo bien
	// * lo hago por que al usuario se mostro un precio y no queiro que se registre otro
	// * me parece que si se envia desde el frontend, esto no es en validate handler hay dos tipos de validates
	/* if total != *order.Total {
		return errors.New("total calculado es diferente a order total")
	} */

	now := time.Now().UTC()
	order.CreatedAt = &now
	order.UpdatedAt = &now

	currency := "usd"
	order.PaymentAgt.Currency = &currency // ? deberia estar en en mobile y backend
	order.PaymentAgt.Status = paymentModel.PaymentPending
	// ! mejor seria pasarlo antes por que aqui se va usar stripe paypal u otro
	order.PaymentAgt.PaymentMethod = paymentModel.MethodStripe
	// *payment intent,
	// * dos flujos que cuando se cree el payment intent se guarde payment en la base de datos
	// * y en este punto tocarái actualizar y si llega el weeebhook de nuevo actualizar
	// * no me parece lo mejor - ver el flujo complero me parece que cambia muchas cosas
	// * despues de crear el payment intent lo adjuntemos en metadata de stripe
	// * y lo obtenemos del weebhook o talves ya viene y no necesitamos adjuntarlo en metadata
	// * entonces esto desde el weebhook junto al pagado estado
	//order.PaymentAgregate.ProviderPaymentID
	order.PaymentAgt.CreatedAt = &now
	order.PaymentAgt.UpdatedAt = &now

	err := s.OrderTx.Create(ctx, order)
	if err != nil {
		return err
	}

	// enviar el correo que se hizo una compra
	return nil
}
