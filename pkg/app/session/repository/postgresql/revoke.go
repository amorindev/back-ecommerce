package postgresql

/*
or update?

func revokeRefreshToken(tokenID string) error {
    query := `
        UPDATE tb_token 
        SET revoked = true 
        WHERE id = $1
    `
    _, err := db.Exec(query, tokenID)
    return err
}
*/