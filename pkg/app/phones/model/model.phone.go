package model

import "time"

// * hay dos formas en phone y addess no importa el orden
// * entonces solo usamos si es default y no piscion
// * al crear un phone se debe actualizar el anterior phone y ponerlo en isdefault false
// * y al nuevomarcarlo como por defecto igual en address usar tx
// * igual para phones
// *  se debe marcar como por defecto
// ! ver que se va a hashear
// * CountryCode // tabla pasises demomento sensillo
// * IsVerified // si es del usuario mandar sms
// * userid desde el token
// TODO Number aqui number ya no debeía ser nulo quitar el puntero
type Phone struct {
	ID             interface{} `json:"id" bson:"_id"`
	UserID         interface{} `json:"user_id" bson:"user_id"`
	Number         *string     `json:"number" bson:"number"`
	CountryCode    *string     `json:"country_code" bson:"country_code"`
	CountryIsoCode *string     `json:"country_iso_code" bson:"country_iso_code"`
	IsDefault      bool        `json:"is_default" bson:"is_default"`
	IsVerified     *bool       `json:"is_verified" bson:"is_verified"`
	CreatedAt      *time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt      *time.Time  `json:"updated_at" bson:"updated_at"`
}

/*
CREATE TABLE users (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME ON UPDATE CURRENT_TIMESTAMP,
    last_login DATETIME,
    is_active BOOLEAN DEFAULT TRUE
);

CREATE TABLE user_phones (
    phone_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    country_code VARCHAR(5) DEFAULT '+1',
    is_verified BOOLEAN DEFAULT FALSE,
    is_primary BOOLEAN DEFAULT FALSE,
    phone_type ENUM('Mobile', 'Home', 'Work', 'Other') DEFAULT 'Mobile',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    UNIQUE KEY unique_phone (user_id, phone_number)
);
*/
/*
type Phone struct {
	ID             bson.ObjectId `bson:"_id"`
	UserID         bson.ObjectId `bson:"user_id"`
	Number         string        `bson:"number"`
	CountryCode    string        `bson:"country_code"`
	CountryIsoCode string        `bson:"country_iso_code"`
	IsDefault      bool          `bson:"is_default"`
	IsVerified     bool          `bson:"is_verified"`
	CreatedAt      time.Time     `bson:"created_at"`
	UpdatedAt      time.Time     `bson:"updated_at"`
}
*/

/*
Key Features:
Security:

Password hashing (never store plain text)

Phone verification system

Timestamp tracking

Flexibility:

Multiple phones per user (in normalized version)

International number support

Performance:

Proper indexing (primary keys, unique constraints)

ON UPDATE timestamps

Extensible:

Easy to add OTP/login functionality

Ready for two-factor authentication

Would you like me to:

Add password reset functionality?

Include OTP verification fields?

Optimize for a specific database engine (MySQL, PostgreSQL, etc.)?

Add indexes for specific query patterns?
*/
