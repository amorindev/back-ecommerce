package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/otp-codes/model"
	userModel "com.fernando/pkg/app/user/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (t *Transaction) SignUpWithOtpUser(ctx context.Context, user *userModel.User, otp *model.OtpCodes) error {

	session, err := t.Client.StartSession()
	if err != nil {
		return fmt.Errorf("user mongo tx - SignUpWithOtp err: %w", err)
	}

	err = session.StartTransaction()
	if err != nil {
		return fmt.Errorf("user mongo tx - SignUpWithOtp err: %w", err)
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

		err = t.RoleRepo.AssignRolesToUser(ctx, user.ID.(string), user.RolesModel)
		if err != nil {
			return err
		}

		otp.UserID = user.ID
		err = t.OtpRepo.Insert(context.Background(), otp)
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
			return fmt.Errorf("user mongo transaction - SignUpWithOtp AbortTransaction err: %w", err)
		}
		return fmt.Errorf("user mongo transaction - SignUpWithOtp err: %w", err)
	}

	return nil
}

// ? cuando diseñar estructuras sepasradas / y cuando UserAgregate - microblog usa de esta manera
// ? será por el tema de  response
type Category struct {
	ID int
}
type Product struct {
	ID         int
	CategoryID int
}
