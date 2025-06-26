package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/otp-codes/model"
	authModel "com.fernando/pkg/app/auth/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (t *Transaction) SignUpWithOtp(ctx context.Context, auth *authModel.Auth, otp *model.OtpCodes) error {

	session, err := t.Client.StartSession()
	if err != nil {
		return fmt.Errorf("auth mongo tx - SignUpWithOtp err: %w", err)
	}

	err = session.StartTransaction()
	if err != nil {
		return fmt.Errorf("auth mongo tx - SignUpWithOtp err: %w", err)
	}

	defer session.EndSession(ctx)

	err = mongo.WithSession(ctx, session, func(ctx context.Context) error {
		//auth.UserID = auth.UserAgregate.ID.(string) // dentro de a funcion auth.ID = atu.User
		err = t.AuthRepo.Insert(ctx, auth)
		if err != nil {
			return err
		}

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
			return fmt.Errorf("auth mongo transaction - SignUpWithOtp AbortTransaction err: %w", err)
		}
		return fmt.Errorf("auth mongo transaction - SignUpWithOtp err: %w", err)
	}

	return nil
}
