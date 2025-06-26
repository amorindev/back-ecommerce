package service

import (
	"context"
	"time"

	authErr "com.fernando/pkg/app/auth/errors"
	authModel "com.fernando/pkg/app/auth/model"
	sessionModel "com.fernando/pkg/app/session/model"
	userErr "com.fernando/pkg/app/user/errors"
	"com.fernando/pkg/app/user/model"
	//authErr "com.fernando/pkg/services/app/auth/errors"
	//authModel "com.fernando/pkg/services/app/auth/model"
	//userAuthErr "com.fernando/pkg/services/app/user/errors"
	//userModel "com.fernando/pkg/services/app/user/model"
	//"com.fernando/pkg/services/app/session/model"
)

// ! RETORNAR LA SESSION
// ! refresh token debe ser hasheado? y como agregarlo a los servicios
// * si creo un service para cada proveedor duplicaría mucho código
// * recuperar los datos de los proveedores foto nombres completos (ver en apple)
// debería separar lo como branch io?
// create el token con auth o user  que pasa si no uno falla token o sign up (no afecta a los datos ACID)
func (s *Service) GoogleSignIn(ctx context.Context, rememberMe bool, providerTokenID string) (*authModel.Auth, error) {
	email, err := s.AuthAdapter.GoogleValidateToken(ctx, providerTokenID)
	if err != nil {
		return nil, err
	}

	// * variales
	now := time.Now().UTC()
	provider := "google"

	// * Verificar si el usuario existe
	user, err := s.UserRepo.GetByEmail(ctx, email)
	if err != nil {
		if err != userErr.ErrUserNotFound {
			return nil, err
		}
	}

	authCreate := &authModel.Auth{
		Provider:  &provider,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	if user == nil {
		// * Si user es nulo no existirá

		//Oauth garantiza email verificado - ver flujo si ocurre algun error deve se r falso
		userCreate := &model.User{
			Email:         email,
			EmailVerified: true, // ? providers aseguran que el email este verificado ?
			CreatedAt:     &now,
			UpdatedAt:     &now,
		}

		// * asignar roles
		userRoles := []string{"USER", "USER-TEST"}
		rolesModel, err := s.RoleRepo.GetByNames(ctx, userRoles)
		if err != nil {
			return nil, err
		}
		userCreate.RolesModel = rolesModel

		// GetByNames se asegura que esten los roles
		userCreate.Roles = userRoles

		// o inicializarlo desde  un constructor en el handler
		// y estoy seguro que ya lo cree y no tengo que verificar si es nulo
		userCreate.AuthProviderCreate = authCreate

		// cambiar el sign up a user user contine a providers
		err = s.AuthTx.SignUpUser(ctx, userCreate)
		if err != nil {
			return nil, err
		}

		session := sessionModel.Session{
			// ! es del user o del auth ?
			UserID:     userCreate.ID.(string),
			RemenberMe: rememberMe,
		}
		err = s.SessionSrv.Create(&session, []string{"ROLE,ROLE-TEST"})

		if err != nil {
			return nil, err
		}

		// * Send email verification (esto garantiza el proveedor)

		// ? donde pasamos el user id al userID del auth en la transacion o repo

		// * no confundir los auth que se crean con auth de los handler  o variables lo mismo con user
		// ! verificar si se retorna eluser
		return authCreate, nil
	}

	// * Si existe el auth
	auth, err := s.AuthRepo.GetByIDProvider(ctx, user.ID.(string), provider)
	if err != nil {
		if err != authErr.ErrAuthNotFound {
			return nil, err
		}
	}

	if auth == nil {
		// ! ver que se cree el otp guiarte sign.up del user
		// ! VERIFICAR POR QUE DEBE SER TRANSACCION CON SESSION ver lo que se necesita
		// * Verificamos si el usurio existe y el auth con el proveedor
		// * Create auth
		// * Create auth y sign in

		// no uso auth.UserAgregate = userCreate
		err = s.AuthRepo.Insert(ctx, authCreate)
		if err != nil {
			return nil, err
		}

		// Si el user con role son embedding ya no hace falta otra consulta
		// * Sign in
		roles, err := s.RoleRepo.GetByUserID(ctx, user.ID.(string))
		if err != nil {
			return nil, err
		}

		user.Roles = roles

		session := sessionModel.Session{
			// ! en este punto user no es nil
			UserID:     user.ID.(string),
			RemenberMe: rememberMe,
		}

		err = s.SessionSrv.Create(&session, []string{"role-test"})

		if err != nil {
			return nil, err
		}

		// user.AuthProviders = []*authModel.Auth{authCreate}
		user.AuthProviderCreate = authCreate

		// si es con el id del usurio el token aqui creas con el user ya que estas verificando que nosea  null
		// ver me parece que es del auth

		// * Send email verification (no no esto garántiza gooogle en que parte)
		return authCreate, nil
	}

	// * sign in
	roles, err := s.RoleRepo.GetByUserID(ctx, user.ID.(string))
	if err != nil {
		return nil, err
	}

	user.Roles = roles

	// !role test?
	session := sessionModel.Session{
		UserID:     user.ID.(string),
		RemenberMe: rememberMe,
	}

	// este si el context de la request por que no afecta a la otomicidad y
	// si se necesita devolver la session al usuarip, no es como create profile, creat eproduct
	// se debería parar como primitivos
	err = s.SessionSrv.Create(&session, []string{"role-test"})
	if err != nil {
		return nil, err
	}

	// dentro de estos if  user  auth no son nulos

	// ! de momento append pero cuidado
	//user.AuthProviders = []*authModel.Auth{authCreate}
	user.AuthProviderCreate = authCreate

	return auth, nil
	// * definir que funciones son para los 2 o tres casos y cuáles no

	// !Asegurar el envíp de email auth = user; no confundir los auth y user variables
	// ! y que en cualquier caso retornen la session

	// * Creamos el auth y el user

	// ! en el handler validar que sea google provider  y tenga idtoken email lo que requiera
	// El cliente obtiene un token JWT desde google/apple
	// El backend verifica este token con Google/Apple
	// Si el token es válido, busca al usuario en la tabla auth o lo registra si es un nuevo usuario
	// Genera un token jwt propio para e usaurio y o guarda en la tabla session.

	// usar el mismo middleware jwt?

	// * Si el email ya existe en User, authenticamos
	// * Si el email no existe, creamos User y Auth automaticamente
	// * soporte para password y oauth
	// * verificar si el outh soporta mobile / se tiene que usar webview?

}
