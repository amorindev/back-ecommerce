package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"com.fernando/cmd/api/message"
	"com.fernando/cmd/api/middlewares"
	"com.fernando/internal/claim"
	deliveryM "com.fernando/pkg/app/ecomm/delivery-orders/model"
	"com.fernando/pkg/app/ecomm/orders/core"
	orderModel "com.fernando/pkg/app/ecomm/orders/model"
	"com.fernando/pkg/app/ecomm/orders/validate"
	paymentModel "com.fernando/pkg/app/ecomm/payment/model"
	pickupM "com.fernando/pkg/app/ecomm/pickup-orders/model"
)

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// ! sacar el userID desde el token
	// ! en todos los handler que requierane  el user id desde el token

	accessTokenClaim, ok := r.Context().Value(middlewares.AccessTokenClaimsIDKey).(*claim.AccessTokenClaims)
	if !ok {
		// bad request o unhotorixed o internal server?
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "failed to parse claims - AccessTokenClaims"})
		return
	}

	if accessTokenClaim.Subject == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "userid-not-found-on-claim"})
		return
	}

	var req core.CreateOrderReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// invalid request body
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	defer r.Body.Close()

	_, err = validate.ValidateCreateOrder(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	// ver como hacer herencia era order dentro de pickup y dentreo de delivery struct

	//orderItems := make([]*orderModel.Order, len(req.ProductItems))
	// hacer un make
	// * deverias verificar que no sea nulo los campos requeridos crear otro validate function
	// * y usarlo dentro
	// ! es necesario verificar par eereos inesperados
	var orderItems []*orderModel.OrderItem
	for _, oItemReq := range req.ProductItems {
		var oItem orderModel.OrderItem
		oItem.ProductItemID = oItemReq.ProductItemID
		oItem.Price = &oItemReq.Price
		oItem.Quantity = &oItemReq.Quantity
		orderItems = append(orderItems, &oItem)
	}

	// * esta bien crearlo aqui para retornarlo?  o donde?
	payment := &paymentModel.Payment{
		Currency:      &req.Payment.Currency,
		PaymentMethod: paymentModel.PaymentMethod(req.Payment.PaymentMethod),
	}

	// * las demás validaciones
	order := &orderModel.Order{
		UserID:       accessTokenClaim.Subject,
		Total:        &req.Total,
		DeliveryType: req.DeliveryType,
		Items:        orderItems,
		PaymentAgt:   payment,
	}

	// * DeliveryType
	if req.DeliveryType == "pickup" {
		pickup := &pickupM.PickupOrder{
			AddressID: req.PickupInfo.AddressID,
			PhoneID:   req.PickupInfo.PhoneID,
			StoreID:   req.PickupInfo.StoreID,
		}
		order.PickupAgt = pickup
	}

	if req.DeliveryType == "delivery" {
		//fmt.Printf("is delivery")
		delivery := &deliveryM.DeliveryOrder{
			PhoneID:   req.DeliveryInfo.PhoneID,
			AddressID: req.DeliveryInfo.AddressID,
			Reference: req.DeliveryInfo.AddressID,
		}
		order.DeliveryAgt = delivery
	}

	//fmt.Printf("ID handler %v\n", order.Items[0].ID)
	err = h.OrderSrv.Create(context.Background(), order)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

/* {
	"user_id": "abc123",
	"items": [
	  {
		"product_id": "p1",
		"quantity": 2,
		"price": 10.0
	  },
	  {
		"product_id": "p2",
		"quantity": 1,
		"price": 25.5
	  }
	]
  }
*/
