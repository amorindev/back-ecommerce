package validate

import (
	"errors"

	"com.fernando/pkg/app/ecomm/orders/core"
)

// * Datos requeridos
// * UserID, Total, Items[], DeliveryType
// * Cada OrderItem debe tener ProductItemID, Quantity, y Price.
// * Si es pickup, validar PhoneID, AddressID, StoreID.
// * Si es delivery, validar PhoneID, AddressID.

// * Datos que NO se guardan directamente en la base de datos (no deben insertarse):
// * Order.Items → es un agregado, posiblemente en otra colección o servicio.
// * Order.PaymentAgregate → no se guarda directamente en la orden.
// * Payment.ProviderPaymentID → lo establece el webhook externo.
// * Payment.Status → lo actualiza el webhook (puede iniciar en pending). me parece que al crear sera pending
// por defecto
// * Datos que deben ser ocultos o no retornados directamente en la respuesta:
// * Payment.ProviderPaymentID (si quieres limitar exposición).
// * Cualquier información sensible como dirección exacta (si aplica por privacidad). pude ser
// ocultar o deslizar  para ver el pago
// * Datos que podrían requerir encriptación
// No se observa necesidad inmediata, pero si se incluyera tarjeta o datos bancarios, deben ir cifrados o nunca almacenarse directamente.

func ValidateCreateOrder(req core.CreateOrderReq) (bool, error) {
	if len(req.ProductItems) == 0 {
		return false, errors.New("order must contain at least one item")
	}
	// deberia ser igual a la suma de las productos de cada item de momento mayor a cero
	if req.Total <= 0 {
		return false, errors.New("total debe ser mayor a cero")
	}
	if req.DeliveryType == "" {
		return false, errors.New("delivery type field is required")
	}

	if req.DeliveryType == "delivery" {
		if req.DeliveryInfo == nil {
			return false, errors.New("delivery field is required because deliverytype is delivery")
		}
		// *validar campos
		if req.DeliveryInfo.AddressID == "" {
			return false, errors.New("address id is required en delivery type")
		}
		if req.DeliveryInfo.PhoneID == "" {
			return false, errors.New("phone id is required en delivery type")
		}
		/* if req.DeliveryInfo.Reference == "" {
			return false, errors.New("reference filed is required en delivery type")
		} */
	}
	if req.DeliveryType == "pickup" {
		if req.PickupInfo == nil {
			return false, errors.New("pickup field is required because delivery type is pickup")
		}
		// * validar campos
		if req.PickupInfo.PhoneID == "" {
			return false, errors.New("phone id is required en delivery type is pickup")
		}
		if req.PickupInfo.AddressID == "" {
			return false, errors.New("address id is required en delivery type is pickup")
		}
		if req.PickupInfo.StoreID == "" {
			return false, errors.New("store id is required en delivery type is pickup")
		}
	}

	if req.Payment == nil {
		return false, errors.New("payment fiel is required")
	}
	if req.Payment.Currency == "" {
		return false, errors.New("payment -> currency fiel is required")
	}
	if req.Payment.PaymentMethod == "" {
		return false, errors.New("payment -> payment_method fiel is required")
	}
	return true, nil
}
