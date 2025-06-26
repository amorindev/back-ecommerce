package v2

//"com.fernando/internal/config"
//mongoClient "com.fernando/internal/mongo"
//pgClient "com.fernando/internal/postgresql"
//resendClient "com.fernando/pkg/resend"
//authAdapter "com.fernando/pkg/services/app/auth/adapter"
//authHandler "com.fernando/pkg/services/app/auth/api/handler"
//authPGtx "com.fernando/pkg/services/app/auth/transaction/mongo"
//userMongo "com.fernando/pkg/services/app/user/repository/mongo"

//authPG "com.fernando/pkg/services/app/auth/repository/postgresql"
//rolePG "com.fernando/pkg/services/app/role/repository/postgresql"

//authMongo "com.fernando/pkg/services/app/auth/repository/mongo"
//authSrv "com.fernando/pkg/services/app/auth/service"
//initConf "com.fernando/pkg/services/app/init"

//roleMongo "com.fernando/pkg/services/app/role/repository/mongo"
//sessionMongo "com.fernando/pkg/services/app/session/repository/mongo"

//sessionPG "com.fernando/pkg/services/app/session/repository/postgresql"
//sessionSrv "com.fernando/pkg/services/app/session/service"
//branchioAdapter "com.fernando/pkg/services/branchio/adapter"
//resendAdapter "com.fernando/pkg/services/email/adapter/resend"
//emailSrv "com.fernando/pkg/services/email/service"

/* func New() http.Handler {
	mongoDBName := os.Getenv("MONGO_DB_NAME")
	if mongoDBName == "" {
		log.Fatal("environment variable MONGO_DB_NAME is not set")
	}

	mux := http.NewServeMux()

	// otra forma para llamar una sola ves a os.getenv
	// crear un objeto y llamarlo desde

	// c, err := claim.GetConfig()
	//if err != nil {
		// no habrría necesidad de que retorne un error si antes de valida que las varibles de
		// entorno esten configuradas
		//log.Fatal("err")
		//}

	// c debería pasarse  igual a Conn, para no
	// depndendiendo del servicio crear nuevos objetos

	// crear database?- donde crear base datos antes de server o despues
	// * Clients
	// si se puede llamar desde fuera del servidor, estaría mejor por que no carga en memoria?
	// igaul el init Run()
	pgConn := pgClient.New()
	mongoConn := mongoClient.New()
	// sacar del .env
	// mongoDB causará conflicto?
	mongoDatabase := mongoConn.DB.Database(mongoDBName)
	// o como la base de datos
	resendClient := resendClient.NewResendClient()

	googleVerifier, err := authAdapter.NewGoogleProvider()
	if err != nil {
		log.Fatal(err)
	}
	appleVerifier, err := authAdapter.NewAppleProvider("","","","")
	if err != nil {
		log.Fatal(err)
	}

	// * Database ping
	pgConn.Ping()
	mongoConn.Ping()

	// * Repositories
	// Postgresql
	//authPGRepo := authPG.NewRepository(pgConn.DB)
	//rolePGRepo := rolePG.NewRepository(pgConn.DB)
	//sessionPGRepo := sessionPG.NewRepository(pgConn.DB)

	// MongoDB
	authMongoRepo := authMongo.NewRepository(mongoConn.DB, mongoDatabase.Collection(config.CollectionAuth), mongoDatabase.Collection(config.CollectionUser))
	userMongoRepo := userMongo.NewRepository(mongoConn.DB, mongoDatabase.Collection(config.CollectionUser), mongoDatabase.Collection(config.CollectionUserRole))
	roleMongoRepo := roleMongo.NewRepository(mongoConn.DB, mongoDatabase.Collection(config.CollectionRole), mongoDatabase.Collection(config.CollectionUserRole))
	sessionMongoRepo := sessionMongo.NewRepository(mongoConn.DB, mongoDatabase.Collection(config.CollectionSessions))

	// * Adapters
	branchAdapter := branchioAdapter.NewAdapter()
	resendAdapter := resendAdapter.NewAdapter(resendClient)
	authAdapter := authAdapter.NewAdapter(googleVerifier.Verifier,appleVerifier)

	// * transaction layer
	authMongoTx := authPGtx.NewTransacion(mongoConn.DB, authMongoRepo, userMongoRepo)

	// * Services
	// cambiar de token a session
	sessionSrv := sessionSrv.NewService(authMongoRepo, roleMongoRepo, sessionMongoRepo)
	emailSrv := emailSrv.NewService(branchAdapter, resendAdapter)
	authSrv := authSrv.NewService(authMongoRepo, roleMongoRepo, userMongoRepo, authAdapter, emailSrv, sessionSrv, authMongoTx)

	// * Initial data
	// si no dependerái de auth service? si la funcion que necesito,
	// verificar por que solo usao una funcion y creo que tranquilamente iria en  config sin
	//usar authSrv
	initSrv := initConf.NewInitConfig()
	initSrv.Run()

	// * Ping Pong

	authHandler.NewHandler(mux, authSrv, sessionSrv)

	return mux
}
*/
