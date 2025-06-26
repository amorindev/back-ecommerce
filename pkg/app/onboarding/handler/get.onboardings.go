package handler

import (
	"encoding/json"
	"net/http"

	"com.fernando/cmd/api/message"
	"com.fernando/pkg/app/onboarding/model"
)

// esta ruta es de prueba ver que ruta se va a usar
func (h *Handler) GetOnboardings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	onboardings, err := h.OnboardingSrv.Get(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	type onboardingResp struct {
		Onboardings []*model.Onboarding `json:"onboardings"`
	}

	resp := onboardingResp{
		Onboardings: onboardings,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
