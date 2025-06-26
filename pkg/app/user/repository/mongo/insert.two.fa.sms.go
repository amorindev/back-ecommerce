package mongo

import (
	"context"
	"errors"
	"fmt"

	"com.fernando/pkg/app/user/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) InsertTwoFaSms(ctx context.Context, twoFaSms *model.UserTwoFaSms) error {
    id := bson.NewObjectID()
    twoFaSms.ID = id

    userObjID, err := bson.ObjectIDFromHex(twoFaSms.UserID.(string))
	if err != nil {
		return errors.New("twofasms- failed to parse to objID")
	}

    phoneObjID, err := bson.ObjectIDFromHex(twoFaSms.PhoneID.(string))
	if err != nil {
		return errors.New("twofasms phoneID - failed to parse to objID")
	}
    twoFaSms.UserID = userObjID
    twoFaSms.PhoneID = phoneObjID

    _, err = r.TwoFaSmsColl.InsertOne(context.Background(), twoFaSms)
    if err != nil {
		return fmt.Errorf("twofasms mongo repo: InsertTwoFaSms err: %w", err)
	}
    twoFaSms.ID = id.Hex()
    twoFaSms.UserID = userObjID.Hex()
    twoFaSms.PhoneID = phoneObjID.Hex()
    //if result.InsertedID == nil ?
    return nil
}
