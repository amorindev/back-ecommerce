package handler

import (
	"encoding/json"
	"net/http"

	"com.fernando/cmd/api/message"
)



func (h Handler) CreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicaction/json")

	type reqBody struct {
		Amount   float64           `json:"amount"`
		Currency string            `json:"currency"`
		Metadata map[string]string `json:"metadata"`
	}
	var req reqBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	defer r.Body.Close()

	// * validar ver que mas se va a validar
	if req.Amount <= 0.0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "amount field is required"})
		return
	}
	if req.Currency == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "currency field is required"})
		return
	}
	// ? como validar metadata cantidad de valores y valores que no necesitamos
	// ? veridfiar si el id existe ? en DB

	// * yano por que la conversion lo haremos en el backend
	// Evitar errores de redondeo al multiplicar. Si trabajas con cálculos financieros
	// aqui o en el servicio
	//montoDecimal, _ := strconv.ParseFloat(req.Amount, 64)

	// que pasa si es peru o tine mas decimales .000 ( buenno no sucede  osi)
	amount := int(req.Amount * 100)
	//amount := int64(math.Round(amountFloat * 100)) cuando usar round

	// ? de donde sacar currency de varibles de entornno?
	pIntent, err := h.PaymentStripeSrv.CreatePaymentIntent(amount, req.Currency, req.Metadata)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	resp := struct {
		ClientSecret string `json:"client_secret"`
	}{
		ClientSecret: pIntent.ClientSecret,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}

// Cuando usas Stripe.instance.initPaymentSheet() y presentPaymentSheet() en Flutter, estás trabajando con Payment Intents, no con Checkout Sessions. En este flujo, los metadatos deben agregarse desde tu backend, al momento de crear el PaymentIntent.