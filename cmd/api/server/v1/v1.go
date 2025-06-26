package v1

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	//"com.fernando/internal/minio"
	"com.fernando/internal/minio"
	mongoClient "com.fernando/internal/mongo"
	"com.fernando/internal/mongo/config"
	resendClient "com.fernando/internal/resend"

	/* "com.fernando/internal/twilio" */

	//"com.fernandopkg\app\otps-codes\repository\mongo\repository.go"
	addressHandler "com.fernando/pkg/app/ecomm/address/handler"
	addressRepository "com.fernando/pkg/app/ecomm/address/repository/mongo"
	addressService "com.fernando/pkg/app/ecomm/address/service"
	addressTransaction "com.fernando/pkg/app/ecomm/address/transaction/mongo"
	categoryHandler "com.fernando/pkg/app/ecomm/category/handler"
	categoryRepository "com.fernando/pkg/app/ecomm/category/repository/mongo"
	categoryService "com.fernando/pkg/app/ecomm/category/service"
	deliveryRepository "com.fernando/pkg/app/ecomm/delivery-orders/repository/mongo"
	orderHandler "com.fernando/pkg/app/ecomm/orders/handler"
	orderRepository "com.fernando/pkg/app/ecomm/orders/repository/mongo"
	orderService "com.fernando/pkg/app/ecomm/orders/service"
	pickupRepository "com.fernando/pkg/app/ecomm/pickup-orders/repository/mongo"
	productConfigRepository "com.fernando/pkg/app/ecomm/product-config/repository/mongo"
	productItemHandler "com.fernando/pkg/app/ecomm/product-item/handler"
	productItemRepository "com.fernando/pkg/app/ecomm/product-item/repository/mongo"
	productItemService "com.fernando/pkg/app/ecomm/product-item/service"
	productHandler "com.fernando/pkg/app/ecomm/products/handler"
	productService "com.fernando/pkg/app/ecomm/products/service"
	storeHandler "com.fernando/pkg/app/ecomm/stores/handler"
	storeRepository "com.fernando/pkg/app/ecomm/stores/repository/mongo"
	storeService "com.fernando/pkg/app/ecomm/stores/service"
	"com.fernando/pkg/app/ecomm/stores/transaction/mongo"
	variationHandler "com.fernando/pkg/app/ecomm/variation/handler"
	variationRepository "com.fernando/pkg/app/ecomm/variation/repository/mongo"
	variationService "com.fernando/pkg/app/ecomm/variation/service"
	onboardingHandler "com.fernando/pkg/app/onboarding/handler"
	onboardingRepository "com.fernando/pkg/app/onboarding/repository/mongo"
	onboardingService "com.fernando/pkg/app/onboarding/service"
	otpMongo "com.fernando/pkg/app/otp-codes/repository/mongo"
	phoneHandler "com.fernando/pkg/app/phones/handler"
	phoneRepository "com.fernando/pkg/app/phones/repository/mongo"
	phoneService "com.fernando/pkg/app/phones/service"
	phoneTransaction "com.fernando/pkg/app/phones/transaction/mongo"
	pProviderRepository "com.fernando/pkg/app/stripe_customer/repository/mongo"

	/* twilioAdapter "com.fernando/pkg/sms/adapter/twilio"
	twilioService "com.fernando/pkg/sms/service" */

	paymentRepository "com.fernando/pkg/app/ecomm/payment/repository/mongo"
	paymentService "com.fernando/pkg/app/ecomm/payment/service"
	paymentHandler "com.fernando/pkg/payments/handler"

	paymentProviderService "com.fernando/pkg/payments/service"

	orderTransaction "com.fernando/pkg/app/ecomm/orders/transaction/mongo"

	authAdapter "com.fernando/pkg/app/auth/adapter"
	authHandler "com.fernando/pkg/app/auth/api/handler"
	authPGtx "com.fernando/pkg/app/auth/transaction/mongo"
	userMongo "com.fernando/pkg/app/user/repository/mongo"

	authMongo "com.fernando/pkg/app/auth/repository/mongo"
	authSrv "com.fernando/pkg/app/auth/service"
	initConf "com.fernando/pkg/app/init"

	productRepository "com.fernando/pkg/app/ecomm/products/repository/mongo"
	productTransaction "com.fernando/pkg/app/ecomm/products/transaction/mongo"

	roleMongo "com.fernando/pkg/app/role/repository/mongo"
	sessionMongo "com.fernando/pkg/app/session/repository/mongo"

	sessionSrv "com.fernando/pkg/app/session/service"
	branchioAdapter "com.fernando/pkg/branchio/adapter"
	resendAdapter "com.fernando/pkg/email/adapter/resend"
	emailService "com.fernando/pkg/email/service"
	minioAdapter "com.fernando/pkg/file-storage/minio/adapter"
	fileStgService "com.fernando/pkg/file-storage/service"
)

