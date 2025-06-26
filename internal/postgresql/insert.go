package postgresql

import (
	"github.com/jackc/pgx/v5"
)

// insert data into the database
func Insert(conn *pgx.Conn) error {

	/* err = insertAdmin(conn, &adminId)
	if err != nil {
		return err
	}
	*/
	/* time.Sleep(5 * time.Second)
	err = setAdminRole(conn, roleId, adminId)
	if err != nil {
		return err
	} */
	return nil
}

// no necesito modificar role id y admin id, no necesito un puntero
/* func setAdminRole(conn *pgx.Conn, roleId uuid.UUID, adminId uuid.UUID) error {
	//fmt.Printf("role id: %s\n", roleId)
	//fmt.Printf("admin id: %s\n", adminId)
	q := `INSERT INTO tb_auth_role (user_id, role_id)
			VALUES ($1, $2);`

	_, err := conn.Exec(context.Background(), q, adminId.String(), roleId.String())
	if err != nil {
		return err
	}
	return nil
}

func insertAdmin(conn *pgx.Conn) error {
	id := uuid.New()
	adminuser := os.Getenv("ADMIN_USER")
	adminpass := os.Getenv("ADMIN_PASS")
	emailverified := true
	now := time.Now()

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(adminpass), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	adminpass = string(passwordHash)

	q := `INSERT INTO tb_auth (id, email, password, email_verified,created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6)`

	_ = conn.QueryRow(context.Background(), q, id, adminuser, adminpass, emailverified, now, now)

	return nil
}
*/
