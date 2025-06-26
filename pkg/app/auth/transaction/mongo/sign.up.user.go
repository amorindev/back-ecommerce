package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/user/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// ! verificar el proveedor si existe con el auth
// * que todos usen context, y el context de WithSession según el caso, sirve el context.Background() ?
// * para no hacer auth.user = nil crear aux, se podría insertar primero los hijos?

func (t *Transaction) SignUpUser(ctx context.Context, user *model.User) error {
	// * Crear el id aqui o en el primero ya que es solo para mongo y no afecta
	// * o crearlo en el primera insercion osea auth, y asignarlo desde el servicio o repo
	// *
	//id := bson.NewObjectID()
	//asignar tambien al auth o al use o a ambos
	//auth.ID = id
	//auth.UserAgregate.ID = id

	session, err := t.Client.StartSession()
	if err != nil {
		return fmt.Errorf("auth mongo tx - SignUp err: %w", err)
	}

	err = session.StartTransaction()
	if err != nil {
		return fmt.Errorf("auth mongo tx - SignUp err: %w", err)
	}

	defer session.EndSession(ctx)

	err = mongo.WithSession(ctx, session, func(ctx context.Context) error {
		// TODO: pasar todo a solo create depués vemos embedding

		err = t.UserRepo.Insert(ctx, user)
		if err != nil {
			return err
		}

		//auth.UserID = auth.UserAgregate.ID.(string) // dentro de a funcion auth.ID = atu.User
		err = t.AuthRepo.Insert(ctx, user.AuthProviderCreate)
		if err != nil {
			return err
		}

		// ? no debería ser solo los ids
		err = t.RoleRepo.AssignRolesToUser(ctx, user.ID.(string), user.RolesModel)
		if err != nil {
			return err
		}

		err = session.CommitTransaction(ctx)
		if err != nil {
			return err
		}

		return nil
		// return session.CommitTransaction(ctx) o usar esto
	})
	if err != nil {
		if err := session.AbortTransaction(context.Background()); err != nil {
			return fmt.Errorf("auth mongo transaction - SignUp AbortTransaction err: %w", err)
		}
		return fmt.Errorf("auth mongo transaction - SignUp err: %w", err)
	}

	return nil
}
