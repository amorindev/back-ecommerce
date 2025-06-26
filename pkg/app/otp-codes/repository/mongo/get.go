package mongo

import (
	"context"
	"errors"

	otpErr "com.fernando/pkg/app/otp-codes/errors"
	"com.fernando/pkg/app/otp-codes/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)


func (r *Repository) Get(ctx context.Context, otpID string) (*model.OtpCodes, error) {
	var otp model.OtpCodes

	objID, err := bson.ObjectIDFromHex(otpID)
	if err != nil {
		return nil, errors.New("auth mongo repo - Get invalid type ID")
	}

	err = r.Collection.FindOne(ctx,bson.D{{Key:"_id",Value:objID}}).Decode(&otp)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, otpErr.ErrOtpNotFound
		}
		return nil, err
	}

	userOID, ok := otp.UserID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("otp-codes mongo repo - Get, failed to convert userID to ObjectID")
	}
	otp.UserID = userOID.Hex()
	
	// parsear a Object y luego a hex o simplemente
	otp.ID = otpID // o objIS.hex
	return &otp, nil
}



/*

func (r *Repository) Get(ctx context.Context, id string) (*model.Auth, error) {

	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("auth mongo repo - Get invalid type ID")
	}
	var auth model.Auth
	// db: 67b93e6f1d2ba978f59abdfe
	fmt.Printf("ID: %v\n", objID)
	// 67b93e6f1d2ba978f59abdfe
	// * probar con el ID
	// ! ahora es mediante en UserID
	err = r.AuthCollection.FindOne(ctx, bson.D{{Key: "_id", Value: objID}}).Decode(&auth)
	fmt.Printf("Auth: %v\n", auth)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, authErr.ErrAuthNotFound
		}
		return nil, err
	}

	// ! cambiar a de idStr to objID
	objID, ok := auth.ID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("auth mongo repo - Get falied to parse ID to string")
	}
	auth.ID = objID.Hex()

	return &auth, nil
}

*/
