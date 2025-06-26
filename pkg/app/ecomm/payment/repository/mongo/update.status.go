package mongo

import (
	"context"
	"fmt"
	"time"

	"com.fernando/pkg/app/ecomm/payment/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) UpdateStatus(ctx context.Context, id string, status model.PaymentStatus) error {
	 // Convertir string a ObjectID
    objectId, err := bson.ObjectIDFromHex(id)
    if err != nil {
        return fmt.Errorf("ID inválido: %w", err)
    }

    // Filtro para buscar por _id
    filtro := bson.M{"_id": objectId}

    // Update usando $set
    update := bson.M{"$set": bson.M{"status": status}}

    // Opcional: timeout de contexto
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()

    // Ejecutar el update
    result, err := r.Collection.UpdateOne(ctx, filtro, update)
    if err != nil {
        return fmt.Errorf("error al actualizar: %w", err)
    }

    if result.MatchedCount == 0 {
        return fmt.Errorf("no se encontró ningún documento con ese ID")
    }

    return nil
}