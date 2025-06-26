package handler

import "net/http"

func (h *Handler) GetWithVarOptions(w http.ResponseWriter, r *http.Request){
	/* w.Header().Set("Content-Type", "application/json")

	h.VariationSrv.

	// retornar token y refreshtoken, junto al auth si se sac de Auth{}
	auth, err := h.AuthService.SignIn(r.Context(), u.Email, u.Password, u.RememberMe)
	if err != nil {
		// puede ser usernot found
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}


	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(auth) */
}