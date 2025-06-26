package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/otp-codes/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) Insert(ctx context.Context, otp *model.OtpCodes) error {
	id := bson.NewObjectID()
	otp.ID = id

	// creo qu no hace falta p or que en todos los repos nos aseguramos que sea string
	otpObjID, err := bson.ObjectIDFromHex(otp.UserID.(string))
	if err != nil {
		return fmt.Errorf("auth mongo repo - Create err: %w", err)
	}
	otp.UserID = otpObjID

	_, err = r.Collection.InsertOne(context.Background(), otp)
	if err != nil {
		return fmt.Errorf("otp mongo repo: Create err: %w", err)
	}
	otp.ID = id.Hex()
	// o reservar en una variable idStr de otp
	otp.UserID = otpObjID.Hex()

	/* authID, ok := auth.ID.(bson.ObjectID)
	if !ok {
		return errors.New("auth mongo repo - Create, failed to parse ID to ObjectId")
	}
	auth.ID = authID.Hex() */
	return nil
}


/* func (r *Repository) Create(ctx context.Context, auth *authModel.Auth) error {
	// ! Asi para todos los que usan transacción

	auth.ID = bson.NewObjectID()

	// * obtener el ID auth y pasarle al UserID
	userObjID, err := bson.ObjectIDFromHex(auth.UserAgregate.ID.(string))
	if err != nil {
		return fmt.Errorf("auth mongo repo - Create err: %w", err)
	}
	auth.UserID = userObjID

	// * descomenta si se desea retornar  el usuario, parece que si
	// ? o hacerlo desde transaccion o service
	userAux := auth.UserAgregate
	auth.UserAgregate = nil

	_, err = r.AuthCollection.InsertOne(ctx, auth)
	if err != nil {
		return fmt.Errorf("auth mongo repo - Create error: %w", err)
	}
	// bueno deveria ser nil tanto para crear como para devolver
	// es por eso su existencia de esta función
	auth.UserAgregate = userAux

	authID, ok := auth.ID.(bson.ObjectID)
	if !ok {
		return errors.New("auth mongo repo - Create, failed to parse ID to ObjectId")
	}
	auth.ID = authID.Hex()

	return nil
	//return errors.New("auth mongo repo - Create unimplement")
}
 */