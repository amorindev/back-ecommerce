package postgresql

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

// ? Crear una función que crea solo un role y desde fuera usar un for para crear varios
// ? o usar CreateMany, esta función debe usar transacción todo o nada?
// * Si se desea retornar los roles junto a su id usa Role{} asigna el id en cada iteracion y parsealo a
// * string, si es asi deberia ser []*modelRole
func (r *Repository) CreateMany(ctx context.Context, names []string) error {
	q := `INSERT INTO tb_role (id, name) VALUES ($1, $2)`

	for _, roleName := range names {
		id := uuid.New()
		_, err := r.Conn.Exec(ctx, q, id, roleName)
		if err != nil {
			return fmt.Errorf("role postgresql repo - CreateMany error: %w", err)
		}
	}
	return nil
}

// ! delete
/*
func (r *Repository) InitRoles(ctx context.Context, roles []string) error {
	q := `INSERT INTO tb_role (id, name)
		VALUES ($1, $2)`
	for _, role := range roles {
		id := uuid.New()

		_, err := r.Conn.Exec(ctx, q, id, role)
		if err != nil {
			return fmt.Errorf("error inserting role: %s error: %w", role, err)
		}
	}
	return nil
}
*/

// TODO: Buscar info
// * para crear roles por ejemplo tengo mi servico y mi repo, tengo solo la lista de nombres de los roles ,debería hacer en for desde servicio con un funcion createRole y le paso el role, o enviar la lista de nombre de los roles y el repositorio hace el for mientras inserta cual es mejor?

//* llamarla desde el repo es muchas llamadas a la base de datos, puede ser meno s eficiente

func (r *Repository) CreateRoles2(ctx context.Context, roleNames []string) error {
	/* tx, err := r.Conn.Begin(ctx)
	if err != nil {
		return err
	}

	// ? esto no va cuando ocurre  algun error ?defer tx.Rollback()
	// ver name
	stmt, err := tx.Prepare(ctx,"insert-role","INSERT INTO tb_role(id, name) VALUES($1,$2)")
	if err != nil {
		return err
	} */


	return nil
}
