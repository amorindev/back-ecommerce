package controller

import (
	"context"
	"errors"

	authPb "com.fernando/pkg/app/auth/grpc/gen"
	/* authModel "com.fernando/pkg/app/auth/model"
	"com.fernando/pkg/app/auth/validate"
	"com.fernando/pkg/app/user/model"
	"google.golang.org/protobuf/types/known/timestamppb" */
)

// ? hay context como http r.context
// * Agragar los demás que faltan como tenlefono dar revisadas a supabase y firebase
func (c *Controller) SignUp(ctx context.Context, req *authPb.SignUpRequestMessage) (*authPb.SignUpResponseMessage, error) {
	//validate.ValidateSignUp2()
	/* _, err := validate.ValidateSignUp(req.Email, req.Password, req.ComfirmPassword)
	if err != nil {
		return nil, err
	}
	// * Traducir la solicitud
	auth := &authModel.Auth{
		Password: req.Password,
		UserAgregate: &model.User{
			Email: req.Email,
		},
	}

	// ! add otp
	// ?  el servicio deberi validar si no viene un campo como go-l
	if err := c.AuthSrv.SignUp(context.Background(), auth,false); err != nil {
		return nil, err
	}

	// * Roles
	rolesPb := make([]*authPb.Role, len(auth.UserAgregate.Roles)-1)
	for i, role := range auth.UserAgregate.RolesModel {
		rolesPb[i] = &authPb.Role{Id: role.ID.(string), Name: role.Name}
	}

	// ! SE ESTA DEVOLVIENDO EL PROVIDER AGREGAR LA .PROTO
	// sign up con retirnar provider? lgo de auth tabla
	// ! es mejor devolver el user en create y no en sign in  por que existe get user
	// * Traducir la respuesta
	// ver siuser es nil?
	resp := &authPb.SignUpResponseMessage{
		Provider: *auth.Provider,
		User: &authPb.User{
			Id:            auth.UserAgregate.ID.(string),
			Email:         auth.UserAgregate.Email,
			EmailVerified: auth.UserAgregate.EmailVerified,
			Roles:         rolesPb,
			// revisar nulos auth user
			CreatedAt: timestamppb.New(*auth.UserAgregate.CreatedAt),
			UpdatedAt: timestamppb.New(*auth.UserAgregate.UpdatedAt),
		},
	}

	return resp, nil */
	return nil, errors.New("auth signup hander unimplement")

}
