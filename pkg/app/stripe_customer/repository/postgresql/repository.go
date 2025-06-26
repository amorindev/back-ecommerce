package postgresql

/* func (db *DB) SaveCustomer(userID, customerID string) error {
    _, err := db.Exec(`INSERT INTO stripe_customers (user_id, customer_id, created_at)
                       VALUES ($1, $2, NOW())`, userID, customerID)
    return err
}

func (db *DB) FindCustomerByUserID(userID string) (*StripeCustomer, error) {
    row := db.QueryRow(`SELECT customer_id FROM stripe_customers WHERE user_id = $1`, userID)
    var customerID string
    err := row.Scan(&customerID)
    if err == sql.ErrNoRows {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }
    return &StripeCustomer{UserID: userID, CustomerID: customerID}, nil
} */
