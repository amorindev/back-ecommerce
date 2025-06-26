package mongo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// parecido updateauth()
func (r *Repository) ConfirmEmail(email string) error {
	filter := bson.M{
		"email":          email,
		"email_verified": false,
	}

	update := bson.M{
		"$set": bson.M{"email_verified": true},
	}

	result, err := r.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		// seria not found pero no se confundiría con el delete cuando no lo encuentra
		// * me parece que no go senior en cada cada cambia el mensaje de error
		return errors.New("no se encontró ningún documento para actualizar")
	}
	// result.MatchedCount estudiar los demas
	//fmt.Printf("Documentos modificados %v\n", result.ModifiedCount)

	return nil

}

// * mira este context
// ? que context aplicar este o el de la transaccion o otro contexto que viene de afura, o
// ? el de la request, por que sabemos que si se pierde la conexion con el usario talvés on
// ? podamos retornar la respuesa pero los datos si son correcto se mantedrán y el usuario
// ? no perderá su trabajo
/*
func UpdateIsActiveToFalseByEmail(client *mongo.Client, email string) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    collection := client.Database("tu_basededatos").Collection("users")

    filter := bson.M{
        "email":     email,
        "is_active": true,
    }

    update := bson.M{
        "$set": bson.M{"is_active": false},
    }

    result, err := collection.UpdateOne(ctx, filter, update)
    if err != nil {
        return err
    }

    if result.ModifiedCount == 0 {
        fmt.Println("No se encontró ningún documento para actualizar.")
    } else {
        fmt.Printf("Documento actualizado. Modificados: %v\n", result.ModifiedCount)
    }

    return nil
}
*/
