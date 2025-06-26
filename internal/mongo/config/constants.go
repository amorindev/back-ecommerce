package config

// *Como hacer modulos en golang y dependencias
// * con vertical slicing cambiamos y agregamos funcionalidades
// * ver como afectan los módulos

// bueno esto es de mongo lo movería a mongo folder
// cambiar a roles a plural
// * me parece que es mejor un archivo cofig.go al lado de la connección
// * 30 tablas
const (
	CollAuth           = "auth_providers"
	CollUsers          = "users"
	CollRoles          = "roles"
	CollRolePermission = "role_permission"
	CollPermissions    = "permissions"
	CollUserRole       = "user_role"
	CollSessions       = "sessions" // dejarlo como session como algo general
	CollOnboarding     = "onboardings"
	CollCoupon         = "coupons" //
	CollPhones         = "phones"  // or UserPhones
	// poner prefijo user o esta demas asi mismo para la entidad
	CollTwoFaSms = "two_fa_sms"
	// aqui tambien se almacenarán informacion del navegador igual a facebook
	//CollRefreshTokens = "refresh_tokens" // o llamarlo solamente session
	CollOtp   = "otp_codes"
	CollUtils = "utils"

	// * Ecomm - mejor que este en su respectivo DDD
	CollCategories      = "categories"
	CollVariations      = "variations"
	CollVarOptions      = "variation_options"
	CollProducts        = "products"
	CollOrderProduct    = "order_product"
	CollProductItems    = "product_items"
	CollProductConfig   = "product_config"
	CollOrders          = "orders"
	CollAssessment      = "assetsments"      // relacion user_producto
	CollProductComments = "product_comments" // es implicito que el usuario realiza el comentario
	CollAddress         = "addresses"
	CollStores          = "stores"
	// para saber a donde se envio o se va envair la orden - relacionar de address y order
	// o delivery y address
	CollDelivery = "delivery"
	CollPickup   = "pickup"

	// ? es reutilizable
	CollPayments = "payments"

	//podria ser impuestos otra tabla ya sea para apple o del pais igv

	CollStripeCustomer = "stripe_customers"
)

// * Quedará pendiente las tablas intermedias de role_permission
// * y user_onboarding muchos a muchos para saber si vio el onboarding
// * ver la relacion user_address por que un usuario puede tener muchas direcciones
// * si se marca como unique entonces que pasa si es una familia
// * puede repetirse ver

/*
CREATE TABLE product_ratings (
    id SERIAL PRIMARY KEY,
    product_id INT REFERENCES products(id),
    user_id INT REFERENCES users(id),
    rating INT CHECK (rating BETWEEN 1 AND 5),
    liked BOOLEAN DEFAULT FALSE,
    UNIQUE(product_id, user_id)
);
*/

/*
CREATE TABLE product_comments (
    id SERIAL PRIMARY KEY,
    product_id INT REFERENCES products(id),
    user_id INT REFERENCES users(id),
    comment TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
*/
