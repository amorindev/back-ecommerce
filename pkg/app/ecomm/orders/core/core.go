package core

// * que validaciones se debe hacer por ejemplo aparete de los nulos, que venga como
// * mpinimo un ProductItem
// * userID desde el token
//?TOtal debe coincidir con lo que se mostro al usuario
type CreateOrderReq struct {
	Total        float64                 `json:"total"`
	DeliveryType string                  `json:"delivery_type"`
	ProductItems []*CreateProductItemReq `json:"product_items"`
	DeliveryInfo *DeliveryInfoReq        `json:"delivery_info"`
	PickupInfo   *PickupInfo             `json:"pickup_info"`
	Payment      *Payment                `json:"payment"`
}

// cada unao de estas entidades no devería ir dentro de su DDD ver
// tratarlo como string o interface{} ver gosenior
type CreateProductItemReq struct {
	ProductItemID string  `json:"product_item_id"`
	Quantity      int     `json:"quantity"`
	Price         float64 `json:"price"`
}

type DeliveryInfoReq struct {
	PhoneID   string `json:"phone_id"`
	AddressID string `json:"address_id"`
	Reference *string `json:"reference"`
}

type PickupInfo struct {
	PhoneID   string `json:"phone_id"`
	AddressID string `json:"address_id"`
	StoreID   string `json:"store_id"`
}

type Payment struct {
	Currency      string `json:"currency"`
	PaymentMethod string `json:"payment_method"`
}
