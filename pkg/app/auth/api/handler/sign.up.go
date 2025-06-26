package handler

import (
	"encoding/json"
	"net/http"

	"com.fernando/cmd/api/message"
	"com.fernando/pkg/app/auth/api/core"

	//authM "com.fernando/pkg/app/auth/model"

	//sessionM "com.fernando/pkg/app/session/model"
	userM "com.fernando/pkg/app/user/model"
)

// ! validar los email con govalidator en todos los handlers que se necesiten
// * verificar para que retorno invalid-email
func (h Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req core.SignUpReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	defer r.Body.Close()

	err = req.IsSignUpValid()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	// crear la funcion, o el constructor sin agregar confirm passsword a la estructura
	// se ve mejor con el constructor
	// ! cambiar pasarlo al auth domain
	// ! crear sign up from  core.SignUpReq o pasarle primitivos
	user := userM.NewUserSignUp(req.Email, req.Name, req.Username, req.Password, req.Phone)

	user.Logger("handler", "signUp")

	ctx := r.Context()
	// de momento trabajaresmos puro OTP hasta ver si es mejor usar tokens paraece mas trabajo
	// veremmos
	req.WithCode = true
	otpID, err := h.AuthService.SignUp(ctx, user, req.WithCode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	// ! cambiar por el Auth response de core
	// debe ser so
	// Session segun firebase pero nosotros lo estamoa agregando despues de que verifique su correo
	// Credentials separado o dentro de User
	type respBody struct {
		User *userM.User `json:"user"`
		//Session     *sessionM.Session `json:"session"`
		//Credentials *authM.Auth       `json:"credentials"`
		OtpID string `json:"otp_id"`
	}

	resp := respBody{
		User:  user,
		OtpID: otpID,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
