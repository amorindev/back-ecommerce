package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/role/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// * de momento no convertioms a string cuando son slices por que no usamos el id en string
/* var roleIDs []string // en el servicio verificar los nulos, o en el repo ?
// ?convertir a ObjectID?
id */
// ? a quie darle la responsabilidad de las tablas intermedias para consultas e inserciones a que servicio?
// ? de momento a role para
// ? [] que sea nil se verifica igual que un puntero *[]
// ! aprender mongo filters y todo lo demás métodos
func (r *Repository) GetByUserID(ctx context.Context, userID string) ([]string, error) {
	// la otra forma se ria con lookup aun no lo he probado
	// Recuerda estamos en tres tablas user user_role y role bueno solo user_role y role
	// * Obtener los ids de los roles relacionados al

	userObjID, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	cursor, err := r.UserRoleColl.Find(ctx, bson.M{"user_id": userObjID})
	if err != nil {
		return nil, fmt.Errorf("role mongo repo - GetByUserID err: %v", err)
	}
	defer cursor.Close(ctx)

	//var roleIDs []string  o esto
	var roleIDs []bson.ObjectID
	for cursor.Next(ctx) {
		var userRole model.UserRole
		if err := cursor.Decode(&userRole); err != nil {
			return nil, fmt.Errorf("role mongo repo - GetByUserID err: %v", err)
		}
		// en este punto se que es de tipo ObjectID, o verificamos con ok
		roleIDs = append(roleIDs, userRole.RoleID.(bson.ObjectID))
	}

	if len(roleIDs) == 0 {
		return nil, nil // usuario sin roles asignados
		//return []string{}, nil // Usuario sin roles asignados por que?
	}

	// * Obtener los nombre de los roles
	filter := bson.M{"_id": bson.M{"$in": roleIDs}}
	roleCursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("role mongo repo - GetByUserID err: %v", err)
	}
	defer roleCursor.Close(ctx)

	var roleNames []string
	for roleCursor.Next(ctx) {
		var role model.Role
		if err := roleCursor.Decode(&role); err != nil {
			return nil, fmt.Errorf("role mongo repo - GetByUserID err: %v", err)
		}
		roleNames = append(roleNames, role.Name)
	}

	return roleNames, nil
}

/* func (r *Repository) GetByUserID2(ctx context.Context, userID string) ([]string, error) {

	// fata FindMany para no hacer el for
	filer := bson.M{"_id": userID}
	// que pasa si tiene la estructura tiene mas campos debería agreagarlos
	projection := bson.M{"name": 1, "_id": 0}

	// puede ser un tipo primitivo?
	var result struct {
		Name string `bson:"name"`
	}

	err := r.Collection.FindOne(ctx, filer, options.FindOne().SetProjection(projection)).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("role mongo repo - GetByUserID err: %w", err)
	}

	// result.Name

	return nil, errors.New("role mongo repository - GetByUserID unimplement")
}

func (r *Repository) Test(ctx context.Context, userID string) ([]string, error) {

	projection := bson.M{"name": 1, "_id": 0}
	// * Find arriba se esta usando FindOne
	cursor, err := r.Collection.Find(context.Background(), bson.M{}, options.Find().SetProjection(projection))
	if err != nil {
		return nil, fmt.Errorf("role mongo repo - GetByUserID err: %w", err)
	}

	// ? se deberia usar el context de la transaccion?
	defer cursor.Close(context.Background())

	type result struct {
		Name string `bson:"name"`
	}
	var names []string
	for cursor.Next(ctx) {
		var r result
		if err := cursor.Decode(&r); err != nil {
			return nil, fmt.Errorf("role mongo repo, Test err: %w", err)
		}
		names = append(names, r.Name)
	}

	// si no incuyes el name quedará con Name con valor por defecto ""
	// projection := bson.M{"name": 1, "_id": 0}

	// si la colleccion tiene más campos como birth_date, pero la estructura solo es name,
	// mongo ignorará los demá´s y no generará erres

	// MongoDB solo llena los campos que están en la estructura Go y en la proyección. Otros campos serán ignorados sin generar errores.

	return names, nil
} */
