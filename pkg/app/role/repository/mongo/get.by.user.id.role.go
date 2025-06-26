package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/role/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// ? convertir los id a string?
func (r *Repository) GetByUserIDRole(ctx context.Context, userID string) ([]model.Role, error) {

	objID, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("role mongo repo - BetByUserIDRole err: %w", err)
	}

	pl := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.M{"user_id": objID}}},
		{{Key: "$lookup", Value: bson.M{
			"from":         "roles",
			"localField":   "role_id",
			"foreignField": "_id",
			"as":           "roles_resp",
		}}},
		//{{Key: "$unwind", Value: "$user_roless"}}, formato plano?
		{{
			Key: "$project", Value: bson.M{
				"_id":        0,
				"roles_resp": 1,
			},
		}},
	}

	cursor, err := r.UserRoleColl.Aggregate(ctx, pl)
	if err != nil {
		return nil, fmt.Errorf("role mongo repo - GetByUserIDRole err: %w", err)
	}

	defer cursor.Close(ctx)

	type roleResponse struct {
		UserRoles []model.Role `bson:"roles_resp"`
	}

	var roles []model.Role
	for cursor.Next(ctx) {
		var roleResp roleResponse
		if err := cursor.Decode(&roleResp); err != nil {
			return nil, fmt.Errorf("role mongo repo - GetByUserIDRole err: %w", err)
		}
		// cambiar id a hex usando? como es un arrglo ahi hay un inconveniente
		roles = append(roles, roleResp.UserRoles...)
	}

	return roles, nil
}

/* var ur []roleResponse
if err := cursor.All(ctx, &ur); err != nil {
	return nil, fmt.Errorf("role mongo repo - GetByUserIDRole err: %w", err)
}
fmt.Printf("Roles: %v\n", ur) */

// ! repasar
/*
/* pl2 := mongo.Pipeline{
		{{Key: "$lookup", Value: bson.M{
			"from":"roles",
			"let": bson.M{"roleIds":"$role_id"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$in": []interface{}{"$_id", "$$roleIds"}}}},
			},
			"as": "user_rolesss",
		}}},
	} */
