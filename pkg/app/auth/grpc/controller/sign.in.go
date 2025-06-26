package controller

import (
	"context"
	"errors"

	authPb "com.fernando/pkg/app/auth/grpc/gen"
	/* "com.fernando/pkg/app/auth/validate"
	"google.golang.org/protobuf/types/known/timestamppb" */
)

// * prototimestamp.AsTime()
// ! REFACTORIZAR DESDE CERO
func (c *Controller) SignIn(ctx context.Context, req *authPb.SignInRequestMessage) (*authPb.SignInResponseMessage, error) {
	// * Traducir
	// * Validar
	/* _, err := validate.ValidateSignIn(req.Email, req.Password, req.RememberMe)
	if err != nil {
		return nil, err
	}

	// * Asignar a un servicio
	// ! ver el tema del context cando es lectura estrictura y si convin ambos
	// el tema si se usn 3 entidades no solo sera auth
	user,session, err := c.AuthSrv.SignIn(ctx, req.Email, req.Password, req.RememberMe)
	if err != nil {
		return nil, err
	}

	// se deberia verificar si es nil el auth y atributos ambos?
	// * Traducir la respuesta
	// TODO: verificar todos los campos que se debe retornar
	createdAtTimestamp := timestamppb.New(*auth.UserAgregate.CreatedAt)
	updatedAtTimestamp := timestamppb.New(*auth.UserAgregate.UpdatedAt)

	// * Traducir Roles

	rolesPb := make([]*authPb.Role, len(auth.UserAgregate.Roles)-1)
	for i, role := range auth.UserAgregate.RolesModel {
		rolesPb[i] = &authPb.Role{Id: role.ID.(string), Name: role.Name}
	}

	resp := &authPb.SignInResponseMessage{
		Provider:     *auth.Provider,
		AccessToken:  auth.AccessToken,
		RefreshToken: auth.RefreshToken,
		User: &authPb.User{
			Id:            auth.UserAgregate.ID.(string),
			Email:         auth.UserAgregate.Email,
			EmailVerified: auth.UserAgregate.EmailVerified,
			CreatedAt:     createdAtTimestamp,
			UpdatedAt:     updatedAtTimestamp,
			Roles:         rolesPb,
			// falta el role y ver si solo sera string slice para modificar la estrucut
		},
	}

	return resp, nil */
	return nil, errors.New("auth sign in hander unimplement")
}