// dos formas pasando el puntero de mux o como este caso retornando htt.Handler
func New() http.Handler {
	mongoDBName := os.Getenv("MONGO_DB_NAME")
	if mongoDBName == "" {
		log.Fatal("environment variable MONGO_DB_NAME is not set")
	}

	mux := http.NewServeMux()

	// otra forma para llamar una sola ves a os.getenv, o viper(es pesado?)
	// crear un objeto y llamarlo desde
	/* c, err := claim.GetConfig()
	if err != nil {
		// no habrría necesidad de que retorne un error si antes de valida que las varibles de
		// entorno esten configuradas
		log.Fatal("err")
		} */
	// c debería pasarse  igual a Conn, para no
	// depndendiendo del servicio crear nuevos objetos

	// crear database?- donde crear base datos antes de server o despues
	// * Clients
	// si se puede llamar desde fuera del servidor, estaría mejor por que no carga en memoria?
	// igaul el init Run()
	// no usaremos postgresql pgConn := pgClient.New()
	mongoConn := mongoClient.New()
	//minioClient := minio.NewClient()
	// sacar del .env
	// mongoDB causará conflicto?
	mongoDB := mongoConn.DB.Database(mongoDBName)
	// o como la base de datos
	resendClient := resendClient.NewClient()
	minioClient := minio.NewClient()
	minioClient.CreateStorage()
	/* twilioClient := twilio.NewClient() */

	//googleVerifier, err := authAdapter.NewGoogleProvider()
	/* _, err := authAdapter.NewGoogleProvider()
	if err != nil {
		log.Fatal(err)
	} */

	/*
		appleVerifier, err := authAdapter.NewAppleProvider("", "", "", "")
		if err != nil {
			log.Fatal(err)
		} */

	// * Database ping
	// ! retorna err
	// ? mas limpio dentro de new etaria mejor ponerlo
	// no usaremos postgresql pgConn.Ping()
	mongoConn.Ping()

	// * Repositoty layer
	// Postgresql
	//authPGRepo := authPG.NewRepository(pgConn.DB)
	//rolePGRepo := rolePG.NewRepository(pgConn.DB)
	//sessionPGRepo := sessionPG.NewRepository(pgConn.DB)

	// MongoDB
	authColl := mongoDB.Collection(config.CollAuth)
	userColl := mongoDB.Collection(config.CollUsers)
	twoFaSmsColl := mongoDB.Collection(config.CollTwoFaSms)
	userRoleColl := mongoDB.Collection(config.CollUserRole)
	roleColl := mongoDB.Collection(config.CollRoles)
	sessionColl := mongoDB.Collection(config.CollSessions)
	otpColl := mongoDB.Collection(config.CollOtp)

	authMongoRepo := authMongo.NewRepository(mongoConn.DB, authColl)
	userMongoRepo := userMongo.NewRepository(mongoConn.DB, userColl, twoFaSmsColl, userRoleColl)
	roleMongoRepo := roleMongo.NewRepository(mongoConn.DB, roleColl, userRoleColl)
	sessionMongoRepo := sessionMongo.NewRepository(mongoConn.DB, sessionColl)
	otpRepository := otpMongo.NewRepository(mongoConn.DB, otpColl)

	//roleMongoRepo.GetByUserID(context.Background(), )

	// * User Phones
	// * me parece mejor que sea dentro de use handler ver lugar correcto aunque esto es DDD
	phoneColl := mongoDB.Collection(config.CollPhones)
	phoneRepo := phoneRepository.NewRepository(mongoConn.DB, phoneColl)
	phoneTx := phoneTransaction.NewTransaction(mongoConn.DB, phoneRepo)
	phoneSrv := phoneService.NewService(phoneRepo, phoneTx)
	phoneHandler.NewHandler(mux, phoneSrv)

	// * Adapters layer
	branchAdapter := branchioAdapter.NewAdapter()
	resendAdapter := resendAdapter.NewAdapter(resendClient)
	//authAdapter := authAdapter.NewAdapter(googleVerifier.Verifier)
	authAdp := authAdapter.NewAdapter()
	//twilioAdp := twilioAdapter.NewTwilioAdapter(twilioClient)

	// * transaction layer
	authMongoTx := authPGtx.NewTransaction(mongoConn.DB, authMongoRepo, userMongoRepo, roleMongoRepo, otpRepository)

	// * Service layer
	// cambiar de token a session
	sessionSrv := sessionSrv.NewSessionSrv(authMongoRepo, userMongoRepo, roleMongoRepo, sessionMongoRepo)
	emailSrv := emailService.NewEmailSrv(branchAdapter, resendAdapter)
	//twilioSrv := twilioService.NewSmsSrv(twilioAdp)

	//authAdapter
	authSrv := authSrv.NewService(userMongoRepo, authMongoRepo, roleMongoRepo, otpRepository, sessionMongoRepo, phoneRepo, sessionSrv, emailSrv, authMongoTx, authAdp)

	minioAdp := minioAdapter.NewAdapter(minioClient.Client)
	fileStgSrv := fileStgService.NewFileStgSrv(minioAdp)

	// ** Register handler
	authHandler.NewHandler(mux, authSrv, sessionSrv)

	// * Product Item Service

	// ** Collections
	productItemColl := mongoDB.Collection(config.CollProductItems)

	// ** Repository layer
	productItemRepo := productItemRepository.NewRepository(mongoConn.DB, productItemColl)
	productItemSrv := productItemService.NewService(productItemRepo, fileStgSrv)

	// ** Register Handler
	productItemHandler.NewHandler(mux, productItemSrv)

	// * ProductConfig
	// * Collection
	productConfigColl := mongoDB.Collection(config.CollProductConfig)

	// *Repository
	productConfigRepo := productConfigRepository.NewRepository(mongoConn.DB, productConfigColl)

	// * Product
	// * Collection
	productsColl := mongoDB.Collection(config.CollProducts)

	// * Repository
	productRepo := productRepository.NewRepository(mongoConn.DB, productsColl)

	// * Transaction
	productTx := productTransaction.NewTransaction(mongoConn.DB, productRepo, productItemRepo, productConfigRepo)

	// * Service
	productSrv := productService.NewService(productRepo, fileStgSrv, productTx)

	productHandler.NewHandler(mux, productSrv)

	// * Categories
	// ** Collections
	categoryColl := mongoDB.Collection(config.CollCategories)

	categoryRepo := categoryRepository.NewCategoryRepo(mongoConn.DB, categoryColl)

	categorySrv := categoryService.NewService(categoryRepo)

	categoryHandler.NewHandler(mux, categorySrv)

	// * Variation
	// ** collections
	variationColl := mongoDB.Collection(config.CollVariations)

	variationRepo := variationRepository.NewRepository(mongoConn.DB, variationColl)

	variationSrv := variationService.NewService(variationRepo)

	variationHandler.NewHandler(mux, variationSrv)

	// * Payment
	// ** Collection
	paymentColl := mongoDB.Collection(config.CollPayments)

	paymentRepo := paymentRepository.NewRepository(mongoConn.DB, paymentColl)

	paymentSrv := paymentService.NewPaymentSrv(paymentRepo)

	// * Pickup
	pickupColl := mongoDB.Collection(config.CollPickup)
	pickupRepo := pickupRepository.NewPickupRepo(mongoConn.DB, pickupColl)

	// * Delivery
	deliveryColl := mongoDB.Collection(config.CollDelivery)
	deliveryRepo := deliveryRepository.NewDeliveryRepo(mongoConn.DB, deliveryColl)

	// * Order
	// ** Collection
	orderColl := mongoDB.Collection(config.CollOrders)
	orderProductColl := mongoDB.Collection(config.CollOrderProduct)

	orderRepo := orderRepository.NewRepository(mongoConn.DB, orderColl, orderProductColl)

	orderTx := orderTransaction.NewTransaction(mongoConn.DB, orderRepo, paymentRepo, pickupRepo, deliveryRepo)

	orderSrv := orderService.NewService(orderRepo, orderTx)

	orderHandler.NewHandler(mux, orderSrv)

	// * StripeCustomer
	stripeCustomerColl := mongoDB.Collection(config.CollStripeCustomer)
	pProviderRepo := pProviderRepository.NewPaymentProviderRepo(mongoConn.DB, stripeCustomerColl)

	// * Payment Stripe Service
	// ! separar por cada proveedor de pagos
	paymentProviderSrv := paymentProviderService.NewService(pProviderRepo)

	paymentHandler.NewHandler(mux, paymentProviderSrv, paymentSrv)

	// * Addresses
	addressColl := mongoDB.Collection(config.CollAddress)
	// ! dejar los nombres asi para su b´usqueda rápida
	addressRepo := addressRepository.NewAddressRepo(mongoConn.DB, addressColl)
	addresstx := addressTransaction.NewTransaction(mongoConn.DB, addressRepo)
	addressSrv := addressService.NewService(addressRepo, addresstx)
	addressHandler.NewHandler(mux, addressSrv)

	// * Stores
	// * Collection
	storeColl := mongoDB.Collection(config.CollStores)
	storeRepo := storeRepository.NewRepository(mongoConn.DB, storeColl)
	storeTx := mongo.NewStoreTx(mongoConn.DB, storeRepo, addressRepo)
	storeSrv := storeService.NewService(storeRepo, storeTx)

	storeHandler.NewHandler(mux, storeSrv)

	// * Onboarding
	onboardingColl := mongoDB.Collection(config.CollOnboarding)
	onboardingRepo := onboardingRepository.NewRepository(mongoConn.DB, onboardingColl)
	onboardingSrv := onboardingService.NewService(onboardingRepo, fileStgSrv)
	onboardingHandler.NewHandler(mux, onboardingSrv)

	// * Initial data
	// si no dependerái de auth service? si la funcion que necesito,
	// verificar por que solo usao una funcion y creo que tranquilamente iria en  config sin
	//usar authSrv
	initSrv := initConf.NewInitConfig()
	initSrv.Run()

	// * sería otr injeccion de dependencias pr cambiar fcilmente de mongo a postgresql
	// * que retornen authService y session service

	// * Ping Pong

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Security-Policy", "default-src 'self'; font-src 'self' https://assets.ngrok.com")
		w.WriteHeader(200)
		type resp struct {
			Msg string `json:"pong"`
		}
		json.NewEncoder(w).Encode(resp{Msg: "pong"})
	})

	mux.HandleFunc("/protected", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "appication/json")
		w.WriteHeader(200)
		type resp struct {
			Msg string `json:"message"`
		}
		json.NewEncoder(w).Encode(resp{Msg: "protected handler"})
	})

	// la idea es donde solo existe un admin - gerente y todo elle ya es para otro repo
	// exclusivo para ecommerce
	// Roles - ver cuales son solo admin
	// Create - /roles                 -  Crear role                - Role assigned successfully" - solo admin
	// Get    - /users/:id/roles       - Obtener roles de un usurio
	// Post   - /users/:id/roles       - Asignar role a un usuario   - "Role assigned successfully"
	// Delete - /users/:id/roles/:role - eliminar role de un usuario - "Role removed successfully"

	return mux
}
